package funciones

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
)

// CalcularMD5 calcula el hash MD5 de un texto ingresado
func CalcularMD5() {
	fmt.Println("\n#️⃣  Cálculo de Hash MD5")
	fmt.Println("=============================================")

	entrada := readEntry("Ingresa el texto para calcular su hash MD5: ")

	// Crear un nuevo hash MD5
	hasher := md5.New()

	// Escribir los datos al hasher
	io.WriteString(hasher, entrada)

	// Calcular el hash y convertirlo a formato hexadecimal
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	fmt.Println("\n✅ Resultado:")
	fmt.Printf("Texto original: %s\n", entrada)
	fmt.Printf("Hash MD5: %s\n", hashString)
	fmt.Printf("Longitud del hash: %d bytes (%d bits)\n", len(hashBytes), len(hashBytes)*8)

	// Mostrar representación binaria
	fmt.Printf("\nRepresentación binaria del hash:\n")
	for i, b := range hashBytes {
		fmt.Printf("%08b", b)
		if (i+1)%4 == 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	fmt.Println("\n📝 Información sobre MD5:")
	fmt.Println("- MD5 (Message-Digest Algorithm 5) fue diseñado por Ronald Rivest en 1991.")
	fmt.Println("- Produce un hash de 128 bits (16 bytes), representado como 32 dígitos hexadecimales.")
	fmt.Println("- Se considera criptográficamente roto y NO debe usarse para propósitos de seguridad.")
	fmt.Println("- Vulnerabilidades conocidas: colisiones, ataques de preimagen y extensión de longitud.")
	fmt.Println("- Uso actual: principalmente para verificación de integridad básica y no crítica.")
}

// CalcularSHA1 calcula el hash SHA-1 de un texto ingresado
func CalcularSHA1() {
	fmt.Println("\n#️⃣  Cálculo de Hash SHA-1")
	fmt.Println("=============================================")

	entrada := readEntry("Ingresa el texto para calcular su hash SHA-1: ")

	// Crear un nuevo hash SHA-1
	hasher := sha1.New()

	// Escribir los datos al hasher
	io.WriteString(hasher, entrada)

	// Calcular el hash y convertirlo a formato hexadecimal
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	fmt.Println("\n✅ Resultado:")
	fmt.Printf("Texto original: %s\n", entrada)
	fmt.Printf("Hash SHA-1: %s\n", hashString)
	fmt.Printf("Longitud del hash: %d bytes (%d bits)\n", len(hashBytes), len(hashBytes)*8)

	// Mostrar representación binaria
	fmt.Printf("\nRepresentación binaria del hash:\n")
	for i, b := range hashBytes {
		fmt.Printf("%08b", b)
		if (i+1)%4 == 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	fmt.Println("\n📝 Información sobre SHA-1:")
	fmt.Println("- SHA-1 (Secure Hash Algorithm 1) fue diseñado por la NSA y publicado en 1995.")
	fmt.Println("- Produce un hash de 160 bits (20 bytes), representado como 40 dígitos hexadecimales.")
	fmt.Println("- Se considera criptográficamente débil desde 2005 y roto desde 2017.")
	fmt.Println("- En 2017, Google demostró la primera colisión práctica de SHA-1.")
	fmt.Println("- No debe usarse para firmas digitales, certificados o propósitos de seguridad.")
	fmt.Println("- Ha sido reemplazado por algoritmos de la familia SHA-2 y SHA-3.")
}

// CalcularSHA512 calcula el hash SHA-512 de un texto ingresado
func CalcularSHA512() {
	fmt.Println("\n#️⃣  Cálculo de Hash SHA-512")
	fmt.Println("=============================================")

	entrada := readEntry("Ingresa el texto para calcular su hash SHA-512: ")

	// Crear un nuevo hash SHA-512
	hasher := sha512.New()

	// Escribir los datos al hasher
	io.WriteString(hasher, entrada)

	// Calcular el hash y convertirlo a formato hexadecimal
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	fmt.Println("\n✅ Resultado:")
	fmt.Printf("Texto original: %s\n", entrada)
	fmt.Printf("Hash SHA-512: %s\n", hashString)
	fmt.Printf("Longitud del hash: %d bytes (%d bits)\n", len(hashBytes), len(hashBytes)*8)

	// Mostrar primeros y últimos bytes en binario
	fmt.Printf("\nPrimeros bytes en binario:\n")
	for i := 0; i < 8; i++ {
		fmt.Printf("%08b ", hashBytes[i])
	}
	fmt.Println("...")
	fmt.Printf("Últimos bytes en binario:\n")
	for i := len(hashBytes) - 8; i < len(hashBytes); i++ {
		fmt.Printf("%08b ", hashBytes[i])
	}
	fmt.Println()

	fmt.Println("\n📝 Información sobre SHA-512:")
	fmt.Println("- SHA-512 es parte de la familia SHA-2, diseñada por la NSA y publicada en 2001.")
	fmt.Println("- Produce un hash de 512 bits (64 bytes), representado como 128 dígitos hexadecimales.")
	fmt.Println("- Es considerablemente más seguro que SHA-1 y MD5.")
	fmt.Println("- Utiliza un tamaño de bloque interno de 1024 bits y operaciones de 64 bits.")
	fmt.Println("- Es ampliamente utilizado en aplicaciones de seguridad, certificados digitales y criptografía.")
	fmt.Println("- Otras variantes incluyen SHA-384, SHA-256 y SHA-224, con diferentes longitudes de salida.")
	fmt.Println("- Hasta la fecha, no se conocen ataques prácticos contra SHA-512.")
}
