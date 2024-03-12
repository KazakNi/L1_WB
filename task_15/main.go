package main

import (
	"strings"
)

/* К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
 Приведите корректный пример реализации.


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

func main() {
	/* Минусы реализации:
		   - глобальная переменная justString, ссылка на которую "держит" в памяти огромный массив байт
		   - срез строки, т.е. произвольного кол-ва байтов, может дать неверный отображение символов в произвольном месте

	     Улучшения:
	     1. Создаем буфер или обнуляем переменную с массивом
	     2. Копируем руны*/

}
func someFunc() string {
	v := createHugeString(1 << 10)
	justString := strings.Clone(v[:100])
	v = nil
	return justString
}
