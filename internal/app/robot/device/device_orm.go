package device

import "gorm.io/gorm"

// 会话设备
type AiDevice struct {
	DB               *gorm.DB `json:"-" gorm:"-"`
	Id               string   `json:"id"`
	MerId            string   `json:"mer_id" `
	StoreId          string   `json:"store_id" `
	Status           int      `json:"status"`
	ImaxId           string   `json:"imax_id"`
	CameraCode       string   `json:"camera_code"`
	VoiceCode        string   `json:"voice_code"`
	SynTime          string   `json:"syn_time"`
	Introducers      string   `json:"introducers"`
	AiMenu           string   `json:"ai_menu"`
	SameCityQuestion string   `json:"same_city_question"`
	DeviceQuestion   int      `json:"device_question"`
	SkillId          string   `json:"skill_id"`
	ImaxIdId         string   `json:"imax_id_id"`
}
