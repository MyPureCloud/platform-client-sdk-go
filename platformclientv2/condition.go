package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Condition
type Condition struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType - The type of the condition.
	VarType *string `json:"type,omitempty"`

	// Inverted - If true, inverts the result of evaluating this Condition. Default is false.
	Inverted *bool `json:"inverted,omitempty"`

	// AttributeName - An attribute name associated with this Condition. Required for a contactAttributeCondition.
	AttributeName *string `json:"attributeName,omitempty"`

	// Value - A value associated with this Condition. This could be text, a number, or a relative time. Not used for a DataActionCondition.
	Value *string `json:"value,omitempty"`

	// ValueType - The type of the value associated with this Condition. Not used for a DataActionCondition.
	ValueType *string `json:"valueType,omitempty"`

	// Operator - An operation with which to evaluate the Condition. Not used for a DataActionCondition.
	Operator *string `json:"operator,omitempty"`

	// Codes - List of wrap-up code identifiers. Required for a wrapupCondition.
	Codes *[]string `json:"codes,omitempty"`

	// Property - A value associated with the property type of this Condition. Required for a contactPropertyCondition.
	Property *string `json:"property,omitempty"`

	// PropertyType - The type of the property associated with this Condition. Required for a contactPropertyCondition.
	PropertyType *string `json:"propertyType,omitempty"`

	// DataAction - The Data Action to use for this condition. Required for a dataActionCondition.
	DataAction *Domainentityref `json:"dataAction,omitempty"`

	// DataNotFoundResolution - The result of this condition if the data action returns a result indicating there was no data. Required for a DataActionCondition.
	DataNotFoundResolution *bool `json:"dataNotFoundResolution,omitempty"`

	// ContactIdField - The input field from the data action that the contactId will be passed to for this condition. Valid for a dataActionCondition.
	ContactIdField *string `json:"contactIdField,omitempty"`

	// CallAnalysisResultField - The input field from the data action that the callAnalysisResult will be passed to for this condition. Valid for a wrapup dataActionCondition.
	CallAnalysisResultField *string `json:"callAnalysisResultField,omitempty"`

	// AgentWrapupField - The input field from the data action that the agentWrapup will be passed to for this condition. Valid for a wrapup dataActionCondition.
	AgentWrapupField *string `json:"agentWrapupField,omitempty"`

	// ContactColumnToDataActionFieldMappings - A list of mappings defining which contact data fields will be passed to which data action input fields for this condition. Valid for a dataActionCondition.
	ContactColumnToDataActionFieldMappings *[]Contactcolumntodataactionfieldmapping `json:"contactColumnToDataActionFieldMappings,omitempty"`

	// Predicates - A list of predicates defining the comparisons to use for this condition. Required for a dataActionCondition.
	Predicates *[]Dataactionconditionpredicate `json:"predicates,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Condition) SetField(field string, fieldValue interface{}) {
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

func (o Condition) MarshalJSON() ([]byte, error) {
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
			if contains(dateTimeFields, fieldName) {
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
	type Alias Condition
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Inverted *bool `json:"inverted,omitempty"`
		
		AttributeName *string `json:"attributeName,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		ValueType *string `json:"valueType,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Codes *[]string `json:"codes,omitempty"`
		
		Property *string `json:"property,omitempty"`
		
		PropertyType *string `json:"propertyType,omitempty"`
		
		DataAction *Domainentityref `json:"dataAction,omitempty"`
		
		DataNotFoundResolution *bool `json:"dataNotFoundResolution,omitempty"`
		
		ContactIdField *string `json:"contactIdField,omitempty"`
		
		CallAnalysisResultField *string `json:"callAnalysisResultField,omitempty"`
		
		AgentWrapupField *string `json:"agentWrapupField,omitempty"`
		
		ContactColumnToDataActionFieldMappings *[]Contactcolumntodataactionfieldmapping `json:"contactColumnToDataActionFieldMappings,omitempty"`
		
		Predicates *[]Dataactionconditionpredicate `json:"predicates,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		Inverted: o.Inverted,
		
		AttributeName: o.AttributeName,
		
		Value: o.Value,
		
		ValueType: o.ValueType,
		
		Operator: o.Operator,
		
		Codes: o.Codes,
		
		Property: o.Property,
		
		PropertyType: o.PropertyType,
		
		DataAction: o.DataAction,
		
		DataNotFoundResolution: o.DataNotFoundResolution,
		
		ContactIdField: o.ContactIdField,
		
		CallAnalysisResultField: o.CallAnalysisResultField,
		
		AgentWrapupField: o.AgentWrapupField,
		
		ContactColumnToDataActionFieldMappings: o.ContactColumnToDataActionFieldMappings,
		
		Predicates: o.Predicates,
		Alias:    (Alias)(o),
	})
}

func (o *Condition) UnmarshalJSON(b []byte) error {
	var ConditionMap map[string]interface{}
	err := json.Unmarshal(b, &ConditionMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := ConditionMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Inverted, ok := ConditionMap["inverted"].(bool); ok {
		o.Inverted = &Inverted
	}
    
	if AttributeName, ok := ConditionMap["attributeName"].(string); ok {
		o.AttributeName = &AttributeName
	}
    
	if Value, ok := ConditionMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if ValueType, ok := ConditionMap["valueType"].(string); ok {
		o.ValueType = &ValueType
	}
    
	if Operator, ok := ConditionMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Codes, ok := ConditionMap["codes"].([]interface{}); ok {
		CodesString, _ := json.Marshal(Codes)
		json.Unmarshal(CodesString, &o.Codes)
	}
	
	if Property, ok := ConditionMap["property"].(string); ok {
		o.Property = &Property
	}
    
	if PropertyType, ok := ConditionMap["propertyType"].(string); ok {
		o.PropertyType = &PropertyType
	}
    
	if DataAction, ok := ConditionMap["dataAction"].(map[string]interface{}); ok {
		DataActionString, _ := json.Marshal(DataAction)
		json.Unmarshal(DataActionString, &o.DataAction)
	}
	
	if DataNotFoundResolution, ok := ConditionMap["dataNotFoundResolution"].(bool); ok {
		o.DataNotFoundResolution = &DataNotFoundResolution
	}
    
	if ContactIdField, ok := ConditionMap["contactIdField"].(string); ok {
		o.ContactIdField = &ContactIdField
	}
    
	if CallAnalysisResultField, ok := ConditionMap["callAnalysisResultField"].(string); ok {
		o.CallAnalysisResultField = &CallAnalysisResultField
	}
    
	if AgentWrapupField, ok := ConditionMap["agentWrapupField"].(string); ok {
		o.AgentWrapupField = &AgentWrapupField
	}
    
	if ContactColumnToDataActionFieldMappings, ok := ConditionMap["contactColumnToDataActionFieldMappings"].([]interface{}); ok {
		ContactColumnToDataActionFieldMappingsString, _ := json.Marshal(ContactColumnToDataActionFieldMappings)
		json.Unmarshal(ContactColumnToDataActionFieldMappingsString, &o.ContactColumnToDataActionFieldMappings)
	}
	
	if Predicates, ok := ConditionMap["predicates"].([]interface{}); ok {
		PredicatesString, _ := json.Marshal(Predicates)
		json.Unmarshal(PredicatesString, &o.Predicates)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Condition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
