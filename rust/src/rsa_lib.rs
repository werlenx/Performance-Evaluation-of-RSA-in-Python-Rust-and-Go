use std::time::{Duration, Instant};
use rsa::{RsaPrivateKey, RsaPublicKey, pkcs8::{EncodePublicKey, LineEnding}};
use rsa::pkcs8::EncodePrivateKey;
use rsa::Pkcs1v15Encrypt;
use rand::rngs::OsRng;

// Estrutura para representar chaves RSA usando biblioteca
#[derive(Debug, Clone)]
pub struct RSALibKey {
    pub private_key: RsaPrivateKey,
    pub public_key: RsaPublicKey,
}

impl RSALibKey {
    pub fn new() -> Self {
        let private_key = RsaPrivateKey::new(&mut OsRng, 2048).expect("Falha ao gerar chave privada");
        let public_key = RsaPublicKey::from(&private_key);
        RSALibKey { private_key, public_key }
    }
}

// Função para criptografar usando biblioteca
pub fn encrypt_lib(message: &[u8], public_key: &RsaPublicKey) -> Result<Vec<u8>, rsa::Error> {
    public_key.encrypt(&mut OsRng, Pkcs1v15Encrypt, message)
}

// Função para descriptografar usando biblioteca
pub fn decrypt_lib(ciphertext: &[u8], private_key: &RsaPrivateKey) -> Result<Vec<u8>, rsa::Error> {
    private_key.decrypt(Pkcs1v15Encrypt, ciphertext)
}

// Função para medir o tempo de geração de chaves
pub fn benchmark_key_generation_lib() -> Duration {
    let start = Instant::now();
    let _key = RSALibKey::new();
    start.elapsed()
}

// Função para medir o tempo de criptografia
pub fn benchmark_encryption_lib(message: &[u8], key: &RSALibKey) -> Duration {
    let start = Instant::now();
    let _ciphertext = encrypt_lib(message, &key.public_key);
    start.elapsed()
}

// Função para medir o tempo de descriptografia
pub fn benchmark_decryption_lib(ciphertext: &[u8], key: &RSALibKey) -> Duration {
    let start = Instant::now();
    let _plaintext = decrypt_lib(ciphertext, &key.private_key);
    start.elapsed()
}

// Função principal para demonstração
pub fn main() {
    println!("=== Implementação RSA com Biblioteca ===");
    
    // Gerar chaves
    println!("Gerando chaves RSA...");
    let key_generation_time = benchmark_key_generation_lib();
    println!("Tempo de geração de chaves: {:?}", key_generation_time);
    
    let key = RSALibKey::new();
    println!("Chave gerada com sucesso!");
    
    // Mensagem de teste
    let message = b"Hello, RSA!";
    println!("Mensagem original: {}", String::from_utf8_lossy(message));
    
    // Criptografar
    let encryption_time = benchmark_encryption_lib(message, &key);
    let ciphertext = encrypt_lib(message, &key.public_key).expect("Falha na criptografia");
    println!("Texto cifrado: {:?}", ciphertext);
    println!("Tempo de criptografia: {:?}", encryption_time);
    
    // Descriptografar
    let decryption_time = benchmark_decryption_lib(&ciphertext, &key);
    let decrypted = decrypt_lib(&ciphertext, &key.private_key).expect("Falha na descriptografia");
    println!("Texto descriptografado: {}", String::from_utf8_lossy(&decrypted));
    println!("Tempo de descriptografia: {:?}", decryption_time);
    
    // Verificar se a descriptografia foi correta
    if message == decrypted.as_slice() {
        println!("✓ Criptografia/Descriptografia bem-sucedida!");
    } else {
        println!("✗ Erro na criptografia/Descriptografia!");
    }
    
    // Salvar chaves em formato PEM (opcional)
    println!("\n=== Informações das Chaves ===");
    println!("Chave pública:");
    println!("{}", key.public_key.to_public_key_pem(LineEnding::LF).unwrap());
    println!("Chave privada:");
    println!("{}", key.private_key.to_pkcs8_pem(LineEnding::LF).unwrap());
} 