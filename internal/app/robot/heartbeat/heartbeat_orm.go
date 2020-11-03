package heartbeat

import "gorm.io/gorm"

type Heartbeat struct {
	DB        *gorm.DB `json:"-" gorm:"-"`
	Id        string   `json:"id"`
	MerId     string   `json:"mer_id"`
	DeviceId  string   `json:"device_id"`
	ImaxId    string   `json:"imax_id"`
	StoreId   string   `json:"store_id"`
	SiteId    string   `json:"site_id"`
	TaskId    string   `json:"task_id"`
	Status    string   `json:"status"`
	CreatedAt int64    `json:"created_at"`
}

func (Heartbeat) TableName() string {
	return "maycms_device_heartbeat"
}
