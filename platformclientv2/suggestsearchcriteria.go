package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Suggestsearchcriteria
type Suggestsearchcriteria struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EndValue - The end value of the range. This field is used for range search types.
	EndValue *string `json:"endValue,omitempty"`

	// Values - A list of values for the search to match against
	Values *[]string `json:"values,omitempty"`

	// StartValue - The start value of the range. This field is used for range search types.
	StartValue *string `json:"startValue,omitempty"`

	// Value - A value for the search to match against
	Value *string `json:"value,omitempty"`

	// Operator - How to apply this search criteria against other criteria
	Operator *string `json:"operator,omitempty"`

	// Group - Groups multiple conditions
	Group *[]Suggestsearchcriteria `json:"group,omitempty"`

	// DateFormat - Set date format for criteria values when using date range search type.  Supports Java date format syntax, example yyyy-MM-dd'T'HH:mm:ss.SSSX.
	DateFormat *string `json:"dateFormat,omitempty"`

	// Fields - Field names to search against
	Fields *[]string `json:"fields,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Suggestsearchcriteria) SetField(field string, fieldValue interface{}) {
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

func (o Suggestsearchcriteria) MarshalJSON() ([]byte, error) {
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
	type Alias Suggestsearchcriteria
	
	return json.Marshal(&struct { 
		EndValue *string `json:"endValue,omitempty"`
		
		Values *[]string `json:"values,omitempty"`
		
		StartValue *string `json:"startValue,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Group *[]Suggestsearchcriteria `json:"group,omitempty"`
		
		DateFormat *string `json:"dateFormat,omitempty"`
		
		Fields *[]string `json:"fields,omitempty"`
		Alias
	}{ 
		EndValue: o.EndValue,
		
		Values: o.Values,
		
		StartValue: o.StartValue,
		
		Value: o.Value,
		
		Operator: o.Operator,
		
		Group: o.Group,
		
		DateFormat: o.DateFormat,
		
		Fields: o.Fields,
		Alias:    (Alias)(o),
	})
}

func (o *Suggestsearchcriteria) UnmarshalJSON(b []byte) error {
	var SuggestsearchcriteriaMap map[string]interface{}
	err := json.Unmarshal(b, &SuggestsearchcriteriaMap)
	if err != nil {
		return err
	}
	
	if EndValue, ok := SuggestsearchcriteriaMap["endValue"].(string); ok {
		o.EndValue = &EndValue
	}
    
	if Values, ok := SuggestsearchcriteriaMap["values"].([]interface{}); ok {
		ValuesString, _ := json.Marshal(Values)
		json.Unmarshal(ValuesString, &o.Values)
	}
	
	if StartValue, ok := SuggestsearchcriteriaMap["startValue"].(string); ok {
		o.StartValue = &StartValue
	}
    
	if Value, ok := SuggestsearchcriteriaMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if Operator, ok := SuggestsearchcriteriaMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Group, ok := SuggestsearchcriteriaMap["group"].([]interface{}); ok {
		GroupString, _ := json.Marshal(Group)
		json.Unmarshal(GroupString, &o.Group)
	}
	
	if DateFormat, ok := SuggestsearchcriteriaMap["dateFormat"].(string); ok {
		o.DateFormat = &DateFormat
	}
    
	if Fields, ok := SuggestsearchcriteriaMap["fields"].([]interface{}); ok {
		FieldsString, _ := json.Marshal(Fields)
		json.Unmarshal(FieldsString, &o.Fields)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Suggestsearchcriteria) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
