package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Fieldlist
type Fieldlist struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// CustomLabels
	CustomLabels *bool `json:"customLabels,omitempty"`

	// InstructionText
	InstructionText *string `json:"instructionText,omitempty"`

	// Key
	Key *string `json:"key,omitempty"`

	// LabelKeys
	LabelKeys *[]string `json:"labelKeys,omitempty"`

	// Params
	Params *map[string]interface{} `json:"params,omitempty"`

	// Repeatable
	Repeatable *bool `json:"repeatable,omitempty"`

	// State
	State *string `json:"state,omitempty"`

	// VarType
	VarType *string `json:"type,omitempty"`

	// Required
	Required *bool `json:"required,omitempty"`

	// Gdpr
	Gdpr *bool `json:"gdpr,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Fieldlist) SetField(field string, fieldValue interface{}) {
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

func (o Fieldlist) MarshalJSON() ([]byte, error) {
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
	type Alias Fieldlist
	
	return json.Marshal(&struct { 
		CustomLabels *bool `json:"customLabels,omitempty"`
		
		InstructionText *string `json:"instructionText,omitempty"`
		
		Key *string `json:"key,omitempty"`
		
		LabelKeys *[]string `json:"labelKeys,omitempty"`
		
		Params *map[string]interface{} `json:"params,omitempty"`
		
		Repeatable *bool `json:"repeatable,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Required *bool `json:"required,omitempty"`
		
		Gdpr *bool `json:"gdpr,omitempty"`
		Alias
	}{ 
		CustomLabels: o.CustomLabels,
		
		InstructionText: o.InstructionText,
		
		Key: o.Key,
		
		LabelKeys: o.LabelKeys,
		
		Params: o.Params,
		
		Repeatable: o.Repeatable,
		
		State: o.State,
		
		VarType: o.VarType,
		
		Required: o.Required,
		
		Gdpr: o.Gdpr,
		Alias:    (Alias)(o),
	})
}

func (o *Fieldlist) UnmarshalJSON(b []byte) error {
	var FieldlistMap map[string]interface{}
	err := json.Unmarshal(b, &FieldlistMap)
	if err != nil {
		return err
	}
	
	if CustomLabels, ok := FieldlistMap["customLabels"].(bool); ok {
		o.CustomLabels = &CustomLabels
	}
    
	if InstructionText, ok := FieldlistMap["instructionText"].(string); ok {
		o.InstructionText = &InstructionText
	}
    
	if Key, ok := FieldlistMap["key"].(string); ok {
		o.Key = &Key
	}
    
	if LabelKeys, ok := FieldlistMap["labelKeys"].([]interface{}); ok {
		LabelKeysString, _ := json.Marshal(LabelKeys)
		json.Unmarshal(LabelKeysString, &o.LabelKeys)
	}
	
	if Params, ok := FieldlistMap["params"].(map[string]interface{}); ok {
		ParamsString, _ := json.Marshal(Params)
		json.Unmarshal(ParamsString, &o.Params)
	}
	
	if Repeatable, ok := FieldlistMap["repeatable"].(bool); ok {
		o.Repeatable = &Repeatable
	}
    
	if State, ok := FieldlistMap["state"].(string); ok {
		o.State = &State
	}
    
	if VarType, ok := FieldlistMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Required, ok := FieldlistMap["required"].(bool); ok {
		o.Required = &Required
	}
    
	if Gdpr, ok := FieldlistMap["gdpr"].(bool); ok {
		o.Gdpr = &Gdpr
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Fieldlist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
