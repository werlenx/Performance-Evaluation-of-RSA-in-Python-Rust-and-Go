package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run . <opção>")
		fmt.Println("Opções:")
		fmt.Println("  manual    - Executar implementação manual do RSA")
		fmt.Println("  lib       - Executar implementação com biblioteca")
		fmt.Println("  benchmark - Executar benchmarks manuais")
		fmt.Println("  test      - Executar benchmarks com testing")
		return
	}
	
	switch os.Args[1] {
	case "manual":
		fmt.Println("Executando implementação manual do RSA...")
		// Executar rsa_manual.go
		runManualRSA()
	case "lib":
		fmt.Println("Executando implementação com biblioteca...")
		// Executar rsa_lib.go
		runLibRSA()
	case "benchmark":
		fmt.Println("Executando benchmarks manuais...")
		// Executar benchmark_manual.go
		runBenchmarkManual()
	case "test":
		fmt.Println("Executando benchmarks com testing...")
		fmt.Println("Para benchmarks detalhados, execute: go test -bench=.")
	default:
		fmt.Println("Opção inválida. Use 'manual', 'lib', 'benchmark' ou 'test'")
	}
}

// runManualRSA executa a implementação manual
func runManualRSA() {
	fmt.Println("=== Implementação Manual do RSA ===")
	
	// Gerar chaves
	fmt.Println("Gerando chaves RSA...")
	keyGenerationTime := benchmarkKeyGeneration(2048)
	fmt.Printf("Tempo de geração de chaves: %v\n", keyGenerationTime)
	
	key := NewRSAKey(2048)
	fmt.Printf("Chave gerada: n=%s, e=%s, d=%s\n", key.N.String(), key.E.String(), key.D.String())
	
	// Mensagem de teste
	message := big.NewInt(12345)
	fmt.Printf("Mensagem original: %s\n", message.String())
	
	// Criptografar
	encryptionTime := benchmarkEncryption(message, key)
	ciphertext := Encrypt(message, key)
	fmt.Printf("Texto cifrado: %s\n", ciphertext.String())
	fmt.Printf("Tempo de criptografia: %v\n", encryptionTime)
	
	// Descriptografar
	decryptionTime := benchmarkDecryption(ciphertext, key)
	decrypted := Decrypt(ciphertext, key)
	fmt.Printf("Texto descriptografado: %s\n", decrypted.String())
	fmt.Printf("Tempo de descriptografia: %v\n", decryptionTime)
	
	// Verificar se a descriptografia foi correta
	if message.Cmp(decrypted) == 0 {
		fmt.Println("✓ Criptografia/Descriptografia bem-sucedida!")
	} else {
		fmt.Println("✗ Erro na criptografia/Descriptografia!")
	}
}

// runLibRSA executa a implementação com biblioteca
func runLibRSA() {
	fmt.Println("=== Implementação RSA com Biblioteca ===")
	
	// Gerar chaves
	fmt.Println("Gerando chaves RSA...")
	keyGenerationTime := benchmarkKeyGenerationLib(2048)
	fmt.Printf("Tempo de geração de chaves: %v\n", keyGenerationTime)
	
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	
	publicKey := &privateKey.PublicKey
	fmt.Println("Chave gerada com sucesso!")
	
	// Mensagem de teste
	message := []byte("Hello, RSA!")
	fmt.Printf("Mensagem original: %s\n", string(message))
	
	// Criptografar
	encryptionTime := benchmarkEncryptionLib(message, publicKey)
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, message)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Texto cifrado: %x\n", ciphertext)
	fmt.Printf("Tempo de criptografia: %v\n", encryptionTime)
	
	// Descriptografar
	decryptionTime := benchmarkDecryptionLib(ciphertext, privateKey)
	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Texto descriptografado: %s\n", string(decrypted))
	fmt.Printf("Tempo de descriptografia: %v\n", decryptionTime)
	
	// Verificar se a descriptografia foi correta
	if string(message) == string(decrypted) {
		fmt.Println("✓ Criptografia/Descriptografia bem-sucedida!")
	} else {
		fmt.Println("✗ Erro na criptografia/Descriptografia!")
	}
}

// runBenchmarkManual executa os benchmarks manuais
func runBenchmarkManual() {
	benchmarkManualRSA()
	benchmarkKeySizes()
	benchmarkMessageSizes()
} 