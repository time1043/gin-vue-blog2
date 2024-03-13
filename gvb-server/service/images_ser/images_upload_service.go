package images_ser

import (
	"fmt"
	"io"
	"mime/multipart"
	"path"
	"strings"
	"time1043/gvb-server/global"
	"time1043/gvb-server/models"
	"time1043/gvb-server/models/ctype"
	"time1043/gvb-server/plugins/qiniu"
	"time1043/gvb-server/utils"
)

var (
	// 图片上传的白名单
	WhiteImageList = []string{
		"jpg", "png", "jpeg", "ico", "tiff", "gif", "svg", "webp",
	}
)

// 有的图片上传成功有的失败 需要给用户返回响应
type FileUploadResponse struct {
	FileName  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 消息
}

// 文件上传的方法：向数据库添加数据、重复不添加
func (ImageService) ImageUploadService(file *multipart.FileHeader) (res FileUploadResponse) {
	filename := file.Filename
	basePath := global.Config.Upload.Path
	filePath := path.Join(basePath, filename)

	res.FileName = filePath

	// 判断尾缀白名单
	nameList := strings.Split(filename, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	if !utils.InList(suffix, WhiteImageList) {
		res.Msg = "非法文件"
		return
	}

	// 判断文件的大小 太大了就失败
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		res.Msg = fmt.Sprintf("图片大小超过设定大小，当前大小为：%.2f MB，设定大小为：%d MB", size, global.Config.Upload.Size)
		return
	}

	// md5检验图片内容
	fileObj, err := file.Open()
	if err != nil {
		global.Log.Error(err)
	}
	byteData, err := io.ReadAll(fileObj)
	imageHash := utils.Md5(byteData)
	// 根据md5值去数据库中查 是否重复
	var bannerModel models.BannerModel
	err = global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
	if err == nil { // 找到重复的
		res.Msg = "图片已存在"
		res.FileName = bannerModel.Path
		return
	}

	fileType := ctype.Local
	res.Msg = "图片上传成功"
	res.IsSuccess = true

	// 是否启用七牛云存储
	if global.Config.QiNiu.Enable {

		filePath, err = qiniu.UploadImage(byteData, filename, global.Config.QiNiu.Predix)
		if err != nil {
			global.Log.Error(err)
			res.Msg = err.Error()
			return
		}
		res.FileName = filePath
		res.Msg = "七牛云存储上传成功"
		fileType = ctype.QiNiu
	}

	// 图片写入数据库
	global.DB.Create(&models.BannerModel{
		Path:      filePath,
		Hash:      imageHash,
		Name:      filename,
		ImageType: fileType,
	})
	return
}
