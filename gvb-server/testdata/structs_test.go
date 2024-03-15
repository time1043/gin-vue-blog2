package testdata

import (
	"fmt"
	"github.com/fatih/structs"
	"testing"
	"time1043/gvb-server/models"
)

type AdvertRequest struct {
	models.MODEL `structs:"-"`
	Title        string `json:"title" binding:"required" msg:"请输入标题" structs:"title"`   // 显示的标题
	Href         string `json:"href" binding:"required,url" msg:"跳转链接非法" `              // 跳转链接
	Images       string `json:"images" binding:"required,url" msg:"图片地址非法" structs:"-"` // 图片
	IsShow       bool   `json:"is_show" msg:"请选择是否展示" `                                 // 是否展示
}

func TestStructs(t *testing.T) {
	u1 := AdvertRequest{
		Title:  "xxx",
		Href:   "xxx",
		Images: "xxx",
		IsShow: true,
	}
	m1 := structs.Map(&u1)
	fmt.Println(m1)
}
