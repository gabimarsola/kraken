# 🐙 Kraken - Gerador de Documentação e PRDs Contextualizados

**Ferramenta automatizada para análise de projetos e geração de documentação técnica, PRDs profissionais e análise Git completa**

[![Go Version](https://img.shields.io/badge/Go-1.25+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

## 📋 Descrição

O Kraken é uma ferramenta poderosa que analisa projetos de software e gera automaticamente:

- 📝 **Documentação técnica de APIs** integrada ao documento de versão
- 📋 **PRDs Profissionais** em português brasileiro com formato de engenharia
- 🌿 **Análise Git completa** (commits, branches, alterações, merges)
- 🎫 **Sistema de Tickets** (DIST/FEAT/BE + numeração automática)
- 🤖 **Conteúdo otimizado para desenvolvedores júnior**
- 🔍 **Detecção automática** de tipo de projeto
- 🧠 **Integração com múltiplas IAs** (OpenAI, Anthropic, Ollama, Gemini)
- 💻 **Compatibilidade com IDEs** (Windsurf, Cursor, VS Code, IntelliJ)

## ✨ Funcionalidades

### 🎯 **Geração Contextualizada de Documentos**

- **Comando único** gera PRD profissional + documentação de versão completa
- **Conteúdo em português brasileiro** otimizado para desenvolvedores júnior
- **Análise Git integrada** com contexto de commits e alterações
- **Saída unificada** em `docs/kraken/`

### � **PRDs Profissionais (Novo Formato)**

- **Formato de engenharia**: Baseado em padrões como ClickBus Platform
- **Sistema de Tickets**: DIST/FEAT/BE + numeração automática
- **9 seções focadas**: Visão Geral, Problema, Objetivos, Requisitos, etc.
- **Português brasileiro**: Todo conteúdo em pt-BR
- **Contexto real**: Baseado no estado atual do projeto

### 🌿 **Análise Git Completa**

- **Branch atual**: Identifica branch em uso
- **Último commit**: Hash e mensagem detalhada
- **Último merge**: Informa merge mais recente
- **Arquivos alterados**: Lista mudanças no último commit
- **Arquivos não commitados**: Detecta mudanças pendentes
- **Contexto real**: Baseado no histórico real do projeto

### � **Documentação de Endpoints Integrada**

- Detecta automaticamente projetos (Go, Node.js, Java/Maven)
- **Integrada ao documento de versão** (não mais arquivo separado)
- Inclui exemplos, códigos de status, tratamento de erros
- Formato Markdown organizado

### 🤖 **Integração com IA**

- **IA Externa**: OpenAI, Anthropic, Ollama, Gemini
- **IA da IDE**: Windsurf, Cursor, VS Code, IntelliJ
- **Prompts especializados** para PRDs profissionais em pt-BR
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
# MÉTODO 1: Copiar executável para o projeto (RECOMENDADO)

# 1. Build do Kraken
make build

# 2. Copiar para o projeto que será analisado
cp ./kraken /caminho/do/seu/projeto/

# 3. Executar no diretório do projeto
cd /caminho/do/seu/projeto/
./kraken

# --------------------------------------------------

# MÉTODO 2: Executar diretamente do repositório Kraken

# 1. Build do Kraken
make build

# 2. Executar no diretório do projeto (usando caminho completo)
/caminho/do/kraken/./kraken

# --------------------------------------------------

# MÉTODO 3: Instalar globalmente (opcional)

# 1. Build do Kraken
make build

# 2. Instalar no sistema
make install

# 3. Executar de qualquer diretório
kraken
```

### 📍 **Como Usar em Seus Projetos (Passo a Passo)**

```bash
# 1. Clone ou build o Kraken
git clone https://github.com/gabimarsola/kraken.git
cd kraken
make build

# 2. Copie o executável para seu projeto
cp ./kraken /caminho/do/seu/projeto/

# 3. Vá para o diretório do seu projeto
cd /caminho/do/seu/projeto/

# 4. Execute o Kraken
./kraken

# 5. Encontre os resultados em docs/kraken/
ls docs/kraken/
# DIST-1234-prd.md
# versao100-doc.md
```

### ⚠️ **Importante**

- **O executável `kraken` deve estar no diretório raiz** do projeto que será analisado
- **Não execute o Kraken de dentro do próprio repositório Kraken** (a menos que queira analisar o Kraken)
- **O Kraken analisa o diretório atual** onde está sendo executado

### 🎯 **Menu Principal**

```
🚀 Kraken - Menu Principal
========================
1. Gerar documentos com IA externa
2. Gerar documentos com IA da IDE
3. Configurar provedor de IA
4. Sair
```

### 🎯 **Fluxo de Uso**

1. **Execute o Kraken**: `./kraken` no diretório do projeto
2. **Escolha a opção 1 ou 2** para gerar documentos
3. **Aguarde o processamento** (análise + geração contextualizada)
4. **Encontre os resultados** em `docs/kraken/`

## 📁 **Estrutura de Saída (Novo Formato)**

```
docs/kraken/
├── DIST-1234-prd.md              # PRD profissional (pt-BR)
└── versao[versao]-doc.md          # Versão + API + Git integrados
```

### 📋 **Estrutura da PRD (9 seções)**

```markdown
# DIST-1234 - Documento de Requisitos do Produto

**Projeto:** Nome do Projeto
**Status:** Draft
**Prioridade:** Medium
**Criado:** 18/03/2026
**Autor:** Equipe de Engenharia

## Visão Geral

## Declaração do Problema

## Objetivos e Metas

## Requisitos Funcionais

## Requisitos Não Funcionais

## Considerações Técnicas

## Dependências

## Riscos e Mitigações
```

### 🌿 **Contexto Git no Documento de Versão**

```markdown
## 🌿 Contexto Git

**Branch Atual:** main
**Último Commit:** d0a4240 - formatação de documento
**Último Merge:** Merge branch 'feature/new-api'

### Arquivos Alterados no Último Commit

- src/prd/contextual_generator.go
- cli/cli.go
- src/prd/prd_template.go
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

| Tipo        | Arquivo de Configuração | Prefixo Ticket | Exemplos                 |
| ----------- | ----------------------- | -------------- | ------------------------ |
| **Go**      | `go.mod`                | `BE`           | Gin, Echo, Fiber         |
| **Node.js** | `package.json`          | `FEAT`         | Express, NestJS, Fastify |
| **Java**    | `pom.xml`               | `DIST`         | Spring Boot, JAX-RS      |

## 🎯 **Exemplo de Uso Completo**

```bash
# 1. Build do Kraken (no repositório do Kraken)
cd /home/user/kraken
make build

# 2. Copie para seu projeto (EXEMPLO: projeto Node.js)
cp ./kraken /home/user/meu-projeto-node/

# 3. Vá para o seu projeto
cd /home/user/meu-projeto-node/

# 4. Execute o Kraken (dentro do seu projeto)
./kraken

# 5. Escolher opção 1 (IA externa) ou 2 (IA da IDE)

# 6. Resultados em docs/kraken/ do seu projeto
ls docs/kraken/
# DIST-1234-prd.md
# versao100-doc.md
```

### 📁 **Estrutura Final do Seu Projeto**

```
meu-projeto-node/
├── src/
├── package.json
├── node_modules/
├── kraken                    # ← Executável do Kraken
└── docs/
    └── kraken/
        ├── DIST-1234-prd.md      # ← PRD gerado
        └── versao100-doc.md       # ← Versão + API + Git
```

## 🔄 **Mudanças Recentes (v2.0)**

### ✅ **Novas Funcionalidades**

- 🎫 **Sistema de Tickets**: DIST/FEAT/BE + numeração automática
- 🇧🇷 **Português Brasileiro**: PRDs 100% em pt-BR
- 🌿 **Análise Git**: Contexto completo de commits e alterações
- 📋 **Formato Profissional**: Baseado em padrões de engenharia
- 📝 **Documentação Integrada**: API dentro do documento de versão
- 🎯 **9 Seções Focadas**: Removidas seções desnecessárias

### 🔄 **Mudanças na Estrutura**

- ❌ **Removido**: Arquivo separado de API
- ❌ **Removido**: Histórias de Usuário, Fora do Escopo, Cronograma
- ✅ **Adicionado**: Contexto Git completo
- ✅ **Adicionado**: Sistema de tickets profissional
- ✅ **Adicionado**: Formato em português brasileiro

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

### Erro: "slice bounds out of range"

```bash
# Isso foi corrigido! Recompile o projeto:
make build
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
- Formato profissional de engenharia de software

---

**🐙 Kraken - Transformando código em documentação profissional contextualizada!**
