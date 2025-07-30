package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"
)

// benchmarkKeyGeneration mede o tempo de geração de chaves usando biblioteca
func benchmarkKeyGeneration(bits int) time.Duration {
	start := time.Now()
	_, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}
	return time.Since(start)
}

// benchmarkEncryption mede o tempo de criptografia usando biblioteca
func benchmarkEncryption(message []byte, publicKey *rsa.PublicKey) time.Duration {
	start := time.Now()
	_, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, message)
	if err != nil {
		panic(err)
	}
	return time.Since(start)
}

// benchmarkDecryption mede o tempo de descriptografia usando biblioteca
func benchmarkDecryption(ciphertext []byte, privateKey *rsa.PrivateKey) time.Duration {
	start := time.Now()
	_, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		panic(err)
	}
	return time.Since(start)
}

// main demonstra a implementação do RSA usando biblioteca
func main() {
	fmt.Println("=== Implementação RSA com Biblioteca ===")
	
	// Gerar chaves
	fmt.Println("Gerando chaves RSA...")
	keyGenerationTime := benchmarkKeyGeneration(2048)
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
	encryptionTime := benchmarkEncryption(message, publicKey)
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, message)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Texto cifrado: %x\n", ciphertext)
	fmt.Printf("Tempo de criptografia: %v\n", encryptionTime)
	
	// Descriptografar
	decryptionTime := benchmarkDecryption(ciphertext, privateKey)
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
	
	// Salvar chaves em formato PEM (opcional)
	fmt.Println("\n=== Informações das Chaves ===")
	
	// Chave pública
	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	}
	fmt.Println("Chave pública:")
	fmt.Println(string(pem.EncodeToMemory(publicKeyPEM)))
	
	// Chave privada
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	fmt.Println("Chave privada:")
	fmt.Println(string(pem.EncodeToMemory(privateKeyPEM)))
} 