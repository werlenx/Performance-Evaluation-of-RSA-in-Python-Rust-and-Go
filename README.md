# Avaliação de Desempenho do RSA em Python, Rust e Go

Este projeto implementa uma avaliação de desempenho das implementações do algoritmo RSA em diferentes linguagens de programação, comparando implementações manuais e com bibliotecas, além de medir o impacto do uso de ferramentas de benchmark.

## Estrutura do Projeto

```
Performance-Evaluation-of-RSA-in-Python-Rust-and-Go/
├── python/                         # Implementações em Python
│   ├── cryptodome/
│   │   ├── NoPytest/
│   │   │   ├── main.py            # Arquivo principal com benchmarks
│   │   │   ├── GenerateKeys.py    # Geração de chaves RSA
│   │   │   ├── Encrypt.py         # Criptografia RSA
│   │   │   ├── Decrypt.py         # Descriptografia RSA
│   │   │   ├── Statistics.py      # Cálculo de estatísticas
│   │   │   ├── MonitorMenCpu.py   # Monitoramento de CPU/Memória
│   │   │   └── full.py            # Implementação completa
│   │   └── txt.txt                # Arquivo de teste
│   └── testes/
│       └── test.py                # Testes adicionais
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
│   ├── benchmark_lib_test.go     # Benchmark com testing
│   └── gotests.go                # Testes adicionais
├── RSA_venv/                     # Ambiente virtual Python
├── playground.py                  # Script de teste
└── README.md                     # Documentação
```

## Implementações

### Python (PyCryptodome)

#### Dependências
- `pycryptodome`: Biblioteca para criptografia RSA
- `tabulate`: Formatação de tabelas
- `psutil`: Monitoramento de recursos do sistema

#### Funcionalidades
- **Geração de Chaves**: RSA.generate() com diferentes tamanhos (2048, 4096 bits)
- **Criptografia**: PKCS1_OAEP para criptografia
- **Descriptografia**: PKCS1_OAEP para descriptografia
- **Monitoramento**: CPU e uso de memória em tempo real
- **Estatísticas**: Média e desvio padrão para múltiplas execuções

#### Execução

```bash
cd python/cryptodome/NoPytest

# Executar benchmark completo
python main.py

# Executar implementação completa
python full.py

# Executar módulos individuais
python GenerateKeys.py
python Encrypt.py
python Decrypt.py
```

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

### 1. Geração de Chaves RSA
- **Python**: RSA.generate() com PyCryptodome (2048, 4096 bits)
- **Rust**: Implementação manual e com biblioteca rsa (2048 bits)
- **Go**: Implementação manual e com crypto/rsa (2048 bits)

### 2. Criptografia e Descriptografia
- **Python**: PKCS1_OAEP para criptografia/descriptografia
- **Rust**: Pkcs1v15Encrypt para criptografia/descriptografia
- **Go**: PKCS1v15 para criptografia/descriptografia

### 3. Medição de Performance
- **Python**: Monitoramento de CPU, memória e tempo com psutil
- **Rust**: Benchmarks manuais e com Criterion
- **Go**: Benchmarks manuais e com testing framework

### 4. Análises Comparativas
- Comparação entre implementações manuais vs bibliotecas
- Análise de diferentes tamanhos de chave
- Estatísticas detalhadas (média, desvio padrão, mínimo, máximo)
- Monitoramento de recursos do sistema (CPU, memória)

## Métricas de Performance

### Python (PyCryptodome)
- **Tempo de Geração de Chaves**: Medido em segundos
- **Tempo de Criptografia**: Medido em segundos
- **Tempo de Descriptografia**: Medido em segundos
- **Uso de CPU**: Percentual durante execução
- **Uso de Memória**: Bytes utilizados
- **Estatísticas**: Média e desvio padrão

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

1. **Comparação entre Linguagens**: Analisar diferenças de performance entre Python, Rust e Go
2. **Impacto das Bibliotecas**: Comparar implementações manuais vs bibliotecas otimizadas
3. **Eficiência dos Benchmarks**: Avaliar diferenças entre benchmarks manuais e com ferramentas especializadas
4. **Monitoramento de Recursos**: Analisar uso de CPU e memória em diferentes implementações
5. **Escalabilidade**: Testar performance com diferentes tamanhos de chave

## Resultados Esperados

- **Python**: Performance mais lenta devido à interpretação, mas com monitoramento detalhado de recursos
- **Rust**: Esperado melhor performance devido à compilação nativa e otimizações
- **Go**: Performance intermediária com garbage collection
- **Bibliotecas**: Melhor performance que implementações manuais
- **Ferramentas de Benchmark**: Resultados mais precisos e estatisticamente significativos

## Características Específicas

### Python (PyCryptodome)
- ✅ Implementação completa com PyCryptodome
- ✅ Monitoramento de CPU e memória em tempo real
- ✅ Estatísticas detalhadas com múltiplas execuções
- ✅ Suporte a diferentes tamanhos de chave (2048, 4096 bits)
- ✅ Formatação de resultados em tabelas

### Rust
- ✅ Implementação manual do RSA
- ✅ Implementação com biblioteca rsa
- ✅ Benchmarks manuais e com Criterion
- ✅ Comparação entre implementações

### Go
- ✅ Implementação manual do RSA
- ✅ Implementação com crypto/rsa
- ✅ Benchmarks manuais e com testing
- ✅ Comparação entre implementações

## Próximos Passos

1. Executar benchmarks em ambiente controlado
2. Coletar dados de performance de todas as implementações
3. Analisar resultados estatisticamente
4. Gerar gráficos comparativos
5. Documentar conclusões e recomendações

## Requisitos

### Python
- Python 3.8+ instalado
- PyCryptodome: `pip install pycryptodome`
- Tabulate: `pip install tabulate`
- Psutil: `pip install psutil`

### Rust
- Rust 1.70+ instalado
- Cargo (gerenciador de pacotes)

### Go
- Go 1.21+ instalado
- Módulos Go habilitados

## Como Contribuir

1. Clone o repositório
2. Instale as dependências para cada linguagem
3. Execute os benchmarks
4. Analise os resultados
5. Documente suas descobertas

## Licença

Este projeto é destinado para fins educacionais e de pesquisa em avaliação de performance de algoritmos criptográficos. 