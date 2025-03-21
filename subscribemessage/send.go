package subscribemessage

import (
	"github.com/medivhzhan/weapp/v3/request"
)

const apiSend = "/cgi-bin/message/subscribe/send"

// SubscribeMessage 订阅消息
type SendRequest struct {
	ToUser           string           `json:"touser"`
	TemplateID       string           `json:"template_id"`
	Page             string           `json:"page,omitempty"`
	MiniprogramState MiniprogramState `json:"miniprogram_state,omitempty"`
	Data             string           `json:"data"`
}

// MiniprogramState 跳转小程序类型
type MiniprogramState = string

// developer为开发版；trial为体验版；formal为正式版；默认为正式版
const (
	MiniprogramStateDeveloper = "developer"
	MiniprogramStateTrial     = "trial"
	MiniprogramStateFormal    = "formal"
)

// 发送订阅消息
func (cli *SubscribeMessage) Send(msg *SendRequest) (*request.CommonError, error) {
	api, err := cli.conbineURI(apiSend, nil)
	if err != nil {
		return nil, err
	}

	res := new(request.CommonError)
	if err := cli.request.Post(api, msg, res); err != nil {
		return nil, err
	}

	return res, nil
}
