package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Actioninput - Input requirements of Action.
type Actioninput struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// InputSchema - JSON Schema that defines the body of the request that the client (edge/architect/postman) is sending to the service, on the /execute path. If the 'flatten' query parameter is omitted or false, this field will be returned. Either inputSchema or inputSchemaFlattened will be returned, not both.
	InputSchema *Jsonschemadocument `json:"inputSchema,omitempty"`

	// InputSchemaFlattened - JSON Schema that defines the body of the request that the client (edge/architect/postman) is sending to the service, on the /execute path. The schema is transformed based on Architect's flattened format. If the 'flatten' query parameter is supplied as true, this field will be returned. Either inputSchema or inputSchemaFlattened will be returned, not both.
	InputSchemaFlattened *Jsonschemadocument `json:"inputSchemaFlattened,omitempty"`

	// InputSchemaUri - The URI of the input schema
	InputSchemaUri *string `json:"inputSchemaUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Actioninput) SetField(field string, fieldValue interface{}) {
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

func (o Actioninput) MarshalJSON() ([]byte, error) {
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
	type Alias Actioninput
	
	return json.Marshal(&struct { 
		InputSchema *Jsonschemadocument `json:"inputSchema,omitempty"`
		
		InputSchemaFlattened *Jsonschemadocument `json:"inputSchemaFlattened,omitempty"`
		
		InputSchemaUri *string `json:"inputSchemaUri,omitempty"`
		Alias
	}{ 
		InputSchema: o.InputSchema,
		
		InputSchemaFlattened: o.InputSchemaFlattened,
		
		InputSchemaUri: o.InputSchemaUri,
		Alias:    (Alias)(o),
	})
}

func (o *Actioninput) UnmarshalJSON(b []byte) error {
	var ActioninputMap map[string]interface{}
	err := json.Unmarshal(b, &ActioninputMap)
	if err != nil {
		return err
	}
	
	if InputSchema, ok := ActioninputMap["inputSchema"].(map[string]interface{}); ok {
		InputSchemaString, _ := json.Marshal(InputSchema)
		json.Unmarshal(InputSchemaString, &o.InputSchema)
	}
	
	if InputSchemaFlattened, ok := ActioninputMap["inputSchemaFlattened"].(map[string]interface{}); ok {
		InputSchemaFlattenedString, _ := json.Marshal(InputSchemaFlattened)
		json.Unmarshal(InputSchemaFlattenedString, &o.InputSchemaFlattened)
	}
	
	if InputSchemaUri, ok := ActioninputMap["inputSchemaUri"].(string); ok {
		o.InputSchemaUri = &InputSchemaUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Actioninput) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
