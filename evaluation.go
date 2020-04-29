package wotlib

// contains function for evaluation based on expanded thing descriptions

// GetPropertyAffordances searches within a set for all property affordances where constraints match
func (s *ExpandedThingDescriptionSet) GetPropertyAffordances(constraint ThingConstraint) []ExpandedPropertyAffordance {
	// some optimization: remove unwanted constraints for filtering things
	strippedThingConstraints := constraint
	strippedThingConstraints.PropertyConstraint = nil

	var result []ExpandedPropertyAffordance

	for _, currTD := range *s {
		if currTD.Fulfills(strippedThingConstraints) {
			if constraint.PropertyConstraint == nil {
				result = append(result, currTD.GetPropertyAffordances(PropertyConstraint{})...)
			} else {
				result = append(result, currTD.GetPropertyAffordances(*constraint.PropertyConstraint)...)
			}
		}
	}

	return result
}

// GetActionAffordances searches within a set for all actions affordances where constraints match
func (s *ExpandedThingDescriptionSet) GetActionAffordances(constraint ThingConstraint) []ExpandedActionAffordance {
	// some optimization: remove unwanted constraints for filtering things
	strippedThingConstraints := constraint
	strippedThingConstraints.ActionConstraint = nil

	var result []ExpandedActionAffordance

	for _, currTD := range *s {
		if currTD.Fulfills(strippedThingConstraints) {
			if constraint.ActionConstraint == nil {
				result = append(result, currTD.GetActionAffordances(ActionConstraint{})...)
			} else {
				result = append(result, currTD.GetActionAffordances(*constraint.ActionConstraint)...)
			}
		}
	}

	return result
}

// GetPropertyAffordances searches for a property affordance with specific criteria
func (t *ExpandedThingDescription) GetPropertyAffordances(constraint PropertyConstraint) []ExpandedPropertyAffordance {
	var result []ExpandedPropertyAffordance

	if len(t.Properties) == 0 {
		return result
	}

	for _, currProperty := range t.Properties {
		if currProperty.Fulfills(constraint) {
			result = append(result, currProperty)
		}
	}

	return result
}

// GetActionAffordances searches for an action affordance with specific criteria
func (t *ExpandedThingDescription) GetActionAffordances(constraint ActionConstraint) []ExpandedActionAffordance {
	var result []ExpandedActionAffordance

	if len(t.Actions) == 0 {
		return result
	}

	for _, currAction := range t.Actions {
		if currAction.Fulfills(constraint) {
			result = append(result, currAction)
		}
	}

	return result
}

// ThingConstraint defines a thing constraint
type ThingConstraint struct {
	ID                 *string
	Type               *[]string
	Name               *string
	PropertyConstraint *PropertyConstraint
	ActionConstraint   *ActionConstraint
}

// Fulfills checks if constraint matches with given element
// A thing description matches if ID, type and name match (if given)
// and if at least one of the PropertyConstraints and one of
// the ActionConstraints matches (if given)
func (t ExpandedThingDescription) Fulfills(c ThingConstraint) bool {
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

		// at least one property has to match
		for _, currProperty := range t.Properties {
			if currProperty.Fulfills(*c.PropertyConstraint) {
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

		// at least one action has to match
		for _, currAction := range t.Actions {
			if currAction.Fulfills(*c.ActionConstraint) {
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
	IsObservable           *bool
}

// Fulfills checks if PropertyConstraint is fulfilled by given ExpandedPropertyAffordance
func (t ExpandedPropertyAffordance) Fulfills(c PropertyConstraint) bool {
	if c.Index != nil && *c.Index != t.Index {
		return false
	}

	if c.Type != nil && !allTypesContained(*c.Type, t.Type) {
		return false
	}

	if c.DataType != nil && *c.DataType != t.DataType.Value() {
		return false
	}

	if c.IsObservable != nil && *c.IsObservable != t.IsObservable.Value() {
		return false
	}

	if c.DataPropertyConstraint != nil {
		matchFound := false

		// at least one property has to match
		for _, currProperty := range t.Properties {
			if currProperty.Fulfills(*c.DataPropertyConstraint) {
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

// ActionConstraint defines an action constraint
type ActionConstraint struct {
	Index           *string
	Type            *[]string
	InputConstraint *InputConstraint
	IsIdempotent    *bool
	IsSafe          *bool
}

// Fulfills checks if ActionConstraint is fulfilled by given ExpandedActionAffordance
func (t ExpandedActionAffordance) Fulfills(c ActionConstraint) bool {
	if c.Index != nil && *c.Index != t.Index {
		return false
	}

	if c.Type != nil && !allTypesContained(*c.Type, t.Type) {
		return false
	}

	if c.IsIdempotent != nil && *c.IsIdempotent != t.IsIdempotent.Value() {
		return false
	}

	if c.IsSafe != nil && *c.IsSafe != t.IsSafe.Value() {
		return false
	}

	if c.InputConstraint != nil && !t.Input.Fulfills(*c.InputConstraint) {
		return false
	}

	return true
}

// InputConstraint defines an input constraint
type InputConstraint struct {
	DataType               *string
	DataPropertyConstraint *DataPropertyConstraint
}

// Fulfills checks if ExpandedInputConstraint matches with given ExpandedDataProperty
func (t ExpandedDataSchemaNode) Fulfills(c InputConstraint) bool {
	elem := t.Value()

	if c.DataType != nil && *c.DataType != elem.DataType.Value() {
		return false
	}

	if c.DataPropertyConstraint != nil {
		matchFound := false

		for _, currProperty := range elem.Properties {
			if currProperty.Fulfills(*c.DataPropertyConstraint) {
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

// DataPropertyConstraint defines a data property constraint
type DataPropertyConstraint struct {
	Index    *string
	Type     *[]string
	DataType *string
}

// Fulfills checks if DataPropertyConstraint is fulfilled by given ExpandedDataProperty
func (t ExpandedDataProperty) Fulfills(c DataPropertyConstraint) bool {
	if c.Index != nil && *c.Index != t.Index {
		return false
	}

	if c.Type != nil && !allTypesContained(*c.Type, t.Type) {
		return false
	}

	if c.DataType != nil && *c.DataType != t.DataType.Value() {
		return false
	}

	return true
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
