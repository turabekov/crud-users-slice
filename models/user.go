package models

type User struct {
	Id       int
	Birthday string
	Name     string
	Surname  string
}

type GetListRequest struct {
	Offset   int
	Limit    int
	Search   string
	FromDate string
	ToDate   string
}
