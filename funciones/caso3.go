package funciones

import (
	"fmt"
	"math/big"
)

func CalcularDiffieHellman() {
	fmt.Println("\n🔐 Protocolo de Intercambio de Claves Diffie-Hellman")
	fmt.Println("=============================================")

	// Paso 1: Parámetros públicos
	fmt.Println("\n1️⃣ Parámetros públicos:")
	p := readNumberBig("Ingresa el número primo p (módulo): ")
	g := readNumberBig("Ingresa el generador g: ")

	// Verificar que p sea primo (para proyectos reales, usar pruebas de primalidad más robustas)
	if !p.ProbablyPrime(20) {
		fmt.Println("⚠️ Advertencia: El número p ingresado probablemente no es primo.")
		fmt.Println("Para seguridad real, p debe ser un número primo grande.")
	}

	// Paso 2: Claves privadas de Alice y Bob
	fmt.Println("\n2️⃣ Claves privadas:")
	fmt.Println("🧑 Alice:")
	a := readNumberBig("Ingresa la clave privada de Alice (a): ")

	fmt.Println("\n👤 Bob:")
	b := readNumberBig("Ingresa la clave privada de Bob (b): ")

	// Paso 3: Cálculos públicos
	fmt.Println("\n3️⃣ Cálculos públicos (intercambiados en canal inseguro):")

	// Alice calcula A = g^a mod p
	A := new(big.Int).Exp(g, a, p)
	fmt.Printf("🧑 Alice envía A = g^a mod p = %d^%d mod %d = %d\n", g, a, p, A)

	// Bob calcula B = g^b mod p
	B := new(big.Int).Exp(g, b, p)
	fmt.Printf("👤 Bob envía B = g^b mod p = %d^%d mod %d = %d\n", g, b, p, B)

	// Paso 4: Cálculo de la clave secreta compartida
	fmt.Println("\n4️⃣ Cálculo de la clave secreta compartida:")

	// Alice calcula s = B^a mod p
	s_alice := new(big.Int).Exp(B, a, p)
	fmt.Printf("🧑 Alice calcula s = B^a mod p = %d^%d mod %d = %d\n", B, a, p, s_alice)

	// Bob calcula s = A^b mod p
	s_bob := new(big.Int).Exp(A, b, p)
	fmt.Printf("👤 Bob calcula s = A^b mod p = %d^%d mod %d = %d\n", A, b, p, s_bob)

	// Verificación
	if s_alice.Cmp(s_bob) == 0 {
		fmt.Println("\n✅ Éxito! Ambos llegaron a la misma clave secreta compartida!")
		fmt.Printf("🔑 Clave secreta compartida: %d\n", s_alice)
	} else {
		fmt.Println("\n❌ Error! Las claves calculadas no coinciden.")
	}

	// Explicación del protocolo
	fmt.Println("\n📚 Explicación del protocolo Diffie-Hellman:")
	fmt.Println("1. Alice y Bob acuerdan públicamente un número primo p y un generador g.")
	fmt.Println("2. Alice elige secretamente un número a, y Bob elige secretamente un número b.")
	fmt.Println("3. Alice calcula A = g^a mod p y lo envía a Bob.")
	fmt.Println("4. Bob calcula B = g^b mod p y lo envía a Alice.")
	fmt.Println("5. Alice calcula la clave secreta: s = B^a mod p.")
	fmt.Println("6. Bob calcula la clave secreta: s = A^b mod p.")
	fmt.Println("7. Ambos obtienen la misma clave secreta: s = g^(a*b) mod p.")
	fmt.Println("8. Un atacante que vea p, g, A y B no puede calcular fácilmente s debido a la")
	fmt.Println("   dificultad del problema del logaritmo discreto.")
}

// CalcularRSA implementa el algoritmo de cifrado RSA
func CalcularRSA() {
	fmt.Println("\n🔐 Algoritmo de Cifrado RSA")
	fmt.Println("=============================================")

	// Paso 1: Generación de claves
	fmt.Println("\n1️⃣ Generación de claves:")

	// Elegir dos números primos
	p := readNumberBig("Ingresa el primer número primo (p): ")
	q := readNumberBig("Ingresa el segundo número primo (q): ")

	// Verificar que p y q sean primos
	if !p.ProbablyPrime(20) || !q.ProbablyPrime(20) {
		fmt.Println("⚠️ Advertencia: Uno o ambos números ingresados probablemente no son primos.")
		fmt.Println("Para seguridad real, p y q deben ser números primos grandes y distintos.")
	}

	// Calcular n = p * q
	n := new(big.Int).Mul(p, q)
	fmt.Printf("\nCalculando n = p * q = %d * %d = %d\n", p, q, n)

	// Calcular φ(n) = (p-1) * (q-1)
	p_minus_1 := new(big.Int).Sub(p, big.NewInt(1))
	q_minus_1 := new(big.Int).Sub(q, big.NewInt(1))
	phi := new(big.Int).Mul(p_minus_1, q_minus_1)
	fmt.Printf("Calculando φ(n) = (p-1) * (q-1) = %d * %d = %d\n", p_minus_1, q_minus_1, phi)

	// Elegir e (exponente público)
	e := readNumberBig("Ingresa el exponente público e (coprimo con φ(n)): ")

	// Verificar que e sea coprimo con φ(n)
	if new(big.Int).GCD(nil, nil, e, phi).Cmp(big.NewInt(1)) != 0 {
		fmt.Printf("❌ Error: e (%d) debe ser coprimo con φ(n) (%d).\n", e, phi)
		return
	}

	// Calcular d (exponente privado): d * e ≡ 1 (mod φ(n))
	d := new(big.Int).ModInverse(e, phi)
	if d == nil {
		fmt.Printf("❌ Error: No se puede calcular el inverso multiplicativo de e (%d) módulo φ(n) (%d).\n", e, phi)
		return
	}
	fmt.Printf("Calculando d = e^(-1) mod φ(n) = %d^(-1) mod %d = %d\n", e, phi, d)

	// Mostrar claves
	fmt.Println("\n🔑 Claves generadas:")
	fmt.Printf("Clave pública (e, n): (%d, %d)\n", e, n)
	fmt.Printf("Clave privada (d, n): (%d, %d)\n", d, n)

	// Paso 2: Cifrado o Descifrado
	fmt.Println("\n2️⃣ ¿Qué deseas hacer?")
	fmt.Println("1. Cifrar un mensaje")
	fmt.Println("2. Descifrar un mensaje")
	opcion := readNumber("Elige una opción (1-2): ")

	if opcion == 1 {
		// Cifrar
		m := readNumberBig("\nIngresa el mensaje a cifrar (número menor que n): ")

		// Verificar que m sea menor que n
		if m.Cmp(n) >= 0 {
			fmt.Printf("❌ Error: El mensaje (%d) debe ser menor que n (%d).\n", m, n)
			return
		}

		// Cifrar: c = m^e mod n
		c := new(big.Int).Exp(m, e, n)

		fmt.Printf("\n✅ Proceso de cifrado:")
		fmt.Printf("\nMensaje original (m): %d", m)
		fmt.Printf("\nMensaje cifrado (c) = m^e mod n = %d^%d mod %d = %d\n", m, e, n, c)

	} else if opcion == 2 {
		// Descifrar
		c := readNumberBig("\nIngresa el mensaje cifrado (número): ")

		// Descifrar: m = c^d mod n
		m := new(big.Int).Exp(c, d, n)

		fmt.Printf("\n✅ Proceso de descifrado:")
		fmt.Printf("\nMensaje cifrado (c): %d", c)
		fmt.Printf("\nMensaje descifrado (m) = c^d mod n = %d^%d mod %d = %d\n", c, d, n, m)

	} else {
		fmt.Println("❌ Opción no válida.")
	}

	// Explicación del algoritmo RSA
	fmt.Println("\n📚 Explicación del algoritmo RSA:")
	fmt.Println("1. Generación de claves:")
	fmt.Println("   a. Elegir dos números primos distintos p y q.")
	fmt.Println("   b. Calcular n = p * q (módulo para cifrar/descifrar).")
	fmt.Println("   c. Calcular φ(n) = (p-1) * (q-1).")
	fmt.Println("   d. Elegir un entero e (1 < e < φ(n)) que sea coprimo con φ(n).")
	fmt.Println("   e. Calcular d tal que d * e ≡ 1 (mod φ(n)).")
	fmt.Println("   f. La clave pública es (e, n) y la clave privada es (d, n).")
	fmt.Println("2. Cifrado: c = m^e mod n (donde m es el mensaje original).")
	fmt.Println("3. Descifrado: m = c^d mod n (donde c es el mensaje cifrado).")
}

// CalcularExponenciacionRapida implementa el algoritmo de exponenciación modular rápida
func CalcularExponenciacionRapida() {
	fmt.Println("\n🔐 Algoritmo de Exponenciación Rápida (Exponenciación Modular)")
	fmt.Println("=============================================")

	base := readNumberBig("Ingresa la base (b): ")
	exponente := readNumberBig("Ingresa el exponente (e): ")
	modulo := readNumberBig("Ingresa el módulo (m): ")

	// Calcular b^e mod m usando el método incorporado en big.Int
	resultado := new(big.Int).Exp(base, exponente, modulo)

	fmt.Printf("\n✅ Resultado: %d^%d mod %d = %d\n", base, exponente, modulo, resultado)

	// Mostrar el proceso paso a paso (implementación explícita para fines educativos)
	fmt.Println("\n📚 Proceso paso a paso con el algoritmo de exponenciación rápida:")

	// Convertir el exponente a binario para la explicación
	fmt.Printf("1. Convertir el exponente %d a binario: %b\n", exponente, exponente)

	// Implementar el algoritmo de exponenciación rápida manualmente para mostrar los pasos
	res := big.NewInt(1)
	b := new(big.Int).Set(base)
	e := new(big.Int).Set(exponente)
	zero := big.NewInt(0)
	one := big.NewInt(1)
	two := big.NewInt(2)

	fmt.Println("2. Inicializar resultado = 1")
	fmt.Println("3. Proceso iterativo: ")
	paso := 1

	for e.Cmp(zero) > 0 {
		// Si el exponente es impar, multiplicar el resultado por la base
		if new(big.Int).And(e, one).Cmp(one) == 0 {
			old_res := new(big.Int).Set(res)
			res = new(big.Int).Mul(res, b)
			res = new(big.Int).Mod(res, modulo)
			fmt.Printf("   Paso %d: Exponente %d es impar, resultado = resultado * base mod m = %d * %d mod %d = %d\n",
				paso, e, old_res, b, modulo, res)
		} else {
			fmt.Printf("   Paso %d: Exponente %d es par, resultado no cambia = %d\n", paso, e, res)
		}

		// Elevar la base al cuadrado
		old_b := new(big.Int).Set(b)
		b = new(big.Int).Mul(b, b)
		b = new(big.Int).Mod(b, modulo)
		fmt.Printf("   Paso %d (cont.): Actualizar base = base^2 mod m = %d^2 mod %d = %d\n", paso, old_b, modulo, b)

		// Dividir el exponente por 2
		e = new(big.Int).Div(e, two)
		fmt.Printf("   Paso %d (cont.): Dividir exponente por 2: %d\n", paso, e)

		paso++
	}

	fmt.Printf("\n✅ Resultado final: %d\n", res)

	// Explicación del algoritmo
	fmt.Println("\n📝 Explicación del algoritmo de exponenciación rápida:")
	fmt.Println("Este algoritmo calcula b^e mod m de forma eficiente con complejidad O(log e).")
	fmt.Println("En lugar de realizar e multiplicaciones, utiliza la representación binaria del exponente.")
	fmt.Println("La idea clave es utilizar la propiedad: b^(2k) = (b^k)^2")
	fmt.Println("")
	fmt.Println("El algoritmo funciona así:")
	fmt.Println("1. Inicializar resultado = 1")
	fmt.Println("2. Convertir el exponente a binario")
	fmt.Println("3. Por cada bit del exponente (de derecha a izquierda):")
	fmt.Println("   a. Si el bit es 1, multiplicar el resultado por la base actual")
	fmt.Println("   b. Elevar la base al cuadrado")
	fmt.Println("")
	fmt.Println("Este enfoque es crucial en criptografía donde se trabaja con exponentes enormes.")
	fmt.Println("Por ejemplo, RSA requiere calcular eficientemente expresiones como m^e mod n,")
	fmt.Println("donde e y n pueden tener cientos o miles de dígitos.")
}
