package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactlistfilterpredicate
type Contactlistfilterpredicate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Column - Contact list column from the ContactListFilter's contactList.
	Column *string `json:"column,omitempty"`

	// ColumnType - The type of data in the contact column.
	ColumnType *string `json:"columnType,omitempty"`

	// Operator - The operator for this ContactListFilterPredicate.
	Operator *string `json:"operator,omitempty"`

	// Value - Value with which to compare the contact's data. This could be text, a number, or a relative time. A value for relative time should follow the format PxxDTyyHzzM, where xx, yy, and zz specify the days, hours and minutes. For example, a value of P01DT08H30M corresponds to 1 day, 8 hours, and 30 minutes from now. To specify a time in the past, include a negative sign before each numeric value. For example, a value of P-01DT-08H-30M corresponds to 1 day, 8 hours, and 30 minutes in the past. You can also do things like P01DT00H-30M, which would correspond to 23 hours and 30 minutes from now (1 day - 30 minutes).
	Value *string `json:"value,omitempty"`

	// VarRange - A range of values. Required for operators BETWEEN and IN.
	VarRange *Contactlistfilterrange `json:"range,omitempty"`

	// Inverted - Inverts the result of the predicate (i.e., if the predicate returns true, inverting it will return false).
	Inverted *bool `json:"inverted,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contactlistfilterpredicate) SetField(field string, fieldValue interface{}) {
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

func (o Contactlistfilterpredicate) MarshalJSON() ([]byte, error) {
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
	type Alias Contactlistfilterpredicate
	
	return json.Marshal(&struct { 
		Column *string `json:"column,omitempty"`
		
		ColumnType *string `json:"columnType,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		VarRange *Contactlistfilterrange `json:"range,omitempty"`
		
		Inverted *bool `json:"inverted,omitempty"`
		Alias
	}{ 
		Column: o.Column,
		
		ColumnType: o.ColumnType,
		
		Operator: o.Operator,
		
		Value: o.Value,
		
		VarRange: o.VarRange,
		
		Inverted: o.Inverted,
		Alias:    (Alias)(o),
	})
}

func (o *Contactlistfilterpredicate) UnmarshalJSON(b []byte) error {
	var ContactlistfilterpredicateMap map[string]interface{}
	err := json.Unmarshal(b, &ContactlistfilterpredicateMap)
	if err != nil {
		return err
	}
	
	if Column, ok := ContactlistfilterpredicateMap["column"].(string); ok {
		o.Column = &Column
	}
    
	if ColumnType, ok := ContactlistfilterpredicateMap["columnType"].(string); ok {
		o.ColumnType = &ColumnType
	}
    
	if Operator, ok := ContactlistfilterpredicateMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Value, ok := ContactlistfilterpredicateMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if VarRange, ok := ContactlistfilterpredicateMap["range"].(map[string]interface{}); ok {
		VarRangeString, _ := json.Marshal(VarRange)
		json.Unmarshal(VarRangeString, &o.VarRange)
	}
	
	if Inverted, ok := ContactlistfilterpredicateMap["inverted"].(bool); ok {
		o.Inverted = &Inverted
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contactlistfilterpredicate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
