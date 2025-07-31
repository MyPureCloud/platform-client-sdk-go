package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerrulesetconfigchangecondition
type Dialerrulesetconfigchangecondition struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// DataAction - A UriReference for a resource
	DataAction *Dialerrulesetconfigchangeurireference `json:"dataAction,omitempty"`

	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

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

	// GetAdditionalProperties
	GetAdditionalProperties *map[string]interface{} `json:"getAdditionalProperties,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dialerrulesetconfigchangecondition) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Dialerrulesetconfigchangecondition) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialerrulesetconfigchangecondition
	
	return json.Marshal(&struct { 
		DataAction *Dialerrulesetconfigchangeurireference `json:"dataAction,omitempty"`
		
		AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
		
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
		
		GetAdditionalProperties *map[string]interface{} `json:"getAdditionalProperties,omitempty"`
		Alias
	}{ 
		DataAction: o.DataAction,
		
		AdditionalProperties: o.AdditionalProperties,
		
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
		
		GetAdditionalProperties: o.GetAdditionalProperties,
		Alias:    (Alias)(o),
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
	
	if AdditionalProperties, ok := DialerrulesetconfigchangeconditionMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
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
	
	if GetAdditionalProperties, ok := DialerrulesetconfigchangeconditionMap["getAdditionalProperties"].(map[string]interface{}); ok {
		GetAdditionalPropertiesString, _ := json.Marshal(GetAdditionalProperties)
		json.Unmarshal(GetAdditionalPropertiesString, &o.GetAdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialerrulesetconfigchangecondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
