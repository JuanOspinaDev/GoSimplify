package simplify

import (
	"errors"
	"strconv"
	"unicode/utf8"
)

// Str convierte un entero a string.
func Str(i int) string {
	return strconv.Itoa(i)
}

// Int convierte un string a entero y retorna 0 si la conversión falla.
func Int(s string) int {
    i, err := strconv.Atoi(s)
    if err != nil {
        return 0 // Retorna 0 si hay error
    }
    return i
}

// Pop elimina y devuelve el último elemento de un slice de enteros.
func Pop(slice *[]int) (int, error) {
	if len(*slice) == 0 {
		return 0, errors.New("cannot pop from an empty slice")
	}
	index := len(*slice) - 1
	element := (*slice)[index]
	*slice = (*slice)[:index]
	return element, nil
}

// Contains verifica si un elemento está en un slice de enteros.
func Contains(slice []int, element int) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

// Reverse invierte un slice de enteros.
func Reverse(slice []int) []int {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

// Max devuelve el valor máximo en un slice de enteros.
func Max(slice []int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("cannot find max of an empty slice")
	}
	max := slice[0]
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max, nil
}

// Min devuelve el valor mínimo en un slice de enteros.
func Min(slice []int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("cannot find min of an empty slice")
	}
	min := slice[0]
	for _, v := range slice {
		if v < min {
			min = v
		}
	}
	return min, nil
}

// Sum suma todos los elementos de un slice de enteros.
func Sum(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

// Pop elimina una sección del string basado en el índice proporcionado y devuelve la parte eliminada y la nueva string.
// Si el índice es positivo, elimina desde el inicio hasta el índice (exclusivo).
// Si el índice es negativo, elimina desde el índice hasta el final.
// Si no se proporciona índice, elimina el último carácter.
func PopString(s string, indices ...int) (popped string, newString string, err error) {
	if len(s) == 0 {
		return "", "", errors.New("cannot pop from an empty string")
	}

	// Determinar el índice de corte
	var index int
	if len(indices) == 0 {
		// No se proporcionó índice, eliminar el último carácter
		index = utf8.RuneCountInString(s) - 1
	} else {
		index = indices[0]
	}

	// Manejar índices negativos
	if index < 0 {
		index = utf8.RuneCountInString(s) + index
	}

	// Validar el índice
	if index < 0 || index >= utf8.RuneCountInString(s) {
		return "", "", errors.New("index out of range")
	}

	// Convertir el string en una slice de runas para manejar correctamente caracteres multibyte
	runes := []rune(s)

	// Separar las partes
	if index >= 0 {
		popped = string(runes[:index])
		newString = string(runes[index:])
	} else {
		popped = string(runes[index:])
		newString = string(runes[:index])
	}

	return popped, newString, nil
}
