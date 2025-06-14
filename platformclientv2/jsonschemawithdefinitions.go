package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Jsonschemawithdefinitions - A JSON Schema document.
type Jsonschemawithdefinitions struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Schema
	Schema *string `json:"$schema,omitempty"`

	// Title
	Title *string `json:"title,omitempty"`

	// Description
	Description *string `json:"description,omitempty"`

	// VarType
	VarType *string `json:"type,omitempty"`

	// Required
	Required *[]string `json:"required,omitempty"`

	// Properties
	Properties *map[string]interface{} `json:"properties,omitempty"`

	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

	// Definitions
	Definitions *map[string]Definition `json:"definitions,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Jsonschemawithdefinitions) SetField(field string, fieldValue interface{}) {
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

func (o Jsonschemawithdefinitions) MarshalJSON() ([]byte, error) {
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
	type Alias Jsonschemawithdefinitions
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Schema *string `json:"$schema,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Required *[]string `json:"required,omitempty"`
		
		Properties *map[string]interface{} `json:"properties,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		
		Definitions *map[string]Definition `json:"definitions,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Schema: o.Schema,
		
		Title: o.Title,
		
		Description: o.Description,
		
		VarType: o.VarType,
		
		Required: o.Required,
		
		Properties: o.Properties,
		
		AdditionalProperties: o.AdditionalProperties,
		
		Definitions: o.Definitions,
		Alias:    (Alias)(o),
	})
}

func (o *Jsonschemawithdefinitions) UnmarshalJSON(b []byte) error {
	var JsonschemawithdefinitionsMap map[string]interface{}
	err := json.Unmarshal(b, &JsonschemawithdefinitionsMap)
	if err != nil {
		return err
	}
	
	if Id, ok := JsonschemawithdefinitionsMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Schema, ok := JsonschemawithdefinitionsMap["$schema"].(string); ok {
		o.Schema = &Schema
	}
    
	if Title, ok := JsonschemawithdefinitionsMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Description, ok := JsonschemawithdefinitionsMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if VarType, ok := JsonschemawithdefinitionsMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Required, ok := JsonschemawithdefinitionsMap["required"].([]interface{}); ok {
		RequiredString, _ := json.Marshal(Required)
		json.Unmarshal(RequiredString, &o.Required)
	}
	
	if Properties, ok := JsonschemawithdefinitionsMap["properties"].(map[string]interface{}); ok {
		PropertiesString, _ := json.Marshal(Properties)
		json.Unmarshal(PropertiesString, &o.Properties)
	}
	
	if AdditionalProperties, ok := JsonschemawithdefinitionsMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	
	if Definitions, ok := JsonschemawithdefinitionsMap["definitions"].(map[string]interface{}); ok {
		DefinitionsString, _ := json.Marshal(Definitions)
		json.Unmarshal(DefinitionsString, &o.Definitions)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Jsonschemawithdefinitions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
