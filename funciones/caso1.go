package funciones

import (
	"fmt"
	"strings"
)

func calcularMCD(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}

	return a
}

func algoritmoExtendidoEuclides(a, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}

	mcd, x1, y1 := algoritmoExtendidoEuclides(b%a, a)

	x := y1 - (b/a)*x1
	y := x1

	return mcd, x, y
}

func CalcularModulo() {
	fmt.Println("\n🧮 Calcular el módulo de dos números (a mod n = b)")
	fmt.Println("=============================================")
	a := readNumber("Ingresa numero a: ")
	n := readNumber("Ingresa numero n: ")

	if n == 0 {
		fmt.Println("❌ E we no se puede hacer mod 0 :C")
		return
	}

	result := ((a % n) + n) % n
	fmt.Printf("\n✅ Resultado: %d mod %d = %d\n", a, n, result)
}

func CalcularInversoAditivo() {
	fmt.Println("\n🧮 Calcular inverso aditivo")
	fmt.Println("=============================================")

	a := readNumber("Ingresa numero a: ")
	n := readNumber("Ingresa numero n: ")

	if n == 0 {
		fmt.Println("❌ E we no se puede hacer mod 0 :C")
		return
	}

	a = (((a % n) + n) % n)
	inverso := (n - a) % n
	fmt.Printf("\n✅ El inverso aditivo de %d en módulo %d es: %d\n", a, n, inverso)
	fmt.Printf("   Comprobación: (%d + %d) mod %d = %d\n", a, inverso, n, (a+inverso)%n)
}

func CalcularInversoXOR() {
	fmt.Println("\n🧮 Calcular inverso de XOR")
	fmt.Println("=============================================")

	a := readNumber("Ingresa numero a: ")

	fmt.Printf("\n✅ El inverso XOR de %d es: %d\n", a, a)
	fmt.Printf("   Comprobación: %d XOR %d = %d\n", a, a, a^a)

	fmt.Printf("\nRepresentación binaria:\n")
	fmt.Printf("   %d en binario: %b\n", a, a)
	fmt.Printf("   %d XOR %d = %b\n", a, a, a^a)
}

func CalcularMCD() {
	fmt.Println("\n🧮 Calcular MCD e indicar si existe inverso multiplicativo")
	fmt.Println("=============================================")

	a := readNumber("Ingresa el primer número (a): ")
	n := readNumber("Ingresa el segundo número (n): ")

	mcd := calcularMCD(a, n)

	fmt.Printf("\n✅ El MCD de %d y %d es: %d\n", a, n, mcd)

	if mcd == 1 {
		fmt.Printf("   Los números %d y %d son coprimos.\n", a, n)
		fmt.Printf("   ✓ EXISTE un inverso multiplicativo de %d en módulo %d.\n", a, n)
	} else {
		fmt.Printf("   Los números %d y %d NO son coprimos.\n", a, n)
		fmt.Printf("   ✗ NO existe un inverso multiplicativo de %d en módulo %d.\n", a, n)
	}
}

func CalcularInversoMultiplicativoTradicional() {
	fmt.Println("\n🧮 Calcular inverso multiplicativo (método tradicional)")
	fmt.Println("=============================================")

	a := readNumber("Ingresa el número a: ")
	n := readNumber("Ingresa el módulo n: ")

	if n <= 0 {
		fmt.Println("❌ Error: El módulo debe ser un número positivo.")
		return
	}

	a = ((a % n) + n) % n

	if calcularMCD(a, n) != 1 {
		fmt.Printf("❌ No existe inverso multiplicativo para %d en módulo %d porque no son coprimos.\n", a, n)
		return
	}

	fmt.Println("\nProceso de cálculo:")

	for x := 1; x < n; x++ {
		producto := (a * x) % n

		if producto == 1 {
			fmt.Printf("%d * %d ≡ %d (mod %d) ✓\n", a, x, producto, n)
			fmt.Printf("\n✅ El inverso multiplicativo de %d en módulo %d es: %d\n", a, n, x)
			return
		} else {
			fmt.Printf("%d * %d ≡ %d (mod %d)\n", a, x, producto, n)
		}
	}

	fmt.Println("❌ Error en el cálculo. No se encontró inverso multiplicativo.")
}

func CalcularInversoMultiplicativoAEE() {
	fmt.Println("\n🧮 Calcular inverso multiplicativo con Algoritmo Extendido de Euclides")
	fmt.Println("=====================================================")

	a := readNumber("Ingresa el número a: ")
	n := readNumber("Ingresa el módulo n: ")

	if n <= 0 {
		fmt.Println("❌ Error: El módulo debe ser un número positivo.")
		return
	}

	a = ((a % n) + n) % n

	mcd, x, _ := algoritmoExtendidoEuclides(a, n)

	if mcd != 1 {
		fmt.Printf("❌ No existe inverso multiplicativo para %d en módulo %d porque no son coprimos.\n", a, n)
		return
	}

	inverso := ((x % n) + n) % n

	fmt.Println("\nTabla de cálculo del Algoritmo Extendido de Euclides:")
	fmt.Println("Ronda\tq\ta\tb\tx\ty")
	fmt.Println("==============================")

	r1, r2 := n, a
	x1, x2 := 0, 1
	y1, y2 := 1, 0
	ronda := 0

	fmt.Printf("%d\t-\t%d\t%d\t%d\t%d\n", ronda, r1, r2, x1, y1)

	for r2 != 0 {
		ronda++
		q := r1 / r2
		r := r1 - q*r2
		x := x1 - q*x2
		y := y1 - q*y2

		fmt.Printf("%d\t%d\t%d\t%d\t%d\t%d\n", ronda, q, r1, r2, x1, y1)

		r1, r2 = r2, r
		x1, x2 = x2, x
		y1, y2 = y2, y
	}

	fmt.Printf("\n✅ El inverso multiplicativo de %d en módulo %d es: %d\n", a, n, inverso)
	fmt.Printf("   Comprobación: %d * %d ≡ %d (mod %d)\n", a, inverso, (a*inverso)%n, n)
	fmt.Printf("   Número de rondas: %d\n", ronda)
}

func CifradoATBASH() {
	fmt.Println("\n🔒 Cifrado ATBASH")
	fmt.Println("=============================================")

	alfabetoNormal := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alfabetoInverso := "ZYXWVUTSRQPONMLKJIHGFEDCBA"

	mensaje := readEntry("Ingresa el mensaje a cifrar/descifrar: ")

	mensaje = strings.ToUpper(mensaje)

	var resultado strings.Builder

	for _, char := range mensaje {
		pos := strings.IndexRune(alfabetoNormal, char)
		if pos != -1 {
			resultado.WriteByte(alfabetoInverso[pos])
		} else {
			resultado.WriteRune(char)
		}
	}

	fmt.Printf("\n✅ Mensaje cifrado/descifrado: %s\n", resultado.String())

	fmt.Println("\nTabla de sustitución ATBASH:")
	fmt.Println("Normal:   " + alfabetoNormal)
	fmt.Println("Inverso:  " + alfabetoInverso)

	fmt.Println("\nNota: En ATBASH, el cifrado y el descifrado son la misma operación.")
}
