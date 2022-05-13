package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerrulesetconfigchangecondition
type Dialerrulesetconfigchangecondition struct { 
	// DataAction - A UriReference for a resource
	DataAction *Dialerrulesetconfigchangeurireference `json:"dataAction,omitempty"`


	// VarType - The type of the condition
	VarType *string `json:"type,omitempty"`


	// Inverted - Indicates whether to evaluate for the opposite of the stated condition; default is false
	Inverted *bool `json:"inverted,omitempty"`


	// AttributeName - An attribute name associated with the condition (applies only to certain rule conditions)
	AttributeName *string `json:"attributeName,omitempty"`


	// Value - A value associated with the condition
	Value *string `json:"value,omitempty"`


	// ValueType - Determines the type of the value associated with the condition
	ValueType *string `json:"valueType,omitempty"`


	// Operator - An operation type for condition evaluation
	Operator *string `json:"operator,omitempty"`


	// Codes - List of wrap-up code identifiers (used only in conditions of type 'wrapupCondition')
	Codes *[]string `json:"codes,omitempty"`


	// PropertyType - Determines the type of the property associated with the condition
	PropertyType *string `json:"propertyType,omitempty"`


	// Property - A value associated with the property type of this condition
	Property *string `json:"property,omitempty"`


	// DataNotFoundResolution - The result of this condition if the data action returns a result indicating there was no data. Required for a DataActionCondition.
	DataNotFoundResolution *bool `json:"dataNotFoundResolution,omitempty"`


	// ContactIdField - The input field from the data action that the contactId will be passed to for this condition. Valid for a dataActionCondition.
	ContactIdField *string `json:"contactIdField,omitempty"`


	// CallAnalysisResultField - The input field from the data action that the callAnalysisResult will be passed to for this condition. Valid for a wrapup dataActionCondition.
	CallAnalysisResultField *string `json:"callAnalysisResultField,omitempty"`


	// AgentWrapupField - The input field from the data action that the agentWrapup will be passed to for this condition. Valid for a wrapup dataActionCondition.
	AgentWrapupField *string `json:"agentWrapupField,omitempty"`


	// ContactColumnToDataActionFieldMappings - A list of mappings defining which contact data fields will be passed to which data action input fields for this condition. Valid for a dataActionCondition.
	ContactColumnToDataActionFieldMappings *[]Dialerrulesetconfigchangecontactcolumntodataactionfieldmapping `json:"contactColumnToDataActionFieldMappings,omitempty"`


	// Predicates - A list of predicates defining the comparisons to use for this condition. Required for a dataActionCondition.
	Predicates *[]Dialerrulesetconfigchangedataactionconditionpredicate `json:"predicates,omitempty"`

}

func (o *Dialerrulesetconfigchangecondition) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialerrulesetconfigchangecondition
	
	return json.Marshal(&struct { 
		DataAction *Dialerrulesetconfigchangeurireference `json:"dataAction,omitempty"`
		
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
		*Alias
	}{ 
		DataAction: o.DataAction,
		
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
		Alias:    (*Alias)(o),
	})
}

func (o *Dialerrulesetconfigchangecondition) UnmarshalJSON(b []byte) error {
	var DialerrulesetconfigchangeconditionMap map[string]interface{}
	err := json.Unmarshal(b, &DialerrulesetconfigchangeconditionMap)
	if err != nil {
		return err
	}
	
	if DataAction, ok := DialerrulesetconfigchangeconditionMap["dataAction"].(map[string]interface{}); ok {
		DataActionString, _ := json.Marshal(DataAction)
		json.Unmarshal(DataActionString, &o.DataAction)
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
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialerrulesetconfigchangecondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
