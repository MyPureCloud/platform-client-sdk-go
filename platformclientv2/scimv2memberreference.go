package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimv2memberreference - Defines a reference to SCIM group members.
type Scimv2memberreference struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType - The type of SCIM resource.
	VarType *string `json:"type,omitempty"`

	// Value - The ID of the group member. Can be \"userId\" or \"groupId\".
	Value *string `json:"value,omitempty"`

	// Ref - The reference URI of the SCIM resource.
	Ref *string `json:"$ref,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Scimv2memberreference) SetField(field string, fieldValue interface{}) {
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

func (o Scimv2memberreference) MarshalJSON() ([]byte, error) {
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
	type Alias Scimv2memberreference
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Value *string `json:"value,omitempty"`
		
		Ref *string `json:"$ref,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		Value: o.Value,
		
		Ref: o.Ref,
		Alias:    (Alias)(o),
	})
}

func (o *Scimv2memberreference) UnmarshalJSON(b []byte) error {
	var Scimv2memberreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &Scimv2memberreferenceMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := Scimv2memberreferenceMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Value, ok := Scimv2memberreferenceMap["value"].(string); ok {
		o.Value = &Value
	}
    
	if Ref, ok := Scimv2memberreferenceMap["$ref"].(string); ok {
		o.Ref = &Ref
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Scimv2memberreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
