package domain

type UserListRequest struct {
	TimeAndPageSearch
}

type SigIn struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}
