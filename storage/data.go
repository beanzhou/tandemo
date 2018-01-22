package storage

type User struct {
	Id   int64
	Name string
	Type string
}

type Relationship struct {
	Id      int64
	UserId  int64
	OtherId int64
	State   string
	Type    string
}
