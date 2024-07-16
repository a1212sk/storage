package storage

import (
	"errors"
	"fmt"
	"github.com/a1212sk/storage/internal/file"
	"github.com/google/uuid"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}
func (s *Storage) Save(filename string, data []byte) (*file.File, error) {
	file, err := file.NewFile(filename, data)
	if err != nil {
		fmt.Println("Error while saving file", err)
		return nil, err
	}
	fmt.Println("Saving file", filename)
	s.files[file.Id] = file
	fmt.Println("Saved", string(data))

	return file, nil
}

func (s *Storage) GetById(id uuid.UUID) (*file.File, error) {
	file, ok := s.files[id]
	if !ok {
		fmt.Println("File not found", id)
		return nil, errors.New("File not found")
	}
	return file, nil
}
