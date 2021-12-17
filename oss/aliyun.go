package oss

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"golang.org/x/xerrors"
)

type service struct {
	Client *oss.Client
	Config *ServiceConfig
}

func (s *service) GetUploadInfo(expireTime int64, uploadDir string) (out *UploadInfo, err error) {
	now := time.Now().Unix()
	expireSyncpoint := now + expireTime
	expire := time.Unix(expireSyncpoint, 0).UTC().Format(time.RFC3339)
	conditionArray := [1][3]string{{"starts-with", "$key", uploadDir}}
	policyDict := map[string]interface{}{
		"expiration": expire,
	}
	policyDict["conditions"] = conditionArray
	var policy []byte
	policy, err = json.Marshal(policyDict)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	policyEncode := base64.StdEncoding.EncodeToString(policy)
	var h []byte
	h, err = hamcData(s.Config.AccessKeySecret, policyEncode)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}

	signResult := base64.StdEncoding.EncodeToString(h)

	out = &UploadInfo{
		AccessID:  s.Config.AccessKeyID,
		Host:      s.Config.UploadHost,
		Policy:    policyEncode,
		Signature: signResult,
		Expire:    strconv.FormatInt(expireSyncpoint, 10),
		Dir:       uploadDir,
		FileName:  strconv.FormatInt(time.Now().UnixNano()/1000, 10),
	}
	return
}

func hamcData(key, data string) (hmacHash []byte, err error) {
	hmac := hmac.New(sha1.New, []byte(key))
	_, err = hmac.Write([]byte(data))
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	hmacHash = hmac.Sum([]byte(""))
	return
}

func (s *service) GetSignURL(bucketName, rawURL string, expireTime int64) (signURL string, err error) {
	urlParse, err := url.Parse(rawURL)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	path := strings.Trim(urlParse.Path, "/")
	bucket, err := s.Client.Bucket(bucketName)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	ossURL, err := bucket.SignURL(path, oss.HTTPGet, expireTime)
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return
	}
	urlSplit := strings.Split(ossURL, "?")[1]
	if strings.Contains(rawURL, "?") {
		signURL = fmt.Sprintf("%s&%s", rawURL, urlSplit)
	} else {
		signURL = fmt.Sprintf("%s?%s", rawURL, urlSplit)
	}
	return
}
