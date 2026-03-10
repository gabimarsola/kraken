# 🐙 Kraken - Gerador de Documentação e PRDs

**Ferramenta automatizada para análise de projetos e geração de documentação técnica e PRDs**

[![Go Version](https://img.shields.io/badge/Go-1.25+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

## 📋 Descrição

O Kraken é uma ferramenta poderosa que analisa projetos de software e gera automaticamente:

- 📝 **Documentação técnica de APIs** com detalhes completos
- 📋 **PRDs (Product Requirements Documents)** estruturados
- 🤖 **Conteúdo otimizado para desenvolvedores júnior**
- 🔍 **Detecção automática** de tipo de projeto
- 🧠 **Integração com múltiplas IAs** (OpenAI, Anthropic, Ollama, Gemini)
- 💻 **Compatibilidade com IDEs** (Windsurf, Cursor, VS Code, IntelliJ)

## ✨ Funcionalidades

### 🎯 **Geração Unificada de Documentos**

- **Comando único** gera ambos: documentação de endpoints E PRD
- **Conteúdo otimizado** para desenvolvedores júnior (explícito, detalhado, sem jargões)
- **Saída unificada** em `docs/kraken/`

### 📝 **Documentação de Endpoints**

- Detecta automaticamente projetos (Go, Node.js, Java/Maven)
- Analisa endpoints e gera documentação completa
- Inclui exemplos, códigos de status, tratamento de erros
- Formato Markdown organizado

### 📋 **PRDs Inteligentes**

- Gerados por IA com base no contexto do projeto
- Estrutura completa: histórias de usuário, requisitos, métricas
- Foco em desenvolvedores júnior
- Formato JSON estruturado

### 🤖 **Integração com IA**

- **IA Externa**: OpenAI, Anthropic, Ollama, Gemini
- **IA da IDE**: Windsurf, Cursor, VS Code, IntelliJ
- **Prompts otimizados** para desenvolvedores júnior
- **Configuração flexível** de provedores

## 🚀 Como Usar

### Instalação e Build

```bash
# Clonar repositório
git clone https://github.com/gabimarsola/kraken.git
cd kraken

# Build (recomendado)
make build

# Ou build manual
go build -o kraken . && chmod +x kraken
```

### Execução

```bash
# Executar o Kraken
./kraken
```

### � **Menu Principal**

```
🚀 Kraken - Menu Principal
========================
1. Gerar documentos com IA externa
2. Gerar documentos com IA da IDE
3. Configurar provedor de IA
4. Sair
```

### 🎯 **Fluxo de Uso**

1. **Execute o Kraken**: `./kraken`
2. **Escolha a opção 1 ou 2** para gerar documentos
3. **Aguarde o processamento** (análise + geração)
4. **Encontre os resultados** em `docs/kraken/`

## 📁 **Estrutura de Saída**

```
docs/kraken/
├── API-DOC.md                    # Documentação de endpoints
├── prd-[nome-projeto].md         # PRD gerado
└── [outros-arquivos].md          # Documentos adicionais
```

## ⚙️ **Configuração de IA**

### Configurar Provedor Externo

```bash
# Opção 3 no menu principal
# Siga o assistente para configurar:
# - OpenAI (API Key)
# - Anthropic (API Key)
# - Ollama (local)
# - Gemini (API Key)
```

### Usar IA da IDE

- **Windsurf**: Detectado automaticamente via `.windsurf/`
- **Cursor**: Detectado automaticamente via `.cursor/`
- **VS Code**: Detectado automaticamente via `.vscode/`
- **IntelliJ**: Detectado automaticamente via `.idea/`

## 🛠️ **Comandos Úteis**

```bash
# Build
make build

# Executar
make run

# Limpar
make clean

# Instalar no sistema
make install
```

## 📊 **Tipos de Projeto Suportados**

| Tipo        | Arquivo de Configuração | Exemplos                 |
| ----------- | ----------------------- | ------------------------ |
| **Go**      | `go.mod`                | Gin, Echo, Fiber         |
| **Node.js** | `package.json`          | Express, NestJS, Fastify |
| **Java**    | `pom.xml`               | Spring Boot, JAX-RS      |

## 🎯 **Exemplo de Uso Completo**

```bash
# 1. Build
make build

# 2. Executar
./kraken

# 3. Escolher opção 1 (IA externa) ou 2 (IA da IDE)

# 4. Resultados em docs/kraken/
ls docs/kraken/
# API-DOC.md
# prd-meu-projeto.md
```

## 🔧 **Solução de Problemas**

### Permissão Negada

```bash
chmod +x kraken
```

### IDE Não Detectada

```bash
# Verifique se o diretório de configuração existe:
# .windsurf/ ou .cursor/ ou .vscode/ ou .idea/
```

### Configuração de IA

```bash
# Configure um provedor via menu (opção 3)
# Ou crie ai_config.json baseado em ai_config.example.json
```

## 📚 **Documentação Adicional**

- [Integração com IDEs](IDE_INTEGRATION.md)
- [Configuração de IA](AI_INTEGRATION.md)
- [Integração Git](GIT_INTEGRATION.md)
- [Padrões de Extração](EXTRACTION_PATTERNS.md)

## 🤝 **Contribuição**

Contribuições são bem-vindas! Por favor:

1. Fork o projeto
2. Crie uma branch (`git checkout -b feature/nova-funcionalidade`)
3. Commit suas mudanças (`git commit -m 'Adiciona nova funcionalidade'`)
4. Push para a branch (`git push origin feature/nova-funcionalidade`)
5. Abra um Pull Request

## 📄 **Licença**

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes.

## 🌟 **Créditos**

- Desenvolvido para automatizar documentação técnica
- Otimizado para times de desenvolvimento
- Foco em desenvolvedores júnior

---

**🐙 Kraken - Transformando código em documentação clara!**
