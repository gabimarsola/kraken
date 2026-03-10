# Sistema de Extração Multi-Formato

O Kraken agora suporta **múltiplas formas de documentar endpoints**, sem depender exclusivamente de comentários customizados.

## 🎯 Extratores Implementados

### 1. **GoCommentExtractor** - Comentários Customizados Go
Extrai endpoints documentados com anotações `@Endpoint`:

```go
// @Endpoint GET /api/users
// @Description Lista todos os usuários
// @Success 200 Lista retornada com sucesso
// @Error 404 Usuários não encontrados
func GetUsers() {}
```

### 2. **GoSwaggerExtractor** - Swagger/Swaggo para Go
Extrai endpoints documentados com anotações Swagger padrão:

```go
// @Summary Lista usuários
// @Description Retorna lista paginada de usuários
// @Success 200 {array} User
// @Failure 500 {object} Error
// @Router /api/users [get]
func GetUsers() {}
```

### 3. **GoRouterExtractor** - Padrões de Roteamento Go
Detecta automaticamente rotas definidas em código (Gorilla Mux, Chi, Gin, Echo):

```go
router.HandleFunc("/api/users", GetUsers).Methods("GET")
r.Get("/api/users", GetUsers)
router.GET("/api/users", GetUsers)
e.GET("/api/users", GetUsers)
```

### 4. **JSExpressExtractor** - Express.js Routes
Detecta rotas Express/Node.js automaticamente:

```javascript
app.get('/api/users', handler)
router.post('/api/users', handler)
app.put('/api/users/:id', handler)
```

### 5. **JSDocExtractor** - JSDoc API Documentation
Extrai documentação JSDoc com anotações `@api`:

```javascript
/**
 * @api {get} /api/users Lista usuários
 * @apiSuccess (200) Lista retornada com sucesso
 * @apiError (404) Usuários não encontrados
 */
function getUsers() {}
```

### 6. **JavaSpringExtractor** - Spring Annotations
Detecta endpoints Spring Boot e JAX-RS:

```java
@GetMapping("/api/users")
public List<User> getUsers() {}

@PostMapping("/api/users")
public User createUser(@RequestBody User user) {}

// JAX-RS
@GET
@Path("/api/users")
public Response getUsers() {}
```

### 7. **JavaSwaggerExtractor** - Swagger Annotations Java
Extrai documentação Swagger/OpenAPI em Java:

```java
@Operation(summary = "Lista usuários")
@ApiResponses({
    @ApiResponse(responseCode = "200", description = "Sucesso"),
    @ApiResponse(responseCode = "404", description = "Não encontrado")
})
@GetMapping("/api/users")
public List<User> getUsers() {}
```

### 8. **SwaggerYAMLExtractor** - Arquivos Swagger/OpenAPI
Lê arquivos `swagger.yaml`, `openapi.yaml`:

```yaml
paths:
  /api/users:
    get:
      summary: Lista usuários
      responses:
        '200':
          description: Sucesso
        '404':
          description: Não encontrado
```

## 🔧 Como Funciona

### Arquitetura

```
AnalyzeEndpoints()
    ↓
ExtractorRegistry (registra todos os extratores)
    ↓
Para cada arquivo no projeto:
    ↓
    Testa qual extrator suporta o arquivo
    ↓
    Aplica TODOS os extratores compatíveis
    ↓
    Remove duplicatas (mesmo METHOD + PATH)
    ↓
Retorna lista unificada de endpoints
```

### Sistema de Prioridade

O sistema aplica **todos os extratores compatíveis** e remove duplicatas. Isso significa:

- Um arquivo `.go` pode ter endpoints detectados por:
  - GoCommentExtractor (comentários `@Endpoint`)
  - GoSwaggerExtractor (comentários Swagger)
  - GoRouterExtractor (código de roteamento)

- Um arquivo `.js` pode ter endpoints detectados por:
  - JSExpressExtractor (código Express)
  - JSDocExtractor (comentários JSDoc)

- Um arquivo `.java` pode ter endpoints detectados por:
  - JavaSpringExtractor (anotações Spring)
  - JavaSwaggerExtractor (anotações Swagger)

### Deduplicação

Endpoints com mesmo **METHOD + PATH** são considerados duplicatas e apenas um é mantido.

## 📝 Exemplos Práticos

Veja a pasta `examples/` para exemplos completos de cada formato:

- `example_go_router.go` - Rotas Gorilla Mux
- `example_go_swagger.go` - Swagger em Go
- `example_express.js` - Express.js
- `example_jsdoc.js` - JSDoc
- `example_spring.java` - Spring Boot
- `swagger.yaml` - OpenAPI YAML

## 🚀 Uso

Execute o Kraken normalmente:

```bash
./kraken
```

O sistema irá:
1. Detectar o tipo de projeto (Go, Node.js, Java)
2. Aplicar **todos os extratores relevantes**
3. Combinar endpoints de múltiplas fontes
4. Gerar README.md unificado

## ✨ Vantagens

### Flexibilidade Total
- Não precisa mudar seu código existente
- Suporta múltiplos padrões simultaneamente
- Detecta endpoints mesmo sem documentação explícita

### Compatibilidade
- Funciona com frameworks populares (Express, Spring, Gin, Echo, etc)
- Suporta padrões de documentação estabelecidos (Swagger, JSDoc)
- Detecta rotas definidas em código

### Extensibilidade
- Fácil adicionar novos extratores
- Interface `Extractor` simples de implementar
- Sistema de registro modular

## 🔌 Adicionando Novos Extratores

Para adicionar suporte a um novo formato:

1. Crie um novo arquivo em `src/extractor/`
2. Implemente a interface `Extractor`:

```go
type MyExtractor struct{}

func (e *MyExtractor) Supports(filePath string) bool {
    // Retorna true se o extrator suporta este arquivo
    return strings.HasSuffix(filePath, ".ext")
}

func (e *MyExtractor) Extract(filePath string, content []byte) ([]structure.Endpoint, error) {
    // Lógica de extração
    var endpoints []structure.Endpoint
    // ... processar content ...
    return endpoints, nil
}
```

3. Registre no `analyzer/endpoint_analyzer.go`:

```go
registry.Register(&extractor.MyExtractor{})
```

## 📊 Formatos de Saída

Independente do formato de entrada, a saída é sempre consistente:

```markdown
# Projeto: nome
Tipo: Go|Node.js|Java/Maven
Versão: x.x.x

## API Endpoints

### METHOD /path
**Descrição:** descrição

#### Cenários de Sucesso
- **CODE**: descrição

#### Cenários de Erro
- **CODE**: descrição
```

## 🎯 Casos de Uso

### Projeto Legado
Detecta endpoints automaticamente mesmo sem documentação:
- Rotas Express em código
- Anotações Spring
- Definições de router

### Projeto Documentado
Extrai documentação rica de:
- Swagger/OpenAPI
- JSDoc
- Comentários customizados

### Projeto Misto
Combina múltiplas fontes:
- Endpoints de código + documentação Swagger
- Rotas Express + JSDoc
- Spring annotations + Swagger

## 🔍 Filtros Automáticos

O sistema ignora automaticamente:
- `node_modules/`
- `vendor/`
- `.git/`

Isso evita processar dependências e melhora a performance.
