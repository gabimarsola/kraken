package structure

const PRDTemplate = `# {{.PRD.Title}}

**Projeto:** {{.ProjectInfo.Name}}  
**Tipo:** {{.ProjectInfo.ProjectType}}  
**Versão:** {{.ProjectInfo.Version}}

## 📋 Introdução/Visão Geral

{{.PRD.Introduction}}

## 🎯 Objetivos

{{range .PRD.Objectives}}
- {{.}}
{{end}}

## 📖 Histórias do Usuário

{{range .PRD.UserStories}}
### {{.ID}}: {{.Title}}

{{.Description}}

#### Critérios de Aceitação

{{range .AcceptanceCriteria}}
- {{.}}
{{end}}

---
{{end}}

## ⚙️ Requisitos Funcionais

{{range .PRD.FunctionalReqs}}
### {{.ID}}: {{.Title}} (Prioridade: {{.Priority}})

{{.Description}}

---
{{end}}

## 🚫 Não-Objetivos (Fora do Escopo)

{{range .PRD.OutOfScope}}
- {{.}}
{{end}}

## 🎨 Considerações de Design

{{range .PRD.DesignConsiderations}}
- {{.}}
{{end}}

## 🔧 Considerações Técnicas

{{range .PRD.TechConsiderations}}
- {{.}}
{{end}}

## 📊 Métricas de Sucesso

{{range .PRD.SuccessMetrics}}
- {{.}}
{{end}}

## ❓ Questões Abertas

{{range .PRD.OpenQuestions}}
- {{.}}
{{end}}
`
