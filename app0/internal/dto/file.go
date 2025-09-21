package dto

// DeleteFileReq
type DeleteFileReq struct {
	ID int64 `json:"id" form:"id" validate:"required"`
}

type ListFileReq struct {
	Name  string `json:"name" form:"name"`
	Type  int    `json:"type" form:"type"`
	Page  int    `json:"page" form:"page"`
	Limit int    `json:"limit" form:"limit"`
}
