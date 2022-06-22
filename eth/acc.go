package eth

import "github.com/ethereum/go-ethereum/crypto"

func NewAcc(pass string) (string, error) {
	// 调用钱包创建
	/*
		if w, err := wallet.NewWallet(""); err != nil {
			fmt.Println("wallet error")
		}
	*/
	// 存储私钥
	//if err := w.StoreKey(pass); err != nil {
	//	fmt.Println("wallet error")
	//}
	//return w.Address.Hex(), nil
	return "", nil
}

func KeccakHash(data []byte) []byte {
	return crypto.Keccak256(data)
}
