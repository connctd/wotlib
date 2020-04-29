package wotlib

import (
	"testing"
)

var iotSchema = SchemaMapping{
	Prefix: SchemaPrefix("iot"),
	IRI:    "http://iotschema.org/",
}

var fulfillTests = []struct {
	Name           string
	TD             []byte
	Constraint     ThingConstraint
	ExpectedResult bool
}{
	{
		Name:           "Empty constraint",
		TD:             testTDOne,
		Constraint:     ThingConstraint{},
		ExpectedResult: true,
	},
	{
		Name: "Search for thing with name",
		TD:   testTDOne,
		Constraint: ThingConstraint{
			Name: asStringPointer("LightOne"),
		},
		ExpectedResult: true,
	},
	{
		Name: "Search with property constraint",
		TD:   testTDOne,
		Constraint: ThingConstraint{
			PropertyConstraint: &PropertyConstraint{
				Type: &[]string{iotSchema.IRIPrefix("SwitchStatus")},
			},
		},
		ExpectedResult: true,
	},
	{
		Name: "Search with action constraint",
		TD:   testTDOne,
		Constraint: ThingConstraint{
			ActionConstraint: &ActionConstraint{
				Type: &[]string{iotSchema.IRIPrefix("TurnOn")},
			},
		},
		ExpectedResult: true,
	},
	{
		Name: "Search for thing with thing and property constraint",
		TD:   testTDOne,
		Constraint: ThingConstraint{
			Type: &[]string{
				iotSchema.IRIPrefix("BinarySwitchControl"),
				iotSchema.IRIPrefix("ColourControl"),
			},
			PropertyConstraint: &PropertyConstraint{
				Type: &[]string{iotSchema.IRIPrefix("SwitchStatus")},
				DataPropertyConstraint: &DataPropertyConstraint{
					Type: &[]string{iotSchema.IRIPrefix("StatusData")},
				},
			},
		},
		ExpectedResult: true,
	},
	{
		Name: "Complex search",
		TD:   testTDOne,
		Constraint: ThingConstraint{
			Type: &[]string{
				iotSchema.IRIPrefix("BinarySwitchControl"),
				iotSchema.IRIPrefix("ColourControl"),
			},
			ID:   asStringPointer("uri:urn:ed2f1fb3-cbf8-479e-99bb-ef9968e5eed6"),
			Name: asStringPointer("LightOne"),
			ActionConstraint: &ActionConstraint{
				Index:        asStringPointer("lamp-setOn"),
				IsSafe:       asBooleanPointer(false),
				IsIdempotent: asBooleanPointer(true),
				Type: &[]string{
					iotSchema.IRIPrefix("TurnOn"),
					iotSchema.IRIPrefix("TurnOff"),
				},
				InputConstraint: &InputConstraint{
					DataType: asStringPointer(SchemaJSON.IRIPrefix("ObjectSchema")),
					DataPropertyConstraint: &DataPropertyConstraint{
						Index: asStringPointer("on"),
						Type: &[]string{
							iotSchema.IRIPrefix("StatusData"),
						},
						DataType: asStringPointer(SchemaJSON.IRIPrefix("BooleanSchema")),
					},
				},
			},
			PropertyConstraint: &PropertyConstraint{
				Index: asStringPointer("lamp-on"),
				Type:  &[]string{iotSchema.IRIPrefix("SwitchStatus")},
				DataPropertyConstraint: &DataPropertyConstraint{
					Index:    asStringPointer("value"),
					Type:     &[]string{iotSchema.IRIPrefix("StatusData")},
					DataType: asStringPointer(SchemaJSON.IRIPrefix("BooleanSchema")),
				},
				DataType:     asStringPointer(SchemaJSON.IRIPrefix("ObjectSchema")),
				IsObservable: asBooleanPointer(false),
			},
		},
		ExpectedResult: true,
	},
}

func TestFulfill(t *testing.T) {
	DefaultContext[iotSchema.Prefix.String()] = iotSchema

	for _, currTest := range fulfillTests {
		t.Run(currTest.Name, func(t *testing.T) {
			expandedTD, err := FromBytes(currTest.TD)
			if err != nil {
				t.Fatalf(currTest.Name+" failed. Failed to build expanded td: %v", err)
			} else {
				res := expandedTD.Fulfills(currTest.Constraint)
				if res != currTest.ExpectedResult {
					t.Fatalf(currTest.Name+" failed. Expected: %t, Got: %t", currTest.ExpectedResult, res)
				}
			}
		})
	}
}

func TestFindPropertyAffordances(t *testing.T) {
	thingConstraint := ThingConstraint{
		Type: &[]string{
			iotSchema.IRIPrefix("BinarySwitchControl"),
			iotSchema.IRIPrefix("ColourControl"),
		},
	}

	expandedTD, err := FromBytes(testTDOne)
	if err != nil {
		t.Fatalf("Failed to build expanded td: %v", err)
	}

	thingMatches := expandedTD.Fulfills(thingConstraint)
	if !thingMatches {
		t.Fatalf("Thing does not match")
	}

	// now searching for property affordances with specific criteria
	propertyConstraint := PropertyConstraint{
		Type: &[]string{
			iotSchema.IRIPrefix("SwitchStatus"),
		},
		DataPropertyConstraint: &DataPropertyConstraint{
			Type: &[]string{
				iotSchema.IRIPrefix("StatusData"),
			},
		},
	}

	result := expandedTD.GetPropertyAffordances(propertyConstraint)
	if len(result) != 1 {
		t.Fatalf("Unexpected result set. Expected: 1, Got: %d", len(result))
	}

	href := result[0].Form.Value().Href
	if href.Value() != "https://api.connctd.io/api/betav1/wot/things/ad4bb62b-4e95-4628-9d8b-3cd412ec140f/components/lamp/properties/on" {
		t.Fatalf("Unexpected href in result set")
	}
}

func TestFindPropertyAffordancesInSet(t *testing.T) {
	constraint := ThingConstraint{
		Type: &[]string{
			iotSchema.IRIPrefix("BinarySwitchControl"),
			iotSchema.IRIPrefix("ColourControl"),
		},
		PropertyConstraint: &PropertyConstraint{
			Type: &[]string{
				iotSchema.IRIPrefix("SwitchStatus"),
			},
			DataPropertyConstraint: &DataPropertyConstraint{
				Type: &[]string{
					iotSchema.IRIPrefix("StatusData"),
				},
			},
		},
	}

	expandedTD, err := FromBytes(testTDOne)
	if err != nil {
		t.Fatalf("Failed to build expanded td: %v", err)
	}

	set := NewExpandedThingDescriptionSet(expandedTD)
	result := set.GetPropertyAffordances(constraint)
	if len(result) != 1 {
		t.Fatalf("Unexpected result set. Expected: 1, Got: %d", len(result))
	}

	href := result[0].Form.Value().Href
	if href.Value() != "https://api.connctd.io/api/betav1/wot/things/ad4bb62b-4e95-4628-9d8b-3cd412ec140f/components/lamp/properties/on" {
		t.Fatalf("Unexpected href in result set")
	}
}

func TestFindActionAffordancesInSet(t *testing.T) {
	constraint := ThingConstraint{
		Type: &[]string{
			iotSchema.IRIPrefix("BinarySwitchControl"),
			iotSchema.IRIPrefix("ColourControl"),
		},
		ActionConstraint: &ActionConstraint{
			Type: &[]string{
				iotSchema.IRIPrefix("TurnOn"),
			},
			InputConstraint: &InputConstraint{
				DataPropertyConstraint: &DataPropertyConstraint{
					Type: &[]string{iotSchema.IRIPrefix("StatusData")},
				},
			},
		},
	}

	expandedTD, err := FromBytes(testTDOne)
	if err != nil {
		t.Fatalf("Failed to build expanded td: %v", err)
	}

	set := NewExpandedThingDescriptionSet(expandedTD)
	result := set.GetActionAffordances(constraint)
	if len(result) != 1 {
		t.Fatalf("Unexpected result set. Expected: 1, Got: %d", len(result))
	}

	href := result[0].Form.Value().Href
	if href.Value() != "https://api.connctd.io/api/betav1/wot/things/ad4bb62b-4e95-4628-9d8b-3cd412ec140f/components/lamp/actions/setOn" {
		t.Fatalf("Unexpected href in result set")
	}
}

func asStringPointer(input string) *string {
	return &input
}

func asBooleanPointer(input bool) *bool {
	return &input
}

var testTDOne = []byte(`{
            "@context": [
                "https://www.w3.org/2019/wot/td/v1",
                {
                    "iot": "http://iotschema.org/",
                    "unit": "http://qudt.org/vocab/unit/",
                    "schema": "http://schema.org/"
                }
			],
			"description": "Generated TD from thing",
            "schema:manufacturer": "LIFX",
            "securityDefinitions": {
                "bearerSecurityScheme": {
                    "in": "header",
                    "alg": "HMAC-SHA256",
                    "name": "Authorization",
                    "format": "jwt",
                    "scheme": "bearer",
                    "authorization": "https://api.connctd.io/oauth2/token"
                }
            },
            "security": [
                "bearerSecurityScheme"
            ],
            "id": "uri:urn:ed2f1fb3-cbf8-479e-99bb-ef9968e5eed6",
            "name": "LightOne",
            "@type": [
                "Thing",
                "iot:ColourControl",
                "iot:DimmerControl",
                "iot:BinarySwitchControl"
            ],
            "title": "LightOne",
            "actions": {
                "lamp-setOn": {
                    "safe": false,
                    "@type": [
                        "ActionAffordance",
                        "iot:TurnOn",
                        "iot:TurnOff"
                    ],
                    "forms": [
                        {
                            "op": "invokeaction",
                            "href": "https://api.connctd.io/api/betav1/wot/things/ad4bb62b-4e95-4628-9d8b-3cd412ec140f/components/lamp/actions/setOn",
                            "contentType": "application/json"
                        }
                    ],
                    "input": {
                        "type": "object",
                        "properties": {
                            "on": {
                                "type": "boolean",
                                "@type": ["iot:StatusData","iot:StateData"]
                            }
                        }
                    },
                    "title": "setOn",
                    "idempotent": true
				},
                "lamp-setColor": {
                    "safe": true,
                    "type": "object",
                    "@type": [
                        "ActionAffordance",
                        "iot:SetColour"
                    ],
                    "forms": [
                        {
                            "op": "invokeaction",
                            "href": "https://api.connctd.io/api/betav1/wot/things/ad4bb62b-4e95-4628-9d8b-3cd412ec140f/components/lamp/actions/setColor",
                            "contentType": "application/json"
                        }
                    ],
                    "input": {
                        "type": "object",
                        "properties": {
                            "red": {
                                "type": "integer",
                                "@type": "iot:RColourData"
                            },
                            "blue": {
                                "type": "integer",
                                "@type": "iot:BColourData"
                            },
                            "green": {
                                "type": "integer",
                                "@type": "iot:GColourData"
                            }
                        }
                    },
                    "title": "setColor",
                    "idempotent": true
                }
			},
            "properties": {
                "lamp-on": {
                    "type": "object",
                    "@type": [
                        "PropertyAffordance",
                        "iot:SwitchStatus"
                    ],
                    "forms": [
                        {
                            "op": "readproperty",
                            "href": "https://api.connctd.io/api/betav1/wot/things/ad4bb62b-4e95-4628-9d8b-3cd412ec140f/components/lamp/properties/on",
                            "contentType": "application/json"
                        }
                    ],
                    "title": "on",
                    "writeable": false,
                    "observable": false,
                    "properties": {
                        "time": {
                            "@id": "dateModified",
                            "type": "string",
                            "@type": "schema:DateTime"
                        },
                        "value": {
                            "type": "boolean",
                            "@type": "iot:StatusData",
                            "schema:dateModified": {
                                "@id": "dateModified"
                            }
                        }
                    }
				},
				"lamp-color": {
                    "type": "object",
                    "@type": [
                        "PropertyAffordance",
                        "iot:CurrentColour"
                    ],
                    "forms": [
                        {
                            "op": "readproperty",
                            "href": "https://api.connctd.io/api/betav1/wot/things/ad4bb62b-4e95-4628-9d8b-3cd412ec140f/components/lamp/properties/color",
                            "contentType": "application/json"
                        }
                    ],
                    "title": "color",
                    "writeable": false,
                    "observable": false,
                    "properties": {
                        "red": {
                            "type": "integer",
                            "@type": "iot:RColourData",
                            "schema:dateModified": {
                                "@id": "dateModified"
                            }
                        },
                        "blue": {
                            "type": "integer",
                            "@type": "iot:BColourData",
                            "schema:dateModified": {
                                "@id": "dateModified"
                            }
                        },
                        "time": {
                            "@id": "dateModified",
                            "type": "string",
                            "@type": "schema:DateTime"
                        },
                        "green": {
                            "type": "integer",
                            "@type": "iot:GColourData",
                            "schema:dateModified": {
                                "@id": "dateModified"
                            }
                        }
                    }
                }
            }
        }
    `)
