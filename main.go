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

// Configuración de la partida según el nivel seleccionado
type GameConfig struct {
	Name        string
	MaxNumber   int
	MaxAttempts int
}

func main() {
	// Inicializar la semilla para los números aleatorios
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	// Bucle para permitir múltiples partidas consecutivas
	for {
		printBanner()

		// 1. Configurar dificultad
		config := selectDifficulty(reader)

		// 2. Ejecutar la partida
		playRound(reader, config)

		// 3. Consultar si desea reiniciar
		if !askPlayAgain(reader) {
			fmt.Println("\n¡Gracias por jugar! Hasta la próxima.")
			break
		}
	}
}

// Banner inicial de presentación
func printBanner() {
	fmt.Println("========================================")
	fmt.Println("   🎮 ¡BIENVENIDO A GUESS THE NUMBER!   ")
	fmt.Println("========================================")

}

// Menú interactivo de selección de dificultad
func selectDifficulty(reader *bufio.Reader) GameConfig {
	for {
		fmt.Println("\nSelecciona un nivel de dificultad:")
		fmt.Println("1) Fácil   (Rango: 1 - 50  | Vidas: 10)")
		fmt.Println("2) Medio   (Rango: 1 - 100 | Vidas: 7)")
		fmt.Println("3) Difícil (Rango: 1 - 200 | Vidas: 5)")
		fmt.Print("Opción (1-3): ")

		input := readCleanInput(reader)
		switch input {
		case "1":
			return GameConfig{Name: "Fácil", MaxNumber: 50, MaxAttempts: 10}
		case "2":
			return GameConfig{Name: "Medio", MaxNumber: 100, MaxAttempts: 7}
		case "3":
			return GameConfig{Name: "Difícil", MaxNumber: 200, MaxAttempts: 5}
		default:
			fmt.Println("Opción no válida. Ingresa 1, 2 o 3.")
		}
	}
}

// Lógica principal de una ronda de juego
func playRound(reader *bufio.Reader, config GameConfig) {
	target := rand.Intn(config.MaxNumber) + 1
	attemptsLeft := config.MaxAttempts

	fmt.Printf("\n[Modo %s] He pensado un número entre 1 y %d.\n", config.Name, config.MaxNumber)
	fmt.Printf("Tienes %d intentos para adivinarlo.\n\n", attemptsLeft)

	for attemptsLeft > 0 {
		fmt.Printf("[Vidas: %d/%d] Tu intento: ", attemptsLeft, config.MaxAttempts)

		input := readCleanInput(reader)
		guess, err := strconv.Atoi(input)

		// Validación: que sea número entero
		if err != nil {
			fmt.Println("Error: Ingresa un número entero válido.")
			continue
		}

		// Validación: dentro de los límites de la dificultad
		if guess < 1 || guess > config.MaxNumber {
			fmt.Printf("Error: El número debe estar entre 1 y %d.\n", config.MaxNumber)
			continue
		}

		// Victoria
		if guess == target {
			score := attemptsLeft * 100
			fmt.Println("\n¡Felicidades! Has adivinado el número secreto.")
			fmt.Printf("Puntuación obtenida: %d puntos (Vidas restantes: %d)\n", score, attemptsLeft)
			return
		}

		attemptsLeft--

		// Pistas de aproximación
		if guess < target {
			fmt.Println("El número secreto es MAYOR.")
		} else {
			fmt.Println("El número secreto es MENOR.")
		}

		if attemptsLeft > 0 {
			fmt.Println("----------------------------------------")
		}
	}

	// Derrota
	fmt.Printf("\nGame Over. Te quedaste sin intentos. El número era: %d\n", target)
}

// Pregunta si se inicia una nueva partida
func askPlayAgain(reader *bufio.Reader) bool {
	for {
		fmt.Print("\n¿Quieres jugar otra ronda? (s/n): ")
		ans := strings.ToLower(readCleanInput(reader))
		if ans == "s" || ans == "si" || ans == "sí" {
			return true
		}
		if ans == "n" || ans == "no" {
			return false
		}
		fmt.Println("Por favor, ingresa 's' para sí o 'n' para no.")
	}
}

// Limpia caracteres de escape y saltos de línea de la consola
func readCleanInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
