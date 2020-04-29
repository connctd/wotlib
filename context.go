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

func (p SchemaPrefix) String() string {
	return string(p)
}

// AppendSchema appends a schema to the default context
func AppendSchema(m SchemaMapping) {
	DefaultContext[m.Prefix.String()] = m.IRI
}

// DefaultContext defines the default context
// Modifications to existing prefixes will corrupt parsing process
var DefaultContext = map[string]interface{}{
	SchemaWoT.Prefix.String():        SchemaWoT.IRI,
	SchemaHypermedia.Prefix.String(): SchemaHypermedia.IRI,
	SchemaRdfType.Prefix.String():    SchemaRdfType.IRI,
	SchemaJSON.Prefix.String():       SchemaJSON.IRI,
}
