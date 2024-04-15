package sequencer

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/0xPolygon/cdk-data-availability/config"
	"github.com/0xPolygon/cdk-data-availability/etherman"
	"github.com/0xPolygon/cdk-data-availability/etherman/smartcontracts/cdkvalidium"
	"github.com/0xPolygon/cdk-data-availability/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

const maxConnectionRetries = 5

// SequencerTracker watches the contract for relevant changes to the sequencer
type SequencerTracker struct {
	client       *etherman.Etherman
	stop         chan struct{}
	timeout      time.Duration
	retry        time.Duration
	addr         common.Address
	url          string
	lock         sync.Mutex
	usePolling   bool
	pollInterval time.Duration
}

// NewSequencerTracker creates a new SequencerTracker
func NewSequencerTracker(cfg config.L1Config, ethClient *etherman.Etherman) (*SequencerTracker, error) {
	log.Info("starting sequencer address tracker")
	addr, err := ethClient.TrustedSequencer()
	if err != nil {
		return nil, err
	}
	log.Infof("current sequencer addr: %s", addr.Hex())
	url, err := ethClient.TrustedSequencerURL()
	if err != nil {
		return nil, err
	}
	log.Infof("current sequencer url: %s", url)
	pollInterval := time.Minute
	if cfg.TrackSequencerPollInterval.Seconds() > 0 {
		pollInterval = cfg.TrackSequencerPollInterval.Duration
	}
	w := &SequencerTracker{
		client:       ethClient,
		stop:         make(chan struct{}),
		timeout:      cfg.Timeout.Duration,
		retry:        cfg.RetryPeriod.Duration,
		addr:         addr,
		url:          url,
		usePolling:   strings.HasPrefix(cfg.RpcURL, "http"),
		pollInterval: pollInterval,
	}
	return w, nil
}

// GetAddr returns the last known address of the Sequencer
func (st *SequencerTracker) GetAddr() common.Address {
	st.lock.Lock()
	defer st.lock.Unlock()
	return st.addr
}

func (st *SequencerTracker) setAddr(addr common.Address) {
	st.lock.Lock()
	defer st.lock.Unlock()
	st.addr = addr
}

// GetUrl returns the last known URL of the Sequencer
func (st *SequencerTracker) GetUrl() string {
	st.lock.Lock()
	defer st.lock.Unlock()
	return st.url
}

func (st *SequencerTracker) setUrl(url string) {
	st.lock.Lock()
	defer st.lock.Unlock()
	st.url = url
}

// Start starts the SequencerTracker
func (st *SequencerTracker) Start() {
	go st.trackAddrChanges()
	go st.trackUrlChanges()
}

func (st *SequencerTracker) trackAddrChanges() {
	addrChan := make(chan common.Address, 1)
	ctx, cancel := context.WithCancel(context.Background())
	if st.usePolling {
		go st.pollAddrChanges(ctx, addrChan)
	} else {
		go st.subscribeOnAddrChanges(ctx, addrChan)
	}
	for {
		select {
		case addr := <-addrChan:
			if st.GetAddr().String() != addr.String() {
				log.Infof("new trusted sequencer address: %v", addr)
				st.setAddr(addr)
			}
		case <-ctx.Done():
			return
		case <-st.stop:
			cancel()
			return
		}
	}
}

func (st *SequencerTracker) pollAddrChanges(ctx context.Context, addrChan chan<- common.Address) {
	ticker := time.NewTicker(st.pollInterval)
	for {
		select {
		case <-ticker.C:
			addr, err := st.client.TrustedSequencer()
			if err != nil {
				log.Errorf("failed to get sequencer addr: %v", err)
				continue
			}
			if err == nil && st.GetAddr().String() != addr.String() {
				addrChan <- addr
			}
		case <-ctx.Done():
			ticker.Stop()
			return
		case <-st.stop:
			ticker.Stop()
			return
		}
	}
}

func (st *SequencerTracker) subscribeOnAddrChanges(ctx context.Context, addrChan chan<- common.Address) {
	events := make(chan *cdkvalidium.CdkvalidiumSetTrustedSequencer)
	defer close(events)

	var sub event.Subscription

	initSubscription := func() {
		err := exponential(func() (err error) {
			ctx, _ := context.WithTimeout(context.Background(), st.timeout)
			opts := &bind.WatchOpts{Context: ctx}
			sub, err = st.client.CDKValidium.WatchSetTrustedSequencer(opts, events)
			if err != nil {
				log.Errorf("error subscribing to trusted sequencer event, retrying: %v", err)
			}
			return err
		}, maxConnectionRetries, st.retry)
		if err != nil {
			log.Fatalf("failed subscribing to trusted sequencer event: %v. Check ws(s) availability.", err)
		}
	}
	initSubscription()

	for {
		select {
		case e := <-events:
			log.Infof("new trusted sequencer address: %v", e.NewTrustedSequencer)
			addrChan <- e.NewTrustedSequencer
		case err := <-sub.Err():
			log.Warnf("subscription error, resubscribing: %v", err)
			initSubscription()
		case <-ctx.Done():
			return
		case <-st.stop:
			if sub != nil {
				sub.Unsubscribe()
			}
			return
		}
	}
}

func (st *SequencerTracker) trackUrlChanges() {
	urlChan := make(chan string, 1)
	ctx, cancel := context.WithCancel(context.Background())
	if st.usePolling {
		go st.pollUrlChanges(ctx, urlChan)
	} else {
		go st.subscribeOnUrlChanges(ctx, urlChan)
	}
	for {
		select {
		case url := <-urlChan:
			if st.GetUrl() != url {
				log.Infof("new trusted sequencer url: %v", url)
				st.setUrl(url)
			}
		case <-ctx.Done():
			return
		case <-st.stop:
			cancel()
			return
		}
	}
}

func (st *SequencerTracker) subscribeOnUrlChanges(ctx context.Context, urlChan chan<- string) {
	events := make(chan *cdkvalidium.CdkvalidiumSetTrustedSequencerURL)
	defer close(events)

	var sub event.Subscription

	initSubscription := func() {
		err := exponential(func() (err error) {
			ctx, _ := context.WithTimeout(context.Background(), st.timeout)
			opts := &bind.WatchOpts{Context: ctx}
			sub, err = st.client.CDKValidium.WatchSetTrustedSequencerURL(opts, events)
			if err != nil {
				log.Errorf("error subscribing to trusted sequencer URL event, retrying: %v", err)
			}
			return err
		}, maxConnectionRetries, st.retry)
		if err != nil {
			log.Fatalf("failed subscribing to trusted sequencer URL event: %v. Check ws(s) availability.", err)
		}
	}
	initSubscription()

	for {
		select {
		case e := <-events:
			urlChan <- e.NewTrustedSequencerURL
		case <-ctx.Done():
			return
		case err := <-sub.Err():
			log.Warnf("subscription error, resubscribing: %v", err)
			initSubscription()
		case <-st.stop:
			if sub != nil {
				sub.Unsubscribe()
			}
			return
		}
	}
}

func (st *SequencerTracker) pollUrlChanges(ctx context.Context, urlChan chan<- string) {
	ticker := time.NewTicker(st.pollInterval)
	for {
		select {
		case <-ticker.C:
			url, err := st.client.TrustedSequencerURL()
			if err != nil {
				log.Errorf("failed to get sequencer URL: %v", err)
				continue
			}
			if err == nil && st.GetUrl() != url {
				urlChan <- url
			}
		case <-ctx.Done():
			ticker.Stop()
			return
		case <-st.stop:
			ticker.Stop()
			return
		}
	}
}

// Stop stops the SequencerTracker
func (st *SequencerTracker) Stop() {
	close(st.stop)
}

func exponential(action func() error, max uint, wait time.Duration) error {
	var err error
	for i := uint(0); i < max; i++ {
		if err = action(); err == nil {
			return nil
		}
		time.Sleep(wait)
		wait *= 2
	}
	return err
}
