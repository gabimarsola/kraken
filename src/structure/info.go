package structure

type Parameter struct {
	Name        string
	Type        string
	Required    bool
	Description string
	Location    string // query, body, path, header
}

type RequestExample struct {
	Language string
	Code     string
}

type Endpoint struct {
	Method          string
	Path            string
	Description     string
	Summary         string // Resumo do funcionamento do código
	Parameters      []Parameter
	RequestExamples []RequestExample
}

type ProjectInfo struct {
	Name        string
	Version     string
	Description string
	ProjectType string
	Endpoints   []Endpoint
}
