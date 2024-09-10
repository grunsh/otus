package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func isLatter(s rune) bool {
	return (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z')
}

func isDigit(s rune) bool {
	return s >= '0' && s <= '9'
}

func Unpack(st string) (string, error) {
	var s strings.Builder
	var tempStr string
	var incr int
	var repeat int // сколько раз писать литерал
	var err error
	i := 0
	for i < len(st) {
		// Вделяем литерал, который повторять 0..9 раз
		if isLatter(rune(st[i])) { //nolint:gocritic,nestif
			tempStr = st[i : i+1]
			incr = 1
		} else if st[i] == '\\' { // Встретили слэш. Анализируем строку вперёд
			if len(st) > i+1 {
				if st[i+1] == 'n' { // У нас тут перевод строки
					tempStr = st[i : i+2]
					incr = 2
				} else { // после сэша не n, ошибка
					return "", ErrInvalidString
				}
			} else { // Слэш за которым строка кончается. Ошибка.
				return "", ErrInvalidString
			}
		} else { // Не буква и не слэш, ошибка
			return "", ErrInvalidString
		}
		// Анализируем строку дальше, есть ли цифра, какая она или строка кончилась
		repeat = 1            // Если цифры далее нет, то 1, если есть переопределится в цифру
		if len(st) > i+incr { // Проверим на конец строки, чтобы заглянуть, а не цифра ли дальше?
			if isDigit(rune(st[i+incr])) { // Не цифра ли следом?
				repeat, err = strconv.Atoi(st[i+incr : i+incr+1])
				if err != nil {
					panic(err)
				}
				i++ // Учитывая, что получили цифру, надо на 1 сдвинуть указатель
			}
		}
		// Записываем в итоговую строку литерал tempStr repeat раз
		s.WriteString(strings.Repeat(tempStr, repeat))
		i += incr // Указатель сдвинем на 1, если буква, на 2 если \n
	}
	return s.String(), nil
}
