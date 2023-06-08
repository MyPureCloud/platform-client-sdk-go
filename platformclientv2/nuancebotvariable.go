package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Nuancebotvariable - Model for a Nuance bot variable
type Nuancebotvariable struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The variable ID
	Id *string `json:"id,omitempty"`

	// Name - The variable display name
	Name *string `json:"name,omitempty"`

	// Description - The variable description
	Description *string `json:"description,omitempty"`

	// Reserved - True if the variable is a reserved variable
	Reserved *bool `json:"reserved,omitempty"`

	// SimpleVariableInfo - The type information for this variable
	SimpleVariableInfo *string `json:"simpleVariableInfo,omitempty"`

	// ComplexGenericVariableInfo - The type information for this variable
	ComplexGenericVariableInfo *Complexvariableinfo `json:"complexGenericVariableInfo,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Nuancebotvariable) SetField(field string, fieldValue interface{}) {
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

func (o Nuancebotvariable) MarshalJSON() ([]byte, error) {
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
	type Alias Nuancebotvariable
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Reserved *bool `json:"reserved,omitempty"`
		
		SimpleVariableInfo *string `json:"simpleVariableInfo,omitempty"`
		
		ComplexGenericVariableInfo *Complexvariableinfo `json:"complexGenericVariableInfo,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Reserved: o.Reserved,
		
		SimpleVariableInfo: o.SimpleVariableInfo,
		
		ComplexGenericVariableInfo: o.ComplexGenericVariableInfo,
		Alias:    (Alias)(o),
	})
}

func (o *Nuancebotvariable) UnmarshalJSON(b []byte) error {
	var NuancebotvariableMap map[string]interface{}
	err := json.Unmarshal(b, &NuancebotvariableMap)
	if err != nil {
		return err
	}
	
	if Id, ok := NuancebotvariableMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := NuancebotvariableMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := NuancebotvariableMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Reserved, ok := NuancebotvariableMap["reserved"].(bool); ok {
		o.Reserved = &Reserved
	}
    
	if SimpleVariableInfo, ok := NuancebotvariableMap["simpleVariableInfo"].(string); ok {
		o.SimpleVariableInfo = &SimpleVariableInfo
	}
    
	if ComplexGenericVariableInfo, ok := NuancebotvariableMap["complexGenericVariableInfo"].(map[string]interface{}); ok {
		ComplexGenericVariableInfoString, _ := json.Marshal(ComplexGenericVariableInfo)
		json.Unmarshal(ComplexGenericVariableInfoString, &o.ComplexGenericVariableInfo)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Nuancebotvariable) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
