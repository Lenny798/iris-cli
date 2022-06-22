package eth

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"iris-cli/wallet"
	"log"
	"math/big"
)

const (
	PxaAddr = "0x36828345e39ADe9b9645aB993d179AA40a98cbf1" // 合约地址
	PxcAddr = "0x9c822d4AE6E5Ec17B8b9a8A2b9032CcFfF0192Ad" // 合约地址
)

// Geth 客户端全局连接句柄
var ethCli *ethclient.Client

//pxa合约全局实例
var instancePXA *Pxa721

// pxc合约全局实例
var instancePXC *Pxc20

func init() {
	cli, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Panic("Failed to ethclient.Dial ", err)
	}
	ethCli = cli
	// pxa 合约实例
	itcPxa, err := NewPxa721(common.HexToAddress(PxaAddr), cli)
	if err != nil {
		log.Panic("NewPxa721 error", err)
	}
	instancePXA = itcPxa

	// pxc 合约全局实例
	pxc20, err := NewPxc20(common.HexToAddress(PxcAddr), cli)
	if err != nil {
		log.Panic("NewPxa20 error", err)
	}
	instancePXC = pxc20
}

func UploadPic(from, pass, to string, tokenId *big.Int) error {
	// 3 设置签名 -- 需要owner 的 keystore文件
	w, err := wallet.LoadWalletByPass(from, "./data", pass)
	if err != nil {
		return err
	}
	auth := w.HdKeyStore.NewTransactOpts()
	// 4 调用
	_, err = instancePXA.UploadMint(auth, common.HexToAddress(to), tokenId)
	if err != nil {
		return err
	}
	return nil
}
