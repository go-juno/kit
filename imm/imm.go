package imm

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	aliImm "github.com/alibabacloud-go/imm-20170906/client"
	"golang.org/x/xerrors"
)

// OfficePreviewInfo Office预览参数
type OfficePreviewInfo struct {
	PreviewURL              string
	RefreshToken            string
	RequestID               string
	AccessToken             string
	RefreshTokenExpiredTime string
	AccessTokenExpiredTime  string
}

// Service 自定义服务
type Service interface {
	GetOfficePreviewInfo(bucketName, Project, fileURL string) (info *OfficePreviewInfo, err error)
}

// ServiceConfig 配置选项
type ServiceConfig struct {
	AccessKeyID     string
	AccessKeySecret string
}

// NewService 初始化自定义oss服务
func NewService(cfg ServiceConfig) (s Service, err error) {
	config := &openapi.Config{
		AccessKeyId:     &cfg.AccessKeyID,
		AccessKeySecret: &cfg.AccessKeySecret,
	}
	client, err := aliImm.NewClient(config)
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
