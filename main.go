package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

// Буквы и символы для генерации пароля (глобальная переменная)
var LetterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-*!")

// Основная структура аккаунта (как "коробка" для данных)
type account struct {
	login    string // ящик для логина
	password string // ящик для пароля
	url      string // ящик для сайта
}

// Расширенная структура аккаунта с датами (наследует account)
type accountWithTimeStamp struct {
	createdAt time.Time // когда создали запись
	updatedAt time.Time // когда обновляли
	account             // встраиваем основную структуру (всё что есть в account)
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// МЕТОД ДЛЯ ГЕНЕРАЦИИ ПАРОЛЯ
// (приклеен к структуре account)
//
// (acc *account) - работаем с ОРИГИНАЛОМ структуры, а не копией
// n int - длина пароля
// НИЧЕГО не возвращаем (нет 'string' в объявлении)
func (acc *account) generatePassword(n int) {
	// 1. Создаём "пустую коробку" для рун (символов)
	newGenPassword := make([]rune, n)

	// 2. Наполняем коробку случайными символами
	for i := range newGenPassword {
		newGenPassword[i] = LetterRunes[rand.IntN(len(LetterRunes))]
	}

	// 3. Превращаем руны в строку и кладём в структуру
	acc.password = string(newGenPassword)
	// return не нужен, потому что метод ничего не возвращает
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// МЕТОД ДЛЯ ВЫВОДА ДАННЫХ
// (приклеен к account)
func (acc *account) outputPassword() {
	// Просто печатаем содержимое структуры
	fmt.Println("🔑 Логин:", acc.login)
	fmt.Println("🔒 Пароль:", acc.password)
	fmt.Println("🌐 Сайт:", acc.url)
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// КОНСТРУКТОР ДЛЯ ОСНОВНОЙ СТРУКТУРЫ (account)
// (это не метод, а обычная функция)
//
// Возвращает:
//   - *account (указатель на структуру)
//   - error (ошибку, если что-то пошло не так)
func newAccount(login, password, urlString string) (*account, error) {
	// Проверяем логин
	if login == "" {
		return nil, errors.New("ЛОГИН НЕ МОЖЕТ БЫТЬ ПУСТЫМ")
	}

	// Проверяем URL (валидный ли он)
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("НЕПРАВИЛЬНЫЙ URL")
	}

	// Создаём аккаунт (с указателем &)
	newAcc := &account{
		login:    login,
		password: password,
		url:      urlString,
	}

	// Если пароль пустой - генерируем автоматически
	if password == "" {
		newAcc.generatePassword(12) // вызываем метод структуры
	}

	return newAcc, nil // возвращаем аккаунт и nil (нет ошибки)
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
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

	// Создаём расширенную структуру:
	newAcc := &accountWithTimeStamp{
		account: account{ // встраиваем основную структуру
			login:    login,
			password: password,
			url:      urlString,
		},
		createdAt: time.Now(), // текущее время создания
		updatedAt: time.Now(), // текущее время обновления
	}

	// Автогенерация пароля если нужно
	if password == "" {
		newAcc.generatePassword(12) // унаследованный метод
	}

	return newAcc, nil
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// ГЛАВНАЯ ФУНКЦИЯ (точка входа в программу)
func main() {
	// Запрашиваем данные у пользователя
	login := promtData("Введите логин:")
	password := promtData("Введите пароль (оставьте пустым для автогенерации):")
	url := promtData("Введите URL сайта:")

	// Пытаемся создать аккаунт
	myAccount, err := newAccountWithTimeStamp(login, password, url)
	// Если была ошибка - печатаем и выходим
	if err != nil {
		fmt.Println("💥 ОШИБКА:", err)
		return // завершаем программу
	}

	// Выводим результат
	myAccount.outputPassword()
	fmt.Println("⏰ Дата создания:", myAccount.createdAt.Format("2006-01-02 15:04:05"))
}

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// ВСПОМОГАТЕЛЬНАЯ ФУНКЦИЯ ДЛЯ ВВОДА ДАННЫХ
func promtData(prompt string) string {
	fmt.Print(prompt + " ")
	var res string
	fmt.Scanln(&res) // читаем ввод пользователя
	return res
}
