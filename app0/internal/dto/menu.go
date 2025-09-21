package dto

import (
	"vmapp/app0/internal/models"
)

// GetMenuReq
type GetMenuReq struct {
	ID int32 `form:"id" json:"id" validate:"required"`
}

// DeleteMenuReq
type DeleteMenuReq struct {
	ID int32 `form:"id" json:"id" validate:"required"`
}

type CreateMenuReq struct {
	Pid         int32  `json:"pid" form:"pid"`
	Title       string `json:"title" form:"title" validate:"required"`
	Type        int8   `json:"type" form:"type" validate:"required"`
	Icon        string `json:"icon" form:"icon"`
	Href        string `json:"href" form:"href"`
	Target      string `json:"target" form:"target"`
	ApiMethod   string `json:"api_method" form:"api_method"`
	ApiUrl      string `json:"api_url" form:"api_url"`
	ApiAuthCode string `json:"api_auth_code" form:"api_auth_code"`
	Sort        int8   `json:"sort" form:"sort"`
	Remark      string `json:"remark" form:"remark"`
}

type UpdateMenuReq struct {
	Id          int32  `json:"id" form:"id" validate:"required"`
	Pid         int32  `json:"pid" form:"pid"`
	Title       string `json:"title" form:"title" validate:"required"`
	Type        int8   `json:"type" form:"type"`
	Icon        string `json:"icon" form:"icon"`
	Href        string `json:"href" form:"href"`
	Target      string `json:"target" form:"target"`
	ApiMethod   string `json:"api_method" form:"api_method"`
	ApiUrl      string `json:"api_url" form:"api_url"`
	ApiAuthCode string `json:"api_auth_code" form:"api_auth_code"`
	Sort        int8   `json:"sort" form:"sort"`
	Remark      string `json:"remark" form:"remark"`
}

type ListMenuResp struct {
	models.Menu
}
