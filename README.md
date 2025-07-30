# Avaliação de Desempenho do RSA em Python, Rust e Go

Este projeto implementa uma avaliação de desempenho das implementações do algoritmo RSA em diferentes linguagens de programação, comparando implementações manuais e com bibliotecas, além de medir o impacto do uso de ferramentas de benchmark.

## Estrutura do Projeto

```
Performance-Evaluation-of-RSA-in-Python-Rust-and-Go/
├── rust/                          # Implementações em Rust
│   ├── Cargo.toml                 # Dependências Rust
│   ├── src/
│   │   ├── main.rs               # Arquivo principal
│   │   ├── rsa_manual.rs         # Implementação manual
│   │   └── rsa_lib.rs            # Implementação com biblioteca
│   └── benches/
│       ├── benchmark_manual.rs    # Benchmark manual
│       └── benchmark_lib.rs      # Benchmark com criterion
├── go/                           # Implementações em Go
│   ├── go.mod                    # Módulo Go
│   ├── main.go                   # Arquivo principal
│   ├── rsa_manual.go             # Implementação manual
│   ├── rsa_lib.go                # Implementação com biblioteca
│   ├── benchmark_manual.go       # Benchmark manual
│   └── benchmark_lib_test.go     # Benchmark com testing
└── python/                       # Implementações em Python (existente)
```

## Implementações

### Rust

#### Dependências
- `rsa`: Biblioteca para criptografia RSA
- `rand`: Geração de números aleatórios
- `criterion`: Framework de benchmark

#### Execução

```bash
cd rust

# Executar implementação manual
cargo run -- manual

# Executar implementação com biblioteca
cargo run -- lib

# Executar benchmarks com criterion
cargo bench
```

### Go

#### Dependências
- `crypto/rsa`: Biblioteca padrão para RSA
- `testing`: Framework de benchmark

#### Execução

```bash
cd go

# Executar implementação manual
go run . manual

# Executar implementação com biblioteca
go run . lib

# Executar benchmarks manuais
go run . benchmark

# Executar benchmarks com testing
go test -bench=.
```

## Funcionalidades Implementadas

### 1. Geração de Chaves RSA (2048 bits)
- Implementação manual usando algoritmos de geração de primos
- Implementação com bibliotecas nativas
- Medição de tempo de geração

### 2. Criptografia e Descriptografia
- Criptografia de mensagens usando chaves públicas
- Descriptografia usando chaves privadas
- Verificação de integridade das operações

### 3. Medição de Performance
- **Benchmarks Manuais**: Usando funções nativas de tempo
- **Benchmarks com Bibliotecas**: Usando frameworks especializados
  - Rust: Criterion
  - Go: testing

### 4. Análises Comparativas
- Comparação entre implementações manuais vs bibliotecas
- Análise de diferentes tamanhos de chave (512, 1024, 2048 bits)
- Análise de diferentes tamanhos de mensagem
- Estatísticas detalhadas (média, desvio padrão, mínimo, máximo)

## Métricas de Performance

### Rust
- **Tempo de Geração de Chaves**: Medido em nanossegundos
- **Tempo de Criptografia**: Operações por segundo
- **Tempo de Descriptografia**: Operações por segundo
- **Estatísticas**: Média, desvio padrão, intervalo de confiança

### Go
- **Tempo de Geração de Chaves**: Medido em nanossegundos
- **Tempo de Criptografia**: Operações por segundo
- **Tempo de Descriptografia**: Operações por segundo
- **Estatísticas**: Média, desvio padrão, mínimo, máximo

## Objetivos da Avaliação

1. **Comparação entre Linguagens**: Analisar diferenças de performance entre Rust, Go e Python
2. **Impacto das Bibliotecas**: Comparar implementações manuais vs bibliotecas otimizadas
3. **Eficiência dos Benchmarks**: Avaliar diferenças entre benchmarks manuais e com ferramentas especializadas
4. **Escalabilidade**: Testar performance com diferentes tamanhos de chave e mensagem

## Resultados Esperados

- **Rust**: Esperado melhor performance devido à compilação nativa e otimizações
- **Go**: Performance intermediária com garbage collection
- **Python**: Performance mais lenta devido à interpretação
- **Bibliotecas**: Melhor performance que implementações manuais
- **Ferramentas de Benchmark**: Resultados mais precisos e estatisticamente significativos

## Próximos Passos

1. Executar benchmarks em ambiente controlado
2. Coletar dados de performance
3. Analisar resultados estatisticamente
4. Gerar gráficos comparativos
5. Documentar conclusões e recomendações

## Requisitos

### Rust
- Rust 1.70+ instalado
- Cargo (gerenciador de pacotes)

### Go
- Go 1.21+ instalado
- Módulos Go habilitados

### Python
- Python 3.8+ instalado
- Bibliotecas de criptografia (PyCryptodome) 