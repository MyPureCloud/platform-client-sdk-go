package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Namedentitytypemechanism
type Namedentitytypemechanism struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Items - The items that define the named entity type.
	Items *[]Namedentitytypeitem `json:"items,omitempty"`

	// Restricted - Whether the named entity type is restricted to the items provided. Default: false
	Restricted *bool `json:"restricted,omitempty"`

	// VarType - The type of the mechanism.
	VarType *string `json:"type,omitempty"`

	// SubType - Subtype of detection mechanism
	SubType *string `json:"subType,omitempty"`

	// MaxLength - The maximum length of the entity resolved value
	MaxLength *int `json:"maxLength,omitempty"`

	// MinLength - The minimum length of the entity resolved value
	MinLength *int `json:"minLength,omitempty"`

	// Examples - Examples for entity detection
	Examples *[]Namedentitytypemechanismexample `json:"examples,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Namedentitytypemechanism) SetField(field string, fieldValue interface{}) {
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

func (o Namedentitytypemechanism) MarshalJSON() ([]byte, error) {
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
	type Alias Namedentitytypemechanism
	
	return json.Marshal(&struct { 
		Items *[]Namedentitytypeitem `json:"items,omitempty"`
		
		Restricted *bool `json:"restricted,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		SubType *string `json:"subType,omitempty"`
		
		MaxLength *int `json:"maxLength,omitempty"`
		
		MinLength *int `json:"minLength,omitempty"`
		
		Examples *[]Namedentitytypemechanismexample `json:"examples,omitempty"`
		Alias
	}{ 
		Items: o.Items,
		
		Restricted: o.Restricted,
		
		VarType: o.VarType,
		
		SubType: o.SubType,
		
		MaxLength: o.MaxLength,
		
		MinLength: o.MinLength,
		
		Examples: o.Examples,
		Alias:    (Alias)(o),
	})
}

func (o *Namedentitytypemechanism) UnmarshalJSON(b []byte) error {
	var NamedentitytypemechanismMap map[string]interface{}
	err := json.Unmarshal(b, &NamedentitytypemechanismMap)
	if err != nil {
		return err
	}
	
	if Items, ok := NamedentitytypemechanismMap["items"].([]interface{}); ok {
		ItemsString, _ := json.Marshal(Items)
		json.Unmarshal(ItemsString, &o.Items)
	}
	
	if Restricted, ok := NamedentitytypemechanismMap["restricted"].(bool); ok {
		o.Restricted = &Restricted
	}
    
	if VarType, ok := NamedentitytypemechanismMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if SubType, ok := NamedentitytypemechanismMap["subType"].(string); ok {
		o.SubType = &SubType
	}
    
	if MaxLength, ok := NamedentitytypemechanismMap["maxLength"].(float64); ok {
		MaxLengthInt := int(MaxLength)
		o.MaxLength = &MaxLengthInt
	}
	
	if MinLength, ok := NamedentitytypemechanismMap["minLength"].(float64); ok {
		MinLengthInt := int(MinLength)
		o.MinLength = &MinLengthInt
	}
	
	if Examples, ok := NamedentitytypemechanismMap["examples"].([]interface{}); ok {
		ExamplesString, _ := json.Marshal(Examples)
		json.Unmarshal(ExamplesString, &o.Examples)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Namedentitytypemechanism) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
