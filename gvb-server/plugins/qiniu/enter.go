package qiniu

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"time"
	"time1043/gvb-server/config"
	"time1043/gvb-server/global"
)

// 获取上传的token
func getToken(q config.QiNiu) string {
	accessKey := q.AccessKey
	secretKey := q.SecretKey
	bucket := q.Bucket
	putPolicy := storage.PutPolicy{Scope: bucket}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}

// 获取上传的配置
func getCfg(q config.QiNiu) storage.Config {
	cfg := storage.Config{}

	// 空间对应的机房
	zone, _ := storage.GetRegionByID(storage.RegionID(q.Zone))
	cfg.Zone = &zone

	cfg.UseHTTPS = false      // 是否使用https域名
	cfg.UseCdnDomains = false // 上传是否使用CDN加速

	return cfg
}

// 上传图片：文件数组、前缀
func UploadImage(data []byte, imageName string, prefix string) (filePath string, err error) {
	if !global.Config.QiNiu.Enable {
		return "", errors.New("请启用七牛云存储")
	}

	q := global.Config.QiNiu
	if q.AccessKey == "" || q.SecretKey == "" {
		return "", errors.New("请配置 AccessKey 和 SecretKey")
	}
	if float64(len(data))/1024/1024 > q.Size {
		return "", errors.New("文件超过设定大小")
	}

	upToken := getToken(q)
	cfg := getCfg(q)

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	dataLen := int64(len(data))

	// 获取当前时间
	now := time.Now().Format("20060102150405")
	key := fmt.Sprintf("%s/%s__%s", prefix, now, imageName)

	formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", q.CDN, ret.Key), nil
}
