package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Agenticvirtualagenttoolinput - Input for a tool.
type Agenticvirtualagenttoolinput struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// TargetName - The unique name that identifies this input parameter within the tool
	TargetName *string `json:"targetName,omitempty"`

	// VarType - Input type name. The valid referenced type depends on the input source.
	VarType *string `json:"type,omitempty"`

	// Source - Source of the input value.
	Source *string `json:"source,omitempty"`

	// Required - Whether this input must be supplied.
	Required *bool `json:"required,omitempty"`

	// FallbackToUser - Whether the virtual agent should ask the user for this input value when it is not available from the configured source.
	FallbackToUser *bool `json:"fallbackToUser,omitempty"`

	// Mapping - Path used to extract this input from a previous tool output. Only valid when source is 'ToolOutput'. The path starts with a tool output type name, may contain only string property names or integer array indexes, and must resolve to a primitive value.
	Mapping *[]interface{} `json:"mapping,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Agenticvirtualagenttoolinput) SetField(field string, fieldValue interface{}) {
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

func (o Agenticvirtualagenttoolinput) MarshalJSON() ([]byte, error) {
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
	type Alias Agenticvirtualagenttoolinput
	
	return json.Marshal(&struct { 
		TargetName *string `json:"targetName,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Source *string `json:"source,omitempty"`
		
		Required *bool `json:"required,omitempty"`
		
		FallbackToUser *bool `json:"fallbackToUser,omitempty"`
		
		Mapping *[]interface{} `json:"mapping,omitempty"`
		Alias
	}{ 
		TargetName: o.TargetName,
		
		VarType: o.VarType,
		
		Source: o.Source,
		
		Required: o.Required,
		
		FallbackToUser: o.FallbackToUser,
		
		Mapping: o.Mapping,
		Alias:    (Alias)(o),
	})
}

func (o *Agenticvirtualagenttoolinput) UnmarshalJSON(b []byte) error {
	var AgenticvirtualagenttoolinputMap map[string]interface{}
	err := json.Unmarshal(b, &AgenticvirtualagenttoolinputMap)
	if err != nil {
		return err
	}
	
	if TargetName, ok := AgenticvirtualagenttoolinputMap["targetName"].(string); ok {
		o.TargetName = &TargetName
	}
    
	if VarType, ok := AgenticvirtualagenttoolinputMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Source, ok := AgenticvirtualagenttoolinputMap["source"].(string); ok {
		o.Source = &Source
	}
    
	if Required, ok := AgenticvirtualagenttoolinputMap["required"].(bool); ok {
		o.Required = &Required
	}
    
	if FallbackToUser, ok := AgenticvirtualagenttoolinputMap["fallbackToUser"].(bool); ok {
		o.FallbackToUser = &FallbackToUser
	}
    
	if Mapping, ok := AgenticvirtualagenttoolinputMap["mapping"].([]interface{}); ok {
		MappingString, _ := json.Marshal(Mapping)
		json.Unmarshal(MappingString, &o.Mapping)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Agenticvirtualagenttoolinput) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
