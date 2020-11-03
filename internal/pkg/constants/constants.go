package constants

const (
	// md5 salt
	MD5Salt = "youke"

	// 百度获取token地址
	AccessTokenAPI = "https://aip.baidubce.com/oauth/2.0/token"

	// 百度对话公私钥
	DiaLogClientID       = "WGV1k7pus62a6FwvOZCPgsoH"
	DialogClientSecret   = "rPZMCpG2LO1sNyCAvDz1bC22g66fjOsW"
	DialogAccessTokenKey = "access_token_" + DiaLogClientID

	// 百度人脸公私钥
	FaceClientID       = "ugjj8iN2cKEG8Q2Vh89TYVAL"
	FaceClientSecret   = "7s18hWNmzeFiGZSOlRGryHGqAebZLvwg"
	FaceAccessTokenKey = "access_token_" + FaceClientID

	GrantType = "client_credentials"

	// 极光推送公私钥
	JPushAppKey = "80318067bebf7a726dabef82"
	JPushSecret = "d1a1a1f779b8772ef87220e2"

	// DMkit
	DMKitAPI = "http://127.0.0.1:8010/search"

	//平台域名
	PLANTFORM_DOMAIN_NAME = "https://www.9youke.com/"

	// 极光推送员工tag前缀
	JPushPrefixStore      = "store_id"
	JPushPrefixPost       = "type"
	JPushPrefixUid        = "uid"
	JPushPrefixPermission = "permission"

	// 状态码
	SUCCESS        = 0
	PARAMETERERROR = 1

	// config表键名（oss）
	OSSEndpoint        = "aliyunoss_endpoint"
	OSSAccessKeyId     = "aliyunoss_access_key_id"
	OSSAccessKeySecret = "aliyunoss_access_key_secret"
	OSSBucketName      = "aliyunoss_bucket"
)
