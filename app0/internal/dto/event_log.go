package dto

type ListEventLogReq struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

type ListSendLogReq struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}
