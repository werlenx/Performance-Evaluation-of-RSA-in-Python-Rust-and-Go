use criterion::{black_box, criterion_group, criterion_main, Criterion};
use rsa_manual::{RSAKey, encrypt, decrypt, generate_rsa_keys};
use rsa_lib::{RSALibKey, encrypt_lib, decrypt_lib};

// Benchmark para implementação manual
fn benchmark_manual_rsa(c: &mut Criterion) {
    let mut group = c.benchmark_group("RSA Manual");
    
    // Benchmark de geração de chaves
    group.bench_function("key_generation", |b| {
        b.iter(|| {
            black_box(RSAKey::new());
        });
    });
    
    // Gerar uma chave para os testes de criptografia
    let key = RSAKey::new();
    let test_message = 12345u128;
    
    // Benchmark de criptografia
    group.bench_function("encryption", |b| {
        b.iter(|| {
            black_box(encrypt(test_message, &key));
        });
    });
    
    // Criptografar uma mensagem para o teste de descriptografia
    let ciphertext = encrypt(test_message, &key);
    
    // Benchmark de descriptografia
    group.bench_function("decryption", |b| {
        b.iter(|| {
            black_box(decrypt(ciphertext, &key));
        });
    });
    
    group.finish();
}

// Benchmark para implementação com biblioteca
fn benchmark_lib_rsa(c: &mut Criterion) {
    let mut group = c.benchmark_group("RSA Biblioteca");
    
    // Benchmark de geração de chaves
    group.bench_function("key_generation", |b| {
        b.iter(|| {
            black_box(RSALibKey::new());
        });
    });
    
    // Gerar uma chave para os testes de criptografia
    let key = RSALibKey::new();
    let test_message = b"Hello, RSA!";
    
    // Benchmark de criptografia
    group.bench_function("encryption", |b| {
        b.iter(|| {
            black_box(encrypt_lib(test_message, &key.public_key));
        });
    });
    
    // Criptografar uma mensagem para o teste de descriptografia
    let ciphertext = encrypt_lib(test_message, &key.public_key).unwrap();
    
    // Benchmark de descriptografia
    group.bench_function("decryption", |b| {
        b.iter(|| {
            black_box(decrypt_lib(&ciphertext, &key.private_key));
        });
    });
    
    group.finish();
}

// Benchmark comparativo entre implementações
fn benchmark_comparison(c: &mut Criterion) {
    let mut group = c.benchmark_group("Comparação RSA");
    
    // Geração de chaves
    group.bench_function("manual_key_gen", |b| {
        b.iter(|| {
            black_box(RSAKey::new());
        });
    });
    
    group.bench_function("lib_key_gen", |b| {
        b.iter(|| {
            black_box(RSALibKey::new());
        });
    });
    
    // Criptografia
    let manual_key = RSAKey::new();
    let lib_key = RSALibKey::new();
    let manual_message = 12345u128;
    let lib_message = b"Hello, RSA!";
    
    group.bench_function("manual_encryption", |b| {
        b.iter(|| {
            black_box(encrypt(manual_message, &manual_key));
        });
    });
    
    group.bench_function("lib_encryption", |b| {
        b.iter(|| {
            black_box(encrypt_lib(lib_message, &lib_key.public_key));
        });
    });
    
    // Descriptografia
    let manual_ciphertext = encrypt(manual_message, &manual_key);
    let lib_ciphertext = encrypt_lib(lib_message, &lib_key.public_key).unwrap();
    
    group.bench_function("manual_decryption", |b| {
        b.iter(|| {
            black_box(decrypt(manual_ciphertext, &manual_key));
        });
    });
    
    group.bench_function("lib_decryption", |b| {
        b.iter(|| {
            black_box(decrypt_lib(&lib_ciphertext, &lib_key.private_key));
        });
    });
    
    group.finish();
}

// Benchmark de diferentes tamanhos de chave (apenas manual)
fn benchmark_key_sizes(c: &mut Criterion) {
    let mut group = c.benchmark_group("Tamanhos de Chave");
    
    let key_sizes = vec![512, 1024, 2048];
    
    for &bits in &key_sizes {
        group.bench_function(&format!("{}_bits", bits), |b| {
            b.iter(|| {
                black_box(generate_rsa_keys(bits));
            });
        });
    }
    
    group.finish();
}

// Configuração do criterion
criterion_group! {
    name = benches;
    config = Criterion::default()
        .sample_size(100)
        .confidence_level(0.95)
        .significance_level(0.05);
    targets = 
        benchmark_manual_rsa,
        benchmark_lib_rsa,
        benchmark_comparison,
        benchmark_key_sizes
}

criterion_main!(benches); 