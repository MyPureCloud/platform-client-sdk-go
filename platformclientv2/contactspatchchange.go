package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactspatchchange
type Contactspatchchange struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Field - A JSONPath string, whose syntax is a strict subset of the JSONPath RFC 9535.  The root of the field string must be \"$.\" indicating a path from the root of the entity. You may only use dot-notation to access named fields. Examples: To select the `firstName` field of a Contact, use: \"$.firstName\".To access object fields, use the top level object field name: \"$.address\". To access nested field names, use the nested field name: \"$.address.city\". Note: trying to patch both nested fields and their parent field is not allowed and will result in a 409 error response.
	Field *string `json:"field,omitempty"`

	// Value - The value which is applied to the selected field for the patch. Acceptable types are String, Integer, Boolean, Array, Map.
	Value *interface{} `json:"value,omitempty"`

	// Action - The action of the operation.UpdateIfEmpty: Update if and only if the current value is emptyUpdate: Update the field unconditionally.UpdateIfExists: Update the field if and only if the existing field is not empty.AppendToCollection: Add new items to a collection, preserving existing data.Remove: Remove the current value unconditionally.RemoveFromCollection: Remove specified value from a collection.
	Action *string `json:"action,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contactspatchchange) SetField(field string, fieldValue interface{}) {
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

func (o Contactspatchchange) MarshalJSON() ([]byte, error) {
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
	type Alias Contactspatchchange
	
	return json.Marshal(&struct { 
		Field *string `json:"field,omitempty"`
		
		Value *interface{} `json:"value,omitempty"`
		
		Action *string `json:"action,omitempty"`
		Alias
	}{ 
		Field: o.Field,
		
		Value: o.Value,
		
		Action: o.Action,
		Alias:    (Alias)(o),
	})
}

func (o *Contactspatchchange) UnmarshalJSON(b []byte) error {
	var ContactspatchchangeMap map[string]interface{}
	err := json.Unmarshal(b, &ContactspatchchangeMap)
	if err != nil {
		return err
	}
	
	if Field, ok := ContactspatchchangeMap["field"].(string); ok {
		o.Field = &Field
	}
    
	if Value, ok := ContactspatchchangeMap["value"].(map[string]interface{}); ok {
		ValueString, _ := json.Marshal(Value)
		json.Unmarshal(ValueString, &o.Value)
	}
	
	if Action, ok := ContactspatchchangeMap["action"].(string); ok {
		o.Action = &Action
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contactspatchchange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
