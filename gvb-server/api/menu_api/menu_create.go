package menu_api

import (
	"github.com/gin-gonic/gin"
	"time1043/gvb-server/global"
	"time1043/gvb-server/models"
	"time1043/gvb-server/models/ctype"
	"time1043/gvb-server/models/res"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	MenuTitle     string      `json:"menu_title" binding:"required" msg:"请完善菜单名称"`      // 菜单名称
	MenuTitleEn   string      `json:"menu_title_en" binding:"required" msg:"请完善菜单英文名称"` // 菜单英文名称
	Slogan        string      `json:"slogan"`                                           // slogan
	Abstract      ctype.Array `json:"abstract"`                                         // 简介
	AbstractTime  int         `json:"abstract_time"`                                    // 简介的切换时间
	BannersTime   int         `json:"banners_time"`                                     // 菜单图片的切换时间 0表示不切换
	Sort          int         `json:"sort" binding:"required" msg:"请输入菜单序号"`            // 菜单的顺序
	ImageSortList []ImageSort `json:"image_sort_list"`                                  // 具体图片的顺序
}

func (MenuApi) MenuCreateView(ctx *gin.Context) {
	var cr MenuRequest
	err := ctx.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, ctx)
		return
	}

	// 重复判断

	// 数据库操作：创建banner数据入库
	menuModel := models.MenuModel{
		MenuTitle:    cr.MenuTitle,
		MenuTitleEn:  cr.MenuTitleEn,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannersTime:  cr.BannersTime,
		Sort:         cr.Sort,
	}
	err = global.DB.Create(&menuModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("菜单添加失败", ctx)
		return
	}
	if len(cr.ImageSortList) == 0 {
		res.OkWithMessage("没有菜单需要添加", ctx)
		return
	}

	var menuBannerList []models.MenuBannerModel
	for _, sort := range cr.ImageSortList {
		// 判断image_id对应的图片是否存在...
		menuBannerList = append(menuBannerList, models.MenuBannerModel{
			MenuID:   menuModel.ID,
			BannerID: sort.ImageID,
			Sort:     sort.Sort,
		})
	}

	// 数据库操作：给第三张表入库
	err = global.DB.Create(&menuBannerList).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("菜单关联失败", ctx)
		return
	}
	res.OkWithMessage("菜单添加成功", ctx)
}
