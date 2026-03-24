package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercontactlistfilterconfigchangefilterpredicate
type Dialercontactlistfilterconfigchangefilterpredicate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Column - The contact list column
	Column *string `json:"column,omitempty"`

	// ColumnType - Whether a contact column is numeric or alphabetic
	ColumnType *string `json:"columnType,omitempty"`

	// Operator - The comparison operator
	Operator *string `json:"operator,omitempty"`

	// Value - The value the predicate applies to
	Value *string `json:"value,omitempty"`

	// VarRange
	VarRange *Dialercontactlistfilterconfigchangefilterrange `json:"range,omitempty"`

	// Inverted - Whether or not to invert to result of evaluating the predicate
	Inverted *bool `json:"inverted,omitempty"`

	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dialercontactlistfilterconfigchangefilterpredicate) SetField(field string, fieldValue interface{}) {
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

func (o Dialercontactlistfilterconfigchangefilterpredicate) MarshalJSON() ([]byte, error) {
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
	type Alias Dialercontactlistfilterconfigchangefilterpredicate
	
	return json.Marshal(&struct { 
		Column *string `json:"column,omitempty"`
		
		ColumnType *string `json:"columnType,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		VarRange *Dialercontactlistfilterconfigchangefilterrange `json:"range,omitempty"`
		
		Inverted *bool `json:"inverted,omitempty"`
		
		AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
		Alias
	}{ 
		Column: o.Column,
		
		ColumnType: o.ColumnType,
		
		Operator: o.Operator,
		
		Value: o.Value,
		
		VarRange: o.VarRange,
		
		Inverted: o.Inverted,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (Alias)(o),
	})
}

func (o *Dialercontactlistfilterconfigchangefilterpredicate) UnmarshalJSON(b []byte) error {
	var DialercontactlistfilterconfigchangefilterpredicateMap map[string]interface{}
	err := json.Unmarshal(b, &DialercontactlistfilterconfigchangefilterpredicateMap)
	if err != nil {
		return err
	}
	
	if Column, ok := DialercontactlistfilterconfigchangefilterpredicateMap["column"].(string); ok {
		o.Column = &Column
	}
    
	if ColumnType, ok := DialercontactlistfilterconfigchangefilterpredicateMap["columnType"].(string); ok {
		o.ColumnType = &ColumnType
	}
    
	if Operator, ok := DialercontactlistfilterconfigchangefilterpredicateMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Value, ok := DialercontactlistfilterconfigchangefilterpredicateMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if VarRange, ok := DialercontactlistfilterconfigchangefilterpredicateMap["range"].(map[string]interface{}); ok {
		VarRangeString, _ := json.Marshal(VarRange)
		json.Unmarshal(VarRangeString, &o.VarRange)
	}
	
	if Inverted, ok := DialercontactlistfilterconfigchangefilterpredicateMap["inverted"].(bool); ok {
		o.Inverted = &Inverted
	}
    
	if AdditionalProperties, ok := DialercontactlistfilterconfigchangefilterpredicateMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialercontactlistfilterconfigchangefilterpredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
