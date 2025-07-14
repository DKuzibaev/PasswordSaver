<<<<<<< HEAD
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

// так смотри теперь это стало методом нашей структуры, теперь мне не обязательно что либо возвращать так как я теперь мутирую нашу структуру
func (acc *Account) generatePassword(n int) {
	newGenPassword := make([]rune, n)
	for i := range newGenPassword {
		newGenPassword[i] = LetterRunes[rand.IntN(len(LetterRunes))]
	}
	acc.Password = string(newGenPassword)
}

/*
передал через указатель что бы явно привязать его к структуре
если передать без указателя у нас каждый раз будет cоздаваться копия myAccount и просто выводить но не запишется в самой переменной для это юзай указатели
т.е мы таким образом будет эконопить память, при маленьких структур нам не обязательно использовать указатель
*/
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
=======
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

// так смотри теперь это стало методом нашей структуры, теперь мне не обязательно что либо возвращать так как я теперь мутирую нашу структуру
func (acc *Account) generatePassword(n int) {
	newGenPassword := make([]rune, n)
	for i := range newGenPassword {
		newGenPassword[i] = LetterRunes[rand.IntN(len(LetterRunes))]
	}
	acc.Password = string(newGenPassword)
}

/*
передал через указатель что бы явно привязать его к структуре
если передать без указателя у нас каждый раз будет cоздаваться копия myAccount и просто выводить но не запишется в самой переменной для это юзай указатели
т.е мы таким образом будет эконопить память, при маленьких структур нам не обязательно использовать указатель
*/
func (acc *Account) OutputPassword() {
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
>>>>>>> f6768fafe261a74e9fb9c5c9f8529acbea46d48e
