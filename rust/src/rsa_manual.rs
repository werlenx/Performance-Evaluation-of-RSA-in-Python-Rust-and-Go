use std::time::{Duration, Instant};
use rand::Rng;

// Estrutura para representar uma chave RSA
#[derive(Debug, Clone)]
pub struct RSAKey {
    pub n: u128,  // módulo
    pub e: u128,  // expoente público
    pub d: u128,  // expoente privado
}

impl RSAKey {
    pub fn new() -> Self {
        let (n, e, d) = generate_rsa_keys(2048);
        RSAKey { n, e, d }
    }
}

// Função para verificar se um número é primo
fn is_prime(n: u128) -> bool {
    if n < 2 {
        return false;
    }
    if n == 2 {
        return true;
    }
    if n % 2 == 0 {
        return false;
    }
    
    let sqrt_n = (n as f64).sqrt() as u128;
    for i in (3..=sqrt_n).step_by(2) {
        if n % i == 0 {
            return false;
        }
    }
    true
}

// Função para gerar um número primo aleatório
fn generate_prime(bits: u32) -> u128 {
    let mut rng = rand::thread_rng();
    let min = 1u128 << (bits - 1);
    let max = (1u128 << bits) - 1;
    
    loop {
        let candidate = rng.gen_range(min..=max);
        if is_prime(candidate) {
            return candidate;
        }
    }
}

// Função para calcular o MDC (Máximo Divisor Comum)
fn gcd(a: u128, b: u128) -> u128 {
    if b == 0 {
        a
    } else {
        gcd(b, a % b)
    }
}

// Função para calcular o inverso modular
fn mod_inverse(a: u128, m: u128) -> Option<u128> {
    let mut t = 0i128;
    let mut new_t = 1i128;
    let mut r = m as i128;
    let mut new_r = a as i128;
    
    while new_r != 0 {
        let quotient = r / new_r;
        let temp_t = t - quotient * new_t;
        let temp_r = r - quotient * new_r;
        
        t = new_t;
        new_t = temp_t;
        r = new_r;
        new_r = temp_r;
    }
    
    if r > 1 {
        return None;
    }
    
    if t < 0 {
        t += m as i128;
    }
    
    Some(t as u128)
}

// Função para exponenciação modular rápida
fn mod_pow(base: u128, exponent: u128, modulus: u128) -> u128 {
    if modulus == 1 {
        return 0;
    }
    
    let mut result = 1;
    let mut base = base % modulus;
    let mut exponent = exponent;
    
    while exponent > 0 {
        if exponent % 2 == 1 {
            result = (result * base) % modulus;
        }
        exponent >>= 1;
        base = (base * base) % modulus;
    }
    
    result
}

// Função para gerar chaves RSA
pub fn generate_rsa_keys(bits: u32) -> (u128, u128, u128) {
    let prime_bits = bits / 2;
    
    // Gerar dois números primos grandes
    let p = generate_prime(prime_bits);
    let q = generate_prime(prime_bits);
    
    let n = p * q;
    let phi = (p - 1) * (q - 1);
    
    // Escolher expoente público e
    let e = 65537; // Valor comum para e
    
    // Calcular expoente privado d
    let d = mod_inverse(e, phi).expect("Não foi possível calcular o inverso modular");
    
    (n, e, d)
}

// Função para criptografar
pub fn encrypt(message: u128, public_key: &RSAKey) -> u128 {
    mod_pow(message, public_key.e, public_key.n)
}

// Função para descriptografar
pub fn decrypt(ciphertext: u128, private_key: &RSAKey) -> u128 {
    mod_pow(ciphertext, private_key.d, private_key.n)
}

// Função para medir o tempo de geração de chaves
pub fn benchmark_key_generation() -> Duration {
    let start = Instant::now();
    let _key = RSAKey::new();
    start.elapsed()
}

// Função para medir o tempo de criptografia
pub fn benchmark_encryption(message: u128, key: &RSAKey) -> Duration {
    let start = Instant::now();
    let _ciphertext = encrypt(message, key);
    start.elapsed()
}

// Função para medir o tempo de descriptografia
pub fn benchmark_decryption(ciphertext: u128, key: &RSAKey) -> Duration {
    let start = Instant::now();
    let _plaintext = decrypt(ciphertext, key);
    start.elapsed()
}

// Função principal para demonstração
pub fn main() {
    println!("=== Implementação Manual do RSA ===");
    
    // Gerar chaves
    println!("Gerando chaves RSA...");
    let key_generation_time = benchmark_key_generation();
    println!("Tempo de geração de chaves: {:?}", key_generation_time);
    
    let key = RSAKey::new();
    println!("Chave gerada: n={}, e={}, d={}", key.n, key.e, key.d);
    
    // Mensagem de teste
    let message = 12345u128;
    println!("Mensagem original: {}", message);
    
    // Criptografar
    let encryption_time = benchmark_encryption(message, &key);
    let ciphertext = encrypt(message, &key);
    println!("Texto cifrado: {}", ciphertext);
    println!("Tempo de criptografia: {:?}", encryption_time);
    
    // Descriptografar
    let decryption_time = benchmark_decryption(ciphertext, &key);
    let decrypted = decrypt(ciphertext, &key);
    println!("Texto descriptografado: {}", decrypted);
    println!("Tempo de descriptografia: {:?}", decryption_time);
    
    // Verificar se a descriptografia foi correta
    if message == decrypted {
        println!("✓ Criptografia/Descriptografia bem-sucedida!");
    } else {
        println!("✗ Erro na criptografia/Descriptografia!");
    }
} 