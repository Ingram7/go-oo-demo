package controller

import "fmt"

type Data interface{}

type Response struct {
	Code    int    `json:"code"`    // 业务响应状态码
	Message string `json:"message"` // 提示信息
	Data    Data   `json:"data"`    // 数据
}

const (
	StatusOK = 0

	StatusBadRequest      = 400
	StatusUnauthenticated = 401
	StatusNotFound        = 404
	StatusRepeatRequest   = 405

	StatusError   = 500 // 错误
	StatusIllegal = 501 // 非法操作

	NotHaveParent = 1000
	TaskAccepted  = 1001
	TaskFinished  = 1002
	TaskOvertime  = 1003
)

var Message = map[int]string{
	StatusOK: "OK",

	StatusBadRequest:      "请求错误",
	StatusUnauthenticated: "请先登陆",
	StatusNotFound:        "未找到该记录",
	StatusRepeatRequest:   "请勿重新请求",

	StatusError:   "服务器出现未知错误，请稍后重试",
	StatusIllegal: "非法操作",

	NotHaveParent: "必须绑定上级才能继续操作",
	TaskAccepted:  "任务已被接单",
	TaskFinished:  "任务已完成",
	TaskOvertime:  "订单已超时，系统自动取消",
}

type Error struct {
	Code    int
	Message string
}

func (e Error) Error() string {
	m := Message[e.Code]

	if len(e.Message) > 0 {
		m = e.Message
	}
	return fmt.Sprintf("%d|%s", e.Code, m)
}
