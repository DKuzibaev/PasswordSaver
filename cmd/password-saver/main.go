package main

import (
	"fmt"
	"os"
	"passwordsaver/internal/account"
	"passwordsaver/internal/encrypter"
	"passwordsaver/internal/files"
	"passwordsaver/internal/output"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.ValultWithDb){
	"1": createAccount,
	"2": findAccountByULR,
	"3": findAccountByLogin,
	"4": deleteAccount,
	"5": findLenOfAccounts,
}

var menuVariants = []string{
	"1. Создать аккаунт",
	"2. Найти аккаунт URL",
	"3. Найти аккаунт по Логину",
	"4. Удалить аккаунт",
	"5. Показать количество сохранённых аккаунтов",
	"6. Выход",
	"Выберите вариант",
}

func main() {
	color.Cyan("Добро пожаловать в менеджер паролей!")
	err := godotenv.Load("configs/.env")
	if err != nil {
		output.PrintError("Ошибка загрузки .env файла")
	}
	ensureVaultFile("data.vault")
	valult := account.NewVault(files.NewJsonDb("data.vault"), *encrypter.NewEncrypter())

Menu:
	for {
		variant := promtData(menuVariants...)
		menuFunc := menu[variant]
		if menuFunc == nil {
			break Menu
		}

		menuFunc(valult)
	}
}

func ensureVaultFile(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.WriteFile(path, []byte("[]"), 0644)
		if err != nil {
			output.PrintError("Не удалось создать файл data.vault: " + err.Error())
		}
	}
}

func createAccount(valult *account.ValultWithDb) {
	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	url := promtData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)

	if err != nil {
		output.PrintError("Неверный формат URL или Логин")
		return
	}
	valult.AddAccount(*myAccount)
}

// Функция вывода с использованием Generic Type
func promtData[T string](prompt ...T) string {
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
	url := promtData("Введите URL для поиска")
	accounts, err := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	NotFound(accounts, err)
}

func findAccountByLogin(vault *account.ValultWithDb) {
	login := promtData("Введите Логин для поиска")
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
	url := promtData("Введите URL для удаления")
	isDeleted := vault.DeleteAccountByURL(url)
	if isDeleted {
		color.Green("Успешно удалено!")
	} else {
		output.PrintError("Аккаунт не найден!")
	}
}

func findLenOfAccounts(vault *account.ValultWithDb) {
	if len(vault.Accounts) == 0 {
		output.PrintError("Аккаунты не найдены!")
		return
	}
	color.Green("Количество аккаунтов: %d", len(vault.Accounts))
}
