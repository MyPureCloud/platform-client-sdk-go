package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Agenticvirtualagentpropertydefinition - Property definition for an object type.
type Agenticvirtualagentpropertydefinition struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - Property name.
	Name *string `json:"name,omitempty"`

	// VarType - Property type name. The valid type depends on the containing type and related fields.
	VarType *string `json:"type,omitempty"`

	// Required - Whether this property must be supplied.
	Required *bool `json:"required,omitempty"`

	// Description - Additional context that helps the virtual agent understand what this property means.
	Description *string `json:"description,omitempty"`

	// Items - Type of items in this array property. Applies when type is array.
	Items *string `json:"items,omitempty"`

	// Mapping - Path used to extract this output data property from a tool output. Only valid for output data properties. The path starts with a tool output type name, may contain only string property names or integer array indexes, and must resolve to a primitive value.
	Mapping *[]interface{} `json:"mapping,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Agenticvirtualagentpropertydefinition) SetField(field string, fieldValue interface{}) {
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

func (o Agenticvirtualagentpropertydefinition) MarshalJSON() ([]byte, error) {
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
	type Alias Agenticvirtualagentpropertydefinition
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Required *bool `json:"required,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Items *string `json:"items,omitempty"`
		
		Mapping *[]interface{} `json:"mapping,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		VarType: o.VarType,
		
		Required: o.Required,
		
		Description: o.Description,
		
		Items: o.Items,
		
		Mapping: o.Mapping,
		Alias:    (Alias)(o),
	})
}

func (o *Agenticvirtualagentpropertydefinition) UnmarshalJSON(b []byte) error {
	var AgenticvirtualagentpropertydefinitionMap map[string]interface{}
	err := json.Unmarshal(b, &AgenticvirtualagentpropertydefinitionMap)
	if err != nil {
		return err
	}
	
	if Name, ok := AgenticvirtualagentpropertydefinitionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := AgenticvirtualagentpropertydefinitionMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Required, ok := AgenticvirtualagentpropertydefinitionMap["required"].(bool); ok {
		o.Required = &Required
	}
    
	if Description, ok := AgenticvirtualagentpropertydefinitionMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Items, ok := AgenticvirtualagentpropertydefinitionMap["items"].(string); ok {
		o.Items = &Items
	}
    
	if Mapping, ok := AgenticvirtualagentpropertydefinitionMap["mapping"].([]interface{}); ok {
		MappingString, _ := json.Marshal(Mapping)
		json.Unmarshal(MappingString, &o.Mapping)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentpropertydefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
