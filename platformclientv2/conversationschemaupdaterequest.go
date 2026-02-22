package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationschemaupdaterequest - Schema update request.
type Conversationschemaupdaterequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Version - The schema's version, a positive integer.
	Version *int `json:"version,omitempty"`

	// Enabled - The schema's enabled/disabled status. A disabled schema cannot be assigned to any other entities, but the data on those entities from the schema still exists.
	Enabled *bool `json:"enabled,omitempty"`

	// JsonSchema - A JSON schema defining the extension to the built-in entity type.
	JsonSchema *Conversationjsonschemarequest `json:"jsonSchema,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationschemaupdaterequest) SetField(field string, fieldValue interface{}) {
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

func (o Conversationschemaupdaterequest) MarshalJSON() ([]byte, error) {
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
	type Alias Conversationschemaupdaterequest
	
	return json.Marshal(&struct { 
		Version *int `json:"version,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		JsonSchema *Conversationjsonschemarequest `json:"jsonSchema,omitempty"`
		Alias
	}{ 
		Version: o.Version,
		
		Enabled: o.Enabled,
		
		JsonSchema: o.JsonSchema,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationschemaupdaterequest) UnmarshalJSON(b []byte) error {
	var ConversationschemaupdaterequestMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationschemaupdaterequestMap)
	if err != nil {
		return err
	}
	
	if Version, ok := ConversationschemaupdaterequestMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if Enabled, ok := ConversationschemaupdaterequestMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if JsonSchema, ok := ConversationschemaupdaterequestMap["jsonSchema"].(map[string]interface{}); ok {
		JsonSchemaString, _ := json.Marshal(JsonSchema)
		json.Unmarshal(JsonSchemaString, &o.JsonSchema)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationschemaupdaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
