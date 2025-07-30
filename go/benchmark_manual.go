package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"time"
)

// runBenchmark executa múltiplas medições e calcula estatísticas
func runBenchmark(operation func(), iterations int, name string) map[string]float64 {
	var times []float64
	
	fmt.Printf("Executando benchmark: %s\n", name)
	
	for i := 0; i < iterations; i++ {
		start := time.Now()
		operation()
		duration := time.Since(start)
		times = append(times, float64(duration.Nanoseconds()))
		
		if (i+1)%10 == 0 {
			fmt.Printf("  Iteração %d/%d\n", i+1, iterations)
		}
	}
	
	// Calcular estatísticas
	var total float64
	for _, t := range times {
		total += t
	}
	mean := total / float64(len(times))
	
	var variance float64
	for _, t := range times {
		variance += (t - mean) * (t - mean)
	}
	variance /= float64(len(times))
	stdDev := math.Sqrt(variance)
	
	min := times[0]
	max := times[0]
	for _, t := range times {
		if t < min {
			min = t
		}
		if t > max {
			max = t
		}
	}
	
	stats := make(map[string]float64)
	stats["mean_ns"] = mean
	stats["std_dev_ns"] = stdDev
	stats["min_ns"] = min
	stats["max_ns"] = max
	stats["total_ns"] = total
	
	return stats
}

// printStats imprime estatísticas
func printStats(stats map[string]float64, operation string) {
	fmt.Printf("\n=== Estatísticas para %s ===\n", operation)
	fmt.Printf("Média: %.2f ns\n", stats["mean_ns"])
	fmt.Printf("Desvio Padrão: %.2f ns\n", stats["std_dev_ns"])
	fmt.Printf("Mínimo: %.2f ns\n", stats["min_ns"])
	fmt.Printf("Máximo: %.2f ns\n", stats["max_ns"])
	fmt.Printf("Total: %.2f ns\n", stats["total_ns"])
}

// benchmarkManualRSA executa benchmark da implementação manual
func benchmarkManualRSA() {
	fmt.Println("=== Benchmark Manual RSA ===")
	fmt.Println("Usando funções nativas de tempo\n")
	
	iterations := 100
	
	// Benchmark de geração de chaves
	keyGenStats := runBenchmark(
		func() {
			_ = NewRSAKey(2048)
		},
		iterations,
		"Geração de Chaves",
	)
	printStats(keyGenStats, "Geração de Chaves")
	
	// Gerar uma chave para os testes de criptografia
	key := NewRSAKey(2048)
	testMessage := big.NewInt(12345)
	
	// Benchmark de criptografia
	encryptionStats := runBenchmark(
		func() {
			_ = Encrypt(testMessage, key)
		},
		iterations,
		"Criptografia",
	)
	printStats(encryptionStats, "Criptografia")
	
	// Criptografar uma mensagem para o teste de descriptografia
	ciphertext := Encrypt(testMessage, key)
	
	// Benchmark de descriptografia
	decryptionStats := runBenchmark(
		func() {
			_ = Decrypt(ciphertext, key)
		},
		iterations,
		"Descriptografia",
	)
	printStats(decryptionStats, "Descriptografia")
	
	// Comparação de performance
	fmt.Println("\n=== Comparação de Performance ===")
	fmt.Printf("Geração de Chaves: %.2f ns (média)\n", keyGenStats["mean_ns"])
	fmt.Printf("Criptografia: %.2f ns (média)\n", encryptionStats["mean_ns"])
	fmt.Printf("Descriptografia: %.2f ns (média)\n", decryptionStats["mean_ns"])
	
	// Verificar se a descriptografia funciona corretamente
	decrypted := Decrypt(ciphertext, key)
	if testMessage.Cmp(decrypted) == 0 {
		fmt.Println("✓ Verificação de integridade: OK")
	} else {
		fmt.Println("✗ Verificação de integridade: FALHOU")
	}
}

// benchmarkKeySizes executa benchmark de diferentes tamanhos de chave
func benchmarkKeySizes() {
	fmt.Println("\n=== Benchmark de Diferentes Tamanhos de Chave ===")
	
	keySizes := []int{512, 1024, 2048}
	iterations := 50
	
	for _, bits := range keySizes {
		fmt.Printf("\n--- Chave de %d bits ---\n", bits)
		
		stats := runBenchmark(
			func() {
				_ = NewRSAKey(bits)
			},
			iterations,
			fmt.Sprintf("Geração de Chave %d bits", bits),
		)
		
		printStats(stats, fmt.Sprintf("Geração de Chave %d bits", bits))
	}
}

// benchmarkMessageSizes executa benchmark de diferentes tamanhos de mensagem
func benchmarkMessageSizes() {
	fmt.Println("\n=== Benchmark de Diferentes Tamanhos de Mensagem ===")
	
	key := NewRSAKey(2048)
	messageSizes := []int{100, 1000, 10000, 100000}
	iterations := 50
	
	for _, size := range messageSizes {
		fmt.Printf("\n--- Mensagem de %d bytes ---\n", size)
		
		testMessage := big.NewInt(int64(size))
		
		encryptionStats := runBenchmark(
			func() {
				_ = Encrypt(testMessage, key)
			},
			iterations,
			fmt.Sprintf("Criptografia %d bytes", size),
		)
		
		printStats(encryptionStats, fmt.Sprintf("Criptografia %d bytes", size))
	}
}

func main() {
	benchmarkManualRSA()
	benchmarkKeySizes()
	benchmarkMessageSizes()
} 