# WoT Lib

This lib was created to simplify work with WoT Thing Descriptions.
It is coupled to W3C's recommendation from 9 April 2020 (https://www.w3.org/TR/2020/REC-wot-thing-description-20200409/) which means it assumes a certain structure. So far only a subset of fields are
implemented.

## Working principle

Read TD -> Expand -> Perform operations

## Capabilities

- Determine if a thing and its sub elements do match with certain criteria
- Search for property affordances with specific constraints
- Search for action affordances with specific constraints

## Example

The following example reads a td from bytes and retrieves all property affordances that
match a given criteria

```
// define additionally used schema
var iotSchema = wotlib.SchemaMapping{
    Prefix: wotlib.SchemaPrefix("iot"),
    IRI:    "http://iotschema.org/",
}

// append it so its used during compaction
wotlib.AppendSchema(iotSchema)

expandedTD, err := wotlib.FromBytes(input)
if err != nil {
    panic(err)
}

// retrieve all property affordances that match given criteria
props := expandedTD.GetPropertyAffordances(wotlib.PropertyConstraint{
    Type: &[]string{iotSchema.IRIPrefix("SwitchStatus")},
})

// prints the href of the first match
fmt.Printf("Result: %s", props[0].Form.Value().Href.Value())
```

In case multiple tds have to be processed append them to a set
and apply the operation

```
set := wotlib.NewExpandedThingDescriptionSet(expandedTD)
set.GetActionAffordances(thingConstraint)
```