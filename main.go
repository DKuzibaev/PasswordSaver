package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

var LetterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-*!")

type account struct {
	login    string
	password string
	url      string
}

// Расширенная структура аккаунта с датами (наследует account)
type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	account
}

// МЕТОД ДЛЯ ГЕНЕРАЦИИ ПАРОЛЯ
func (acc *account) generatePassword(n int) {
	newGenPassword := make([]rune, n)
	for i := range newGenPassword {
		newGenPassword[i] = LetterRunes[rand.IntN(len(LetterRunes))]
	}
	acc.password = string(newGenPassword)
}

// МЕТОД ДЛЯ ВЫВОДА ДАННЫХ
func (acc *account) outputPassword() {
	// Просто печатаем содержимое структуры
	fmt.Println("🔑 Логин:", acc.login)
	fmt.Println("🔒 Пароль:", acc.password)
	fmt.Println("🌐 Сайт:", acc.url)
}

// КОНСТРУКТОР ДЛЯ ОСНОВНОЙ СТРУКТУРЫ (account)
func newAccount(login, password, urlString string) (*account, error) {
	// Проверяем логин
	if login == "" {
		return nil, errors.New("ЛОГИН НЕ МОЖЕТ БЫТЬ ПУСТЫМ")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("НЕПРАВИЛЬНЫЙ URL")
	}

	newAcc := &account{
		login:    login,
		password: password,
		url:      urlString,
	}

	if password == "" {
		newAcc.generatePassword(12)
	}

	return newAcc, nil
}

// КОНСТРУКТОР ДЛЯ РАСШИРЕННОЙ СТРУКТУРЫ (accountWithTimeStamp)
func newAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) {
	// Те же проверки что и выше
	if login == "" {
		return nil, errors.New("ЛОГИН НЕ МОЖЕТ БЫТЬ ПУСТЫМ")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("НЕПРАВИЛЬНЫЙ URL")
	}

	newAcc := &accountWithTimeStamp{
		account: account{
			login:    login,
			password: password,
			url:      urlString,
		},
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	if password == "" {
		newAcc.generatePassword(12)
	}

	return newAcc, nil
}

// ГЛАВНАЯ ФУНКЦИЯ (точка входа в программу)
func main() {
	login := promtData("Введите логин:")
	password := promtData("Введите пароль (оставьте пустым для автогенерации):")
	url := promtData("Введите URL сайта:")

	myAccount, err := newAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("💥 ОШИБКА:", err)
		return
	}

	myAccount.outputPassword()
	fmt.Println("⏰ Дата создания:", myAccount.createdAt.Format("2006-01-02 15:04:05"))
}

// ВСПОМОГАТЕЛЬНАЯ ФУНКЦИЯ ДЛЯ ВВОДА ДАННЫХ
func promtData(prompt string) string {
	fmt.Print(prompt + " ")
	var res string
	fmt.Scanln(&res)
	return res
}
