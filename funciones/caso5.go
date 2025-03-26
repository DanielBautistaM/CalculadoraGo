package funciones

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

// CodificarDecodificarBinario permite convertir entre texto y representación binaria
func CodificarDecodificarBinario() {
	fmt.Println("\n💻 Codificación/Decodificación Binaria")
	fmt.Println("=============================================")

	fmt.Println("1. Codificar (texto -> binario)")
	fmt.Println("2. Decodificar (binario -> texto)")

	opcion := readNumber("Selecciona una opción (1-2): ")

	if opcion == 1 {
		// Codificar: texto -> binario
		texto := readEntry("Ingresa el texto a codificar en binario: ")

		var resultado strings.Builder

		fmt.Println("\nProceso de codificación:")
		fmt.Println("Carácter | ASCII/Unicode | Binario")
		fmt.Println("------------------------------------")

		for _, char := range texto {
			binario := fmt.Sprintf("%08b", char)
			fmt.Printf("   %c    |      %d       | %s\n", char, char, binario)

			resultado.WriteString(binario)
			resultado.WriteString(" ")
		}

		fmt.Println("\n✅ Resultado:")
		fmt.Printf("Texto original: %s\n", texto)
		fmt.Printf("Representación binaria: %s\n", resultado.String())

	} else if opcion == 2 {
		// Decodificar: binario -> texto
		binario := readEntry("Ingresa el código binario a decodificar (separado por espacios): ")

		// Dividir la entrada por espacios
		partes := strings.Split(binario, " ")

		var resultado strings.Builder
		var errores []string

		fmt.Println("\nProceso de decodificación:")
		fmt.Println("Binario | Decimal | Carácter")
		fmt.Println("---------------------------")

		for _, parte := range partes {
			// Ignorar entradas vacías
			if parte == "" {
				continue
			}

			// Convertir de binario a entero
			valor, err := strconv.ParseInt(parte, 2, 64)
			if err != nil {
				errores = append(errores, parte)
				continue
			}

			fmt.Printf(" %s  |   %d    |    %c\n", parte, valor, valor)
			resultado.WriteRune(rune(valor))
		}

		fmt.Println("\n✅ Resultado:")
		if len(errores) > 0 {
			fmt.Printf("⚠️ Advertencia: Algunos códigos binarios no pudieron ser decodificados: %s\n", strings.Join(errores, ", "))
		}
		fmt.Printf("Texto decodificado: %s\n", resultado.String())

	} else {
		fmt.Println("❌ Opción no válida.")
	}

	fmt.Println("\n📝 Información sobre codificación binaria:")
	fmt.Println("- Cada carácter se representa mediante su valor ASCII/Unicode")
	fmt.Println("- En la codificación básica, cada carácter se convierte a 8 bits (1 byte)")
	fmt.Println("- Caracteres extendidos (como emojis) pueden requerir más bytes")
	fmt.Println("- La codificación binaria es la base de todas las representaciones digitales")
}

// CodificarDecodificarHexa permite convertir entre texto y representación hexadecimal
func CodificarDecodificarHexa() {
	fmt.Println("\n💻 Codificación/Decodificación Hexadecimal")
	fmt.Println("=============================================")

	fmt.Println("1. Codificar (texto -> hexadecimal)")
	fmt.Println("2. Decodificar (hexadecimal -> texto)")

	opcion := readNumber("Selecciona una opción (1-2): ")

	if opcion == 1 {
		// Codificar: texto -> hexadecimal
		texto := readEntry("Ingresa el texto a codificar en hexadecimal: ")

		// Codificar a hexadecimal
		hexadecimal := hex.EncodeToString([]byte(texto))

		fmt.Println("\nProceso de codificación:")
		fmt.Println("Carácter | ASCII/Unicode | Hexadecimal")
		fmt.Println("--------------------------------------")

		for i, char := range texto {
			fmt.Printf("   %c    |      %d       |     %s\n", char, char, hexadecimal[i*2:(i*2)+2])
		}

		fmt.Println("\n✅ Resultado:")
		fmt.Printf("Texto original: %s\n", texto)
		fmt.Printf("Representación hexadecimal: %s\n", hexadecimal)

	} else if opcion == 2 {
		// Decodificar: hexadecimal -> texto
		hexadecimal := readEntry("Ingresa el código hexadecimal a decodificar (sin espacios): ")

		// Eliminar posibles espacios y preparar entrada
		hexadecimal = strings.ReplaceAll(hexadecimal, " ", "")

		// Decodificar hexadecimal
		bytes, err := hex.DecodeString(hexadecimal)
		if err != nil {
			fmt.Println("❌ Error: El código hexadecimal proporcionado no es válido.")
			return
		}

		texto := string(bytes)

		fmt.Println("\nProceso de decodificación:")
		fmt.Println("Hexadecimal | Decimal | Carácter")
		fmt.Println("-------------------------------")

		for i := 0; i < len(hexadecimal); i += 2 {
			if i+1 < len(hexadecimal) {
				byteStr := hexadecimal[i : i+2]
				byteVal, _ := strconv.ParseInt(byteStr, 16, 64)

				if i/2 < len(texto) {
					fmt.Printf("    %s     |   %d    |    %c\n", byteStr, byteVal, texto[i/2])
				}
			}
		}

		fmt.Println("\n✅ Resultado:")
		fmt.Printf("Texto decodificado: %s\n", texto)

	} else {
		fmt.Println("❌ Opción no válida.")
	}

	fmt.Println("\n📝 Información sobre codificación hexadecimal:")
	fmt.Println("- La representación hexadecimal usa 16 símbolos: 0-9 y A-F")
	fmt.Println("- Cada dígito hexadecimal representa exactamente 4 bits (medio byte)")
	fmt.Println("- Dos dígitos hexadecimales (00-FF) representan 1 byte completo (8 bits)")
	fmt.Println("- Es más compacta que la representación binaria y muy usada en informática")
	fmt.Println("- Común en direcciones de memoria, códigos de colores RGB, hashes, etc.")
}

// CodificarDecodificarBase64 permite convertir entre texto y codificación Base64
func CodificarDecodificarBase64() {
	fmt.Println("\n💻 Codificación/Decodificación Base64")
	fmt.Println("=============================================")

	fmt.Println("1. Codificar (texto -> Base64)")
	fmt.Println("2. Decodificar (Base64 -> texto)")

	opcion := readNumber("Selecciona una opción (1-2): ")

	if opcion == 1 {
		// Codificar: texto -> Base64
		texto := readEntry("Ingresa el texto a codificar en Base64: ")

		// Codificar a Base64
		base64String := base64.StdEncoding.EncodeToString([]byte(texto))

		fmt.Println("\n✅ Resultado:")
		fmt.Printf("Texto original: %s\n", texto)
		fmt.Printf("Codificación Base64: %s\n", base64String)

		// Mostrar información sobre la longitud
		fmt.Printf("\nLongitud original: %d bytes\n", len(texto))
		fmt.Printf("Longitud en Base64: %d caracteres\n", len(base64String))

	} else if opcion == 2 {
		// Decodificar: Base64 -> texto
		base64String := readEntry("Ingresa el código Base64 a decodificar: ")

		// Decodificar Base64
		bytes, err := base64.StdEncoding.DecodeString(base64String)
		if err != nil {
			fmt.Println("❌ Error: El código Base64 proporcionado no es válido.")
			return
		}

		texto := string(bytes)

		fmt.Println("\n✅ Resultado:")
		fmt.Printf("Texto decodificado: %s\n", texto)

		// Mostrar información sobre la longitud
		fmt.Printf("\nLongitud en Base64: %d caracteres\n", len(base64String))
		fmt.Printf("Longitud decodificada: %d bytes\n", len(texto))

	} else {
		fmt.Println("❌ Opción no válida.")
	}

	fmt.Println("\n📝 Información sobre codificación Base64:")
	fmt.Println("- Base64 utiliza 64 caracteres imprimibles: A-Z, a-z, 0-9, + y /")
	fmt.Println("- Se usa = como carácter de relleno (padding) al final si es necesario")
	fmt.Println("- Cada 3 bytes (24 bits) de datos se representan como 4 caracteres Base64")
	fmt.Println("- Incrementa el tamaño aproximadamente un 33% respecto al original")
	fmt.Println("- Es útil para transferir datos binarios a través de medios que solo soportan texto")
	fmt.Println("- Común en correos electrónicos, datos incrustados en HTML/CSS, API REST, etc.")
}
