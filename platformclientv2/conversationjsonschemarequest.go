package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationjsonschemarequest - A JSON Schema for create/update requests.
type Conversationjsonschemarequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Schema - The JSON Schema specification link. The only value currently supported is \"http://json-schema.org/draft-04/schema#\".
	Schema *string `json:"$schema,omitempty"`

	// Title - The title of the schema. Must be unique across all enabled Custom Attributes schemas.
	Title *string `json:"title,omitempty"`

	// Description - The schema description.
	Description *string `json:"description,omitempty"`

	// Required - The list of required schema properties. All fields are optional unless listed. New fields added after initial schema creation must be optional before being able to update to required.
	Required *[]string `json:"required,omitempty"`

	// Properties - The map of schema properties and their limits.
	Properties *map[string]interface{} `json:"properties,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationjsonschemarequest) SetField(field string, fieldValue interface{}) {
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

func (o Conversationjsonschemarequest) MarshalJSON() ([]byte, error) {
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
	type Alias Conversationjsonschemarequest
	
	return json.Marshal(&struct { 
		Schema *string `json:"$schema,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Required *[]string `json:"required,omitempty"`
		
		Properties *map[string]interface{} `json:"properties,omitempty"`
		Alias
	}{ 
		Schema: o.Schema,
		
		Title: o.Title,
		
		Description: o.Description,
		
		Required: o.Required,
		
		Properties: o.Properties,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationjsonschemarequest) UnmarshalJSON(b []byte) error {
	var ConversationjsonschemarequestMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationjsonschemarequestMap)
	if err != nil {
		return err
	}
	
	if Schema, ok := ConversationjsonschemarequestMap["$schema"].(string); ok {
		o.Schema = &Schema
	}
    
	if Title, ok := ConversationjsonschemarequestMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Description, ok := ConversationjsonschemarequestMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Required, ok := ConversationjsonschemarequestMap["required"].([]interface{}); ok {
		RequiredString, _ := json.Marshal(Required)
		json.Unmarshal(RequiredString, &o.Required)
	}
	
	if Properties, ok := ConversationjsonschemarequestMap["properties"].(map[string]interface{}); ok {
		PropertiesString, _ := json.Marshal(Properties)
		json.Unmarshal(PropertiesString, &o.Properties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationjsonschemarequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
