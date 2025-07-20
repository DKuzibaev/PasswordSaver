package main

import (
	"fmt"
	"passwordsaver/account"
	"passwordsaver/files"
	"passwordsaver/output"
	"strings"

	"github.com/fatih/color"
)

var menu = map[string]func(*account.ValultWithDb){
	"1": createAccount,
	"2": findAccountByULR,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

func main() {
	color.Cyan("Добро пожаловать в менеджер паролей!")
	valult := account.NewVault(files.NewJsonDb("data.json"))
	// valult := account.NewVault(cloud.NewCloudDb("https://a.ru"))
Menu:
	for {
		variant := promtData([]string{
			"1. Создать аккаунт",
			"2. Найти аккаунт URL",
			"3. Найти аккаунт по Логину",
			"4. Удалить аккаунт",
			"5. Выход",
			"Выберите вариант",
		})
		menuFunc := menu[variant]
		if menuFunc == nil {
			break Menu
		}

		menuFunc(valult)
		// switch variant {
		// case "1":
		// 	createAccount(valult)
		// case "2":
		// 	findAccount(valult)
		// case "3":
		// 	deleteAccount(valult)
		// default:
		// 	color.Green("Хорошего Вам дня!")
		// 	break Menu
		// }
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

func findAccountByULR(vault *account.ValultWithDb) {
	url := promtData([]string{"Введите URL для поиска"})
	accounts, err := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	NotFound(accounts, err)
}

func findAccountByLogin(vault *account.ValultWithDb) {
	login := promtData([]string{"Введите Логин для поиска"})
	accounts, err := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	NotFound(accounts, err)
}


func NotFound(accs []account.Account, err error) { // Изменили тип на срез аккаунтов
    if err != nil {
        output.PrintError("Неверный формат URL или Логин")
        return
    }
    if len(accs) == 0 {
        output.PrintError("Аккаунт не найден!")
        return // Добавили return, чтобы не выводить пустой список
    }
    for _, acc := range accs {
        acc.Output()
    }
}

func deleteAccount(vault *account.ValultWithDb) {
	url := promtData([]string{"Введите URL для удаления"})
	isDeleted := vault.DeleteAccountByURL(url)
	if isDeleted {
		color.Green("Успешно удалено!")
	} else {
		output.PrintError("Аккаунт не найден!")
	}
}
