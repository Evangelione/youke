package handler

import (
	"encoding/json"
	"fmt"
	"sync"
	"yk/internal/app/robot"
	"yk/internal/app/robot/websocket"
	"yk/internal/pkg/infra"
	"yk/internal/pkg/ws"

	"github.com/gin-gonic/gin"
)

type WebSocketHandler struct {
	infra.BaseHandler
}

var WL = make(websocket.WsList)

func (w WebSocketHandler) Index(c *gin.Context) {}

// swagger:route POST /websockets/connections 长连接 MerchantCreateParams
//
// 创建商户.
//
// Responses:
// 200: ServerSuccess
// 400: ClientError
// 401: ClientError
// 500: ServerError
func (w WebSocketHandler) Create(c *gin.Context) {
	// 1.参数校验
	var params websocket.CreateParams
	err := w.BindParams(c, &params.Body)
	if err != nil {
		w.InvalidParameter(c, err.Error())
		return
	}

	// 升级成websocket协议
	upgrade, err := ws.Socket.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		w.HandleError(c, err)
		return
	}

	// 创建长连接
	WL.Add(params.Body.SId, params.Body.UId, websocket.Ws{
		Conn: upgrade,
		Lock: new(sync.Mutex),
	})

	defer func() {
		if err := upgrade.Close(); err != nil {
			return
		}
	}()

	for {
		//读取ws中的数据
		_, message, err := upgrade.ReadMessage()
		if err != nil {
			robot.Logger().Error(err.Error())
			break
		}

		var status websocket.RobotStatus
		err = json.Unmarshal(message, &status)

		if err == nil {
			//保存状态
			fmt.Println(status)
			//wsService.SaveStatus(params.Body.SId, status.Status, status.SiteId, status.TaskId)
		}

		//if string(message) == "OrderCommit" {
		//	WsSend(sId, uId)
		//}

		//写入ws数据
		err = upgrade.WriteJSON(gin.H{
			"err": 0,
			"msg": string(message),
		})
		if err != nil {
			break
		}
	}
}
