package main

import (
	"fmt"

	"github.com/connctd/wotlib"
)

func main() {
	input := []byte(`{
            "@context": [
                "https://www.w3.org/2019/wot/td/v1",
                {
                    "iot": "http://iotschema.org/",
                    "unit": "http://qudt.org/vocab/unit/",
                    "schema": "http://schema.org/"
                }
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
                    "safe": true,
                    "type": "object",
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
                },
                "lamp-enablePulse": {
                    "safe": true,
                    "type": "object",
                    "@type": [
                        "ActionAffordance"
                    ],
                    "forms": [
                        {
                            "op": "invokeaction",
                            "href": "https://api.connctd.io/api/betav1/wot/things/ad4bb62b-4e95-4628-9d8b-3cd412ec140f/components/lamp/actions/enablePulse",
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
                            },
                            "cycles": {
                                "type": "number"
                            },
                            "period": {
                                "type": "number"
                            },
                            "fromcolor": {
                                "type": "string"
                            }
                        }
                    },
                    "title": "enablePulse",
                    "idempotent": true
                },
                "lamp-setDimmlevel": {
                    "safe": true,
                    "type": "object",
                    "@type": [
                        "ActionAffordance",
                        "iot:SetDimmer"
                    ],
                    "forms": [
                        {
                            "op": "invokeaction",
                            "href": "https://api.connctd.io/api/betav1/wot/things/ad4bb62b-4e95-4628-9d8b-3cd412ec140f/components/lamp/actions/setDimmlevel",
                            "contentType": "application/json"
                        }
                    ],
                    "input": {
                        "type": "object",
                        "properties": {
                            "dimmlevel": {
                                "type": "number",
                                "unit": "unit:PERCENT",
                                "@type": "iot:DimmerData"
                            }
                        }
                    },
                    "title": "setDimmlevel",
                    "idempotent": true
                },
                "lamp-adjustedSetOn": {
                    "safe": true,
                    "type": "object",
                    "@type": [
                        "ActionAffordance",
                        "iot:TurnOn",
                        "iot:TurnOff"
                    ],
                    "forms": [
                        {
                            "op": "invokeaction",
                            "href": "https://api.connctd.io/api/betav1/wot/things/ad4bb62b-4e95-4628-9d8b-3cd412ec140f/components/lamp/actions/adjustedSetOn",
                            "contentType": "application/json"
                        }
                    ],
                    "input": {
                        "type": "object",
                        "properties": {
                            "on": {
                                "type": "boolean",
                                "@type": "iot:StatusData"
                            },
                            "transitionspeed": {
                                "type": "number",
                                "unit": "unit:SEC",
                                "@type": "iot:TransitionTimeData"
                            }
                        }
                    },
                    "title": "adjustedSetOn",
                    "idempotent": true
                }
            },
            "security": [
                "bearerSecurityScheme"
            ],
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
                        "iot: CurrentColour"
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
                },
                "lamp-dimmlevel": {
                    "type": "object",
                    "@type": [
                        "PropertyAffordance",
                        "iot:CurrentDimmer"
                    ],
                    "forms": [
                        {
                            "op": "readproperty",
                            "href": "https://api.connctd.io/api/betav1/wot/things/ad4bb62b-4e95-4628-9d8b-3cd412ec140f/components/lamp/properties/dimmlevel",
                            "contentType": "application/json"
                        }
                    ],
                    "title": "dimmlevel",
                    "writeable": false,
                    "observable": false,
                    "properties": {
                        "time": {
                            "@id": "dateModified",
                            "type": "string",
                            "@type": "schema:DateTime"
                        },
                        "value": {
                            "type": "number",
                            "unit": "unit:PERCENT",
                            "@type": "iot:DimmerData",
                            "schema:dateModified": {
                                "@id": "dateModified"
                            }
                        }
                    }
                }
            },
            "description": "Generated TD from thing",
            "referenceId": "ad4bb62b-4e95-4628-9d8b-3cd412ec140f",
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
            }
        }
    `)

	expanded, err := wotlib.FromBytes(input)
	if err != nil {
		fmt.Println("Failed expand", err)
		return
	}

	b, err := expanded.Compact()
	if err != nil {
		fmt.Println("Failed to compact", err)
		return
	}

	fmt.Printf("%s", string(b))
}
