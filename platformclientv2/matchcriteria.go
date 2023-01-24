package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Matchcriteria - Defines a simple matching condition
type Matchcriteria struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// JsonPath - The Goessner json path of the field to match
	JsonPath *string `json:"jsonPath,omitempty"`

	// Operator - The type of operation to perform for matching check
	Operator *string `json:"operator,omitempty"`

	// Value - The value to match on. Only one of value and values can be included
	Value *interface{} `json:"value,omitempty"`

	// Values - The list of values to match on. Only one of value and values can be included
	Values *[]map[string]interface{} `json:"values,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Matchcriteria) SetField(field string, fieldValue interface{}) {
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

func (o Matchcriteria) MarshalJSON() ([]byte, error) {
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
	type Alias Matchcriteria
	
	return json.Marshal(&struct { 
		JsonPath *string `json:"jsonPath,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Value *interface{} `json:"value,omitempty"`
		
		Values *[]map[string]interface{} `json:"values,omitempty"`
		Alias
	}{ 
		JsonPath: o.JsonPath,
		
		Operator: o.Operator,
		
		Value: o.Value,
		
		Values: o.Values,
		Alias:    (Alias)(o),
	})
}

func (o *Matchcriteria) UnmarshalJSON(b []byte) error {
	var MatchcriteriaMap map[string]interface{}
	err := json.Unmarshal(b, &MatchcriteriaMap)
	if err != nil {
		return err
	}
	
	if JsonPath, ok := MatchcriteriaMap["jsonPath"].(string); ok {
		o.JsonPath = &JsonPath
	}
    
	if Operator, ok := MatchcriteriaMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Value, ok := MatchcriteriaMap["value"].(map[string]interface{}); ok {
		ValueString, _ := json.Marshal(Value)
		json.Unmarshal(ValueString, &o.Value)
	}
	
	if Values, ok := MatchcriteriaMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Matchcriteria) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
