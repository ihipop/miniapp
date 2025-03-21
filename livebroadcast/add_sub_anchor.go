package livebroadcast

import "github.com/medivhzhan/weapp/v3/request"

const apiAddSubAnchor = "/wxaapi/broadcast/room/addsubanchor"

type AddSubAnchorRequest struct {
	// 必传	房间ID
	RoomId int64 `json:"roomId"`
	// 必传	用户微信号
	Username string `json:"username"`
}

// 添加主播副号
func (cli *LiveBroadcast) AddSubAnchor(req *AddSubAnchorRequest) (*request.CommonError, error) {

	api, err := cli.conbineURI(apiAddSubAnchor, nil)
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
