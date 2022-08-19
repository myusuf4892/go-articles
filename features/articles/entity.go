package articles

type Core struct {
	ID         int
	Title      string
	CategoryID int
	Category   Category
}

type Category struct {
	ID   int
	Name string
}

type Business interface {
	AddPost(dataPost Core) (response string, err error)
	GetPost() (dataPost []Core, err error)
}

type Data interface {
	Insert(dataPost Core) (row int, err error)
	Get() (dataPost []Core, err error)
}
