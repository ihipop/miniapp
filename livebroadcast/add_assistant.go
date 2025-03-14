package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiAddAssistant = "/wxaapi/broadcast/room/addassistant"

type AddAssistantRequest struct {
	// 必传	房间ID
	RoomId int64 `json:"roomId"`
	// 必传	用户数组
	Users []*Assistant `json:"users"`
}

type Assistant struct {
	// 必传	用户微信号
	Username string `json:"username"`
	// 必传	用户昵称
	Nickname string `json:"nickname"`
}

// 添加管理直播间小助手
func (cli *LiveBroadcast) AddAssistant(req *AddAssistantRequest) (*request.CommonError, error) {

	api, err := cli.conbineURI(apiAddAssistant, nil)
	if err != nil {
		return nil, err
	}

	res := new(request.CommonError)
	err = cli.request.Post(api, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
