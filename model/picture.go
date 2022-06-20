package model

type Picture struct {
	Title       string `json:"title"`       // 图片原名称
	ContentPath string `json:"contentPath"` // 图片保存路径
	ContentHash string `json:"contentHash"` // 图片hash
	Address     string `json:"address"`     // 图片上传用户地址
	TokenId     string `json:"tokenId"`     // 图片token id
}
