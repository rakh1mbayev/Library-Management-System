package dal

type BookDal interface {
	GetAll()
	Save()
}
type BookData struct {
	filePath string
}

func NewBookRepo(filePath string) *BookData {
	return &BookData{filePath: filePath}
}

func (bookData *BookData) GetAll() {

}

func (bookData *BookData) Save() {

}
