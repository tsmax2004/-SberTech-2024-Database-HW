package model

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type Task struct {
	Id           uint64
	Title        string `objectbox:"index:hash64"`
	Text         string
	DateCreated  int64
	DateFinished int64
}
