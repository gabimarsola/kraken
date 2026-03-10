package structure

const RouteTemplate = `# Rota: {{.RouteName}}

**Projeto:** {{.ProjectName}}  
**Tipo:** {{.ProjectType}}  
**Versão:** {{.Version}}

## Endpoints da Rota {{.RouteName}}

{{range .Endpoints}}
### {{.Method}} {{.Path}}

**Descrição:** {{.Description}}

{{if .Summary}}
**Funcionamento:** {{.Summary}}

{{end}}
{{if .Parameters}}
#### Parâmetros Necessários

| Nome | Tipo | Obrigatório | Localização | Descrição |
|------|------|-------------|-------------|-----------|
{{range .Parameters}}| {{.Name}} | {{.Type}} | {{if .Required}}Sim{{else}}Não{{end}} | {{.Location}} | {{.Description}} |
{{end}}
{{end}}

{{if .RequestExamples}}
#### Exemplos de Requisição

{{range .RequestExamples}}
**{{.Language}}:**

` + "```bash\n" + `{{.Code}}
` + "```\n" + `
{{end}}
{{end}}

---
{{end}}
`

const MdTemplate = `# Projeto: {{.Name}}
Tipo: {{.ProjectType}}
Versão: {{.Version}}

## Descrição
{{.Description}}

## API Endpoints

{{if .Endpoints}}
{{range .Endpoints}}
### {{.Method}} {{.Path}}

**Descrição:** {{.Description}}

{{if .Summary}}
**Funcionamento:** {{.Summary}}

{{end}}
{{if .Parameters}}
#### Parâmetros Necessários

| Nome | Tipo | Obrigatório | Localização | Descrição |
|------|------|-------------|-------------|-----------|
{{range .Parameters}}| {{.Name}} | {{.Type}} | {{if .Required}}Sim{{else}}Não{{end}} | {{.Location}} | {{.Description}} |
{{end}}
{{end}}

{{if .RequestExamples}}
#### Exemplos de Requisição

{{range .RequestExamples}}
**{{.Language}}:**

` + "```bash\n" + `{{.Code}}
` + "```\n" + `
{{end}}
{{end}}

---
{{end}}
{{else}}
Nenhum endpoint documentado encontrado.
{{end}}
`
