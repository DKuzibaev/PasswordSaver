package account

import (
	"encoding/json"
	"net/url"
	"passwordsaver/files"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ValultWithDb struct {
	Vault
	db files.JsonDb
}

// Конструктор структуры Valult
func NewVault(db *files.JsonDb) *ValultWithDb {
	file, err := db.Read()
	if err != nil {
		return &ValultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: *db,
		}
	}
	var valult Vault
	err = json.Unmarshal(file, &valult)
	if err != nil {
		color.Red("Не удалось разобрать файл data.json")
		return &ValultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db: *db,
		}
	}

	return &ValultWithDb{
		Vault: valult,
		db:    *db,
	}
}

// Сохранение и обновление аккаунтов
func (v *ValultWithDb) save() {

	v.UpdatedAt = time.Now()
	data, err := v.Vault.ToByteSlice()

	if err != nil {
		color.Red("Не удалось переобразовать")
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
