package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerrulesetconfigchangecondition
type Dialerrulesetconfigchangecondition struct { 
	// VarType
	VarType *string `json:"type,omitempty"`


	// Inverted
	Inverted *bool `json:"inverted,omitempty"`


	// AttributeName
	AttributeName *string `json:"attributeName,omitempty"`


	// Value
	Value *string `json:"value,omitempty"`


	// ValueType
	ValueType *string `json:"valueType,omitempty"`


	// Operator
	Operator *string `json:"operator,omitempty"`


	// Codes
	Codes *[]string `json:"codes,omitempty"`


	// PropertyType
	PropertyType *string `json:"propertyType,omitempty"`


	// Property
	Property *string `json:"property,omitempty"`


	// DataNotFoundResolution
	DataNotFoundResolution *bool `json:"dataNotFoundResolution,omitempty"`


	// ContactIdField
	ContactIdField *string `json:"contactIdField,omitempty"`


	// CallAnalysisResultField
	CallAnalysisResultField *string `json:"callAnalysisResultField,omitempty"`


	// AgentWrapupField
	AgentWrapupField *string `json:"agentWrapupField,omitempty"`


	// ContactColumnToDataActionFieldMappings
	ContactColumnToDataActionFieldMappings *[]Dialerrulesetconfigchangecontactcolumntodataactionfieldmapping `json:"contactColumnToDataActionFieldMappings,omitempty"`


	// Predicates
	Predicates *[]Dialerrulesetconfigchangedataactionconditionpredicate `json:"predicates,omitempty"`


	// DataAction
	DataAction *Dialerrulesetconfigchangeurireference `json:"dataAction,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialerrulesetconfigchangecondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
