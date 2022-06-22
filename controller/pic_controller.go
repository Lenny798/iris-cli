package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/spf13/cast"
	"iris-cli/core"
	"iris-cli/eth"
	"iris-cli/model"
	"iris-cli/repo"
	"iris-cli/util"
	"math/big"
)

type PicController struct {
	// iris 框架自动为每个请求都绑定上下文对象：可作为接受参数
	Context iris.Context
	// session 对象： 储存session信息
	Session *sessions.Session

	// admin功能实体：引入Service 接口
	UserRepo repo.UserRepo
}

func NewPicController() *PicController {
	return &PicController{}
}

func (dc *PicController) Add() (result *core.Result) {
	pic := &model.Picture{}
	file, header, err := dc.Context.FormFile("fileName")
	if err != nil {
		return result
	}
	src, err := header.Open()
	defer src.Close()

	// 2.2获取token id
	tokenId := util.NewTokenId()
	pic.TokenId = cast.ToString(tokenId)
	// 准备工作做好之后，就可以把读到的数据写入服务器的文件中，同时利用文件内容计算一下hash值

	// 2。3计算hash值
	fData := make([]byte, header.Size)
	read, _ := src.Read(fData)
	hash := eth.KeccakHash(fData)
	pic.ContentHash = cast.ToString(hash)
	// TODO 写入文件 把fData写入服务器文件
	// 3 获取到用户address(注册账户时，通过钱包生成的一个hash)
	address := "hash"
	pic.Address = address
	// 4.操作mysql 存储
	//picRepo.Add()
	// 5.操作以太坊
	eth.UploadPic(pic.Address, pass, pic.Address, big.NewInt(tokenId))
	return result
}

func (dc *PicController) GetPic() (result *core.Result) {
	// TODO 查看图片直接从mysql 中获取
	return result
}
