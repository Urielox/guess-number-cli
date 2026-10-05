package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 1. Inicialización del juego
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	maxAttempts := 7
	attemptsLeft := maxAttempts
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("========================================")
	fmt.Println("   🎮 ¡BIENVENIDO A GUESS THE NUMBER!   ")
	fmt.Println("========================================")
	fmt.Printf("He pensado un número entre 1 y 100.\nTienes %d vidas para acertar.\n\n", maxAttempts)

	// 2. Bucle principal del juego (Game Loop)
	for attemptsLeft > 0 {
		fmt.Printf("[Vidas: %d/%d] Introduce tu intento: ", attemptsLeft, maxAttempts)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error al leer entrada. Intenta de nuevo.")
			continue
		}

		// Limpiar espacios y saltos de línea
		cleanedInput := strings.TrimSpace(input)
		guess, err := strconv.Atoi(cleanedInput)
		if err != nil {
			fmt.Println("⚠️  Por favor, escribe un número entero válido.")
			continue
		}

		// 3. Lógica de comparación
		if guess == target {
			fmt.Println("\n🎉 ¡EXCELENTE! Has adivinado el número secreto.")
			fmt.Printf("Te quedaron %d vidas de reserva.\n", attemptsLeft)
			return
		}

		attemptsLeft--

		if guess < target {
			fmt.Println("🔼 El número secreto es MAYOR.")
		} else {
			fmt.Println("🔽 El número secreto es MENOR.")
		}

		if attemptsLeft > 0 {
			fmt.Println("----------------------------------------")
		}
	}

	// 4. Fin de la partida si se acaban las vidas
	fmt.Printf("\n💀 Game Over. El número secreto era: %d\n", target)
}