package ding

// Service 自定义服务
type Service interface {
	// TODO
}

// ServiceConfig 配置选项
type ServiceConfig struct {
	AccessKeyID     string
	AccessKeySecret string
	UserAgent       string
}

// NewService 初始化自定义oss服务
func NewService(cfg ServiceConfig) (s Service, err error) {

	s = &service{
		Config: &cfg,
	}
	return
}
