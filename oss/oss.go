package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"golang.org/x/xerrors"
)

// UploadInfo 上传文件参数
type UploadInfo struct {
	AccessID  string `json:"access_id"`
	Host      string `json:"host"`
	Policy    string `json:"policy"`
	Signature string `json:"signature"`
	Expire    string `json:"expire"`
	Dir       string `json:"dir"`
	FileName  string `json:"file_name"`
}

// Service Oss自定义服务
type Service interface {
	// GetUploadInfo 获取上传文件夹参数
	GetUploadInfo(expireTime int64, uploadDir string) (out *UploadInfo, err error)
	// GetSignURL 获取签名URL
	GetSignURL(bucket, rawURL string, expireTime int64) (signURL string, err error)
}

// ServiceConfig Oss 配置选项
type ServiceConfig struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	ClientOption    []oss.ClientOption
	UploadHost      string
}

// NewService 初始化自定义oss服务
func NewService(cfg ServiceConfig) (s Service, err error) {

	client, err := oss.New(cfg.Endpoint, cfg.AccessKeyID, cfg.AccessKeySecret, cfg.ClientOption...)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	s = &service{
		Config: &cfg,
		Client: client,
	}
	return
}
