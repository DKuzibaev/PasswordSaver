package account

import (
	"encoding/json"
	"passwordsaver/internal/encrypter"
	"passwordsaver/internal/output"
	"strings"
	"time"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type Db interface {
	ByteReader
	ByteWriter
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ValultWithDb struct {
	Vault
	db  Db
	enc encrypter.Encrypter
}

// Конструктор структуры Valult
func NewVault(db Db, enc encrypter.Encrypter) *ValultWithDb {
	file, err := db.Read()
	if err != nil {
		// Ошибка чтения файла — создаём пустой Vault
		return &ValultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}

	// Проверяем длину данных перед дешифровкой
	if len(file) == 0 || len(file) < enc.NonceSize() {
		// Пустой или слишком короткий файл — создаём пустой Vault
		return &ValultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}

	decryptedData, err := enc.Decrypt(file)
	if err != nil {
		panic("Failed to decrypt data: " + err.Error())
	}

	var valult Vault
	err = json.Unmarshal(decryptedData, &valult)
	if err != nil {
		output.PrintError("не удалось разобрать файл data.vault")
		return &ValultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}

	return &ValultWithDb{
		Vault: valult,
		db:    db,
		enc:   enc,
	}
}

// Сохранение и обновление аккаунтов
func (v *ValultWithDb) save() {
	v.UpdatedAt = time.Now()
	data, err := v.Vault.ToByteSlice()
	encryptedData := v.enc.Encrypt(data)
	if err != nil {
		output.PrintError(err)
	}
	v.db.Write(encryptedData)
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

// Поиск аккаунта по URL
func (v *ValultWithDb) FindAccounts(str string, cheker func(Account, string) bool) ([]Account, error) {
	var accounts []Account
	for _, acc := range v.Accounts {
		isMatch := cheker(acc, str)
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

// показать все аккаунты
func (v *ValultWithDb) ShowAll() ([]Account, error) {
	var accounts []Account
	for _, acc := range v.Accounts {
		accounts = append(accounts, acc)
	}
	return accounts, nil
}

// Группировка аккаунтов по тегам
func (v *ValultWithDb) GroupByTag() map[string][]Account {
	// создаём пустую карту: ключ — тег, значение — список аккаунтов
	tagGroup := make(map[string][]Account)

	// бежим по всем аккаунтам в хранилище
	for _, acc := range v.Accounts {
		// добавляем аккаунт в нужную группу по тегу
		tagGroup[acc.Tag] = append(tagGroup[acc.Tag], acc)
	}

	// возвращаем сгруппированные аккаунты
	return tagGroup
}
