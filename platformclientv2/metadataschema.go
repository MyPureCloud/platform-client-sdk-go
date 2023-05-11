package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Metadataschema - A description of the contents of a data gathering interface for an accelerator
type Metadataschema struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Title - title for the data gathering page
	Title *string `json:"title,omitempty"`

	// Description - description of the data being gathered on this page
	Description *string `json:"description,omitempty"`

	// VarType - type of data being gathered
	VarType *string `json:"type,omitempty"`

	// Properties - list of properties for which input is to be gathered, bother required and optional
	Properties *[]map[string]Metadataproperty `json:"properties,omitempty"`

	// Required - list of required properties
	Required *[]string `json:"required,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Metadataschema) SetField(field string, fieldValue interface{}) {
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

func (o Metadataschema) MarshalJSON() ([]byte, error) {
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
	type Alias Metadataschema
	
	return json.Marshal(&struct { 
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Properties *[]map[string]Metadataproperty `json:"properties,omitempty"`
		
		Required *[]string `json:"required,omitempty"`
		Alias
	}{ 
		Title: o.Title,
		
		Description: o.Description,
		
		VarType: o.VarType,
		
		Properties: o.Properties,
		
		Required: o.Required,
		Alias:    (Alias)(o),
	})
}

func (o *Metadataschema) UnmarshalJSON(b []byte) error {
	var MetadataschemaMap map[string]interface{}
	err := json.Unmarshal(b, &MetadataschemaMap)
	if err != nil {
		return err
	}
	
	if Title, ok := MetadataschemaMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Description, ok := MetadataschemaMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if VarType, ok := MetadataschemaMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Properties, ok := MetadataschemaMap["properties"].([]interface{}); ok {
		PropertiesString, _ := json.Marshal(Properties)
		json.Unmarshal(PropertiesString, &o.Properties)
	}
	
	if Required, ok := MetadataschemaMap["required"].([]interface{}); ok {
		RequiredString, _ := json.Marshal(Required)
		json.Unmarshal(RequiredString, &o.Required)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Metadataschema) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
