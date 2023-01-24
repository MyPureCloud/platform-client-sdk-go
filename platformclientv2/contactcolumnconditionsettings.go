package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactcolumnconditionsettings
type Contactcolumnconditionsettings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ColumnName - The name of the contact list column to evaluate.
	ColumnName *string `json:"columnName,omitempty"`

	// Operator - The operator to use when comparing values.
	Operator *string `json:"operator,omitempty"`

	// Value - The value to compare against the contact's data.
	Value *string `json:"value,omitempty"`

	// ValueType - The data type the value should be treated as.
	ValueType *string `json:"valueType,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contactcolumnconditionsettings) SetField(field string, fieldValue interface{}) {
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

func (o Contactcolumnconditionsettings) MarshalJSON() ([]byte, error) {
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
	type Alias Contactcolumnconditionsettings
	
	return json.Marshal(&struct { 
		ColumnName *string `json:"columnName,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		ValueType *string `json:"valueType,omitempty"`
		Alias
	}{ 
		ColumnName: o.ColumnName,
		
		Operator: o.Operator,
		
		Value: o.Value,
		
		ValueType: o.ValueType,
		Alias:    (Alias)(o),
	})
}

func (o *Contactcolumnconditionsettings) UnmarshalJSON(b []byte) error {
	var ContactcolumnconditionsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &ContactcolumnconditionsettingsMap)
	if err != nil {
		return err
	}
	
	if ColumnName, ok := ContactcolumnconditionsettingsMap["columnName"].(string); ok {
		o.ColumnName = &ColumnName
	}
    
	if Operator, ok := ContactcolumnconditionsettingsMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Value, ok := ContactcolumnconditionsettingsMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if ValueType, ok := ContactcolumnconditionsettingsMap["valueType"].(string); ok {
		o.ValueType = &ValueType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contactcolumnconditionsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
