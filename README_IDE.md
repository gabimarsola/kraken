# 🚀 Kraken com IA da IDE - Guia Rápido

## ✨ Novidade: Geração Automática de PRDs com IA da IDE!

O Kraken agora detecta automaticamente sua IDE (Windsurf, Cursor, VS Code) e usa a IA integrada para gerar PRDs completos sem necessidade de configuração!

### 🎯 Principais Vantagens

- **Zero Configuração**: Funciona imediatamente após detectar a IDE
- **100% Gratuito**: Sem API keys, sem custos
- **Contexto Real**: Usa o entendimento do seu projeto pela IDE
- **Privacidade**: Dados não saem do seu ambiente

### 🖥️ IDEs Suportadas

| IDE | Status | IA Integrada | Setup |
|-----|--------|--------------|-------|
| **Windsurf** | ✅ | Nativa | Automático |
| **Cursor** | ✅ | Nativa | Automático |
| **VS Code** | ⚠️ | Extensões | Requer Copilot |

### 🚀 Como Usar (3 Passos)

#### 1. Abra seu projeto na IDE
```bash
# Seu projeto já aberto no Windsurf/Cursor/VS Code
# O diretório da IDE (.windsurf, .cursor, .vscode) é criado automaticamente
```

#### 2. Execute o Kraken
```bash
./kraken
```

#### 3. Escolha "Criar PRD com IA da IDE"
```
🚀 Kraken - Menu Principal
========================
1. Gerar documentação de endpoints
2. Criar PRD interativo
3. Criar PRD com IA externa
4. Criar PRD com IA da IDE  ← ESCOLHA ESTA
5. Configurar provedor de IA
6. Sair
```

### 📋 Exemplo de Resultado

```bash
🤖 Detectado IDE: Windsurf
📍 Workspace: /home/user/project
🚀 Gerando PRD com IA integrada da IDE...
✅ PRD gerado com sucesso usando IA do Windsurf!
📄 Arquivo criado: tasks/prd-sistema-de-autenticacao.md
```

### 🎯 O que é Gerado

O PRD inclui automaticamente:
- ✅ Título e introdução contextual
- ✅ Objetivos principais do projeto
- ✅ Histórias de usuário com critérios de aceitação
- ✅ Requisitos funcionais priorizados
- ✅ Considerações de design e técnicas
- ✅ Métricas de sucesso
- ✅ Questões importantes a resolver

### 🔧 Se a IDE Não For Detectada

```bash
❌ Erro: IDE não detectada ou não suportada
💡 Solução rápida:
mkdir .windsurf  # Para Windsurf
# ou
mkdir .cursor    # Para Cursor
# ou  
mkdir .vscode    # Para VS Code
```

### 🆚 Comparação: IA da IDE vs IA Externa

| Característica | IA da IDE | IA Externa |
|----------------|-----------|------------|
| **Setup** | ✅ Automático | ❌ Requer API keys |
| **Custo** | ✅ Gratuito | ❌ Pago |
| **Contexto** | ✅ Projeto real | ⚠️ Genérico |
| **Privacidade** | ✅ Local | ❌ Externo |

### 🎖️ Por que Usar IA da IDE?

1. **Contexto Real**: A IDE já entende seu código, dependências e padrões
2. **Zero Configuração**: Sem cadastros, sem API keys, sem custos
3. **Privacidade**: Seus dados ficam no seu ambiente
4. **Velocidade**: Processamento local e imediato
5. **Qualidade**: IA treinada no contexto do seu projeto

### 📖 Documentação Completa

- 📖 [Guia Detalhado de Integração](IDE_INTEGRATION.md)
- 🤖 [Integração com LLMs Externos](AI_INTEGRATION.md)
- 📋 [Estrutura de PRDs](docs/prd-structure.md)

### 🎉 Comece Agora!

```bash
# 1. Abra seu projeto no Windsurf/Cursor
# 2. Execute:
./kraken
# 3. Escolha opção 4
# 4. Pronto! PRD gerado automaticamente 🚀
```

**Transforme seu desenvolvimento com IA contextual e sem esforço!** 🎯
