# 🚀 Guia Rápido de Uso - Kraken

## ⚡ Start Rápido

```bash
# 1. Build
make build

# 2. Executar
./kraken

# 3. Escolher opção 1 ou 2
# 4. Encontrar resultados em docs/kraken/
```

## 📋 Menu Principal

| Opção | Descrição | Quando Usar |
|-------|-----------|-------------|
| **1** | Gerar documentos com IA externa | Tem API Key (OpenAI, Anthropic, etc.) |
| **2** | Gerar documentos com IA da IDE | Usa Windsurf, Cursor, VS Code ou IntelliJ |
| **3** | Configurar provedor de IA | Primeira vez ou mudar de provedor |
| **4** | Sair | Encerrar o programa |

## 🎯 Fluxo Completo

### Passo 1: Build
```bash
make build
```

### Passo 2: Executar
```bash
./kraken
```

### Passo 3: Escolher Opção
- **Opção 1**: IA externa (requer configuração prévia)
- **Opção 2**: IA da IDE (detecta automaticamente)

### Passo 4: Aguardar Processamento
```
🔍 Detectando tipo de projeto...
✅ Projeto detectado: go
📖 Parseando informações do projeto...
📋 Gerando documentação...
🤖 Processando com IA...
✅ Documentos gerados com sucesso!
```

### Passo 5: Ver Resultados
```bash
ls docs/kraken/
# API-DOC.md
# prd-nome-projeto.md
```

## ⚙️ Configuração de IA

### Primeira Vez
1. Execute `./kraken`
2. Escolha **opção 3**
3. Selecione o provedor:
   - OpenAI (necessita API Key)
   - Anthropic (necessita API Key)
   - Ollama (local)
   - Gemini (necessita API Key)
4. Siga as instruções

### IA da IDE (Sem Configuração)
- **Windsurf**: Detectado via `.windsurf/`
- **Cursor**: Detectado via `.cursor/`
- **VS Code**: Detectado via `.vscode/`
- **IntelliJ**: Detectado via `.idea/`

## 📁 Arquivos Gerados

### Documentação de Endpoints (`API-DOC.md`)
- Lista completa de endpoints
- Parâmetros e exemplos
- Códigos de status
- Tratamento de erros

### PRD (`prd-[projeto].md`)
- Objetivos do projeto
- Histórias de usuário
- Requisitos funcionais
- Métricas de sucesso

## 🔧 Problemas Comuns

### Permissão Negada
```bash
chmod +x kraken
```

### IDE Não Detectada
```bash
# Verifique diretórios:
ls -la .windsurf/ .cursor/ .vscode/ .idea/
```

### Erro de API Key
```bash
# Reconfigure:
./kraken
# Escolha opção 3
```

## 🎉 Dicas

- **Use `make build`** sempre para garantir permissões
- **Documentação unificada** em `docs/kraken/`
- **Conteúdo otimizado** para desenvolvedores júnior
- **Sempre funcione** mesmo sem IA (fallback inteligente)

---

**Pronto! Em 3 comandos você tem documentação completa! 🐙**
