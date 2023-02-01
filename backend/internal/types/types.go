// Code generated by goctl. DO NOT EDIT.
package types

type UserInfo struct {
	Id        int64  `json:"id,optional"`
	Account   string `json:"account"`
	Signature string `json:"signature"`
	Email     string `json:"email,optional"`
	Status    string `json:"status,optional"`
}

type TaskInfo struct {
	Id      int64  `json:"id,optional"`
	Account string `json:"account"`
	Chain   int32  `json:"chain"`
	Address string `json:"address"`
	Type    string `json:"type"`
	Status  string `json:"status,optional"`
}

type LoginRequest struct {
	UserInfo
}

type ChangeEmailRequest struct {
	UserInfo
}

type CreateTaskRequest struct {
	TaskInfo
}

type TaskFilterInfo struct {
	Id      int64  `form:"id,optional"`
	Account string `form:"account,optional"`
	Chain   int32  `form:"chain,optional"`
	Address string `form:"address,optional"`
	Type    string `form:"type,optional"`
	Status  string `form:"status,optional"`
}

type ReadTaskRequest struct {
	TaskFilterInfo
}

type DeleteTaskRequest struct {
	IdArray []int64 `json:"id_array"`
	Account string  `json:"account"`
}

type CommonResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ReadTaskResponse struct {
	CommonResponse
	Data []TaskInfo `json:"data,optional"`
}
