package verification

import (
	"encoding/hex"
	"testing"
)

func TestVerifyProof(t *testing.T) {
	type args struct {
		hash       string
		verifyKey  []byte
		proofBytes []byte
	}
	var vk []byte
	vk, _ = hex.DecodeString("1f1504803bdfb7c81c7247fcbd11eb81ef02db53a99b5d39e7b2d6e0257f9144211d59634275e925cd099fa7d9d9121a3d2f8cc28a82bc0819a1bcbb4d9b64b311f54739ba18bbf61947d2d91c6e2a5208ff9c74671ea7a1e7f531bd3027c97c27d7f72a2f9d5337b7925c845c5fbdfe8487df71aa55d9bdd9d0c01efddf59ed2eb2275afe626e9752eb3165cccb21b70eb6887011e3b5ed84b2d6d65c17ca1714b881a84764bc55f884600958220ff7808b22319fa880e71f178aa25b0a963119b2910cc9b3d77f195fc26d7986af3ac5ff7d51f48602d6257077934eb8738a2691a5819de91c0e8f25aeac32bc4e06c7e9c9913b60b5ad9be46b6aff6cdfa80280b3c5e1208bb901484ba0fa3f2a156b8ecd1a27c332c06711ad11b9d8267112b25cf9b659a92c49a153ab29a36a97f2f19806f9dedc2fe1eb6acdc6a329621dbd26de182d9f3f009231fae06764b5c83a6211f12b8a62ce50182bb3f91ddd0b96ca3e88cd5e22843bfd13c453d776be8cd25fa5e82e88b7a185d44dd67091129a7d3a2671a2dc03f9f5b97e2a17403a89cf26870d102fb3549f1b574b5f77263b0d16efde0d0020dfe8e5a8185df6ffd31bb08f921882be5cf99eff6a022f014dcc3a7890e6ddea8277498798d0ee37be019c4e5bb636b64166fadf6b0c420d9bd725c9168f0804dc01df9a457f4f49b0aef53b49b6fe22d3e5fd5b08ec7f118e54589fc366d0f0816b1f5c6bb0b0d0b7b1d08b208d3abfbca8e23906ca9609479308c6d1c1c81f1cbe2ab2417048f6b88adce2f671022581488f6715cb3f0000000202c398350c5aba23aa3d06ea702d4ec07e322f0adc6ee281e9c03365ce4305e80a0bdc411143718622de3676b0fdb017f47c1dc106ee6b7b35ade63d732a217d0d759ab4a5cc3cb586e91b82fbc3801b92e05d667ea2f015de79c4bcb4cdbda1060f25b02ced44315f9c96cde52ce44c1787c6d156fe617e7eb2655b331c0431")
	var proof []byte
	proof, _ = hex.DecodeString("1d3dd51cf2baef780851a75d164e60a92f1c52b3b9107912495327a42f6bd8de0ab8b51224fc9e94a65f1e9592ae378ac6bc00b8e49b0373bbfed5cab5f25e5f02311175e0db827c53cac5473f82b9ac5602e4dd9b6bb98333a4dc229e729f600b9c9f79b456ac38e265cc60c986cc3f3c96edd9eec3a822e3b840856d3c70a4008ec5ac14454a68b1dba507cceb375ef269245741a5c1e4bd1952927c6b4cdc2948dbe55a9b7f0cf33ad79607fb45fa5a5acd92675ccc2f9284935de07836511b380ace4c984b7786fabf29f2c5832c2ce5f5c3a924b8c24c3f8887938db90708982276f63a20d6015edf125fa4445f5043b505700b02a90fc8a10070380861")
	var tests = []struct {
		name       string
		verifyPara args
		want       bool
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "success example 1",
			verifyPara: args{"8674594860895598770446879254410848023850744751986836044725552747672873438975",
				vk,
				proof,
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VerifyProof(tt.verifyPara.hash, tt.verifyPara.verifyKey, tt.verifyPara.proofBytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyProof() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VerifyProof() got = %v, want %v", got, tt.want)
			}
		})
	}
}
