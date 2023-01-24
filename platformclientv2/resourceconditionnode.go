package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Resourceconditionnode
type Resourceconditionnode struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VariableName
	VariableName *string `json:"variableName,omitempty"`

	// Conjunction
	Conjunction *string `json:"conjunction,omitempty"`

	// Operator
	Operator *string `json:"operator,omitempty"`

	// Operands
	Operands *[]Resourceconditionvalue `json:"operands,omitempty"`

	// Terms
	Terms *[]Resourceconditionnode `json:"terms,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Resourceconditionnode) SetField(field string, fieldValue interface{}) {
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

func (o Resourceconditionnode) MarshalJSON() ([]byte, error) {
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
	type Alias Resourceconditionnode
	
	return json.Marshal(&struct { 
		VariableName *string `json:"variableName,omitempty"`
		
		Conjunction *string `json:"conjunction,omitempty"`
		
		Operator *string `json:"operator,omitempty"`
		
		Operands *[]Resourceconditionvalue `json:"operands,omitempty"`
		
		Terms *[]Resourceconditionnode `json:"terms,omitempty"`
		Alias
	}{ 
		VariableName: o.VariableName,
		
		Conjunction: o.Conjunction,
		
		Operator: o.Operator,
		
		Operands: o.Operands,
		
		Terms: o.Terms,
		Alias:    (Alias)(o),
	})
}

func (o *Resourceconditionnode) UnmarshalJSON(b []byte) error {
	var ResourceconditionnodeMap map[string]interface{}
	err := json.Unmarshal(b, &ResourceconditionnodeMap)
	if err != nil {
		return err
	}
	
	if VariableName, ok := ResourceconditionnodeMap["variableName"].(string); ok {
		o.VariableName = &VariableName
	}
    
	if Conjunction, ok := ResourceconditionnodeMap["conjunction"].(string); ok {
		o.Conjunction = &Conjunction
	}
    
	if Operator, ok := ResourceconditionnodeMap["operator"].(string); ok {
		o.Operator = &Operator
	}
    
	if Operands, ok := ResourceconditionnodeMap["operands"].([]interface{}); ok {
		OperandsString, _ := json.Marshal(Operands)
		json.Unmarshal(OperandsString, &o.Operands)
	}
	
	if Terms, ok := ResourceconditionnodeMap["terms"].([]interface{}); ok {
		TermsString, _ := json.Marshal(Terms)
		json.Unmarshal(TermsString, &o.Terms)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Resourceconditionnode) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
