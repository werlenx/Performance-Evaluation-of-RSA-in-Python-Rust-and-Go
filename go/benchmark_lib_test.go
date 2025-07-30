package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"math/big"
	"testing"
)

// BenchmarkManualKeyGeneration testa geração de chaves manual
func BenchmarkManualKeyGeneration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewRSAKey(2048)
	}
}

// BenchmarkManualEncryption testa criptografia manual
func BenchmarkManualEncryption(b *testing.B) {
	key := NewRSAKey(2048)
	message := big.NewInt(12345)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Encrypt(message, key)
	}
}

// BenchmarkManualDecryption testa descriptografia manual
func BenchmarkManualDecryption(b *testing.B) {
	key := NewRSAKey(2048)
	message := big.NewInt(12345)
	ciphertext := Encrypt(message, key)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Decrypt(ciphertext, key)
	}
}

// BenchmarkLibKeyGeneration testa geração de chaves com biblioteca
func BenchmarkLibKeyGeneration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkLibEncryption testa criptografia com biblioteca
func BenchmarkLibEncryption(b *testing.B) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		b.Fatal(err)
	}
	publicKey := &privateKey.PublicKey
	message := []byte("Hello, RSA!")
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, message)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkLibDecryption testa descriptografia com biblioteca
func BenchmarkLibDecryption(b *testing.B) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		b.Fatal(err)
	}
	publicKey := &privateKey.PublicKey
	message := []byte("Hello, RSA!")
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, message)
	if err != nil {
		b.Fatal(err)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkKeySizes testa diferentes tamanhos de chave (manual)
func BenchmarkKeySizes(b *testing.B) {
	keySizes := []int{512, 1024, 2048}
	
	for _, bits := range keySizes {
		b.Run(fmt.Sprintf("%d_bits", bits), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = NewRSAKey(bits)
			}
		})
	}
}

// BenchmarkMessageSizes testa diferentes tamanhos de mensagem (manual)
func BenchmarkMessageSizes(b *testing.B) {
	key := NewRSAKey(2048)
	messageSizes := []int{100, 1000, 10000}
	
	for _, size := range messageSizes {
		b.Run(fmt.Sprintf("%d_bytes", size), func(b *testing.B) {
			testMessage := big.NewInt(int64(size))
			
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = Encrypt(testMessage, key)
			}
		})
	}
}

// BenchmarkComparison compara implementações manual vs biblioteca
func BenchmarkComparison(b *testing.B) {
	// Geração de chaves
	b.Run("manual_key_gen", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = NewRSAKey(2048)
		}
	})
	
	b.Run("lib_key_gen", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, err := rsa.GenerateKey(rand.Reader, 2048)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	
	// Criptografia
	b.Run("manual_encryption", func(b *testing.B) {
		key := NewRSAKey(2048)
		message := big.NewInt(12345)
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Encrypt(message, key)
		}
	})
	
	b.Run("lib_encryption", func(b *testing.B) {
		privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			b.Fatal(err)
		}
		publicKey := &privateKey.PublicKey
		message := []byte("Hello, RSA!")
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, message)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
	
	// Descriptografia
	b.Run("manual_decryption", func(b *testing.B) {
		key := NewRSAKey(2048)
		message := big.NewInt(12345)
		ciphertext := Encrypt(message, key)
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = Decrypt(ciphertext, key)
		}
	})
	
	b.Run("lib_decryption", func(b *testing.B) {
		privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			b.Fatal(err)
		}
		publicKey := &privateKey.PublicKey
		message := []byte("Hello, RSA!")
		ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, message)
		if err != nil {
			b.Fatal(err)
		}
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
} 