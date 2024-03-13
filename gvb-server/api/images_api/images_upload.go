package images_api

import (
	"github.com/gin-gonic/gin"
	"io/fs"
	"os"
	"time1043/gvb-server/global"
	"time1043/gvb-server/models/res"
	"time1043/gvb-server/service"
	"time1043/gvb-server/service/images_ser"
)

// 上传图片 返回图片的url
func (ImagesApi) ImageUploadView(ctx *gin.Context) {
	form, err := ctx.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), ctx)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("不存在文件", ctx)
		return
	}

	// 判断路径是否存在 不存在则创建
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}

	var resList []images_ser.FileUploadResponse
	for _, file := range fileList {
		// 上传文件 封装成ServiceApp
		ServiceRes := service.ServiceApp.ImageService.ImageUploadService(file)
		if !ServiceRes.IsSuccess {
			resList = append(resList, ServiceRes)
			continue
		}

		if !global.Config.QiNiu.Enable {
			// 本地文件写入 (拦截的前面都干过了)
			err = ctx.SaveUploadedFile(file, ServiceRes.FileName)
			if err != nil {
				global.Log.Error(err)
				ServiceRes.Msg = err.Error()
				ServiceRes.IsSuccess = false
				resList = append(resList, ServiceRes)
				continue
			}
		}

		resList = append(resList, ServiceRes)
	}

	res.OkWithData(resList, ctx)

}
