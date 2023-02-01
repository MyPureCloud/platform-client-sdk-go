package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimconfigresourcetype - Defines a SCIM resource.
type Scimconfigresourcetype struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The ID of the SCIM resource. Set by the service provider. \"caseExact\" is set to \"true\". \"mutability\" is set to \"readOnly\". \"returned\" is set to \"always\".
	Id *string `json:"id,omitempty"`

	// Schemas - The list of supported schemas.
	Schemas *[]string `json:"schemas,omitempty"`

	// Name - The name of the resource type.
	Name *string `json:"name,omitempty"`

	// Description - The description of the resource type.
	Description *string `json:"description,omitempty"`

	// Schema - The URI of the primary or base schema for the resource type.
	Schema *string `json:"schema,omitempty"`

	// SchemaExtensions - The list of schema extensions for the resource type.
	SchemaExtensions *[]Scimconfigresourcetypeschemaextension `json:"schemaExtensions,omitempty"`

	// Endpoint - The HTTP-addressable endpoint of the resource type. Appears after the base URL.
	Endpoint *string `json:"endpoint,omitempty"`

	// Meta - The metadata of the SCIM resource. Only \"location\" and \"resourceType\" are set for \"ResourceType\" resources.
	Meta *Scimmetadata `json:"meta,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Scimconfigresourcetype) SetField(field string, fieldValue interface{}) {
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

func (o Scimconfigresourcetype) MarshalJSON() ([]byte, error) {
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
	type Alias Scimconfigresourcetype
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Schemas *[]string `json:"schemas,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Schema *string `json:"schema,omitempty"`
		
		SchemaExtensions *[]Scimconfigresourcetypeschemaextension `json:"schemaExtensions,omitempty"`
		
		Endpoint *string `json:"endpoint,omitempty"`
		
		Meta *Scimmetadata `json:"meta,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Schemas: o.Schemas,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Schema: o.Schema,
		
		SchemaExtensions: o.SchemaExtensions,
		
		Endpoint: o.Endpoint,
		
		Meta: o.Meta,
		Alias:    (Alias)(o),
	})
}

func (o *Scimconfigresourcetype) UnmarshalJSON(b []byte) error {
	var ScimconfigresourcetypeMap map[string]interface{}
	err := json.Unmarshal(b, &ScimconfigresourcetypeMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ScimconfigresourcetypeMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Schemas, ok := ScimconfigresourcetypeMap["schemas"].([]interface{}); ok {
		SchemasString, _ := json.Marshal(Schemas)
		json.Unmarshal(SchemasString, &o.Schemas)
	}
	
	if Name, ok := ScimconfigresourcetypeMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := ScimconfigresourcetypeMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Schema, ok := ScimconfigresourcetypeMap["schema"].(string); ok {
		o.Schema = &Schema
	}
    
	if SchemaExtensions, ok := ScimconfigresourcetypeMap["schemaExtensions"].([]interface{}); ok {
		SchemaExtensionsString, _ := json.Marshal(SchemaExtensions)
		json.Unmarshal(SchemaExtensionsString, &o.SchemaExtensions)
	}
	
	if Endpoint, ok := ScimconfigresourcetypeMap["endpoint"].(string); ok {
		o.Endpoint = &Endpoint
	}
    
	if Meta, ok := ScimconfigresourcetypeMap["meta"].(map[string]interface{}); ok {
		MetaString, _ := json.Marshal(Meta)
		json.Unmarshal(MetaString, &o.Meta)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Scimconfigresourcetype) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
