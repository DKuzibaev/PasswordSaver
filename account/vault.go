package account

import (
	"encoding/json"
	"fmt"
	"passwordsaver/files"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Конструктор структуры Valult
func NewVault() *Vault {
	file, err := files.ReadFile("data.json")
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	var valult Vault
	err = json.Unmarshal(file, &valult)
	if err != nil {
		color.Red("Не удалось разобрать файл data.json")
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}

	return &valult
}

// Функция добавления аккаунта
func (valult *Vault) AddAccount(acc Account) {
	valult.Accounts = append(valult.Accounts, acc)
	valult.UpdatedAt = time.Now()
	data, err := valult.ToByteSlice()

	if err != nil {
		color.Red("Не удалось переобразовать")
	}

	files.WriteFile(data, "data.json")

}

func FindAccountByURL(vault *Vault, url string) ([]string, error) {
	var results []string

	for _, item := range vault.Accounts {
		if strings.Contains(item.Url, url) {
			// Формируем строку результата, например, "login:password"
			result := fmt.Sprintf("%s:%s", item.Login, item.Password)
			results = append(results, result)
		}
	}

	if len(results) == 0 {
		return results, fmt.Errorf("no accounts found for URL containing: %s", url)
	}

	return results, nil
}

// Методо преобразования структуры в массив Byte
func (valult *Vault) ToByteSlice() ([]byte, error) {
	file, err := json.Marshal(valult) // методо преобразование в массив json

	if err != nil {
		return nil, err
	}

	return file, nil
}
