<<<<<<< HEAD
package files

import (
	"fmt"
	"os"
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
		fmt.Println(err)
		return nil, err
	}

	return data, err
}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.filename)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Запись успешна!")
}
=======
package files

import (
	"fmt"
	"os"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return data, err
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Запись успешна!")
}
>>>>>>> f6768fafe261a74e9fb9c5c9f8529acbea46d48e
