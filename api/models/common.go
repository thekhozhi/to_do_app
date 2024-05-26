package models

type PrimaryKey struct {
	ID string `json:"id"`
}

type GetListRequest struct {
	Page      int`json:"page"`
	Limit 	  int`json:"limit"`
	Search string`json:"search"`
}

type UpdateResponse struct {
	ID		  string`json:"id"`
	UpdatedAt string`json:"updated_at"`
}