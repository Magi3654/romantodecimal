package main

import (
	"fmt"
)

// Función para convertir números romanos a arábigos
func romanoAAarabigo(entrada string) int {
	// Definimos los valores de cada letra romana
	valores := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0 // Inicializamos el total en 0

	for i := 0; i < len(entrada); i++ {
		valorActual := valores[entrada[i]] // Obtenemos el valor del carácter actual

		// Verificamos si hay un siguiente carácter
		if i+1 < len(entrada) {
			valorSiguiente := valores[entrada[i+1]]
			// Si el valor actual es menor que el siguiente, restamos
			if valorActual < valorSiguiente {
				total -= valorActual
			} else {
				total += valorActual
			}
		} else {
			// Si no hay siguiente carácter, simplemente sumamos
			total += valorActual
		}
	}

	return total // Devolvemos el total calculado
}

// Función para convertir números arábigos a romanos
func arabigoARomano(numero int) string {
	// Definimos los valores de los números romanos en orden descendente
	valores := []struct {
		valor   int
		simbolo string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	resultado := "" // Inicializamos el resultado como una cadena vacía

	for _, valor := range valores {
		// Mientras el número sea mayor o igual al valor, añadimos el símbolo
		for numero >= valor.valor {
			resultado += valor.simbolo
			numero -= valor.valor // Restamos el valor del número
		}
	}

	return resultado // Devolvemos el resultado en formato romano
}

func main() {
	// Ejemplo de uso
	entradaRomano := "MMMIX" // Este es el número romano para 3999
	resultadoArabigo := romanoAAarabigo(entradaRomano)
	fmt.Printf("El valor de la entrada '%s' es: %d\n", entradaRomano, resultadoArabigo)

	entradaArabigo := 395 // Este es el número arábigo que queremos convertir
	resultadoRomano := arabigoARomano(entradaArabigo)
	fmt.Printf("El valor de la entrada '%d' en romano es: %s\n", entradaArabigo, resultadoRomano)
}