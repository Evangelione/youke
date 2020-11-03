package baidu

import (
	"errors"
	"yk/internal/pkg/constants"
	"yk/internal/pkg/infra"

	"github.com/imroc/req"
)

const (
	faceSearch          = "https://aip.baidubce.com/rest/2.0/face/v3/search"
	faceAdd             = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/add"
	faceDetect          = "https://aip.baidubce.com/rest/2.0/face/v3/detect"
	faceGroup           = "9youke"
	imageType           = "BASE64"
	matchUserIsNotFound = 222207
)

// 人脸搜索
func (BaiDu) FaceSearch(base64 string) (*FaceSearchModel, error) {
	r, err := req.Post(faceSearch+"?access_token="+infra.GetRedisValue(constants.FaceAccessTokenKey), req.Param{
		"image":         base64,
		"image_type":    imageType,
		"group_id_list": faceGroup,
	})

	// 访问api出错
	if err != nil {
		return nil, err
	}

	// 转换成结构体
	var res FaceSearchWrapper

	// 结构体转换出错
	if err = r.ToJSON(&res); err != nil {
		return nil, err
	}

	// 人脸库中无此人脸
	if res.ErrorCode == matchUserIsNotFound {
		return &res.Result, errors.New(res.ErrorMsg)
	}

	// 人脸库查询出错
	if res.ErrorCode != 0 {
		return nil, errors.New(res.ErrorMsg)
	}

	// 查询成功
	return &res.Result, nil
}

// 人脸注册
func (BaiDu) FaceAdd(base64, faceId string) (*FaceAddModel, error) {
	r, err := req.Post(faceAdd+"?access_token="+infra.GetRedisValue(constants.FaceAccessTokenKey), req.Param{
		"image":      base64,
		"image_type": imageType,
		"group_id":   faceGroup,
		"user_id":    faceId,
	})

	// 访问api出错
	if err != nil {
		return nil, err
	}

	// 转换成结构体
	var res FaceAddWrapper

	// 结构体转换出错
	if err = r.ToJSON(&res); err != nil {
		return nil, err
	}

	// 人脸库注册出错
	if res.ErrorCode != 0 {
		return nil, errors.New(res.ErrorMsg)
	}

	// 注册成功
	return &res.Result, nil
}

// 人脸检测
func (BaiDu) FaceDetect(base64 string) (*FaceDetectWrapper, error) {
	r, err := req.Post(faceDetect+"?access_token="+infra.GetRedisValue(constants.FaceAccessTokenKey), req.Param{
		"image":      base64,
		"image_type": imageType,
		"face_field": "age,beauty,expression,face_shape,gender,glasses,landmark,landmark150,race,quality,eye_status,emotion,face_type,mask,spoofing",
	})

	// 访问api出错
	if err != nil {
		return nil, err
	}

	// 转换成结构体
	var res FaceDetectWrapper

	// 结构体转换出错
	if err = r.ToJSON(&res); err != nil {
		return nil, err
	}

	// 人脸检测出错
	if res.ErrorCode != 0 {
		return nil, errors.New(res.ErrorMsg)
	}

	// 将整个查询结果转成bytes
	bytes, err := r.ToBytes()

	// 转成bytes出错
	if err != nil {
		return nil, err
	}

	// 将bytes转成字符串存入结果中，用于存储数据库
	res.Result.FaceList[0].ResultStr = string(bytes)

	// 检测成功
	return &res, nil
}
