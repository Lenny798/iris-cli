package util

import (
	"math/big"
	"time"
)

// TokenReader token 生成器
type TokenReader struct {
	fc func() int64
}

// 全局结构
var tr *TokenReader

// init 自动初始化
func init() {
	tr = NewTokenReader(NextId())
}

// NewTokenReader 构造token生成器
func NewTokenReader(f func() int64) *TokenReader {
	return &TokenReader{fc: f}
}

// 闭包函数
func NextId() func() int64 {
	var index int64
	return func() int64 {
		index++
		if index > 10000 {
			index = 0
		}
		return index
	}
}

func NewTokenId() int64 {
	val1 := big.NewInt(tr.fc())
	val2 := big.NewInt(0)
	val2, _ = val2.SetString(time.Now().Format("20060102150405"), 10)
	return val2.Int64()*10000 + val1.Int64()
}
