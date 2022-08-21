package categories

type Core struct {
	ID   int
	Name string
}

type Business interface {
	AddCategory(dataCategory Core) (response string, err error)
	GetCategory() (dataCategory []Core, err error)
}

type Data interface {
	Insert(dataCategory Core) (row int, err error)
	Get() (dataCategory []Core, err error)
}
