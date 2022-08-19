package categories

type Core struct {
	ID   int
	Name string
}

type Business interface {
	AddCtgy(dataCtgy Core) (response string, err error)
	GetCtgy() (dataCtgy Core, err error)
}

type Data interface {
	Insert(dataCtgy Core) (row int, err error)
	Get() (dataCtgy []Core, err error)
}
