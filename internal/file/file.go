package file

import "github.com/google/uuid"

type File struct {
	Name string
	Data []byte
	Id   uuid.UUID
}

func NewFile(filename string, data []byte) (*File, error) {
	return &File{
		Name: filename,
		Data: data,
		Id:   uuid.New(),
	}, nil
}
