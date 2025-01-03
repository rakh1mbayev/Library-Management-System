package services

import (
	"Library-Management-System/internal/dal"
	"Library-Management-System/internal/models"
)

type BookService interface {
	Create(book *models.Book) error
	Get() (*models.Book, error)
	GetById(id string) (*models.Book, error)
	Update(id string, item *models.Book) error
	Delete(id string) error
}
type FileBookService struct {
	dataAccess dal.BookDal
}

func NewFileBookService(dataAccess dal.BookDal) *FileBookService {
	return &FileBookService{dataAccess: dataAccess}
}

func (f *FileBookService) Create(book *models.Book) error {
	
	return nil
}

func (f *FileBookService) GetById(id string) (*models.Book, error) {
	book := models.Book{}
	return &book, nil
}
func (f *FileBookService) Get() (*[]models.Book, error) {
	book := []models.Book{}
	return &book, nil
}

func (f *FileBookService) Update(id string, item *models.Book) error {
	return nil
}

func (f *FileBookService) Delete(id string) error {
	return nil
}
