package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Agenticvirtualagenttypedefinition - Type definition used by a virtual agent. The applicable fields depend on the type value and related fields.
type Agenticvirtualagenttypedefinition struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - Type name.
	Name *string `json:"name,omitempty"`

	// Description - Additional context that helps the virtual agent understand what this type is used for.
	Description *string `json:"description,omitempty"`

	// Direction - Intended direction of use for this type.
	Direction *string `json:"direction,omitempty"`

	// VarType - Type value. The applicable fields depend on this value and related fields.
	VarType *string `json:"type,omitempty"`

	// UserUtteranceSubstring - Whether values of this string type must be copied as a contiguous substring from recent user messages.
	UserUtteranceSubstring *bool `json:"userUtteranceSubstring,omitempty"`

	// Undisclosed - Whether values of this string type are hidden from the virtual agent and represented as opaque identifiers. Only valid when type is string.
	Undisclosed *bool `json:"undisclosed,omitempty"`

	// Properties - Properties of this object type. Applies when type is object.
	Properties *[]Agenticvirtualagentpropertydefinition `json:"properties,omitempty"`

	// Items - Type of items in this array type. Applies when type is array.
	Items *string `json:"items,omitempty"`

	// StatusCodes - HTTP 4xx or 5xx status codes this error type can handle. Applies when type is DataActionHttpError.
	StatusCodes *[]int `json:"statusCodes,omitempty"`

	// DefaultInstruction - Default instruction for how the virtual agent should handle this error type when a tool references it without its own error instruction. Applies when type is DataActionHttpError.
	DefaultInstruction *string `json:"defaultInstruction,omitempty"`

	// Enum - Allowed enum values. Applies to enum types.
	Enum *[]string `json:"enum,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Agenticvirtualagenttypedefinition) SetField(field string, fieldValue interface{}) {
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

func (o Agenticvirtualagenttypedefinition) MarshalJSON() ([]byte, error) {
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
	type Alias Agenticvirtualagenttypedefinition
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		UserUtteranceSubstring *bool `json:"userUtteranceSubstring,omitempty"`
		
		Undisclosed *bool `json:"undisclosed,omitempty"`
		
		Properties *[]Agenticvirtualagentpropertydefinition `json:"properties,omitempty"`
		
		Items *string `json:"items,omitempty"`
		
		StatusCodes *[]int `json:"statusCodes,omitempty"`
		
		DefaultInstruction *string `json:"defaultInstruction,omitempty"`
		
		Enum *[]string `json:"enum,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		Description: o.Description,
		
		Direction: o.Direction,
		
		VarType: o.VarType,
		
		UserUtteranceSubstring: o.UserUtteranceSubstring,
		
		Undisclosed: o.Undisclosed,
		
		Properties: o.Properties,
		
		Items: o.Items,
		
		StatusCodes: o.StatusCodes,
		
		DefaultInstruction: o.DefaultInstruction,
		
		Enum: o.Enum,
		Alias:    (Alias)(o),
	})
}

func (o *Agenticvirtualagenttypedefinition) UnmarshalJSON(b []byte) error {
	var AgenticvirtualagenttypedefinitionMap map[string]interface{}
	err := json.Unmarshal(b, &AgenticvirtualagenttypedefinitionMap)
	if err != nil {
		return err
	}
	
	if Name, ok := AgenticvirtualagenttypedefinitionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := AgenticvirtualagenttypedefinitionMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Direction, ok := AgenticvirtualagenttypedefinitionMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if VarType, ok := AgenticvirtualagenttypedefinitionMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if UserUtteranceSubstring, ok := AgenticvirtualagenttypedefinitionMap["userUtteranceSubstring"].(bool); ok {
		o.UserUtteranceSubstring = &UserUtteranceSubstring
	}
    
	if Undisclosed, ok := AgenticvirtualagenttypedefinitionMap["undisclosed"].(bool); ok {
		o.Undisclosed = &Undisclosed
	}
    
	if Properties, ok := AgenticvirtualagenttypedefinitionMap["properties"].([]interface{}); ok {
		PropertiesString, _ := json.Marshal(Properties)
		json.Unmarshal(PropertiesString, &o.Properties)
	}
	
	if Items, ok := AgenticvirtualagenttypedefinitionMap["items"].(string); ok {
		o.Items = &Items
	}
    
	if StatusCodes, ok := AgenticvirtualagenttypedefinitionMap["statusCodes"].([]interface{}); ok {
		StatusCodesString, _ := json.Marshal(StatusCodes)
		json.Unmarshal(StatusCodesString, &o.StatusCodes)
	}
	
	if DefaultInstruction, ok := AgenticvirtualagenttypedefinitionMap["defaultInstruction"].(string); ok {
		o.DefaultInstruction = &DefaultInstruction
	}
    
	if Enum, ok := AgenticvirtualagenttypedefinitionMap["enum"].([]interface{}); ok {
		EnumString, _ := json.Marshal(Enum)
		json.Unmarshal(EnumString, &o.Enum)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Agenticvirtualagenttypedefinition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
