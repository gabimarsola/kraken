# 🖥️ Integração com IA da IDE (Windsurf, Cursor, VS Code)

## Visão Geral

O Kraken agora detecta automaticamente a IDE que você está usando e utiliza a IA integrada da própria IDE para gerar PRDs, eliminando a necessidade de configurar APIs externas.

## IDEs Suportadas

### 1. Windsurf ✅
- **Detecção**: Presença do diretório `.windsurf`
- **IA**: Integrada nativamente
- **Vantagens**: Sem configuração necessária, IA contextual do projeto

### 2. Cursor ✅  
- **Detecção**: Presença do diretório `.cursor`
- **IA**: Integrada nativamente
- **Vantagens**: IA especializada em desenvolvimento

### 3. VS Code ⚠️
- **Detecção**: Presença do diretório `.vscode`
- **IA**: Requer extensões (Copilot, etc.)
- **Limitações**: IA não nativa, pode precisar configuração adicional

## Como Funciona

### 1. Detecção Automática
```bash
./kraken
# O sistema detecta automaticamente:
# - Windsurf: busca .windsurf/
# - Cursor: busca .cursor/
# - VS Code: busca .vscode/
```

### 2. Uso da IA da IDE
```bash
# Menu → Opção 4: Criar PRD com IA da IDE
🤖 Detectado IDE: Windsurf
📍 Workspace: /path/to/project
🚀 Gerando PRD com IA integrada da IDE...
```

### 3. Geração Automática
- **Análise do Projeto**: Extrai informações do projeto atual
- **Contexto da IDE**: Usa o contexto e entendimento do projeto pela IDE
- **Geração de PRD**: Cria PRD completo sem input manual
- **Formatação**: Salva em `tasks/prd-nome-da-funcionalidade.md`

## Vantagens da IA da IDE

### 🎯 Contexto do Projeto
- A IDE já conhece a estrutura do seu código
- Entende as dependências e padrões utilizados
- Tem acesso ao histórico de alterações

### 🚀 Sem Configuração
- Não precisa de API keys
- Sem cadastro em serviços externos
- Funciona imediatamente após detecção

### 🔒 Privacidade
- Dados não saem do seu ambiente
- Processamento local (na maioria dos casos)
- Sem compartilhamento com terceiros

### 💰 Gratuito
- Sem custos por uso
- Sem limites de rate limiting
- Disponível imediatamente

## Exemplo de Uso com Windsurf

### 1. Ambiente Windsurf
```bash
# Seu projeto já está aberto no Windsurf
# O diretório .windsurf existe automaticamente
ls -la
# .windsurf/
# src/
# go.mod
# ...
```

### 2. Executar o Kraken
```bash
./kraken
# Menu:
# 1. Gerar documentação de endpoints
# 2. Criar PRD interativo
# 3. Criar PRD com IA externa
# 4. Criar PRD com IA da IDE  ← ESCOLHA ESTA
# 5. Configurar provedor de IA
# 6. Sair
```

### 3. Resultado
```bash
🤖 Detectado IDE: Windsurf
📍 Workspace: /home/user/project
🚀 Gerando PRD com IA integrada da IDE...
✅ PRD gerado com sucesso usando IA do Windsurf!
📄 Arquivo criado: tasks/prd-sistema-de-autenticacao.md
```

## Detecção Manual

Se a IDE não for detectada automaticamente:

### Windsurf
```bash
# Criar diretório manualmente
mkdir .windsurf
echo "windsurf" > .windsurf/ide-marker
```

### Cursor
```bash
# Criar diretório manualmente
mkdir .cursor
echo "cursor" > .cursor/ide-marker
```

### VS Code
```bash
# Criar diretório manualmente
mkdir .vscode
echo "vscode" > .vscode/ide-marker
```

## Comparação: IA da IDE vs IA Externa

| Característica | IA da IDE | IA Externa |
|----------------|-----------|------------|
| **Setup** | Automático | Requer API keys |
| **Custo** | Gratuito | Pago |
| **Contexto** | Projeto local | Genérico |
| **Privacidade** | Local | Externo |
| **Velocidade** | Rápida | Depende da API |
| **Qualidade** | Contextual | Variável |

## Troubleshooting

### IDE Não Detectada
```bash
❌ Erro ao criar gerador IDE: IDE não detectada ou não suportada
💡 Dicas para detectar a IDE:
   - Windsurf: Verifique se o diretório .windsurf existe
   - Cursor: Verifique se o diretório .cursor existe
   - VS Code: Verifique se o diretório .vscode existe
```

**Solução**:
```bash
# Para Windsurf
mkdir -p .windsurf

# Para Cursor  
mkdir -p .cursor

# Para VS Code
mkdir -p .vscode
```

### IA Não Disponível
```bash
❌ IDE VS Code não possui IA integrada disponível
```

**Solução**: Instale extensões como GitHub Copilot ou use opção de IA externa.

### Erro de Geração
```bash
❌ Erro ao gerar PRD com IA da IDE
```

**Soluções**:
- Verifique se a IDE está rodando
- Confirme se a IA está ativa na IDE
- Tente reiniciar a IDE

## Configuração Avançada

### Variáveis de Ambiente
```bash
# Opcional: forçar detecção de IDE específica
export KRAKEN_IDE=windsurf
export KRAKEN_IDE=cursor
export KRAKEN_IDE=vscode
```

### Arquivos de Configuração
```bash
# .kraken-ide.json
{
  "preferred_ide": "windsurf",
  "fallback_to_external": true,
  "ai_model": "default"
}
```

## Exemplo de PRD Gerado

O PRD gerado pela IA da IDE inclui:

```markdown
# Sistema de Autenticação

**Projeto:** kraken  
**Tipo:** Go  
**Versão:** 1.25.7

## 📋 Introdução/Visão Geral

Implementação de sistema completo de autenticação...

## 🎯 Objetivos

- Proporcionar experiência de login segura e intuitiva
- Implementar gestão completa de perfis de usuário
- Garantir segurança das informações sensíveis

## 📖 Histórias do Usuário

### US001: Cadastro de Novo Usuário
Como um novo usuário, eu quero criar uma conta...

#### Critérios de Aceitação
- Email válido e único
- Senha com mínimo 8 caracteres
- Verificação por email

[...]
```

## Melhores Práticas

### 1. Mantenha a IDE Atualizada
- Use versões recentes da IDE
- Mantenha extensões de IA atualizadas

### 2. Projeto Bem Estruturado
- Código organizado facilita o entendimento da IA
- Comentários e documentação ajudam o contexto

### 3. Revisão do PRD
- Sempre revise o PRD gerado
- Ajuste conforme necessidades específicas
- Adicione detalhes do negócio se necessário

### 4. Iteração
- Use a geração como ponto de partida
- Refine com conhecimento específico
- Combine com input manual quando preciso

## Futuras Melhorias

- [ ] Suporte para mais IDEs (JetBrains, etc.)
- [ ] Detecção automática de capacidades de IA
- [ ] Integração com múltiplas IAs simultaneamente
- [ ] Personalização de prompts por projeto
- [ ] Histórico de PRDs gerados
- [ ] Comparação entre diferentes IAs
