package account

import (
	"encoding/json"
	"passwordsaver/output"
	"strings"
	"time"
)

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

// Поиск аккаунта по URL
func (v *ValultWithDb) FindAccounts(str string, cheker func(Account, string) bool) ([]Account, error) {
	var accounts []Account
	for _, acc := range v.Accounts {
		isMatch := cheker(acc, str)
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
