package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Dialerrulesetconfigchangecondition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialerrulesetconfigchangecondition

	

	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Inverted *bool `json:"inverted,omitempty"`
		
		AttributeName *string `json:"attributeName,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		ValueType *string `json:"valueType,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Codes *[]string `json:"codes,omitempty"`
		
		PropertyType *string `json:"propertyType,omitempty"`
		
		Property *string `json:"property,omitempty"`
		
		DataNotFoundResolution *bool `json:"dataNotFoundResolution,omitempty"`
		
		ContactIdField *string `json:"contactIdField,omitempty"`
		
		CallAnalysisResultField *string `json:"callAnalysisResultField,omitempty"`
		
		AgentWrapupField *string `json:"agentWrapupField,omitempty"`
		
		ContactColumnToDataActionFieldMappings *[]Dialerrulesetconfigchangecontactcolumntodataactionfieldmapping `json:"contactColumnToDataActionFieldMappings,omitempty"`
		
		Predicates *[]Dialerrulesetconfigchangedataactionconditionpredicate `json:"predicates,omitempty"`
		
		DataAction *Dialerrulesetconfigchangeurireference `json:"dataAction,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		VarType: u.VarType,
		
		Inverted: u.Inverted,
		
		AttributeName: u.AttributeName,
		
		Value: u.Value,
		
		ValueType: u.ValueType,
		
		Operator: u.Operator,
		
		Codes: u.Codes,
		
		PropertyType: u.PropertyType,
		
		Property: u.Property,
		
		DataNotFoundResolution: u.DataNotFoundResolution,
		
		ContactIdField: u.ContactIdField,
		
		CallAnalysisResultField: u.CallAnalysisResultField,
		
		AgentWrapupField: u.AgentWrapupField,
		
		ContactColumnToDataActionFieldMappings: u.ContactColumnToDataActionFieldMappings,
		
		Predicates: u.Predicates,
		
		DataAction: u.DataAction,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialerrulesetconfigchangecondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
