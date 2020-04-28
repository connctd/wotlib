package wotlib

// SchemaPrefix like wot, schema...
type SchemaPrefix string

// well know schemata
var (
	SchemaWoT = SchemaMapping{
		Prefix: SchemaPrefix("wot"),
		IRI:    "https://www.w3.org/2019/wot/td#",
	}

	SchemaHypermedia = SchemaMapping{
		Prefix: SchemaPrefix("hypermedia"),
		IRI:    "https://www.w3.org/2019/wot/hypermedia#",
	}

	SchemaRdfType = SchemaMapping{
		Prefix: SchemaPrefix("rdftype"),
		IRI:    "http://www.w3.org/1999/02/22-rdf-syntax-ns#",
	}

	SchemaJSON = SchemaMapping{
		Prefix: SchemaPrefix("jsonschema"),
		IRI:    "https://www.w3.org/2019/wot/json-schema#",
	}
)

// SchemaMapping defines a prefix iri mapping
type SchemaMapping struct {
	Prefix SchemaPrefix
	IRI    string
}

// IRIPrefix builds a prefixed node name for referencing elements in
// a thing description. Eg. SchemaWoT.IRIPrefix(PropertyAffordance) would
// return "https://...#PropertyAffordance"
func (s SchemaMapping) IRIPrefix(nodeID string) string {
	return string(s.IRI) + nodeID
}

// IRIContext defines a special type of context
// We are assuming such type to simplify working with tds
type IRIContext map[SchemaPrefix]string

// ContextContainer wraps the context
type ContextContainer struct {
	Context map[string]interface{} `json:"@context"`
}

// DefaultContext defines the default context
// Modifications to existing prefixes might corrupt parsing process
var DefaultContext = map[string]interface{}{
	string(SchemaWoT.Prefix):        SchemaWoT.IRI,
	string(SchemaHypermedia.Prefix): SchemaHypermedia.IRI,
	string(SchemaRdfType.Prefix):    SchemaRdfType.IRI,
	string(SchemaJSON.Prefix):       SchemaJSON.IRI,
}
