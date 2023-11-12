package UI

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

//Clean de Linux
func Clean() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

// Esperar : esperar los segundos indicados
func Esperar(s int) {
	switch s {
	case 1:
		time.Sleep(1 * time.Second)
	case 2:
		time.Sleep(2 * time.Second)
	case 3:
		time.Sleep(3 * time.Second)
	case 4:
		time.Sleep(4 * time.Second)
	case 5:
		time.Sleep(5 * time.Second)
	default:
		time.Sleep(0 * time.Second)
	}
}

func maxLenString(s []string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		if len(s[i]) > max {
			max = len(s[i])
		}
	}
	return max
}

func separador(symbol string, cadenas []string) {
	fmt.Println()
	fmt.Print("  ")
	max := maxLenString(cadenas)
	for i := 0; i < max; i++ {
		fmt.Print(symbol)
	}
	fmt.Println("\n")
}

// BloqueDeTexto : imprime las cadenas ingresadas
func BloqueDeTexto(cadenas ...string) {
	separador(cadenas[0], cadenas)
	for i := 1; i < len(cadenas); i++ {
		fmt.Println("  " + cadenas[i])
	}
	separador(cadenas[0], cadenas)
}
