package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Agenticvirtualagentbuiltininputinstruction - A built-in input instruction for a tool. Built-in instructions are stored separately from inputInstructions and merged into them at the given index when the version is published.
type Agenticvirtualagentbuiltininputinstruction struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType - The kind of built-in instruction. A tool can have at most one built-in instruction of each type.
	VarType *string `json:"type,omitempty"`

	// Index - Position of this instruction in the tool's merged input instruction list. Must be unique per tool and less than the combined count of inputInstructions and builtinInputInstructions.
	Index *int `json:"index,omitempty"`

	// Override - Author-provided instruction text that replaces the default. Omit or set to null to use the default. Maximum 4000 characters.
	Override *string `json:"override,omitempty"`

	// VarDefault - The default instruction text. Used when no override is set. Maximum 4000 characters.
	VarDefault *string `json:"default,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Agenticvirtualagentbuiltininputinstruction) SetField(field string, fieldValue interface{}) {
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

func (o Agenticvirtualagentbuiltininputinstruction) MarshalJSON() ([]byte, error) {
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
	type Alias Agenticvirtualagentbuiltininputinstruction
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Index *int `json:"index,omitempty"`
		
		Override *string `json:"override,omitempty"`
		
		VarDefault *string `json:"default,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		Index: o.Index,
		
		Override: o.Override,
		
		VarDefault: o.VarDefault,
		Alias:    (Alias)(o),
	})
}

func (o *Agenticvirtualagentbuiltininputinstruction) UnmarshalJSON(b []byte) error {
	var AgenticvirtualagentbuiltininputinstructionMap map[string]interface{}
	err := json.Unmarshal(b, &AgenticvirtualagentbuiltininputinstructionMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := AgenticvirtualagentbuiltininputinstructionMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Index, ok := AgenticvirtualagentbuiltininputinstructionMap["index"].(float64); ok {
		IndexInt := int(Index)
		o.Index = &IndexInt
	}
	
	if Override, ok := AgenticvirtualagentbuiltininputinstructionMap["override"].(string); ok {
		o.Override = &Override
	}
    
	if VarDefault, ok := AgenticvirtualagentbuiltininputinstructionMap["default"].(string); ok {
		o.VarDefault = &VarDefault
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentbuiltininputinstruction) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
