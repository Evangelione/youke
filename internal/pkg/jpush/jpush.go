package jpush

import (
	"encoding/json"
	"errors"
	"yk/internal/pkg/constants"

	jPushClient "github.com/ylywyn/jpush-api-go-client"
)

func JPush(platform string, jad jPushClient.Audience, notice jPushClient.Notice, msg jPushClient.Message) error {
	// 1.构建 jPushClient.PayLoad
	payload := jPushClient.NewPushPayLoad()

	// 2.构建要推送的平台 jPushClient.Platform
	var pf jPushClient.Platform
	if err := pf.Add(platform); err != nil {
		return err
	}
	// 全部广播 ad.All()
	payload.SetPlatform(&pf)

	// 3.构建接收听众 jPushClient.Audience -> jad
	payload.SetAudience(&jad)

	// 4.构建通知 jPushClient.Notice -> notice，或者消息： jPushClient.Message -> message
	payload.SetNotice(&notice)
	payload.SetMessage(&msg)

	bytes, _ := payload.ToBytes()

	// 5.构建 PushClient，发出推送
	c := jPushClient.NewPushClient(constants.JPushSecret, constants.JPushAppKey)
	backStr, err := c.Send(bytes)
	if err != nil {
		return err
	}

	var backFail RespFail
	err = json.Unmarshal([]byte(backStr), &backFail)
	if err != nil {
		return err
	}

	if backFail.Error.Message != "" {
		return errors.New(backFail.Error.Message)
	}

	return nil
}
