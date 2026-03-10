# 🔍 Integração com Git - PRDs Baseados em Alterações

## Visão Geral

O Kraken agora pode analisar automaticamente as alterações do seu projeto Git (arquivos modificados ou último commit) e gerar PRDs contextualizados baseados no que você está desenvolvendo.

## 🚀 Funcionalidades

### 📊 Análise Inteligente de Alterações
- **Detecta arquivos modificados**: Identifica todos os arquivos alterados no working directory
- **Analisa último commit**: Se não há alterações pendentes, analisa o último commit
- **Ignora arquivos irrelevantes**: Filtra automaticamente arquivos que não afetam o PRD
- **Extrai informações do código**: Identifica funções, classes e estruturas importantes

### 🎯 Geração Contextualizada
- **Títulos inteligentes**: Gera títulos baseados nos padrões dos arquivos alterados
- **Histórias de usuário específicas**: Cria histórias baseadas nas funcionalidades implementadas
- **Requisitos funcionais**: Define requisitos baseados no código adicionado/modificado
- **Considerações técnicas**: Identifica aspectos técnicos relevantes das alterações

## 🌟 Como Usar

### 1. Faça suas alterações no projeto
```bash
# Exemplo: Adicionar nova funcionalidade
echo 'func NewEndpoint() { ... }' >> src/api/handler.go
git add src/api/handler.go
```

### 2. Execute o Kraken
```bash
./kraken
```

### 3. Escolha a opção 5
```
🚀 Kraken - Menu Principal
========================
1. Gerar documentação de endpoints
2. Criar PRD interativo
3. Criar PRD com IA externa
4. Criar PRD com IA da IDE
5. Criar PRD baseado em alterações Git  ← ESCOLHA ESTA
6. Configurar provedor de IA
7. Sair
```

### 4. PRD Gerado Automaticamente
```bash
🔍 Analisando alterações do projeto...
📊 Encontradas 3 alterações
🌿 Branch: feature/new-api
⚠️ Analisando alterações não commitadas
✅ PRD gerado com sucesso!
📄 Arquivo criado: tasks/prd-sistema-de-api-branch-feature-new-api.md
```

## 📋 Exemplo de Análise

### Entrada: Arquivos Modificados
```
📁 src/api/auth.go (Modified)
   🔧 Funções: LoginUser, ValidateToken, RefreshToken
   📏 45 linhas

📁 src/models/user.go (Added)
   🔧 Funções: User, CreateUser, ValidateUser
   📏 28 linhas

📁 config/database.go (Modified)
   📏 12 linhas
```

### Saída: PRD Contextualizado
```markdown
# Sistema de API e Autenticação

## 📋 Introdução/Visão Geral
Desenvolvimento em andamento com 3 arquivos modificados. Esta atualização foca em implementar as funcionalidades identificadas nas alterações recentes do projeto.

## 🎯 Objetivos
- Implementar as alterações identificadas no código
- Manter compatibilidade com o sistema existente
- Garantir qualidade e performance do código
- Melhorar segurança da autenticação
- Expandir funcionalidades da API

## 📖 Histórias do Usuário

### US001: Implementação de auth.go
Como usuário do sistema, eu quero que as funcionalidades do arquivo src/api/auth.go sejam implementadas/atualizadas para melhorar a experiência e funcionalidade do sistema.

#### Critérios de Aceitação
- O código deve compilar sem erros
- As funcionalidades devem funcionar conforme esperado
- O código deve seguir os padrões do projeto
- Todos os testes devem passar
- Cobertura de testes adequada

[...]
```

## 🔍 Detalhes da Análise Git

### O que é Analisado

#### ✅ Arquivos Incluídos
- **Código fonte**: `.go`, `.js`, `.ts`, `.py`, `.java`, etc.
- **Configurações**: Arquivos de configuração relevantes
- **Documentação**: Arquivos de documentação técnicos
- **Testes**: Arquivos de teste (com prioridade especial)

#### ❌ Arquivos Ignorados
- **Versionamento**: `.git/`, `.gitignore`
- **Dependências**: `node_modules/`, `vendor/`
- **IDE**: `.vscode/`, `.windsurf/`, `.cursor/`
- **Build**: `dist/`, `build/`, `*.exe`
- **Logs**: `*.log`, `*.tmp`
- **READMEs**: `README.md`, arquivos `.md` gerais

### Tipos de Alterações Detectadas

| Status | Descrição | Tratamento |
|--------|-----------|------------|
| **Added** | Arquivos novos | Análise completa do conteúdo |
| **Modified** | Arquivos modificados | Comparação e análise |
| **Deleted** | Arquivos removidos | Registro da remoção |
| **Renamed** | Arquivos renomeados | Tratamento como novo |
| **Untracked** | Não commitados | Incluídos na análise |

### Informações Extraídas

#### Para Arquivos Go
```go
// Detecta automaticamente:
func LoginUser(user User) (token string, err error) { ... }
func ValidateToken(token string) (valid bool) { ... }
func RefreshToken(oldToken string) (newToken string, err error) { ... }
```

#### Para Outras Linguagens
- **JavaScript/TypeScript**: Funções, classes, exports
- **Python**: Funções, classes, métodos
- **Java**: Classes, métodos, interfaces

## 🎯 Geração Inteligente

### Títulos Contextuais
O sistema analisa os padrões nos arquivos para gerar títulos relevantes:

| Padrões Detectados | Título Gerado |
|-------------------|---------------|
| `auth`, `login`, `user` | "Sistema de Autenticação" |
| `api`, `endpoint`, `router` | "Sistema de API" |
| `db`, `database`, `model` | "Sistema de Dados" |
| `ui`, `frontend`, `view` | "Sistema de Interface" |

### Prioridades Automáticas
- **Alta**: Arquivos principais (`main.go`, arquivos de API)
- **Média**: Arquivos de lógica de negócio
- **Baixa**: Testes, configuração, documentação

### Histórias de Usuario Contextuais
Cada arquivo modificado gera uma história de usuário específica:
- **Título**: Baseado no nome do arquivo
- **Descrição**: Focada na funcionalidade implementada
- **Critérios**: Específicos para o tipo de alteração

## 🔄 Fluxos de Trabalho

### Fluxo 1: Desenvolvimento Ativo
1. **Desenvolver**: Faça alterações no código
2. **Analisar**: Execute opção 5 do Kraken
3. **PRD**: Gerado baseado nas alterações atuais
4. **Refinar**: Ajuste o PRD se necessário
5. **Commit**: Commit com PRD documentado

### Fluxo 2: Pós-Commit
1. **Commit**: Faça commit das alterações
2. **Analisar**: Execute opção 5 (detectará último commit)
3. **PRD**: Gerado baseado no commit
4. **Documentar**: Use PRD para documentação da release

### Fluxo 3: Code Review
1. **PR**: Crie Pull Request
2. **Analisar**: Gere PRD das alterações
3. **Review**: Use PRD como base para review
4. **Aprovar**: Documente decisões no PRD

## 📊 Exemplos Práticos

### Exemplo 1: Nova API Endpoint
```bash
# Desenvolvedor adiciona novo endpoint
echo 'func GetUserProfile(c *gin.Context) { ... }' >> src/api/user.go
git add src/api/user.go

# Gerar PRD
./kraken
# Opção 5

# Resultado:
# 📄 tasks/prd-sistema-de-api-branch-feature-user-profile.md
# - História de usuário para GetUserProfile
# - Requisitos funcionais da API
# - Considerações de segurança
```

### Exemplo 2: Refatoração
```bash
# Desenvolvedor refatora código
mv src/old_handler.go src/new_handler.go
git add -A

# Gerar PRD
./kraken
# Opção 5

# Resultado:
# 📄 tasks/prd-atualizacao-do-sistema-branch-refactor-handlers.md
# - História sobre refatoração
# - Requisitos de compatibilidade
# - Considerações de migração
```

### Exemplo 3: Correção de Bugs
```bash
# Desenvolvedor corrige bug
echo '// Fix: Added validation for empty email' >> src/auth.go
git add src/auth.go

# Gerar PRD
./kraken
# Opção 5

# Resultado:
# 📄 tasks/prd-sistema-de-autenticacao-branch-fix-email-validation.md
# - História sobre correção do bug
# - Requisitos de validação
# - Critérios de teste
```

## ⚙️ Configuração Avançada

### Personalizar Arquivos Ignorados
```go
// Em src/analyzer/git_analyzer.go
func (g *GitAnalyzer) shouldIgnoreFile(filePath string) bool {
    ignorePatterns := []string{
        // Adicione seus padrões personalizados aqui
        "seu/padrão/personalizado",
    }
    // ...
}
```

### Personalizar Análise de Código
```go
// Adicione extratores para outras linguagens
func (g *GitPRDGenerator) extractPythonFunctions(content string) []string {
    // Implementação para Python
}
```

## 🔧 Troubleshooting

### Erro: "não é um repositório git"
```bash
❌ Erro ao analisar Git: não é um repositório git
```
**Solução**: Inicialize o repositório Git
```bash
git init
git add .
git commit -m "Initial commit"
```

### Erro: "Nenhuma alteração encontrada"
```bash
ℹ️ Nenhuma alteração encontrada para analisar
```
**Solução**: Faça alterações ou verifique se há commits recentes
```bash
# Verificar status
git status

# Verificar commits
git log --oneline -5
```

### Apenas arquivos irrelevantes detectados
**Solução**: Verifique se os arquivos estão sendo ignorados incorretamente
```bash
# Verificar .gitignore
cat .gitignore

# Verificar se há arquivos de código
find . -name "*.go" -o -name "*.js" -o -name "*.py"
```

## 🎖️ Benefícios

### 🎯 Contexto Real
- **Baseado no que você está fazendo**: Não especulações, mas análise real do código
- **Relevância garantida**: Cada requisito vem de uma alteração real
- **Tempo atual**: Reflete o estado atual do desenvolvimento

### ⚡ Eficiência
- **Zero digitação**: Não precisa descrever o que fez
- **Análise automática**: Extração inteligente das informações
- **Foco no essencial**: Apenas o que realmente mudou

### 🔄 Iterativo
- **Acompanha o desenvolvimento**: Use a cada alteração significativa
- **Histórico documentado**: PRDs evoluem com o código
- **Feedback contínuo**: Documentação viva do projeto

### 🧠 Inteligente
- **Padrões reconhecidos**: Identifica tipos comuns de alterações
- **Prioridades automáticas**: Classifica importância das mudanças
- **Integração com IA**: Usa IA da IDE quando disponível

## 🚀 Dicas de Uso

### Melhores Práticas
1. **Use após alterações significativas**: Não para cada pequeno change
2. **Combine com IA da IDE**: Para análises mais profundas
3. **Revise o PRD gerado**: Ajuste conforme necessário
4. **Documente decisões**: Adicione contexto que a IA não captura
5. **Use em code reviews**: Base objetiva para discussões

### Quando Usar
- ✅ **Nova funcionalidade**: Após implementar features completas
- ✅ **Refatoração**: Antes/durante grandes refatorações
- ✅ **Bug fixes**: Para documentar correções importantes
- ✅ **Code reviews**: Como base para discussões técnicas
- ✅ **Documentação**: Para gerar docs de release

### Quando Não Usar
- ❌ **Alterações triviais**: Comentários, formatação
- ❌ **Arquivos de configuração**: Mudanças simples de config
- ❌ **Dependências**: Atualizações de packages
- ❌ **Documentação**: Apenas mudanças em README

## 📈 Futuras Melhorias

- [ ] Suporte para mais linguagens (Rust, Kotlin, Swift)
- [ ] Análise de dependências entre arquivos
- [ ] Detecção automática de breaking changes
- [ ] Integração com sistemas de CI/CD
- [ ] Comparação entre branches
- [ ] Análise de métricas de código
- [ ] Geração automática de testes baseados no PRD
