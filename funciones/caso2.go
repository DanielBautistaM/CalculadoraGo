package funciones

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func CifradoModulo27() {
	fmt.Println("\n🔒 Cifrado Módulo 27")
	fmt.Println("=============================================")

	alfabeto := "ABCDEFGHIJKLMNÑOPQRSTUVWXYZ"

	mensaje := readEntry("Ingresa el mensaje a cifrar: ")
	clave := readNumber("Ingresa la clave (desplazamiento): ")

	mensaje = strings.ToUpper(mensaje)

	opcion := readEntry("¿Quieres cifrar o descifrar? (C/D): ")

	var resultado strings.Builder

	if strings.ToUpper(opcion) == "D" {
		clave = -clave
	}

	clave = ((clave % 27) + 27) % 27

	for _, char := range mensaje {
		idx := strings.IndexRune(alfabeto, char)
		if idx != -1 {
			nuevaPos := (idx + clave) % 27
			resultado.WriteRune(rune(alfabeto[nuevaPos]))
		} else {
			resultado.WriteRune(char)
		}
	}

	if strings.ToUpper(opcion) == "C" {
		fmt.Printf("\n✅ Mensaje cifrado: %s\n", resultado.String())
		fmt.Printf("\nLa fórmula aplicada fue: C = (P + %d) mod 27\n", clave)
	} else {
		fmt.Printf("\n✅ Mensaje descifrado: %s\n", resultado.String())
		fmt.Printf("\nLa fórmula aplicada fue: P = (C - %d) mod 27\n", clave)
	}

	fmt.Println("\nTabla de sustitución:")
	fmt.Println("Original:  " + alfabeto)
	cifrado := alfabeto[clave:] + alfabeto[:clave]
	fmt.Println("Cifrado:   " + cifrado)
}

func CifradoCesar() {
	fmt.Println("\n🔒 Cifrado César")
	fmt.Println("=============================================")

	alfabeto := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	mensaje := readEntry("Ingresa el mensaje a cifrar: ")
	clave := readNumber("Ingresa la clave (desplazamiento entre 1-25): ")

	clave = ((clave % 26) + 26) % 26
	if clave == 0 {
		clave = 3
		fmt.Println("Usando cifrado César clásico con clave 3")
	}

	mensaje = strings.ToUpper(mensaje)

	opcion := readEntry("¿Quieres cifrar o descifrar? (C/D): ")

	var resultado strings.Builder

	if strings.ToUpper(opcion) == "D" {
		clave = (26 - clave) % 26
	}

	for _, char := range mensaje {
		if strings.ContainsRune(alfabeto, char) {
			pos := strings.IndexRune(alfabeto, char)

			nuevaPos := (pos + clave) % 26

			resultado.WriteRune(rune(alfabeto[nuevaPos]))
		} else {
			resultado.WriteRune(char)
		}
	}

	if strings.ToUpper(opcion) == "C" {
		fmt.Printf("\n✅ Mensaje cifrado: %s\n", resultado.String())
		fmt.Printf("\nLa fórmula aplicada fue: C = (P + %d) mod 26\n", clave)
	} else {
		fmt.Printf("\n✅ Mensaje descifrado: %s\n", resultado.String())
		fmt.Printf("\nLa fórmula aplicada fue: P = (C + %d) mod 26\n", clave)
	}

	fmt.Println("\nTabla de sustitución:")
	fmt.Println("Original:  " + alfabeto)
	cifrado := alfabeto[clave:] + alfabeto[:clave]
	fmt.Println("Cifrado:   " + cifrado)
}

func CifradoVernam() {
	fmt.Println("\n🔒 Cifrado Vernam (XOR)")
	fmt.Println("=============================================")

	mensaje := readEntry("Ingresa el mensaje a cifrar: ")
	clave := readEntry("Ingresa la clave (debe ser al menos igual de larga que el mensaje): ")

	if len(clave) < len(mensaje) {
		fmt.Println("⚠️ Advertencia: La clave debe ser al menos igual de larga que el mensaje.")
		fmt.Println("   Se repetirá la clave para cubrir el mensaje completo (esto reduce la seguridad).")
		for len(clave) < len(mensaje) {
			clave += clave
		}
	}

	clave = clave[:len(mensaje)]

	fmt.Println("Clave efectiva:", clave)

	opcion := readEntry("¿Quieres cifrar o descifrar? (C/D): ")

	var resultado strings.Builder

	fmt.Println("\nProceso de cifrado XOR bit a bit:")
	fmt.Println("Mensaje | Clave | Resultado")
	fmt.Println("-------------------------")

	for i := 0; i < len(mensaje); i++ {
		resultadoChar := byte(mensaje[i]) ^ byte(clave[i])
		resultado.WriteByte(resultadoChar)

		if i < 5 {
			fmt.Printf("%08b | %08b | %08b (%c XOR %c = %c)\n",
				byte(mensaje[i]), byte(clave[i]), resultadoChar,
				mensaje[i], clave[i], resultadoChar)
		}
	}

	if len(mensaje) > 5 {
		fmt.Println("... (mostrando solo los primeros 5 caracteres)")
	}

	resultadoHex := ""
	for i := 0; i < resultado.Len(); i++ {
		resultadoHex += fmt.Sprintf("%02x", resultado.String()[i])
	}

	if strings.ToUpper(opcion) == "C" {
		fmt.Printf("\n✅ Mensaje cifrado (hex): %s\n", resultadoHex)
	} else {
		fmt.Printf("\n✅ Mensaje descifrado (hex): %s\n", resultadoHex)
		fmt.Printf("Mensaje en texto: %s\n", resultado.String())
	}

	fmt.Println("\nNota: El cifrado Vernam con XOR es perfectamente seguro solo cuando:")
	fmt.Println("1. La clave es completamente aleatoria")
	fmt.Println("2. La clave es del mismo tamaño que el mensaje")
	fmt.Println("3. La clave nunca se reutiliza")
}

func CifradoAtbash() {
	fmt.Println("\n🔒 Cifrado ATBASH")
	fmt.Println("=============================================")

	// ATBASH funciona con un alfabeto invertido
	alfabetoNormal := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alfabetoInvertido := "ZYXWVUTSRQPONMLKJIHGFEDCBA"

	mensaje := readEntry("Ingresa el mensaje a cifrar/descifrar: ")
	mensaje = strings.ToUpper(mensaje)

	var resultado strings.Builder

	// ATBASH es un cifrado recíproco, la operación de cifrado y descifrado
	// es la misma (para cualquiera de los dos, se sustituye cada letra por su opuesta)
	for _, char := range mensaje {
		idx := strings.IndexRune(alfabetoNormal, char)
		if idx != -1 {
			// Sustituir con la letra correspondiente del alfabeto invertido
			resultado.WriteRune(rune(alfabetoInvertido[idx]))
		} else {
			// Si no es una letra, mantenerla igual
			resultado.WriteRune(char)
		}
	}

	fmt.Printf("\n✅ Mensaje procesado con ATBASH: %s\n", resultado.String())
	fmt.Println("\nNota: El cifrado ATBASH es recíproco, significa que")
	fmt.Println("      el proceso de cifrado y descifrado es el mismo.")

	// Mostrar la tabla de sustitución
	fmt.Println("\nTabla de sustitución completa:")
	fmt.Println("Original | Cifrado")
	fmt.Println("-------------------")
	for i := 0; i < len(alfabetoNormal); i++ {
		fmt.Printf("   %c    |    %c\n", alfabetoNormal[i], alfabetoInvertido[i])
	}
}

func CifradoTransposicionColumnar() {
	fmt.Println("\n🔒 Cifrado Transposición Columnar")
	fmt.Println("=============================================")

	mensaje := readEntry("Ingresa el mensaje a cifrar: ")
	clave := readEntry("Ingresa la clave (número): ")

	if len(clave) == 0 {
		fmt.Println("❌ Error: La clave no puede estar vacía.")
		return
	}
	opcion := readEntry("¿Quieres cifrar o descifrar? (C/D): ")

	if strings.ToUpper(opcion) == "C" {
		// CIFRADO

		// Crear la matriz
		numColumnas := len(clave)
		numFilas := (len(mensaje) + numColumnas - 1) / numColumnas // Redondeamos hacia arriba

		// Matriz para almacenar el mensaje
		matriz := make([][]rune, numFilas)
		for i := range matriz {
			matriz[i] = make([]rune, numColumnas)
			for j := range matriz[i] {
				matriz[i][j] = ' ' // Inicializamos con espacios
			}
		}

		// Llenar la matriz con el mensaje
		k := 0
		for i := 0; i < numFilas; i++ {
			for j := 0; j < numColumnas && k < len(mensaje); j++ {
				matriz[i][j] = rune(mensaje[k])
				k++
			}
		}

		// Mostrar la matriz
		fmt.Println("\nMatriz de transposición:")
		fmt.Print("  ")
		for j := 0; j < numColumnas; j++ {
			fmt.Printf(" %c ", clave[j])
		}
		fmt.Println()

		for i := 0; i < numFilas; i++ {
			fmt.Printf("%d ", i+1)
			for j := 0; j < numColumnas; j++ {
				fmt.Printf(" %c ", matriz[i][j])
			}
			fmt.Println()
		}

		// Determinar el orden de las columnas basado en la clave
		orden := obtenerOrdenColumnas(clave)

		// Leer las columnas en el orden determinado por la clave
		var resultado strings.Builder
		for _, idx := range orden {
			for i := 0; i < numFilas; i++ {
				if matriz[i][idx] != ' ' {
					resultado.WriteRune(matriz[i][idx])
				}
			}
		}

		fmt.Printf("\n✅ Mensaje cifrado: %s\n", resultado.String())
		fmt.Println("\nOrden de lectura de columnas basado en la clave:")
		for j, idx := range orden {
			fmt.Printf("%c(%d) ", clave[idx], j+1)
		}
		fmt.Println()

	} else {
		// DESCIFRADO - Implementación simplificada
		fmt.Println("\n❌ La funcionalidad de descifrado está en desarrollo.")
		fmt.Println("Puedes descifrar manualmente siguiendo el proceso inverso:")
		fmt.Println("1. Ordena las letras de la clave alfabéticamente")
		fmt.Println("2. Distribuye el mensaje cifrado en columnas según este orden")
		fmt.Println("3. Lee la matriz por filas para obtener el mensaje original")
	}
}

// Función auxiliar para obtener el orden de las columnas basado en la clave
func obtenerOrdenColumnas(clave string) []int {
	// Crear pares (carácter, índice)
	type Par struct {
		Char rune
		Idx  int
	}

	pares := make([]Par, len(clave))
	for i, c := range clave {
		pares[i] = Par{Char: c, Idx: i}
	}

	// Ordenar por carácter
	sort.Slice(pares, func(i, j int) bool {
		return pares[i].Char < pares[j].Char
	})

	// Extraer los índices ordenados
	orden := make([]int, len(clave))
	for i, par := range pares {
		orden[i] = par.Idx
	}

	return orden
}

// CifradoAfin implementa el cifrado afín: C = (aP + b) mod m
func CifradoAfin() {
	fmt.Println("\n🔒 Cifrado Afín")
	fmt.Println("=============================================")

	// Alfabeto estándar de 26 letras (inglés)
	alfabeto := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	m := 26 // Tamaño del alfabeto

	mensaje := readEntry("Ingresa el mensaje a cifrar/descifrar: ")
	a := readNumber("Ingresa el valor de 'a' (debe ser coprimo con 26): ")
	b := readNumber("Ingresa el valor de 'b': ")

	// Convertimos a mayúsculas
	mensaje = strings.ToUpper(mensaje)

	// Verificar que 'a' sea coprimo con 'm'
	if calcularMCD(a, m) != 1 {
		fmt.Printf("❌ Error: 'a' (%d) debe ser coprimo con %d.\n", a, m)
		fmt.Println("   Valores válidos para 'a': 1, 3, 5, 7, 9, 11, 15, 17, 19, 21, 23, 25")
		return
	}

	opcion := readEntry("¿Quieres cifrar o descifrar? (C/D): ")

	var resultado strings.Builder

	if strings.ToUpper(opcion) == "C" {
		// CIFRADO: C = (aP + b) mod m
		for _, char := range mensaje {
			idx := strings.IndexRune(alfabeto, char)
			if idx != -1 {
				// Aplicar la fórmula de cifrado afín
				nuevaPos := (a*idx + b) % m
				resultado.WriteRune(rune(alfabeto[nuevaPos]))
			} else {
				// Si no es una letra, mantenerla igual
				resultado.WriteRune(char)
			}
		}

		fmt.Printf("\n✅ Mensaje cifrado: %s\n", resultado.String())
		fmt.Printf("\nFórmula aplicada: C = (%d*P + %d) mod %d\n", a, b, m)

	} else {
		// DESCIFRADO: P = a^(-1) * (C - b) mod m

		// Calcular el inverso multiplicativo de 'a'
		var aInverso int

		// Buscar el inverso multiplicativo de 'a' mod m
		for i := 1; i < m; i++ {
			if (a*i)%m == 1 {
				aInverso = i
				break
			}
		}

		for _, char := range mensaje {
			idx := strings.IndexRune(alfabeto, char)
			if idx != -1 {
				// Aplicar la fórmula de descifrado afín
				// P = a^(-1) * (C - b) mod m
				temp := (idx - b) % m
				if temp < 0 {
					temp += m // Asegurarse de que sea positivo
				}
				nuevaPos := (aInverso * temp) % m
				resultado.WriteRune(rune(alfabeto[nuevaPos]))
			} else {
				// Si no es una letra, mantenerla igual
				resultado.WriteRune(char)
			}
		}

		fmt.Printf("\n✅ Mensaje descifrado: %s\n", resultado.String())
		fmt.Printf("\nFórmula aplicada: P = %d * (C - %d) mod %d\n", aInverso, b, m)
	}

	// Mostrar la tabla de sustitución
	fmt.Println("\nTabla de sustitución completa:")
	fmt.Println("Original | Cifrado")
	fmt.Println("-------------------")
	for i := 0; i < m; i++ {
		cifrado := (a*i + b) % m
		fmt.Printf("   %c    |    %c\n", alfabeto[i], alfabeto[cifrado])
	}
}

// CifradoSustitucion implementa el cifrado de sustitución simple con un alfabeto personalizado
func CifradoSustitucion() {
	fmt.Println("\n🔒 Cifrado de Sustitución Simple")
	fmt.Println("=============================================")

	// Alfabeto estándar
	alfabetoNormal := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	mensaje := readEntry("Ingresa el mensaje a cifrar/descifrar: ")

	// Convertimos a mayúsculas
	mensaje = strings.ToUpper(mensaje)

	// Opciones para el alfabeto de sustitución
	fmt.Println("\nSelecciona una opción para el alfabeto de sustitución:")
	fmt.Println("1. Ingresar un alfabeto personalizado")
	fmt.Println("2. Generar un alfabeto aleatorio")
	fmt.Println("3. Usar una palabra clave para generar el alfabeto")

	opcionAlfabeto := readNumber("Selecciona una opción (1-3): ")

	var alfabetoCifrado string

	switch opcionAlfabeto {
	case 1:
		// Alfabeto personalizado
		alfabetoCifrado = readEntry("Ingresa el alfabeto de sustitución (26 letras sin repetir): ")
		alfabetoCifrado = strings.ToUpper(alfabetoCifrado)

		// Verificar que tenga 26 letras y no haya repetidas
		if len(alfabetoCifrado) != 26 {
			fmt.Println("❌ Error: El alfabeto debe tener exactamente 26 letras.")
			return
		}

		// Verificar que no haya letras repetidas
		for i, c := range alfabetoCifrado {
			if strings.IndexRune(alfabetoCifrado[i+1:], c) >= 0 {
				fmt.Printf("❌ Error: La letra '%c' aparece más de una vez en el alfabeto.\n", c)
				return
			}
		}

	case 2:
		// Generar alfabeto aleatorio
		fmt.Println("Generando alfabeto aleatorio...")

		// Convertir a slice para poder mezclarlo
		letras := []rune(alfabetoNormal)

		// Semilla aleatoria
		rand.Seed(time.Now().UnixNano())

		// Mezclar el slice
		rand.Shuffle(len(letras), func(i, j int) {
			letras[i], letras[j] = letras[j], letras[i]
		})

		alfabetoCifrado = string(letras)

	case 3:
		// Usar palabra clave
		clave := readEntry("Ingresa una palabra clave (sin letras repetidas): ")
		clave = strings.ToUpper(clave)

		// Eliminar letras repetidas de la clave
		var claveSinRepetir strings.Builder
		for _, c := range clave {
			if strings.IndexRune(claveSinRepetir.String(), c) == -1 && strings.ContainsRune(alfabetoNormal, c) {
				claveSinRepetir.WriteRune(c)
			}
		}

		// Construir el alfabeto cifrado
		alfabetoCifrado = claveSinRepetir.String()

		// Añadir el resto de las letras que no están en la clave
		for _, c := range alfabetoNormal {
			if strings.IndexRune(alfabetoCifrado, c) == -1 {
				alfabetoCifrado += string(c)
			}
		}

	default:
		fmt.Println("❌ Opción no válida. Usando alfabeto aleatorio.")

		// Generar alfabeto aleatorio como en el caso 2
		letras := []rune(alfabetoNormal)
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(letras), func(i, j int) {
			letras[i], letras[j] = letras[j], letras[i]
		})

		alfabetoCifrado = string(letras)
	}

	// Mostrar los alfabetos
	fmt.Println("\nAlfabeto normal:  ", alfabetoNormal)
	fmt.Println("Alfabeto cifrado:", alfabetoCifrado)

	// Cifrar o descifrar
	opcion := readEntry("¿Quieres cifrar o descifrar? (C/D): ")

	var resultado strings.Builder

	if strings.ToUpper(opcion) == "C" {
		// CIFRADO
		for _, char := range mensaje {
			idx := strings.IndexRune(alfabetoNormal, char)
			if idx != -1 {
				// Sustituir con la letra correspondiente del alfabeto cifrado
				resultado.WriteRune(rune(alfabetoCifrado[idx]))
			} else {
				// Si no es una letra, mantenerla igual
				resultado.WriteRune(char)
			}
		}

		fmt.Printf("\n✅ Mensaje cifrado: %s\n", resultado.String())

	} else {
		// DESCIFRADO
		for _, char := range mensaje {
			idx := strings.IndexRune(alfabetoCifrado, char)
			if idx != -1 {
				// Sustituir con la letra correspondiente del alfabeto normal
				resultado.WriteRune(rune(alfabetoNormal[idx]))
			} else {
				// Si no es una letra, mantenerla igual
				resultado.WriteRune(char)
			}
		}

		fmt.Printf("\n✅ Mensaje descifrado: %s\n", resultado.String())
	}

	// Mostrar la tabla de sustitución completa
	fmt.Println("\nTabla de sustitución completa:")
	fmt.Println("Normal  | Cifrado")
	fmt.Println("----------------")
	for i := 0; i < len(alfabetoNormal); i++ {
		fmt.Printf("   %c    |    %c\n", alfabetoNormal[i], alfabetoCifrado[i])
	}
}
