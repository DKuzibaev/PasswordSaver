<<<<<<< HEAD
package main

import (
	"bufio"
	"fmt"
	"os"
	"passwordsaver/account"
	"passwordsaver/files"
	"strings"

	"github.com/fatih/color"
)

func main() {
	color.Cyan("Добро пожаловать в менеджер паролей!")
	valult := account.NewVault(files.NewJsonDb("data.json"))
Menu:
	for {
		variant := getMenu(valult)
		switch variant {
		case 1:
			createAccount(valult)
		case 2:
			findAccout(valult)
		case 3:
			deleteAccout(valult)
		default:
			break Menu
		}
	}
}

func getMenu(valult *account.ValultWithDb) int {
	color.Cyan("----------------------------------------")
	color.Cyan("Выберите пункт: ")
	color.Cyan("1. Создать аккаунт")
	color.Cyan("2. Найти аккаунт")
	color.Cyan("3. Удалить аккаунт")
	color.Cyan("4. Выход")
	color.Cyan("----------------------------------------")

	var userInput int
	if _, err := fmt.Scan(&userInput); err != nil {
		color.Red("Ошибка ввода! Введите число от 1 до 4!")
		bufio.NewReader(os.Stdin).ReadString('\n')
		return 0
	}
	bufio.NewReader(os.Stdin).ReadString('\n')
	return userInput
}

func createAccount(valult *account.ValultWithDb) {
	login := promtData("Введите логин:")
	password := promtData("Введите пароль:")
	url := promtData("Введите URL:")

	myAccount, err := account.NewAccount(login, password, url)

	if err != nil {
		fmt.Printf("ОШИБКА: %s\n", err)
		return
	}
	valult.AddAccount(*myAccount)
}

func promtData(prompt string) string {
	fmt.Print(prompt + " ")
	reader := bufio.NewReader(os.Stdin)
	res, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return ""
	}
	return strings.TrimSpace(strings.Trim(res, "\r\n"))
}

func findAccout(vault *account.ValultWithDb) {
	url := promtData("Введите URL для поиска:")
	accounts, err := vault.FindAccountByURL(url)
	if err != nil {
		color.Red("Ошибка: %s\n", err)
		return
	}
	if len(accounts) == 0 {
		color.Red("Аккаунт не найден!")
	}
	for _, acc := range accounts {
		acc.Output()
	}
}

func deleteAccout(vault *account.ValultWithDb) {
	url := promtData("Введите URL для удаления:")
	isDeleted := vault.DeleteAccountByURL(url)
	if isDeleted {
		color.Green("Успешно удалено!")
	} else {
		color.Red("Аккаунт не найден!")
	}

}
=======
package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"passwordsaver/account"
	"strings"

	"github.com/fatih/color"
)

func main() {
	color.Cyan("Добро пожаловать в менеджер паролей!")

Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount()
		case 2:
			findAccout()
		case 3:
			deleteAccout()
		default:
			break Menu
		}
	}
}

func getMenu() int {
	color.Cyan("----------------------------------------")
	color.Cyan("Выберите пункт: ")
	color.Cyan("1. Создать аккаунт")
	color.Cyan("2. Найти аккаунт")
	color.Cyan("3. Удалить аккаунт")
	color.Cyan("4. Выход")
	color.Cyan("----------------------------------------")

	var userInput int
	if _, err := fmt.Scan(&userInput); err != nil {
		color.Red("Ошибка ввода! Введите число от 1 до 4!")
		bufio.NewReader(os.Stdin).ReadString('\n')
		return 0
	}
	bufio.NewReader(os.Stdin).ReadString('\n')
	return userInput
}

func createAccount() {
	login := promtData("Введите логин:")
	password := promtData("Введите пароль:")
	url := promtData("Введите URL:")

	myAccount, err := account.NewAccount(login, password, url)

	if err != nil {
		fmt.Printf("ОШИБКА: %s\n", err)
		return
	}

	valult := account.NewVault()
	valult.AddAccount(*myAccount)
}

func promtData(prompt string) string {
	fmt.Print(prompt + " ")
	reader := bufio.NewReader(os.Stdin)
	res, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return ""
	}
	return strings.TrimSpace(strings.Trim(res, "\r\n"))
}

func findAccout() {
	// запросить url
	fmt.Print("Введите URL: ")
	var urlToFind string
	fmt.Scan(&urlToFind)
	_, err := url.ParseRequestURI(urlToFind)
	if err != nil {
		fmt.Println(err)
	}
	// поиск
	acc, err := account.FindAccountByURL(account.NewVault(), urlToFind)
	if err != nil {
		color.Red("Ошибка")
	}
	fmt.Println(acc)
	// вывод
}

func deleteAccout() {}
>>>>>>> f6768fafe261a74e9fb9c5c9f8529acbea46d48e
