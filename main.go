package main

import (
	"calculadora/funciones"
	"fmt"
	"os/exec"
	"runtime"

	"github.com/eiannone/keyboard"
)

func cleanmenu() {
	var err error

	if runtime.GOOS == "windows" {
		err = exec.Command("cmd", "/c", "cls").Run()
	} else {
		err = exec.Command("clear").Run()
	}
	if err != nil {
		fmt.Print("\033[H\033[2J")
	}
}

func menuPrincipal(selectoption int) {
	cleanmenu()

	menuprincipal := `****************
88888b.d88b.  .d88b. 88888b. 888  888 
888 "888 "88bd8P  Y8b888 "88b888  888 
888  888  88888888888888  888888  888 
888  888  888Y8b.    888  888Y88b 888 
888  888  888 "Y8888 888  888 "Y88888 
	
`
	fmt.Println(menuprincipal)
	options := []string{
		"🧮 Operaciones matemáticas modulares",
		"🔒 Criptográfica Clásica",
		"🔐 Criptografía Moderna",
		"#️⃣  Algoritmos Hash",
		"💻 Codificación",
		"🚪 Salir",
	}

	for i, option := range options {
		if i+1 == selectoption {
			fmt.Printf("> %v <\n", option)
		} else {
			fmt.Printf("  %v\n", option)
		}
	}

	fmt.Println("\n****************")
	fmt.Println("Usa ↑↓ para navegar y Enter para seleccionar")
	fmt.Println("Salir con 'q' we :D")
	fmt.Println("****************")
}

func menuSecondary(menuNumber int, selectoption int) {
	cleanmenu()
	menusecondary := `****************
.d8888b  888  888 888888b.  888b     d888 8888888888 888b    888 888  888 
88K      888  888 888  "88b 8888b   d8888 888        8888b   888 888  888 
"Y8888b. 888  888 888888K   8888b.d888888 8888888    88888b  888 888  888 
     X88 888  888 888  "88b 888Y88888P888 888        888Y88b 888 Y88b 888 
 88888P' 'Y888888 888888P'  888 Y888P 888 8888888888 888 Y88888  "Y88888
 
 `

	fmt.Println(menusecondary)

	var options []string

	switch menuNumber {
	case 1:
		options = []string{
			"1️⃣ Calcular el módulo de dos números a mod n = b",
			"2️⃣ Calcular inverso aditivo",
			"3️⃣ Calcular inverso de XOR",
			"4️⃣ Calcular máximo común divisor (MCD) e indicar si existe el inverso multiplicativo",
			"5️⃣ Calcular inverso multiplicativo por método tradicional visto en clase",
			"6️⃣ Calcular inverso multiplicativo aplicando el Algoritmo Extendido de Euclides AEE",
			"🔙 Volver al menú principal",
		}
	case 2:
		options = []string{
			"1️⃣ Cifrado Módulo 27",
			"2️⃣ Cifrado César",
			"3️⃣ Cifrado Venam",
			"4️⃣ Cifrado ATBASH",
			"5️⃣ Cifrado Transposición columnar simple",
			"6️⃣ Cifrado Afín",
			"7️⃣ Cifrado de Sustitución",
			"🔙 Volver al menú principal",
		}
	case 3:
		options = []string{
			"1️⃣ Calcular Diffie Hellman",
			"2️⃣ Calcular RSA",
			"3️⃣ Calcular Algoritmo de exponenciación rápida",
			"🔙 Volver al menú principal",
		}
	case 4:
		options = []string{
			"1️⃣ Calcular md5",
			"2️⃣ Calcular SHA1",
			"3️⃣ Calcular SHA512",
			"🔙 Volver al menú principal",
		}
	case 5:
		options = []string{
			"1️⃣ Codificar/Decodificar Binario",
			"2️⃣ Codificar/Decodificar Hexa",
			"3️⃣ Codificar/Decodificar Base64",
			"🔙 Volver al menú principal",
		}
	}

	for i, option := range options {
		if i+1 == selectoption {
			fmt.Printf("> %v <\n", option)
		} else {
			fmt.Printf("  %v\n", option)
		}
	}

	fmt.Println("\n****************")
	fmt.Println("Usa ↑↓ para navegar y Enter para seleccionar")
	fmt.Println("Salir con 'q' we :D")
	fmt.Println("****************")
}

func handleSubmenu(menuNumber int) {
	selectedOption := 0
	maxOption := 0

	switch menuNumber {
	case 1:
		maxOption = 6
	case 2:
		maxOption = 7
	case 3, 4, 5:
		maxOption = 3
	}

	maxOption++

	for {
		// Sumamos 1 a selectedOption para corregir el índice en menuSecondary
		menuSecondary(menuNumber, selectedOption+1)

		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			fmt.Println("error en el teclado we :D", err)
			return
		}

		if key == keyboard.KeyArrowUp {
			selectedOption = (selectedOption - 1 + maxOption) % maxOption
		} else if key == keyboard.KeyArrowDown {
			selectedOption = (selectedOption + 1) % maxOption
		} else if key == keyboard.KeyEnter {
			if selectedOption == maxOption-1 {
				return
			}

			cleanmenu()

			switch menuNumber {
			case 1:
				switch selectedOption {
				case 0:
					funciones.CalcularModulo()
				case 1:
					funciones.CalcularInversoAditivo()
				case 2:
					funciones.CalcularInversoXOR()
				case 3:
					funciones.CalcularMCD()
				case 4:
					funciones.CalcularInversoMultiplicativoTradicional()
				case 5:
					funciones.CalcularInversoMultiplicativoAEE()
				}
			case 2:
				switch selectedOption {
				case 0:
					funciones.CifradoModulo27()
				case 1:
					funciones.CifradoCesar()
				case 2:
					funciones.CifradoVernam()
				case 3:
					funciones.CifradoAtbash()
				case 4:
					funciones.CifradoTransposicionColumnar()
				case 5:
					funciones.CifradoAfin()
				case 6:
					funciones.CifradoSustitucion()
				}
			case 3:
				switch selectedOption {
				case 0:
					funciones.CalcularDiffieHellman()
				case 1:
					funciones.CalcularRSA()
				case 2:
					funciones.CalcularExponenciacionRapida()
				}
			case 4:
				switch selectedOption {
				case 0:
					funciones.CalcularMD5()
				case 1:
					funciones.CalcularSHA1()
				case 2:
					funciones.CalcularSHA512()
				}
			case 5:
				switch selectedOption {
				case 0:
					funciones.CodificarDecodificarBinario()
				case 1:
					funciones.CodificarDecodificarHexa()
				case 2:
					funciones.CodificarDecodificarBase64()
				}
			}

			fmt.Println("\nPresiona cualquier tecla para continuar...")
			keyboard.GetSingleKey()
		} else if char == 'q' || char == 'Q' {
			cleanmenu()
			fmt.Println("Adios we :D")
			return
		}
	}
}

func menus() {

	if err := keyboard.Open(); err != nil {
		fmt.Println("error en el teclado we :D", err)
		return
	}
	defer keyboard.Close()

	selectedOption := 1
	maxOption := 6

	for {
		menuPrincipal(selectedOption)

		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			fmt.Println("error en el teclado we :D", err)
			return
		}

		if key == keyboard.KeyArrowUp {
			selectedOption = (selectedOption - 1)
			if selectedOption < 1 {
				selectedOption = maxOption
			}
		} else if key == keyboard.KeyArrowDown {
			selectedOption = (selectedOption + 1)
			if selectedOption > maxOption {
				selectedOption = 1
			}
		} else if key == keyboard.KeyEnter {
			if selectedOption == 6 {
				cleanmenu()
				fmt.Println("Adios we :D")
				return
			}
			handleSubmenu(selectedOption)
		} else if char == 'q' || char == 'Q' {
			cleanmenu()
			fmt.Println("Adios we :D")
			return
		}
	}

}

func main() {
	menus()
}
