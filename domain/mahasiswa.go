package domain

type Mahasiswa struct {
	ID      int `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name    string
	NPM     string
	Address string
	Major   string
}

type MahasiswaUseCase interface {
	GetAll() ([]Mahasiswa, error)
	GetByID(id int) (Mahasiswa, error)
	Create(mahasiswa Mahasiswa) error
	Update(id int, mahasiswa Mahasiswa) error
	Delete(id int) error
}
