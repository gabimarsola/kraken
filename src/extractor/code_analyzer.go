package extractor

import (
	"regexp"
	"strings"
)

// AnalyzeMethodCode analisa o código de um método e gera um resumo do funcionamento
func AnalyzeMethodCode(lines []string, startLine int, language string) string {
	switch language {
	case "java":
		return analyzeJavaMethod(lines, startLine)
	case "typescript", "javascript":
		return analyzeJSMethod(lines, startLine)
	case "go":
		return analyzeGoMethod(lines, startLine)
	default:
		return ""
	}
}

func analyzeJavaMethod(lines []string, startLine int) string {
	var summary []string
	methodBody := extractMethodBody(lines, startLine, "{", "}")
	
	// Detectar operações de banco de dados
	if containsAny(methodBody, []string{"repository.", "Repository", ".save(", ".findBy", ".delete(", "entityManager"}) {
		summary = append(summary, "Realiza operações no banco de dados")
	}
	
	// Detectar validações
	if containsAny(methodBody, []string{"validate", "Validator", "if (", "throw new", "Exception"}) {
		summary = append(summary, "Valida dados de entrada")
	}
	
	// Detectar chamadas de serviço
	if containsAny(methodBody, []string{"service.", "Service", ".call(", ".execute("}) {
		summary = append(summary, "Chama serviços externos ou internos")
	}
	
	// Detectar transformações de dados
	if containsAny(methodBody, []string{"map(", "stream()", "collect(", "DTO", "toDto", "fromDto"}) {
		summary = append(summary, "Transforma/mapeia dados entre objetos")
	}
	
	// Detectar autenticação/autorização
	if containsAny(methodBody, []string{"Authentication", "authorize", "hasRole", "Principal", "SecurityContext"}) {
		summary = append(summary, "Verifica autenticação/autorização")
	}
	
	// Detectar retornos
	if containsAny(methodBody, []string{"return ResponseEntity.ok", "return ResponseEntity.status"}) {
		summary = append(summary, "Retorna resposta HTTP com status apropriado")
	} else if containsAny(methodBody, []string{"return", "ResponseEntity"}) {
		summary = append(summary, "Retorna dados processados")
	}
	
	if len(summary) == 0 {
		return "Processa requisição e retorna resposta"
	}
	
	return strings.Join(summary, " • ")
}

func analyzeJSMethod(lines []string, startLine int) string {
	var summary []string
	methodBody := extractMethodBody(lines, startLine, "{", "}")
	
	// Detectar operações de banco de dados
	if containsAny(methodBody, []string{"await", "find(", "findOne(", "save(", "update(", "delete(", "create(", "repository"}) {
		summary = append(summary, "Realiza operações assíncronas no banco de dados")
	}
	
	// Detectar validações
	if containsAny(methodBody, []string{"validate", "validator", "if (", "throw new", "Error"}) {
		summary = append(summary, "Valida dados de entrada")
	}
	
	// Detectar chamadas de API/serviços
	if containsAny(methodBody, []string{"fetch(", "axios.", "http.", "this.httpService", "await this."}) {
		summary = append(summary, "Chama APIs ou serviços externos")
	}
	
	// Detectar transformações de dados
	if containsAny(methodBody, []string{".map(", ".filter(", ".reduce(", "DTO", "toDto", "plainToClass"}) {
		summary = append(summary, "Transforma/mapeia dados")
	}
	
	// Detectar autenticação
	if containsAny(methodBody, []string{"@UseGuards", "AuthGuard", "JwtService", "authenticate"}) {
		summary = append(summary, "Verifica autenticação")
	}
	
	// Detectar retornos
	if containsAny(methodBody, []string{"return", "res.json", "res.send"}) {
		summary = append(summary, "Retorna dados processados")
	}
	
	if len(summary) == 0 {
		return "Processa requisição de forma assíncrona"
	}
	
	return strings.Join(summary, " • ")
}

func analyzeGoMethod(lines []string, startLine int) string {
	var summary []string
	methodBody := extractMethodBody(lines, startLine, "{", "}")
	
	// Detectar operações de banco de dados
	if containsAny(methodBody, []string{"db.", "Query(", "Exec(", "QueryRow(", "gorm"}) {
		summary = append(summary, "Realiza operações no banco de dados")
	}
	
	// Detectar validações
	if containsAny(methodBody, []string{"if err", "validate", "Validate", "errors."}) {
		summary = append(summary, "Valida dados e trata erros")
	}
	
	// Detectar chamadas HTTP
	if containsAny(methodBody, []string{"http.Get", "http.Post", "client.Do", "Request"}) {
		summary = append(summary, "Faz requisições HTTP")
	}
	
	// Detectar encoding/decoding
	if containsAny(methodBody, []string{"json.Marshal", "json.Unmarshal", "Encode", "Decode"}) {
		summary = append(summary, "Serializa/deserializa dados JSON")
	}
	
	// Detectar retornos
	if containsAny(methodBody, []string{"w.Write", "json.NewEncoder", "return"}) {
		summary = append(summary, "Retorna resposta HTTP")
	}
	
	if len(summary) == 0 {
		return "Processa requisição HTTP"
	}
	
	return strings.Join(summary, " • ")
}

// extractMethodBody extrai o corpo de um método entre delimitadores
func extractMethodBody(lines []string, startLine int, openDelim, closeDelim string) string {
	var body strings.Builder
	braceCount := 0
	started := false
	
	for i := startLine; i < len(lines) && i < startLine+100; i++ {
		line := lines[i]
		
		for _, char := range line {
			if string(char) == openDelim {
				braceCount++
				started = true
			} else if string(char) == closeDelim {
				braceCount--
			}
		}
		
		if started {
			body.WriteString(line)
			body.WriteString("\n")
		}
		
		if started && braceCount == 0 {
			break
		}
	}
	
	return body.String()
}

// containsAny verifica se o texto contém alguma das strings
func containsAny(text string, patterns []string) bool {
	textLower := strings.ToLower(text)
	for _, pattern := range patterns {
		if strings.Contains(textLower, strings.ToLower(pattern)) {
			return true
		}
	}
	return false
}

// ExtractReturnType tenta extrair o tipo de retorno de um método
func ExtractReturnType(lines []string, startLine int, language string) string {
	if startLine >= len(lines) {
		return ""
	}
	
	switch language {
	case "java":
		return extractJavaReturnType(lines, startLine)
	case "typescript":
		return extractTSReturnType(lines, startLine)
	case "go":
		return extractGoReturnType(lines, startLine)
	default:
		return ""
	}
}

func extractJavaReturnType(lines []string, startLine int) string {
	// Procura por padrões como: public ResponseEntity<Type> methodName(
	returnTypeRegex := regexp.MustCompile(`(?:public|private|protected)\s+(\w+(?:<[^>]+>)?)\s+\w+\s*\(`)
	
	for i := startLine; i < startLine+5 && i < len(lines); i++ {
		if matches := returnTypeRegex.FindStringSubmatch(lines[i]); len(matches) > 1 {
			return matches[1]
		}
	}
	return ""
}

func extractTSReturnType(lines []string, startLine int) string {
	// Procura por padrões como: async methodName(): Promise<Type>
	returnTypeRegex := regexp.MustCompile(`\):\s*(?:Promise<)?([^{>\s]+)`)
	
	for i := startLine; i < startLine+5 && i < len(lines); i++ {
		if matches := returnTypeRegex.FindStringSubmatch(lines[i]); len(matches) > 1 {
			return matches[1]
		}
	}
	return ""
}

func extractGoReturnType(lines []string, startLine int) string {
	// Procura por padrões como: func methodName() (Type, error)
	returnTypeRegex := regexp.MustCompile(`\)\s*\(([^)]+)\)`)
	
	for i := startLine; i < startLine+5 && i < len(lines); i++ {
		if matches := returnTypeRegex.FindStringSubmatch(lines[i]); len(matches) > 1 {
			return matches[1]
		}
	}
	return ""
}
