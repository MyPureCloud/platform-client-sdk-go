package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Responseassetfilter
type Responseassetfilter struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EndValue - The end value of the range. This field is used for range search types. Accepts numbers and date in ISO8601 format
	EndValue *string `json:"endValue,omitempty"`

	// Values - A list of values for the search to match against
	Values *[]string `json:"values,omitempty"`

	// StartValue - The start value of the range. This field is used for range search types. Accepts numbers and date in ISO8601 format
	StartValue *string `json:"startValue,omitempty"`

	// Fields - Field name to search against. Allowed Values: divisionId, name, contentLength, contentType, dateCreated
	Fields *[]string `json:"fields,omitempty"`

	// Value - A value for the search to match against
	Value *string `json:"value,omitempty"`

	// VarType - How to apply this search criteria against other criteria. Filter type supported for each field:- name:[STARTS_WITH, TERM], divisionId:[TERM, TERMS], contentLength:[RANGE, GREATER_THAN_EQUAL_TO, LESS_THAN_EQUAL_TO], contentType:[STARTS_WITH, TERM] dateCreated:[DATE_RANGE]
	VarType *string `json:"type,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Responseassetfilter) SetField(field string, fieldValue interface{}) {
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

func (o Responseassetfilter) MarshalJSON() ([]byte, error) {
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
	type Alias Responseassetfilter
	
	return json.Marshal(&struct { 
		EndValue *string `json:"endValue,omitempty"`
		
		Values *[]string `json:"values,omitempty"`
		
		StartValue *string `json:"startValue,omitempty"`
		
		Fields *[]string `json:"fields,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		Alias
	}{ 
		EndValue: o.EndValue,
		
		Values: o.Values,
		
		StartValue: o.StartValue,
		
		Fields: o.Fields,
		
		Value: o.Value,
		
		VarType: o.VarType,
		Alias:    (Alias)(o),
	})
}

func (o *Responseassetfilter) UnmarshalJSON(b []byte) error {
	var ResponseassetfilterMap map[string]interface{}
	err := json.Unmarshal(b, &ResponseassetfilterMap)
	if err != nil {
		return err
	}
	
	if EndValue, ok := ResponseassetfilterMap["endValue"].(string); ok {
		o.EndValue = &EndValue
	}
    
	if Values, ok := ResponseassetfilterMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	
	if StartValue, ok := ResponseassetfilterMap["startValue"].(string); ok {
		o.StartValue = &StartValue
	}
    
	if Fields, ok := ResponseassetfilterMap["fields"].([]interface{}); ok {
		FieldsString, _ := json.Marshal(Fields)
		json.Unmarshal(FieldsString, &o.Fields)
	}
	
	if Value, ok := ResponseassetfilterMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if VarType, ok := ResponseassetfilterMap["type"].(string); ok {
		o.VarType = &VarType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Responseassetfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
