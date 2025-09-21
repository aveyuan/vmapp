package dto

import (
	"vmapp/app0/internal/models"
)

// GetRoleReq
type GetRoleReq struct {
	ID int32 `form:"id" json:"id" validate:"required"`
}

// DeleteRoleReq
type DeleteRoleReq struct {
	ID int32 `form:"id" json:"id" validate:"required"`
}

type CreateRoleReq struct {
	RoleName  string `json:"role_name" form:"role_name" validate:"required"`
	RoleCode  string `json:"role_code" form:"role_code" validate:"required"`
	MenuIds   string `json:"menu_ids" form:"menu_ids"`
	IndexMenu string `json:"index_menu" form:"index_menu"`
	Remark    string `json:"remark" form:"remark"`
}

type UpdateRoleReq struct {
	Id        int32  `json:"id" form:"id" validate:"required"`
	RoleName  string `json:"role_name" form:"role_name" validate:"required"`
	RoleCode  string `json:"role_code" form:"role_code" validate:"required"`
	MenuIds   string `json:"menu_ids" form:"menu_ids"`
	IndexMenu string `json:"index_menu" form:"index_menu"`
	Remark    string `json:"remark" form:"remark"`
}

type ListRoleResp struct {
	models.Role
}
