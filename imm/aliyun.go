package imm

import (
	"fmt"
	"net/url"
	"strings"

	aliImm "github.com/alibabacloud-go/imm-20170906/client"
	"golang.org/x/xerrors"
)

type service struct {
	Client *aliImm.Client
	Config *ServiceConfig
}

func (s *service) GetOfficePreviewInfo(bucketName, Project, fileURL string) (info *OfficePreviewInfo, err error) {
	// 转化为oss
	var urlParse *url.URL
	urlParse, err = url.Parse(fileURL)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	path := strings.Trim(urlParse.Path, "/")
	srcURI := fmt.Sprintf("oss://%s/%s", bucketName, path)
	req := &aliImm.GetOfficePreviewURLRequest{
		Project: &Project,
		SrcUri:  &srcURI,
	}

	result, err := s.Client.GetOfficePreviewURL(req)
	// index := strings.LastIndex(path, "/")
	// filename := "preview"
	// if index > 0 {
	// 	filename = path[index+1:]
	// }
	// user := `{"ID": "user1","Name": "test-user1"}`
	// Permission := `{"Readonly": true}`
	// notifyEndpoint := ""
	// bs := md5.Sum([]byte(srcURI))
	// fileID := hex.EncodeToString(bs[:])
	// file := fmt.Sprintf(`[{"Modifier": {"ID": "user1", "Name": "test-user1"},`+
	// 	` "Name": "%s", "Creator": {"ID": "user1", "Name": "test-user1"},`+
	// 	` "SrcUri": "%s", "Version": %v, "TgtUri": "%s"}]`, filename, srcURI, 1, srcURI)
	// req := &aliImm.GetWebofficeURLRequest{
	// 	File:           &file,
	// 	FileID:         &fileID,
	// 	NotifyEndpoint: &notifyEndpoint,
	// 	Permission:     &Permission,
	// 	Project:        &Project,
	// 	User:           &user,
	// }

	// result, err := s.Client.GetWebofficeURL(req)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	info = &OfficePreviewInfo{
		PreviewURL:              *result.Body.PreviewURL,
		RefreshToken:            *result.Body.RefreshToken,
		RequestID:               *result.Body.RequestId,
		AccessToken:             *result.Body.AccessToken,
		RefreshTokenExpiredTime: *result.Body.RefreshTokenExpiredTime,
		AccessTokenExpiredTime:  *result.Body.AccessTokenExpiredTime,
	}
	return
}
