package output

import (
	"github.com/fatih/color"
)

func PrintError(value any) {
	switch t := value.(type) {
	case string:
		color.Red(t)
	case int:
		color.Red("КОД ОШИБКИ: %d", t)
	case error:
		color.Red(t.Error())
	default:
		color.Red("НЕИЗВЕСТНЫЙ ТИП ОШИБКИ")
	}

}
