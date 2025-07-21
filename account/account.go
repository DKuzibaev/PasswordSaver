package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var LetterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-*!")

type Account struct {
	Login     string    `json:"login"` // Это структурыне теги для json
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Метод для генерации пароля
func (acc *Account) generatePassword(n int) {
	newGenPassword := make([]rune, n)
	for i := range newGenPassword {
		newGenPassword[i] = LetterRunes[rand.IntN(len(LetterRunes))]
	}
	acc.Password = string(newGenPassword)
}

// Метод для добавления аккаунта в Vault
func (acc *Account) Output() {
	color.Green("Ваш логин: " + acc.Login)
	color.Red("Ваш пароль: " + acc.Password)
	color.Yellow("Ваш URL: " + acc.Url)
}

// Функция конструктор структуры
func NewAccount(login, password, urlString string) (*Account, error) {

	if login == "" {
		return nil, errors.New("INVALID LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID URL")
	}

	newAcc := &Account{
		Login:     login,
		Password:  password,
		Url:       urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if password == "" {
		newAcc.generatePassword(12)
	}

	return newAcc, nil

}
