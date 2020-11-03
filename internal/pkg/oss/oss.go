package oss

import (
	"fmt"
	"os"
	"yk/internal/app/enterprise/config"
	"yk/internal/pkg/constants"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type ObjectKey string

const (
	Merchant ObjectKey = "merchant"
	Store    ObjectKey = "store"
	GroupBuy ObjectKey = "groupBuy"
	Portrait ObjectKey = "portrait"
	Robot    ObjectKey = "robot"
	Activity ObjectKey = "activity"
	Logo     ObjectKey = "logo"
	Avatar   ObjectKey = "avatar"
)

type ossConf struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
}

var bucket *oss.Bucket

func GetOSSBucket() *oss.Bucket {
	if bucket != nil {
		return bucket
	}
	// Endpoint以杭州为例，其它Region请按实际情况填写。
	endpoint, err := config.FindOne(constants.OSSEndpoint)
	if err != nil {
		handleError(err)
		return nil
	}

	accessKeyId, err := config.FindOne(constants.OSSAccessKeyId)
	if err != nil {
		handleError(err)
		return nil
	}

	accessKeySecret, err := config.FindOne(constants.OSSAccessKeySecret)
	if err != nil {
		handleError(err)
		return nil
	}

	bucketName, err := config.FindOne(constants.OSSBucketName)
	if err != nil {
		handleError(err)
		return nil
	}

	conf := ossConf{
		Endpoint:        endpoint.Value,
		AccessKeyId:     accessKeyId.Value,
		AccessKeySecret: accessKeySecret.Value,
		BucketName:      bucketName.Value,
	}
	// 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。

	// 创建OSSClient实例。
	client, err := oss.New(conf.Endpoint, conf.AccessKeyId, conf.AccessKeySecret)
	if err != nil {
		handleError(err)
	}

	bucket, err := client.Bucket(bucketName.Value)
	if err != nil {
		handleError(err)
	}

	return bucket
}

func UploadLocalFileToOSS(key, filePath string) error {
	ossBucket := GetOSSBucket()
	err := ossBucket.PutObjectFromFile(key, filePath)

	if err != nil {
		return err
	}

	return nil
}

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}
