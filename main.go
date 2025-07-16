package main

import (
	"fmt"
	"passwordsaver/account"
	"passwordsaver/files"
	"passwordsaver/output"

	"github.com/fatih/color"
)

func main() {
	color.Cyan("Добро пожаловать в менеджер паролей!")
	valult := account.NewVault(files.NewJsonDb("data.json"))
	// valult := account.NewVault(cloud.NewCloudDb("https://a.ru"))
Menu:
	for {
		variant := promtData([]string{
			"1. Создать аккаунт",
			"2. Найти аккаунт",
			"3. Удалить аккаунт",
			"4. Выход",
			"Выберите вариант",
		})
		switch variant {
		case "1":
			createAccount(valult)
		case "2":
			findAccout(valult)
		case "3":
			deleteAccout(valult)
		default:
			color.Green("Хорошего Вам дня!")
			break Menu
		}
	}
}

func createAccount(valult *account.ValultWithDb) {
	login := promtData([]string{"Введите логин"})
	password := promtData([]string{"Введите пароль"})
	url := promtData([]string{"Введите URL"})

	myAccount, err := account.NewAccount(login, password, url)

	if err != nil {
		output.PrintError("Неверный формат URL или Логин")
		return
	}
	valult.AddAccount(*myAccount)
}

// Функция вывода с использованием Generic Type
func promtData[T any](prompt []T) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			str := fmt.Sprint(line)
			color.Cyan(str)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}

func findAccout(vault *account.ValultWithDb) {
	url := promtData([]string{"Введите URL для поиска"})
	accounts, err := vault.FindAccountByURL(url)
	if err != nil {
		output.PrintError("Неверный формат URL")
		return
	}
	if len(accounts) == 0 {
		output.PrintError("Аккаунт не найден!")
	}
	for _, acc := range accounts {
		acc.Output()
	}
}

func deleteAccout(vault *account.ValultWithDb) {
	url := promtData([]string{"Введите URL для удаления"})
	isDeleted := vault.DeleteAccountByURL(url)
	if isDeleted {
		color.Green("Успешно удалено!")
	} else {
		output.PrintError("Аккаунт не найден!")
	}
}
