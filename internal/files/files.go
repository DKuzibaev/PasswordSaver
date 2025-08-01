package files

import (
	"os"
	"passwordsaver/internal/output"

	"github.com/fatih/color"
)

type JsonDb struct {
	filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)

	if err != nil {
		output.PrintError(err)
		return nil, err
	}

	return data, err
}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.filename)

	if err != nil {
		output.PrintError(err)
	}

	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		output.PrintError(err)
		return
	}

	color.Green("Запись успешна!")
}
