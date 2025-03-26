// helpers.go
package funciones

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func readEntry(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	entry, _ := reader.ReadString('\n')
	return strings.TrimSpace(entry)
}

func readNumber(message string) int {
	for {
		entryStr := readEntry(message)
		num, err := strconv.Atoi(entryStr)
		if err == nil {
			return num
		}
		fmt.Println("❌ E we esto no es un numero valido :C")
	}
}

func readNumberBig(message string) *big.Int {
	for {
		entryStr := readEntry(message)
		num := new(big.Int)
		_, success := num.SetString(entryStr, 10)
		if success {
			return num
		}
		fmt.Println("❌ Error: Por favor ingresa un número entero válido.")
	}
}
