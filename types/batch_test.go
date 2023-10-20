package types

import (
	"crypto/ecdsa"
	"testing"

	"github.com/0xPolygon/cdk-data-availability/rpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testBatchCases = []Batch{
	{
		L2Data:         common.Hex2Bytes(""),
		GlobalExitRoot: common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		Timestamp:      1685130880,
		Coinbase:       common.HexToAddress("0x1f1CCb74C0e2e84E657f4C4b7f1F2CDC3D2114d5"),
	},
	{
		L2Data:         rpc.ArgBytes(common.Hex2Bytes("f9056880808301ba5a945904257b03bb7bf01a704a233090cc904e1c164580b905442cffd02e0000000000000000000000000000000000000000000000000000000000000000ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5b4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d3021ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba85e58769b32a1beaf1ea27375a44095a0d1fb664ce2dd358e7fcbfb78c26a193440eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d887c22bd8750d34016ac3c66b5ff102dacdd73f6b014e710b51e8022af9a1968ffd70157e48063fc33c97a050f7f640233bf646cc98d9524c6b92bcf3ab56f839867cc5f7f196b93bae1e27e6320742445d290f2263827498b54fec539f756afcefad4e508c098b9a7e1d8feb19955fb02ba9675585078710969d3440f5054e0f9dc3e7fe016e050eff260334f18a5d4fe391d82092319f5964f2e2eb7c1c3a5f8b13a49e282f609c317a833fb8d976d11517c571d1221a265d25af778ecf8923490c6ceeb450aecdc82e28293031d10c7d73bf85e57bf041a97360aa2c5d99cc1df82d9c4b87413eae2ef048f94b4d3554cea73d92b0f7af96e0271c691e2bb5c67add7c6caf302256adedf7ab114da0acfe870d449a3a489f781d659e8beccda7bce9f4e8618b6bd2f4132ce798cdc7a60e7e1460a7299e3c6342a579626d22733e50f526ec2fa19a22b31e8ed50f23cd1fdf94c9154ed3a7609a2f1ff981fe1d3b5c807b281e4683cc6d6315cf95b9ade8641defcb32372f1c126e398ef7a5a2dce0a8a7f68bb74560f8f71837c2c2ebbcbf7fffb42ae1896f13f7c7479a0b46a28b6f55540f89444f63de0378e3d121be09e06cc9ded1c20e65876d36aa0c65e9645644786b620e2dd2ad648ddfcbf4a7e5b1a3a4ecfe7f64667a3f0b7e2f4418588ed35a2458cffeb39b93d26f18d2ab13bdce6aee58e7b99359ec2dfd95a9c16dc00d6ef18b7933a6f8dc65ccb55667138776f7dea101070dc8796e3774df84f40ae0c8229d0d6069e5c8f39a7c299677a09d367fc7b05e3bc380ee652cdc72595f74c7b1043d0e1ffbab734648c838dfb0527d971b602bc216c9619ef0abf5ac974a1ed57f4050aa510dd9c74f508277b39d7973bb2dfccc5eeb0618db8cd74046ff337f0a7bf2c8e03e10f642c1886798d71806ab1e888d9e5ee87d0838c5655cb21c6cb83313b5a631175dff4963772cce9108188b34ac87c81c41e662ee4dd2dd7b2bc707961b1e646c4047669dcb6584f0d8d770daf5d7e7deb2e388ab20e2573d171a88108e79d820e98f26c0b84aa8b2f4aa4968dbb818ea32293237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d7358448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a9000000000000000000000000000000000000000000000000000000000000000095b76b43196f04d39cc922a7933240fa6cfc3fb1502ecfdf693e0ad078f668e100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000001f1ccb74c0e2e84e657f4c4b7f1f2cdc3d2114d5000000000000000000000000000000000000000000000000002386f26fc100000000000000000000000000000000000000000000000000000000000000000520000000000000000000000000000000000000000000000000000000000000000082084a8080213c87d62b9620bccb33eed6e22861ce4f112ec89f40a66111bbec222bfb313e37dc77d691d1436916b3dd8519eb3e572dd2962c7317732c38443d86390f988d1c")),
		GlobalExitRoot: common.HexToHash("0xbd6ad630a17775e238d412044d37868838687ba4c6346b917551f4f60def4a2a"),
		Timestamp:      1685349728,
		Coinbase:       common.HexToAddress("0x1f1CCb74C0e2e84E657f4C4b7f1F2CDC3D2114d5"),
	},
	{
		L2Data:         common.Hex2Bytes("f905688080830f4240945904257b03bb7bf01a704a233090cc904e1c164580b905442cffd02e0000000000000000000000000000000000000000000000000000000000000000ad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5b4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d3021ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba85e58769b32a1beaf1ea27375a44095a0d1fb664ce2dd358e7fcbfb78c26a193440eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d887c22bd8750d34016ac3c66b5ff102dacdd73f6b014e710b51e8022af9a1968ffd70157e48063fc33c97a050f7f640233bf646cc98d9524c6b92bcf3ab56f839867cc5f7f196b93bae1e27e6320742445d290f2263827498b54fec539f756afcefad4e508c098b9a7e1d8feb19955fb02ba9675585078710969d3440f5054e0f9dc3e7fe016e050eff260334f18a5d4fe391d82092319f5964f2e2eb7c1c3a5f8b13a49e282f609c317a833fb8d976d11517c571d1221a265d25af778ecf8923490c6ceeb450aecdc82e28293031d10c7d73bf85e57bf041a97360aa2c5d99cc1df82d9c4b87413eae2ef048f94b4d3554cea73d92b0f7af96e0271c691e2bb5c67add7c6caf302256adedf7ab114da0acfe870d449a3a489f781d659e8beccda7bce9f4e8618b6bd2f4132ce798cdc7a60e7e1460a7299e3c6342a579626d22733e50f526ec2fa19a22b31e8ed50f23cd1fdf94c9154ed3a7609a2f1ff981fe1d3b5c807b281e4683cc6d6315cf95b9ade8641defcb32372f1c126e398ef7a5a2dce0a8a7f68bb74560f8f71837c2c2ebbcbf7fffb42ae1896f13f7c7479a0b46a28b6f55540f89444f63de0378e3d121be09e06cc9ded1c20e65876d36aa0c65e9645644786b620e2dd2ad648ddfcbf4a7e5b1a3a4ecfe7f64667a3f0b7e2f4418588ed35a2458cffeb39b93d26f18d2ab13bdce6aee58e7b99359ec2dfd95a9c16dc00d6ef18b7933a6f8dc65ccb55667138776f7dea101070dc8796e3774df84f40ae0c8229d0d6069e5c8f39a7c299677a09d367fc7b05e3bc380ee652cdc72595f74c7b1043d0e1ffbab734648c838dfb0527d971b602bc216c9619ef0abf5ac974a1ed57f4050aa510dd9c74f508277b39d7973bb2dfccc5eeb0618db8cd74046ff337f0a7bf2c8e03e10f642c1886798d71806ab1e888d9e5ee87d0838c5655cb21c6cb83313b5a631175dff4963772cce9108188b34ac87c81c41e662ee4dd2dd7b2bc707961b1e646c4047669dcb6584f0d8d770daf5d7e7deb2e388ab20e2573d171a88108e79d820e98f26c0b84aa8b2f4aa4968dbb818ea32293237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d7358448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a9000000000000000000000000000000000000000000000000000000000000000095b76b43196f04d39cc922a7933240fa6cfc3fb1502ecfdf693e0ad078f668e100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000001f1ccb74c0e2e84e657f4c4b7f1f2cdc3d2114d5000000000000000000000000000000000000000000000000002386f26fc100000000000000000000000000000000000000000000000000000000000000000520000000000000000000000000000000000000000000000000000000000000000082084a80801ca5229fb6abcf5febe08c28268626cd1aa9adce9069410bbdf5cb958d52f23a15077324cb889bf4e37063540fcf4d1b9c7f523d54f2182fdc217de4e7a82b491b"),
		GlobalExitRoot: common.HexToHash("0xbd6ad630a17775e238d412044d37868838687ba4c6346b917551f4f60def4a2a"),
		Timestamp:      1685374947,
		Coinbase:       common.HexToAddress("0x1f1CCb74C0e2e84E657f4C4b7f1F2CDC3D2114d5"),
	},
	{
		L2Data:         rpc.ArgBytes(common.Hex2Bytes("e1808082520894400f2117d43f95f63a8e06dcf813dd746c4b5aba808082084a8080a6c23693c74f1492b1b1b580d0333e980532dcd0c7407a8c2134e28065bc8ceb273fd0b759cf1194ab221bb1dcada4b6f43251ae9909621152c450b062b4a3431ce1018082520894400f2117d43f95f63a8e06dcf813dd746c4b5aba808082084a80801c15bf03602f2907d0ab1fc027b98104c7625a36008b53c93a44191fadc68c305d9d9d65fdc2a108c71cd586ca7d58ec6da1de735aef6a0434196f966f777b191c")),
		GlobalExitRoot: common.HexToHash("0xbd6ad630a17775e238d412044d37868838687ba4c6346b917551f4f60def4a2a"),
		Timestamp:      1685374947,
		Coinbase:       common.HexToAddress("0x1f1CCb74C0e2e84E657f4C4b7f1F2CDC3D2114d5"),
	},
	{
		L2Data:         rpc.ArgBytes(common.Hex2Bytes("e102808252089445be386709a8f3a056651bbb32752e0c3bebd8f0808082084a8080b429fc3ab2e3a947dd69921696779a6acf85abeb858b2e918ed047df27ed389c2407a7550ae4df61c05eb4f4cf2a410ff064ca73e3437074f7685d6f849799921c")),
		GlobalExitRoot: common.HexToHash("0xbd6ad630a17775e238d412044d37868838687ba4c6346b917551f4f60def4a2a"),
		Timestamp:      1685375541,
		Coinbase:       common.HexToAddress("0x1f1CCb74C0e2e84E657f4C4b7f1F2CDC3D2114d5"),
	},
}

func TestBatchSigning(t *testing.T) {
	privKeys := []*ecdsa.PrivateKey{}
	for i := 0; i < 5; i++ {
		pk, err := crypto.GenerateKey()
		require.NoError(t, err)
		privKeys = append(privKeys, pk)
	}
	for _, c := range testBatchCases {
		for _, pk := range privKeys {
			signedBatch, err := c.Sign(pk)
			require.NoError(t, err)
			actualAddr, err := signedBatch.Signer()
			require.NoError(t, err)
			expectedAddr := crypto.PubkeyToAddress(pk.PublicKey)
			assert.Equal(t, expectedAddr, actualAddr)
		}
	}
}
