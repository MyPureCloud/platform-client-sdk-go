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

func (o *Dialerrulesetconfigchangecondition) MarshalJSON() ([]byte, error) {
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
		VarType: o.VarType,
		
		Inverted: o.Inverted,
		
		AttributeName: o.AttributeName,
		
		Value: o.Value,
		
		ValueType: o.ValueType,
		
		Operator: o.Operator,
		
		Codes: o.Codes,
		
		PropertyType: o.PropertyType,
		
		Property: o.Property,
		
		DataNotFoundResolution: o.DataNotFoundResolution,
		
		ContactIdField: o.ContactIdField,
		
		CallAnalysisResultField: o.CallAnalysisResultField,
		
		AgentWrapupField: o.AgentWrapupField,
		
		ContactColumnToDataActionFieldMappings: o.ContactColumnToDataActionFieldMappings,
		
		Predicates: o.Predicates,
		
		DataAction: o.DataAction,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialerrulesetconfigchangecondition) UnmarshalJSON(b []byte) error {
	var DialerrulesetconfigchangeconditionMap map[string]interface{}
	err := json.Unmarshal(b, &DialerrulesetconfigchangeconditionMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := DialerrulesetconfigchangeconditionMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Inverted, ok := DialerrulesetconfigchangeconditionMap["inverted"].(bool); ok {
		o.Inverted = &Inverted
	}
	
	if AttributeName, ok := DialerrulesetconfigchangeconditionMap["attributeName"].(string); ok {
		o.AttributeName = &AttributeName
	}
	
	if Value, ok := DialerrulesetconfigchangeconditionMap["value"].(string); ok {
		o.Value = &Value
	}
	
	if ValueType, ok := DialerrulesetconfigchangeconditionMap["valueType"].(string); ok {
		o.ValueType = &ValueType
	}
	
	if Operator, ok := DialerrulesetconfigchangeconditionMap["operator"].(string); ok {
		o.Operator = &Operator
	}
	
	if Codes, ok := DialerrulesetconfigchangeconditionMap["codes"].([]interface{}); ok {
		CodesString, _ := json.Marshal(Codes)
		json.Unmarshal(CodesString, &o.Codes)
	}
	
	if PropertyType, ok := DialerrulesetconfigchangeconditionMap["propertyType"].(string); ok {
		o.PropertyType = &PropertyType
	}
	
	if Property, ok := DialerrulesetconfigchangeconditionMap["property"].(string); ok {
		o.Property = &Property
	}
	
	if DataNotFoundResolution, ok := DialerrulesetconfigchangeconditionMap["dataNotFoundResolution"].(bool); ok {
		o.DataNotFoundResolution = &DataNotFoundResolution
	}
	
	if ContactIdField, ok := DialerrulesetconfigchangeconditionMap["contactIdField"].(string); ok {
		o.ContactIdField = &ContactIdField
	}
	
	if CallAnalysisResultField, ok := DialerrulesetconfigchangeconditionMap["callAnalysisResultField"].(string); ok {
		o.CallAnalysisResultField = &CallAnalysisResultField
	}
	
	if AgentWrapupField, ok := DialerrulesetconfigchangeconditionMap["agentWrapupField"].(string); ok {
		o.AgentWrapupField = &AgentWrapupField
	}
	
	if ContactColumnToDataActionFieldMappings, ok := DialerrulesetconfigchangeconditionMap["contactColumnToDataActionFieldMappings"].([]interface{}); ok {
		ContactColumnToDataActionFieldMappingsString, _ := json.Marshal(ContactColumnToDataActionFieldMappings)
		json.Unmarshal(ContactColumnToDataActionFieldMappingsString, &o.ContactColumnToDataActionFieldMappings)
	}
	
	if Predicates, ok := DialerrulesetconfigchangeconditionMap["predicates"].([]interface{}); ok {
		PredicatesString, _ := json.Marshal(Predicates)
		json.Unmarshal(PredicatesString, &o.Predicates)
	}
	
	if DataAction, ok := DialerrulesetconfigchangeconditionMap["dataAction"].(map[string]interface{}); ok {
		DataActionString, _ := json.Marshal(DataAction)
		json.Unmarshal(DataActionString, &o.DataAction)
	}
	
	if AdditionalProperties, ok := DialerrulesetconfigchangeconditionMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialerrulesetconfigchangecondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
