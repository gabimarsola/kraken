package extractor

import "kraken/src/structure"

type Extractor interface {
	Extract(filePath string, content []byte) ([]structure.Endpoint, error)
	Supports(filePath string) bool
}

type ExtractorRegistry struct {
	extractors []Extractor
}

func NewExtractorRegistry() *ExtractorRegistry {
	return &ExtractorRegistry{
		extractors: make([]Extractor, 0),
	}
}

func (r *ExtractorRegistry) Register(extractor Extractor) {
	r.extractors = append(r.extractors, extractor)
}

func (r *ExtractorRegistry) ExtractFromFile(filePath string, content []byte) []structure.Endpoint {
	var allEndpoints []structure.Endpoint

	for _, extractor := range r.extractors {
		if extractor.Supports(filePath) {
			endpoints, err := extractor.Extract(filePath, content)
			if err == nil && len(endpoints) > 0 {
				allEndpoints = append(allEndpoints, endpoints...)
			}
		}
	}

	return allEndpoints
}
