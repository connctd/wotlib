package wotlib

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/piprate/json-gold/ld"
)

// Default json-ld options
var (
	DefaultJSONDLDOptions = ld.NewJsonLdOptions("")
)

// FromResponse tries to extract an expanded wot td from a
// response object
func FromResponse(resp *http.Response) (ExpandedThingDescription, error) {
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ExpandedThingDescription{}, err
	}

	return FromBytes(bytes)
}

// FromBytes expands input bytes and converts it to ExpandedThingDescription
func FromBytes(b []byte) (ExpandedThingDescription, error) {
	proc := ld.NewJsonLdProcessor()

	// lib is expecting a map
	var asMap map[string]interface{}
	err := json.Unmarshal(b, &asMap)
	if err != nil {
		return ExpandedThingDescription{}, err
	}

	expandedObj, err := proc.Expand(asMap, DefaultJSONDLDOptions)
	if err != nil {
		return ExpandedThingDescription{}, err
	}

	// first we need to convert the map into a byte arr again
	expandedBytes, err := json.Marshal(expandedObj)
	if err != nil {
		return ExpandedThingDescription{}, err
	}

	// now we can properly convert it to a td
	var td []ExpandedThingDescription
	err = json.Unmarshal(expandedBytes, &td)
	if err != nil {
		return ExpandedThingDescription{}, err
	}

	return td[0], nil
}

// Compact compacts the thing description
func (e *ExpandedThingDescription) Compact() (json.RawMessage, error) {
	compactedBytes, err := compact(e)
	if err != nil {
		return nil, err
	}

	return json.RawMessage(compactedBytes), nil
}

// Compact compacts an expanded property affordance
func (e *ExpandedPropertyAffordance) Compact() (json.RawMessage, error) {
	compactedBytes, err := compact(e)
	if err != nil {
		return nil, err
	}

	return json.RawMessage(compactedBytes), nil
}

// Compact compacts an expanded action affordance
func (e *ExpandedActionAffordance) Compact() (json.RawMessage, error) {
	compactedBytes, err := compact(e)
	if err != nil {
		return nil, err
	}

	return json.RawMessage(compactedBytes), nil
}

func compact(e interface{}) ([]byte, error) {
	proc := ld.NewJsonLdProcessor()

	expandedBytes, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}

	var expandedMap map[string]interface{}
	if err := json.Unmarshal(expandedBytes, &expandedMap); err != nil {
		return nil, err
	}

	compactedMap, err := proc.Compact(expandedMap, map[string]interface{}{"@context": DefaultContext}, DefaultJSONDLDOptions)
	if err != nil {
		return nil, err
	}

	return json.Marshal(compactedMap)
}
