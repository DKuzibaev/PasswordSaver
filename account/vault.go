package account

import (
	"encoding/json"
	"net/url"
	"passwordsaver/output"
	"strings"
	"time"
)

// Это мой интерфейс (либо провайдер) по сути каждый интерфейс который с ним будет связан должен в обязательном порядке его создать
// не нужно прям расписывать как эти методы должны работать, просто укажи что метод интерфейса должен вернуть или получить как в примере ниже!
// Важный момент интерфейс не надо явно связывать к пакетом или структурой!

// Интерфейсы можено разбивать так же как и структуры и докидывать в другой интерфейс
type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type Db interface {
	ByteReader
	ByteWriter
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ValultWithDb struct {
	Vault
	db Db
}

// Конструктор структуры Valult
func NewVault(db Db) *ValultWithDb {
	file, err := db.Read()
	if err != nil {
		return &ValultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}
	var valult Vault
	err = json.Unmarshal(file, &valult)
	if err != nil {
		output.PrintError("Не удалось разобрать файл data.json")
		return &ValultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}

	return &ValultWithDb{
		Vault: valult,
		db:    db,
	}
}

// Сохранение и обновление аккаунтов
func (v *ValultWithDb) save() {

	v.UpdatedAt = time.Now()
	data, err := v.Vault.ToByteSlice()

	if err != nil {
		output.PrintError(err)
	}

	v.db.Write(data)
}

// Функция добавления аккаунта
func (v *ValultWithDb) AddAccount(acc Account) {
	v.Accounts = append(v.Accounts, acc)
	v.save()
}

// Методо преобразования структуры в массив Byte
func (v *Vault) ToByteSlice() ([]byte, error) {
	file, err := json.Marshal(v) // методо преобразование в массив json
	if err != nil {
		return nil, err
	}
	return file, nil
}

// Поиск аккаунта
func (v *ValultWithDb) FindAccountByURL(urlString string) ([]Account, error) {
	var accounts []Account
	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, err
	}
	for _, acc := range v.Accounts {
		isMatch := strings.Contains(acc.Url, urlString)
		if isMatch {
			accounts = append(accounts, acc)
		}
	}
	return accounts, nil
}

// Поиск удаление аккаунта
func (v *ValultWithDb) DeleteAccountByURL(urlString string) bool {
	var accounts []Account
	isDeleted := false
	for _, acc := range v.Accounts {
		isMatch := strings.Contains(acc.Url, urlString)
		if !isMatch {
			accounts = append(accounts, acc)
			continue
		}
		isDeleted = true
	}
	v.Accounts = accounts
	v.save()
	return isDeleted
}

/*
Удаление аккаунта по индексу, но лучшее так не делать потому что range может наебнуться
func (v *Vault) DeleteAccountByURL(urlString string) bool {
	var accounts []Account
	isDeleted := false
	for i, acc := range v.Accounts {
		isMatch := strings.Contains(acc.Url, urlString)
		if !isMatch {
			v.Accounts = append(v.Accounts[:i], v.Accounts[i+1:]...)
			accounts = append(accounts, acc)
			continue
		}
		isDeleted = true
	}
	v.Accounts = accounts
	v.UpdatedAt = time.Now()
	data, err := v.ToByteSlice()

	if err != nil {
		color.Red("Не удалось переобразовать")
	}

	files.WriteFile(data, "data.json")

	return isDeleted
}
*/
