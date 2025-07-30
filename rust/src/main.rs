mod rsa_manual;
mod rsa_lib;

use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();
    
    if args.len() < 2 {
        println!("Uso: cargo run -- <opção>");
        println!("Opções:");
        println!("  manual    - Executar implementação manual do RSA");
        println!("  lib       - Executar implementação com biblioteca");
        println!("  benchmark - Executar benchmarks");
        return;
    }
    
    match args[1].as_str() {
        "manual" => {
            println!("Executando implementação manual do RSA...");
            rsa_manual::main();
        }
        "lib" => {
            println!("Executando implementação com biblioteca...");
            rsa_lib::main();
        }
        "benchmark" => {
            println!("Executando benchmarks...");
            // Para benchmarks detalhados, use: cargo bench
            println!("Para benchmarks detalhados, execute: cargo bench");
        }
        _ => {
            println!("Opção inválida. Use 'manual', 'lib' ou 'benchmark'");
        }
    }
} 