use std::time::{Duration, Instant};
use std::collections::HashMap;
use rsa_manual::{RSAKey, encrypt, decrypt, generate_rsa_keys};

// Função para executar múltiplas medições e calcular estatísticas
fn run_benchmark<F>(func: F, iterations: usize, name: &str) -> HashMap<String, f64>
where
    F: Fn() -> Duration,
{
    let mut times = Vec::new();
    
    println!("Executando benchmark: {}", name);
    
    for i in 0..iterations {
        let duration = func();
        times.push(duration.as_nanos() as f64);
        
        if (i + 1) % 10 == 0 {
            println!("  Iteração {}/{}", i + 1, iterations);
        }
    }
    
    // Calcular estatísticas
    let total: f64 = times.iter().sum();
    let mean = total / times.len() as f64;
    
    let variance: f64 = times.iter()
        .map(|&x| (x - mean).powi(2))
        .sum::<f64>() / times.len() as f64;
    let std_dev = variance.sqrt();
    
    let min = times.iter().fold(f64::INFINITY, |a, &b| a.min(b));
    let max = times.iter().fold(f64::NEG_INFINITY, |a, &b| a.max(b));
    
    let mut stats = HashMap::new();
    stats.insert("mean_ns".to_string(), mean);
    stats.insert("std_dev_ns".to_string(), std_dev);
    stats.insert("min_ns".to_string(), min);
    stats.insert("max_ns".to_string(), max);
    stats.insert("total_ns".to_string(), total);
    
    stats
}

// Função para imprimir estatísticas
fn print_stats(stats: &HashMap<String, f64>, operation: &str) {
    println!("\n=== Estatísticas para {} ===", operation);
    println!("Média: {:.2} ns", stats["mean_ns"]);
    println!("Desvio Padrão: {:.2} ns", stats["std_dev_ns"]);
    println!("Mínimo: {:.2} ns", stats["min_ns"]);
    println!("Máximo: {:.2} ns", stats["max_ns"]);
    println!("Total: {:.2} ns", stats["total_ns"]);
}

// Função principal do benchmark manual
pub fn main() {
    println!("=== Benchmark Manual RSA ===");
    println!("Usando funções nativas de tempo\n");
    
    let iterations = 100;
    
    // Benchmark de geração de chaves
    let key_gen_stats = run_benchmark(
        || {
            let start = Instant::now();
            let _key = RSAKey::new();
            start.elapsed()
        },
        iterations,
        "Geração de Chaves"
    );
    print_stats(&key_gen_stats, "Geração de Chaves");
    
    // Gerar uma chave para os testes de criptografia
    let key = RSAKey::new();
    let test_message = 12345u128;
    
    // Benchmark de criptografia
    let encryption_stats = run_benchmark(
        || {
            let start = Instant::now();
            let _ciphertext = encrypt(test_message, &key);
            start.elapsed()
        },
        iterations,
        "Criptografia"
    );
    print_stats(&encryption_stats, "Criptografia");
    
    // Criptografar uma mensagem para o teste de descriptografia
    let ciphertext = encrypt(test_message, &key);
    
    // Benchmark de descriptografia
    let decryption_stats = run_benchmark(
        || {
            let start = Instant::now();
            let _plaintext = decrypt(ciphertext, &key);
            start.elapsed()
        },
        iterations,
        "Descriptografia"
    );
    print_stats(&decryption_stats, "Descriptografia");
    
    // Comparação de performance
    println!("\n=== Comparação de Performance ===");
    println!("Geração de Chaves: {:.2} ns (média)", key_gen_stats["mean_ns"]);
    println!("Criptografia: {:.2} ns (média)", encryption_stats["mean_ns"]);
    println!("Descriptografia: {:.2} ns (média)", decryption_stats["mean_ns"]);
    
    // Verificar se a descriptografia funciona corretamente
    let decrypted = decrypt(ciphertext, &key);
    if test_message == decrypted {
        println!("✓ Verificação de integridade: OK");
    } else {
        println!("✗ Verificação de integridade: FALHOU");
    }
}

// Função para benchmark de diferentes tamanhos de chave
pub fn benchmark_key_sizes() {
    println!("\n=== Benchmark de Diferentes Tamanhos de Chave ===");
    
    let key_sizes = vec![512, 1024, 2048];
    let iterations = 50;
    
    for &bits in &key_sizes {
        println!("\n--- Chave de {} bits ---", bits);
        
        let stats = run_benchmark(
            || {
                let start = Instant::now();
                let _keys = generate_rsa_keys(bits);
                start.elapsed()
            },
            iterations,
            &format!("Geração de Chave {} bits", bits)
        );
        
        print_stats(&stats, &format!("Geração de Chave {} bits", bits));
    }
}

// Função para benchmark de diferentes tamanhos de mensagem
pub fn benchmark_message_sizes() {
    println!("\n=== Benchmark de Diferentes Tamanhos de Mensagem ===");
    
    let key = RSAKey::new();
    let message_sizes = vec![100, 1000, 10000, 100000];
    let iterations = 50;
    
    for &size in &message_sizes {
        println!("\n--- Mensagem de {} bytes ---", size);
        
        let test_message = size as u128;
        
        let encryption_stats = run_benchmark(
            || {
                let start = Instant::now();
                let _ciphertext = encrypt(test_message, &key);
                start.elapsed()
            },
            iterations,
            &format!("Criptografia {} bytes", size)
        );
        
        print_stats(&encryption_stats, &format!("Criptografia {} bytes", size));
    }
} 