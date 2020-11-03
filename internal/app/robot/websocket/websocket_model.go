package websocket

import (
	"sync"

	"github.com/gorilla/websocket"
)

// swagger:parameters WebsocketCreateParams
type CreateParams struct {
	// in: body
	Body struct {
		SId string `json:"s_id" validate:"required,number"`
		UId string `json:"u_id" validate:"required,number"`
	}
}

// 长连接实例
type Ws struct {
	Conn *websocket.Conn
	Lock *sync.Mutex
}

type WsList map[string]map[string]*Ws

func (wl *WsList) Add(sid, uid string, ws Ws) {
	if (*wl)[sid] == nil {
		(*wl)[sid] = make(map[string]*Ws)
	}
	(*wl)[sid][uid] = &ws
}

func (wl *WsList) Remove(sid, uid string) {
	delete((*wl)[sid], uid)
}

func (wl *WsList) RemoveSid(sid string) {
	delete(*wl, sid)
}

func (wl *WsList) GetWsList(sid string) map[string]*Ws {
	return (*wl)[sid]
}

func (wl *WsList) GetWs(sid, uid string) *Ws {
	return (*wl)[sid][uid]
}

type RobotStatus struct {
	Status string `json:"status"`  //机器人状态
	SiteId string `json:"site_id"` //当前位置id
	TaskId string `json:"task_id"` //任务id
}

type SendJSON struct {
	Action string      `json:"action"` //动作
	Data   interface{} `json:"data"`   //附带信息
}
