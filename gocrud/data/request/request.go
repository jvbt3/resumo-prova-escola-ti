package request

type CreateUsersRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UpdateUsersRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
