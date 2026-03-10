# Projeto: kraken

Tipo: Go
Versão: 1.25.7

## Descrição

Ferramenta para análise de projetos Go que gera documentação de endpoints e PRDs (Product Requirements Documents) de forma automatizada.

## Funcionalidades

### 📝 Geração de Documentação de Endpoints

- Detecta automaticamente o tipo de projeto (Go, Node.js, Java/Maven)
- Analisa os endpoints da API
- Gera documentação separada por rota no diretório `docs/kraken/`

### 📋 Criação de PRDs (Product Requirements Documents)

- Interface interativa para criação de PRDs completos
- Segue as melhores práticas do cursor-prd-task-rules
- Gera PRDs estruturados no diretório `tasks/`

## Como Usar

### Build e Execução

#### Opção 1: Script de Instalação (Mais fácil)

```bash
# Instalar com permissões automáticas
./install.sh
```

#### Opção 2: Usando Make (Recomendado)

```bash
# Build com permissões corretas
make build

# Executar
make run
# ou
./kraken
```

#### Opção 3: Build Manual

```bash
# Build
go build -o kraken .

# Dar permissão de execução (importante!)
chmod +x kraken

# Executar
./kraken
```

### 🔧 Solução de Problemas de Permissão

Se o arquivo `kraken` não tiver permissão de execução:

```bash
# Corrigir permissões
chmod +x kraken

# Verificar permissões
ls -la kraken
# Deve mostrar: -rwxrwx-x (ou similar com 'x' para executável)
```

**Dica:** Use sempre `make build` para garantir permissões corretas automaticamente.

### Menu de Opções

1. **Gerar documentação de endpoints** - Analisa o projeto e cria documentação completa da API
2. **Criar PRD** - Inicia o processo interativo de criação de PRD
3. **Sair** - Encerra o programa

### Estrutura de Arquivos Gerados

#### Documentação de Endpoints

```
docs/kraken/
├── NOME-ROTA-DOC.md
└── ...
```

#### PRDs

```
tasks/
├── prd-nome-da-funcionalidade.md
└── ...
```

## Estrutura de um PRD Gerado

Os PRDs gerados seguem a estrutura padrão:

- **Introdução/Visão Geral**
- **Objetivos**
- **Histórias do Usuário** (com critérios de aceitação)
- **Requisitos Funcionais** (com prioridades)
- **Não-Objetivos** (fora do escopo)
- **Considerações de Design**
- **Considerações Técnicas**
- **Métricas de Sucesso**
- **Questões Abertas**

## Exemplo de Uso

### Criar um PRD

1. Execute `./kraken`
2. Escolha a opção 2 (Criar PRD)
3. Siga o assistente interativo:
   - Informe o título da funcionalidade
   - Descreva a funcionalidade
   - Adicione objetivos, histórias de usuário, requisitos, etc.
4. O PRD será gerado em `tasks/prd-nome-da-funcionalidade.md`

### Gerar Documentação

1. Execute `./kraken`
2. Escolha a opção 1 (Gerar documentação)
3. A documentação será criada em `docs/kraken/`

## Requisitos

- Go 1.25.7 ou superior
- Projeto com um dos arquivos de configuração reconhecidos:
  - `go.mod` (projetos Go)
  - `package.json` (projetos Node.js)
  - `pom.xml` (projetos Java/Maven)

## API Endpoints

### GET /api/hello

**Descrição:** Retorna uma mensagem de saudação personalizada

#### Cenários de Sucesso

- **200**: Mensagem de saudação retornada com sucesso

#### Cenários de Erro

- **400**: Parâmetro 'name' não fornecido ou vazio

- **500**: Erro interno ao processar a requisição

---

### GET /api/status

**Descrição:** Retorna o status da API

#### Cenários de Sucesso

- **200**: Status da API retornado com sucesso

---

### POST /api/echo

**Descrição:** Retorna o texto enviado no corpo da requisição

#### Cenários de Sucesso

- **200**: Texto retornado com sucesso

#### Cenários de Erro

- **400**: Corpo da requisição vazio ou inválido

- **413**: Texto muito grande (máximo 1000 caracteres)

---
