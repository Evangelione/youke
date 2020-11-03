package baidu

type BaiDu struct {
}

// 人脸搜索
type FaceSearchWrapper struct {
	ErrorCode int             `json:"error_code"`
	ErrorMsg  string          `json:"error_msg"`
	LogID     int             `json:"log_id"`
	Timestamp int             `json:"timestamp"`
	Cached    int             `json:"cached"`
	Result    FaceSearchModel `json:"result"`
}

type FaceSearchModel struct {
	FaceToken string           `json:"face_token"`
	UserList  []FaceSearchUser `json:"user_list"`
}

type FaceSearchUser struct {
	GroupID  string  `json:"group_id"`
	UserID   string  `json:"user_id"`
	UserInfo string  `json:"user_info"`
	Score    float64 `json:"score"`
}

// 人脸添加
type FaceAddWrapper struct {
	ErrorCode int          `json:"error_code"`
	ErrorMsg  string       `json:"error_msg"`
	LogID     int          `json:"log_id"`
	Timestamp int          `json:"timestamp"`
	Cached    int          `json:"cached"`
	Result    FaceAddModel `json:"result"`
}

type FaceAddModel struct {
	FaceToken string          `json:"face_token"`
	Location  FaceAddLocation `json:"location"`
}

type FaceAddLocation struct {
	Left     float64 `json:"left"`
	Top      float64 `json:"top"`
	Width    float64 `json:"width"`
	Height   float64 `json:"height"`
	Rotation int     `json:"rotation"`
}

// 人脸检测
type FaceDetectWrapper struct {
	ErrorCode int             `json:"error_code"`
	ErrorMsg  string          `json:"error_msg"`
	LogID     int             `json:"log_id"`
	Timestamp int             `json:"timestamp"`
	Cached    int             `json:"cached"`
	Result    FaceDetectModel `json:"result"`
}

type FaceDetectModel struct {
	FaceNum  int              `json:"face_num"`
	FaceList []FaceDetectItem `json:"face_list"`
}

type FaceDetectItem struct {
	FaceToken string            `json:"face_token"`
	Age       int               `json:"age"`
	Gender    FaceDetectGender  `json:"gender"`
	Quality   FaceDetectQuality `json:"quality"`
	ResultStr string            `json:"result_str"`
}

type FaceDetectGender struct {
	Type        string  `json:"type"`
	Probability float64 `json:"probability"`
}

type FaceDetectQuality struct {
	Blur float64 `json:"blur"`
}
