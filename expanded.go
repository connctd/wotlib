package wotlib

// ExpandedThingDescriptionSet set of thing descriptions with some convenience functions
type ExpandedThingDescriptionSet map[string]ExpandedThingDescription

// NewExpandedThingDescriptionSet creates a new set
func NewExpandedThingDescriptionSet(tds ...ExpandedThingDescription) ExpandedThingDescriptionSet {
	s := ExpandedThingDescriptionSet{}

	for i := range tds {
		s.Append(tds[i])
	}

	return s
}

// Append appends an expanded td to the list
func (s *ExpandedThingDescriptionSet) Append(td ExpandedThingDescription) {
	(*s)[td.ID] = td
}

// Remove removes a td from the list
func (s *ExpandedThingDescriptionSet) Remove(id string) {
	delete(*s, id)
}

// Get retrieves a td by id
func (s *ExpandedThingDescriptionSet) Get(id string) ExpandedThingDescription {
	return (*s)[id]
}

// ExpandedThingDescription reflects a thing description in its expanded format
// Note: currently this lib only supports a small sub set of fields
type ExpandedThingDescription struct {
	ID         string                       `json:"@id"`
	Type       []string                     `json:"@type,omitempty"`
	Name       StringNode                   `json:"https://www.w3.org/2019/wot/td#name"`
	Actions    []ExpandedActionAffordance   `json:"https://www.w3.org/2019/wot/td#hasActionAffordance"`
	Properties []ExpandedPropertyAffordance `json:"https://www.w3.org/2019/wot/td#hasPropertyAffordance"`
}

// ExpandedActionAffordance defines an expanded action affordance within a td
type ExpandedActionAffordance struct {
	Name         StringNode             `json:"https://www.w3.org/2019/wot/td#name"`
	Type         []string               `json:"@type,omitempty"`
	Form         ExpandedFormNode       `json:"https://www.w3.org/2019/wot/td#hasForm"`
	Input        ExpandedDataSchemaNode `json:"https://www.w3.org/2019/wot/td#hasInputSchema"`
	IsIdempotent BooleanNode            `json:"https://www.w3.org/2019/wot/td#isIdempotent"`
	IsSafe       BooleanNode            `json:"https://www.w3.org/2019/wot/td#isSafe"`
}

// ExpandedPropertyAffordance defines an expanded property affordance within a td
type ExpandedPropertyAffordance struct {
	Name         StringNode             `json:"https://www.w3.org/2019/wot/td#name"`
	Type         []string               `json:"@type,omitempty"`
	DataType     IDNode                 `json:"http://www.w3.org/1999/02/22-rdf-syntax-ns#type"`
	Properties   []ExpandedDataProperty `json:"https://www.w3.org/2019/wot/json-schema#properties"`
	Form         ExpandedFormNode       `json:"https://www.w3.org/2019/wot/td#hasForm"`
	IsObservable BooleanNode            `json:"https://www.w3.org/2019/wot/td#isObservable"`
}

// ExpandedDataSchemaNode is an array of expanded data schema
type ExpandedDataSchemaNode []ExpandedDataSchema

// Value returns the first element inside the schema node
// or an empty schema if no such element exists
func (s ExpandedDataSchemaNode) Value() ExpandedDataSchema {
	if len(s) == 0 {
		return ExpandedDataSchema{}
	}

	return s[0]
}

// ExpandedDataSchema can be inside an input param of an action affordance or
// inside a property
type ExpandedDataSchema struct {
	DataType   IDNode                 `json:"http://www.w3.org/1999/02/22-rdf-syntax-ns#type"`
	Properties []ExpandedDataProperty `json:"https://www.w3.org/2019/wot/json-schema#properties"`
}

// ExpandedDataProperty is part of a data schema
type ExpandedDataProperty struct {
	Name     StringNode `json:"https://www.w3.org/2019/wot/json-schema#propertyName"`
	Type     []string   `json:"@type,omitempty"`
	DataType IDNode     `json:"http://www.w3.org/1999/02/22-rdf-syntax-ns#type"`
}

// ExpandedFormNode is an array of expanded forms
type ExpandedFormNode []ExpandedForm

// Value returns the first element inside the form node
// or an empty form if no such element exists
func (s ExpandedFormNode) Value() ExpandedForm {
	if len(s) == 0 {
		return ExpandedForm{}
	}

	return s[0]
}

// ExpandedForm is part of actions and properties and describes how to resolve an entity
type ExpandedForm struct {
	ContentType StringNode `json:"https://www.w3.org/2019/wot/hypermedia#forContentType"`
	Op          IDNode     `json:"https://www.w3.org/2019/wot/hypermedia#hasOperationType"`
	Href        IDNode     `json:"https://www.w3.org/2019/wot/hypermedia#hasTarget"`
}

// StringNode defines an array of string values
type StringNode []StringValue

// Value returns the first element inside the string node
// or an empty string if no such element exists
func (s StringNode) Value() string {
	if len(s) == 0 {
		return ""
	}

	return s[0].Value
}

// BooleanNode defines an array of boolean values
type BooleanNode []BooleanValue

// Value returns the first element inside the boolean node
// or false if no such element exists
func (b BooleanNode) Value() bool {
	if len(b) == 0 {
		return false
	}

	return b[0].Value
}

// IDNode defines an array of id values
type IDNode []IDValue

// Value returns the first element inside the string node
// or an empty string if no such element exists
func (s IDNode) Value() string {
	if len(s) == 0 {
		return ""
	}

	return s[0].ID
}

// StringValue describes a string value
type StringValue struct {
	Value string `json:"@value"`
}

// BooleanValue describes a boolean value
type BooleanValue struct {
	Value bool `json:"@value"`
}

// IDValue describes an id value
type IDValue struct {
	ID string `json:"@id"`
}
