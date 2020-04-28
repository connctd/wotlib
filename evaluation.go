package wotlib

// contains function for evaluation based on expanded thing descriptions

// FindPropertyAffordances searches for a property affordance with specific criteria
func (t *ExpandedThingDescription) FindPropertyAffordances() []ExpandedPropertyAffordance {
	return []ExpandedPropertyAffordance{}
}

// ThingConstraint defines a thing constraint
type ThingConstraint struct {
	ID                 *string
	Type               *[]string
	Name               *string
	PropertyConstraint *PropertyConstraint
	ActionConstraint   *ActionConstraint
}

// Matches checks if constraint matches with given element
// A thing description matches if ID, type and name match (if given)
// and if at least one of the PropertyConstraints and one of
// the ActionConstraints matches (if given)
func (c ThingConstraint) Matches(t ExpandedThingDescription) bool {
	if c.ID != nil && *c.ID != t.ID {
		return false
	}

	if c.Type != nil && !allTypesContained(*c.Type, t.Type) {
		return false
	}

	if c.Name != nil && *c.Name != t.Name.Value() {
		return false
	}

	if c.PropertyConstraint != nil {
		matchFound := false
		for _, currProperty := range t.Properties {
			if c.PropertyConstraint.Matches(currProperty) {
				matchFound = true
				break
			}
		}

		if !matchFound {
			return false
		}
	}

	if c.ActionConstraint != nil {
		matchFound := false
		for _, currAction := range t.Actions {
			if c.ActionConstraint.Matches(currAction) {
				matchFound = true
				break
			}
		}

		if !matchFound {
			return false
		}
	}

	return true
}

// PropertyConstraint defines a property constraint
type PropertyConstraint struct {
	Index                  *string
	Type                   *[]string
	DataType               *string
	DataPropertyConstraint *DataPropertyConstraint
	isObservable           *bool
}

// Matches checks if PropertyConstraint is fulfilled by given ExpandedPropertyAffordance
func (c PropertyConstraint) Matches(t ExpandedPropertyAffordance) bool {
	// TODO
	return true
}

// ActionConstraint defines an action constraint
type ActionConstraint struct {
	Index           *string
	Type            *[]string
	InputConstraint *InputConstraint
	IsIdempotent    *bool
	IsSafe          *bool
}

// Matches checks if ActionConstraint is fulfilled by given ExpandedActionAffordance
func (c ActionConstraint) Matches(t ExpandedActionAffordance) bool {
	// TODO
	return true
}

// InputConstraint defines an input constraint
type InputConstraint struct {
	DataType               *string
	DataPropertyConstraint *DataPropertyConstraint
}

// DataPropertyConstraint defines a data property constraint
type DataPropertyConstraint struct {
	Index    *string
	Type     *[]string
	DataType *string
}

func allTypesContained(requiredTypes []string, givenTypes []string) bool {
	for _, r := range requiredTypes {
		found := false
		for _, g := range givenTypes {
			if r == g {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}
