package service

import (
	"context"
	"mime/multipart"

	"github.com/qiniu/api.v7/v7/auth/qbox"

	"github.com/spf13/viper"

	"github.com/qiniu/api.v7/v7/storage"
)

func UploadFile(file multipart.File, fileSize int64) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope: viper.GetString("qiniu.bucket"),
	}
	mac := qbox.NewMac(viper.GetString("qiniu.access-key"), viper.GetString("qiniu.secret-key"))
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}
	putExtra := storage.PutExtra{Params: map[string]string{"x:name": "github logo"}}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	url := viper.GetString("qiniu.img-path") + "/" + ret.Key
	return url, err
}

//
//func UpLoadFile(file multipart.File, fileSize int64) (string, int) {
//	putPolicy := storage.PutPolicy{
//		Scope: viper.GetString("qiniu.bucket"),
//	}
//	mac := qbox.NewMac(viper.GetString("qiniu.access-key"), viper.GetString("qiniu.secret-key"))
//	upToken := putPolicy.UploadToken(mac)
//	cfg := storage.Config{
//		Zone:          &storage.ZoneHuadong,
//		UseCdnDomains: false,
//		UseHTTPS:      false,
//	}
//
//	putExtra := storage.PutExtra{}
//
//	formUploader := storage.NewFormUploader(&cfg)
//	ret := storage.PutRet{}
//
//	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
//	if err != nil {
//		return "", errmsg.ERROR
//	}
//	url := viper.GetString("qiniu.img-path") + "/" + ret.Key
//	return url, errmsg.Succes
//
//}
