package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Skillexpressionvalidationresult - Result of skill expression validation
type Skillexpressionvalidationresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Valid - Whether the expression is valid
	Valid *bool `json:"valid,omitempty"`

	// Expression - Normalized SpEL expression (null if validation failed)
	Expression *string `json:"expression,omitempty"`

	// Skills - List of skill references extracted from the expression (empty if no skills found and/or invalid expression)
	Skills *[]Skillreference `json:"skills,omitempty"`

	// Errors - List of validation errors (empty if valid)
	Errors *[]Skillexpressionvalidationerror `json:"errors,omitempty"`

	// Hint - Optional hint message (e.g., if expression is non-optimal or system is near capacity)
	Hint *string `json:"hint,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Skillexpressionvalidationresult) SetField(field string, fieldValue interface{}) {
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

func (o Skillexpressionvalidationresult) MarshalJSON() ([]byte, error) {
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
	type Alias Skillexpressionvalidationresult
	
	return json.Marshal(&struct { 
		Valid *bool `json:"valid,omitempty"`
		
		Expression *string `json:"expression,omitempty"`
		
		Skills *[]Skillreference `json:"skills,omitempty"`
		
		Errors *[]Skillexpressionvalidationerror `json:"errors,omitempty"`
		
		Hint *string `json:"hint,omitempty"`
		Alias
	}{ 
		Valid: o.Valid,
		
		Expression: o.Expression,
		
		Skills: o.Skills,
		
		Errors: o.Errors,
		
		Hint: o.Hint,
		Alias:    (Alias)(o),
	})
}

func (o *Skillexpressionvalidationresult) UnmarshalJSON(b []byte) error {
	var SkillexpressionvalidationresultMap map[string]interface{}
	err := json.Unmarshal(b, &SkillexpressionvalidationresultMap)
	if err != nil {
		return err
	}
	
	if Valid, ok := SkillexpressionvalidationresultMap["valid"].(bool); ok {
		o.Valid = &Valid
	}
    
	if Expression, ok := SkillexpressionvalidationresultMap["expression"].(string); ok {
		o.Expression = &Expression
	}
    
	if Skills, ok := SkillexpressionvalidationresultMap["skills"].([]interface{}); ok {
		SkillsString, _ := json.Marshal(Skills)
		json.Unmarshal(SkillsString, &o.Skills)
	}
	
	if Errors, ok := SkillexpressionvalidationresultMap["errors"].([]interface{}); ok {
		ErrorsString, _ := json.Marshal(Errors)
		json.Unmarshal(ErrorsString, &o.Errors)
	}
	
	if Hint, ok := SkillexpressionvalidationresultMap["hint"].(string); ok {
		o.Hint = &Hint
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Skillexpressionvalidationresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
