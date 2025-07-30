package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

// RSAKey representa uma chave RSA
type RSAKey struct {
	N *big.Int // módulo
	E *big.Int // expoente público
	D *big.Int // expoente privado
}

// NewRSAKey gera uma nova chave RSA
func NewRSAKey(bits int) *RSAKey {
	p := generatePrime(bits / 2)
	q := generatePrime(bits / 2)
	
	n := new(big.Int).Mul(p, q)
	phi := new(big.Int).Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))
	
	// Expoente público e = 65537
	e := big.NewInt(65537)
	
	// Calcular expoente privado d
	d := new(big.Int).ModInverse(e, phi)
	if d == nil {
		panic("Não foi possível calcular o inverso modular")
	}
	
	return &RSAKey{
		N: n,
		E: e,
		D: d,
	}
}

// generatePrime gera um número primo aleatório
func generatePrime(bits int) *big.Int {
	for {
		// Gerar número aleatório
		candidate, err := rand.Prime(rand.Reader, bits)
		if err != nil {
			panic(err)
		}
		
		// Verificar se é primo (rand.Prime já garante isso)
		return candidate
	}
}

// modPow calcula base^exponent mod modulus usando exponenciação rápida
func modPow(base, exponent, modulus *big.Int) *big.Int {
	if modulus.Cmp(big.NewInt(1)) == 0 {
		return big.NewInt(0)
	}
	
	result := big.NewInt(1)
	base = new(big.Int).Mod(base, modulus)
	exp := new(big.Int).Set(exponent)
	
	for exp.Cmp(big.NewInt(0)) > 0 {
		if new(big.Int).And(exp, big.NewInt(1)).Cmp(big.NewInt(1)) == 0 {
			result = new(big.Int).Mod(new(big.Int).Mul(result, base), modulus)
		}
		exp.Rsh(exp, 1)
		base = new(big.Int).Mod(new(big.Int).Mul(base, base), modulus)
	}
	
	return result
}

// Encrypt criptografa uma mensagem usando a chave pública
func Encrypt(message *big.Int, key *RSAKey) *big.Int {
	return modPow(message, key.E, key.N)
}

// Decrypt descriptografa uma mensagem usando a chave privada
func Decrypt(ciphertext *big.Int, key *RSAKey) *big.Int {
	return modPow(ciphertext, key.D, key.N)
}

// benchmarkKeyGeneration mede o tempo de geração de chaves
func benchmarkKeyGeneration(bits int) time.Duration {
	start := time.Now()
	_ = NewRSAKey(bits)
	return time.Since(start)
}

// benchmarkEncryption mede o tempo de criptografia
func benchmarkEncryption(message *big.Int, key *RSAKey) time.Duration {
	start := time.Now()
	_ = Encrypt(message, key)
	return time.Since(start)
}

// benchmarkDecryption mede o tempo de descriptografia
func benchmarkDecryption(ciphertext *big.Int, key *RSAKey) time.Duration {
	start := time.Now()
	_ = Decrypt(ciphertext, key)
	return time.Since(start)
}

// main demonstra a implementação manual do RSA
func main() {
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