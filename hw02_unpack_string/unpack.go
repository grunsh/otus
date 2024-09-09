package hw02_unpack_string

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(st string) (string, error) {
	var s strings.Builder
	var ErrInvalidString error = errors.New("invalid string")
	i := 0
	for i < len(st) {
		// Если буква, то ...
		if (st[i] >= 'a' && st[i] <= 'z') || (st[i] >= 'A' && st[i] <= 'Z') {
			if len(st) > i+1 { // Надо загялнуть на символ дальше, потому проверим, не выйдем ли за пределы строки
				if st[i+1] >= '0' && st[i+1] <= '9' { // Если следом за буквой цифра, то пишем повторы в строку
					if r, err := strconv.Atoi(st[i+1 : i+2]); err == nil {
						s.WriteString(strings.Repeat(st[i:i+1], r))
						i += 2
					}
				} else { // Текущая буква и следом НЕ цыфра, росто добавляем букву результирующую строку.
					s.WriteString(st[i : i+1])
					i++
				}
			} else { // Строка за текущей буквой заканчивается. Цифры точно нет, повторять нечего, пишем в результат.
				s.WriteString(st[i : i+1])
				i++
			}
		} else if st[i] == '\\' { // Встретили слэш. Анализируем строку вперёд
			if len(st) > i+1 {
				if st[i+1] == 'n' { // У нас тут перевод строки
					if len(st) > i+2 {
						if st[i+2] >= '0' && st[i+2] <= '9' { // Перевод строки надо повторить st[i+2] раз, ибо следом цифра
							if r, err := strconv.Atoi(st[i+2 : i+3]); err == nil {
								s.WriteString(strings.Repeat(st[i:i+2], r))
								i += 3 // У нас тут \n и цифра, двигаем указатель сразу на 3
							}
						} else { // Нет цыфры за \n, поэтому просто пишем один \n и двигаем указатель на 2
							s.WriteString(st[i : i+2])
							i += 2
						}
					} else { // На \n строка заканчивается. Пишем его в итог и двигаем указатель на 2
						s.WriteString(st[i : i+2])
						i += 2
					}
				} else {
					return "", ErrInvalidString
				}
			}
		} else {
			return "", ErrInvalidString
		}
	}
	return "", nil
}
