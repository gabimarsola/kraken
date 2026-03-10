# 🤖 Integração com LLMs para Geração Automática de PRDs

## Visão Geral

O Kraken agora suporta geração automática de PRDs (Product Requirements Documents) usando diversos modelos de linguagem (LLMs), eliminando a necessidade de input manual do usuário.

## Provedores Suportados

### 1. OpenAI (GPT-4)
- **Modelo**: GPT-4
- **Requisito**: API Key da OpenAI
- **Setup**: 
  ```bash
  export OPENAI_API_KEY="sua-api-key"
  # ou configure via menu do Kraken
  ```

### 2. Anthropic (Claude)
- **Modelo**: Claude 3 Sonnet
- **Requisito**: API Key da Anthropic
- **Setup**:
  ```bash
  export ANTHROPIC_API_KEY="sua-api-key"
  # ou configure via menu do Kraken
  ```

### 3. Ollama (Local)
- **Modelos**: Llama2, CodeLlama, Mistral, etc.
- **Requisito**: Ollama instalado localmente
- **Setup**:
  ```bash
  # Instalar Ollama
  curl -fsSL https://ollama.ai/install.sh | sh
  
  # Baixar modelo
  ollama pull llama2
  
  # Iniciar servidor
  ollama serve
  ```

### 4. Google Gemini
- **Modelo**: Gemini Pro
- **Requisito**: API Key do Google AI
- **Setup**:
  ```bash
  export GEMINI_API_KEY="sua-api-key"
  # ou configure via menu do Kraken
  ```

## Como Usar

### 1. Configurar Provedor de IA
```bash
./kraken
# Escolha opção 4: Configurar provedor de IA
# Selecione o provedor desejado
# Informe as credenciais necessárias
```

### 2. Gerar PRD Automático
```bash
./kraken
# Escolha opção 3: Criar PRD com IA
# O sistema analisará o projeto e gerará o PRD automaticamente
```

## Fluxo de Geração Automática

1. **Análise do Projeto**: O Kraken extrai informações do projeto (nome, tipo, endpoints)
2. **Envio para IA**: Envia dados estruturados para o LLM escolhido
3. **Processamento**: A IA analisa e gera um PRD completo
4. **Formatação**: O sistema converte a resposta para o formato padrão
5. **Geração de Arquivo**: Salva o PRD em `tasks/prd-nome-da-funcionalidade.md`

## Estrutura do PRD Gerado

O LLM é instruído a gerar PRDs com todas as seções padrão:

- **Título e Introdução**
- **Objetivos Principais**
- **Histórias de Usuário** (com critérios de aceitação)
- **Requisitos Funcionais** (com prioridades)
- **Fora do Escopo**
- **Considerações de Design**
- **Considerações Técnicas**
- **Métricas de Sucesso**
- **Questões Abertas**

## Vantagens da Geração com IA

### 🚀 Rapidez
- PRDs gerados em segundos vs horas manuais
- Sem necessidade de digitação extensiva

### 🎯 Consistência
- Formato padronizado sempre
- Qualidade consistente entre documentos

### 🧠 Inteligência
- Análise inteligente dos endpoints existentes
- Recomendações baseadas nas melhores práticas

### 📈 Escalabilidade
- Gere múltiplos PRDs rapidamente
- Ideal para projetos com muitas funcionalidades

## Configuração Avançada

### Variáveis de Ambiente
```bash
# OpenAI
export OPENAI_API_KEY="sk-..."

# Anthropic
export ANTHROPIC_API_KEY="sk-ant-..."

# Gemini
export GEMINI_API_KEY="AIza..."
```

### Arquivo de Configuração
O sistema salva a configuração em `ai_config.json`:
```json
{
  "provider": "openai",
  "config": {
    "api_key": "sua-api-key"
  }
}
```

## Exemplo de Uso com Ollama (Gratuito)

### 1. Instalar e Configurar Ollama
```bash
# Instalar
curl -fsSL https://ollama.ai/install.sh | sh

# Baixar modelo
ollama pull llama2

# Iniciar servidor
ollama serve
```

### 2. Configurar no Kraken
```bash
./kraken
# Opção 4 → Provedor 3 (Ollama)
# URL: http://localhost:11434
# Modelo: llama2
```

### 3. Gerar PRD
```bash
./kraken
# Opção 3 → Criar PRD com IA
```

## Troubleshooting

### Erro: "API key não encontrada"
- Verifique se a variável de ambiente está configurada
- Use a opção 4 do menu para configurar manualmente

### Erro: "Conexão recusada" (Ollama)
- Verifique se o Ollama está rodando: `ollama serve`
- Confirme a URL: `http://localhost:11434`

### Resposta vazia ou incompleta
- Tente novamente (às vezes a API pode falhar)
- Verifique os limites de taxa do provedor
- Considere usar um modelo diferente

## Melhores Práticas

### 1. Projetos Bem Estruturados
- Mantenha endpoints bem documentados
- Use nomes descritivos nas funções
- Organize o código de forma lógica

### 2. Revisão Humana
- Sempre revise o PRD gerado
- Ajuste conforme necessário
- Adicione detalhes específicos do negócio

### 3. Iteração
- Use a geração como ponto de partida
- Refine com conhecimento específico
- Combine geração automática com input manual

## Comparação de Provedores

| Provedor | Custo | Qualidade | Velocidade | Setup |
|----------|-------|-----------|------------|-------|
| OpenAI | $$ | Alta | Rápida | Médio |
| Anthropic | $$ | Alta | Rápida | Médio |
| Ollama | Gratuito | Média | Média | Fácil |
| Gemini | $ | Alta | Rápida | Médio |

## Futuras Melhorias

- [ ] Suporte para mais modelos
- [ ] Geração de documentação técnica completa
- [ ] Integração com GitHub Issues
- [ ] Tradução automática de PRDs
- [ ] Análise de similaridade entre PRDs
