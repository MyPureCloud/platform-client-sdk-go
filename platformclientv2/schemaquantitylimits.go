package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Schemaquantitylimits
type Schemaquantitylimits struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// MinFieldNameCharacters - The minimum number of schema field name characters allowed.
	MinFieldNameCharacters *int `json:"minFieldNameCharacters,omitempty"`

	// MaxFieldNameCharacters - The maximum number of schema field name characters allowed.
	MaxFieldNameCharacters *int `json:"maxFieldNameCharacters,omitempty"`

	// MinFieldDescriptionCharacters - The minimum number of schema field description characters allowed.
	MinFieldDescriptionCharacters *int `json:"minFieldDescriptionCharacters,omitempty"`

	// MaxFieldDescriptionCharacters - The maximum number of schema field description characters allowed.
	MaxFieldDescriptionCharacters *int `json:"maxFieldDescriptionCharacters,omitempty"`

	// MinSchemaNameCharacters - The minimum number of schema name characters allowed.
	MinSchemaNameCharacters *int `json:"minSchemaNameCharacters,omitempty"`

	// MaxSchemaNameCharacters - The maximum number of schema name characters allowed.
	MaxSchemaNameCharacters *int `json:"maxSchemaNameCharacters,omitempty"`

	// MinSchemaDescriptionCharacters - The minimum number of schema description characters allowed.
	MinSchemaDescriptionCharacters *int `json:"minSchemaDescriptionCharacters,omitempty"`

	// MaxSchemaDescriptionCharacters - The maximum number of schema description characters allowed.
	MaxSchemaDescriptionCharacters *int `json:"maxSchemaDescriptionCharacters,omitempty"`

	// MaxNumberOfSchemasPerOrg - The maximum number of schema allowed per org.
	MaxNumberOfSchemasPerOrg *int `json:"maxNumberOfSchemasPerOrg,omitempty"`

	// MaxNumberOfFieldsPerSchema - The maximum number of schema fields allowed per schema.
	MaxNumberOfFieldsPerSchema *int `json:"maxNumberOfFieldsPerSchema,omitempty"`

	// MaxNumberOfFieldsPerOrg - The maximum number of schema fields allowed per organization across all of their schemas.
	MaxNumberOfFieldsPerOrg *int `json:"maxNumberOfFieldsPerOrg,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Schemaquantitylimits) SetField(field string, fieldValue interface{}) {
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

func (o Schemaquantitylimits) MarshalJSON() ([]byte, error) {
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
	type Alias Schemaquantitylimits
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		MinFieldNameCharacters *int `json:"minFieldNameCharacters,omitempty"`
		
		MaxFieldNameCharacters *int `json:"maxFieldNameCharacters,omitempty"`
		
		MinFieldDescriptionCharacters *int `json:"minFieldDescriptionCharacters,omitempty"`
		
		MaxFieldDescriptionCharacters *int `json:"maxFieldDescriptionCharacters,omitempty"`
		
		MinSchemaNameCharacters *int `json:"minSchemaNameCharacters,omitempty"`
		
		MaxSchemaNameCharacters *int `json:"maxSchemaNameCharacters,omitempty"`
		
		MinSchemaDescriptionCharacters *int `json:"minSchemaDescriptionCharacters,omitempty"`
		
		MaxSchemaDescriptionCharacters *int `json:"maxSchemaDescriptionCharacters,omitempty"`
		
		MaxNumberOfSchemasPerOrg *int `json:"maxNumberOfSchemasPerOrg,omitempty"`
		
		MaxNumberOfFieldsPerSchema *int `json:"maxNumberOfFieldsPerSchema,omitempty"`
		
		MaxNumberOfFieldsPerOrg *int `json:"maxNumberOfFieldsPerOrg,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		MinFieldNameCharacters: o.MinFieldNameCharacters,
		
		MaxFieldNameCharacters: o.MaxFieldNameCharacters,
		
		MinFieldDescriptionCharacters: o.MinFieldDescriptionCharacters,
		
		MaxFieldDescriptionCharacters: o.MaxFieldDescriptionCharacters,
		
		MinSchemaNameCharacters: o.MinSchemaNameCharacters,
		
		MaxSchemaNameCharacters: o.MaxSchemaNameCharacters,
		
		MinSchemaDescriptionCharacters: o.MinSchemaDescriptionCharacters,
		
		MaxSchemaDescriptionCharacters: o.MaxSchemaDescriptionCharacters,
		
		MaxNumberOfSchemasPerOrg: o.MaxNumberOfSchemasPerOrg,
		
		MaxNumberOfFieldsPerSchema: o.MaxNumberOfFieldsPerSchema,
		
		MaxNumberOfFieldsPerOrg: o.MaxNumberOfFieldsPerOrg,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Schemaquantitylimits) UnmarshalJSON(b []byte) error {
	var SchemaquantitylimitsMap map[string]interface{}
	err := json.Unmarshal(b, &SchemaquantitylimitsMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SchemaquantitylimitsMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := SchemaquantitylimitsMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if MinFieldNameCharacters, ok := SchemaquantitylimitsMap["minFieldNameCharacters"].(float64); ok {
		MinFieldNameCharactersInt := int(MinFieldNameCharacters)
		o.MinFieldNameCharacters = &MinFieldNameCharactersInt
	}
	
	if MaxFieldNameCharacters, ok := SchemaquantitylimitsMap["maxFieldNameCharacters"].(float64); ok {
		MaxFieldNameCharactersInt := int(MaxFieldNameCharacters)
		o.MaxFieldNameCharacters = &MaxFieldNameCharactersInt
	}
	
	if MinFieldDescriptionCharacters, ok := SchemaquantitylimitsMap["minFieldDescriptionCharacters"].(float64); ok {
		MinFieldDescriptionCharactersInt := int(MinFieldDescriptionCharacters)
		o.MinFieldDescriptionCharacters = &MinFieldDescriptionCharactersInt
	}
	
	if MaxFieldDescriptionCharacters, ok := SchemaquantitylimitsMap["maxFieldDescriptionCharacters"].(float64); ok {
		MaxFieldDescriptionCharactersInt := int(MaxFieldDescriptionCharacters)
		o.MaxFieldDescriptionCharacters = &MaxFieldDescriptionCharactersInt
	}
	
	if MinSchemaNameCharacters, ok := SchemaquantitylimitsMap["minSchemaNameCharacters"].(float64); ok {
		MinSchemaNameCharactersInt := int(MinSchemaNameCharacters)
		o.MinSchemaNameCharacters = &MinSchemaNameCharactersInt
	}
	
	if MaxSchemaNameCharacters, ok := SchemaquantitylimitsMap["maxSchemaNameCharacters"].(float64); ok {
		MaxSchemaNameCharactersInt := int(MaxSchemaNameCharacters)
		o.MaxSchemaNameCharacters = &MaxSchemaNameCharactersInt
	}
	
	if MinSchemaDescriptionCharacters, ok := SchemaquantitylimitsMap["minSchemaDescriptionCharacters"].(float64); ok {
		MinSchemaDescriptionCharactersInt := int(MinSchemaDescriptionCharacters)
		o.MinSchemaDescriptionCharacters = &MinSchemaDescriptionCharactersInt
	}
	
	if MaxSchemaDescriptionCharacters, ok := SchemaquantitylimitsMap["maxSchemaDescriptionCharacters"].(float64); ok {
		MaxSchemaDescriptionCharactersInt := int(MaxSchemaDescriptionCharacters)
		o.MaxSchemaDescriptionCharacters = &MaxSchemaDescriptionCharactersInt
	}
	
	if MaxNumberOfSchemasPerOrg, ok := SchemaquantitylimitsMap["maxNumberOfSchemasPerOrg"].(float64); ok {
		MaxNumberOfSchemasPerOrgInt := int(MaxNumberOfSchemasPerOrg)
		o.MaxNumberOfSchemasPerOrg = &MaxNumberOfSchemasPerOrgInt
	}
	
	if MaxNumberOfFieldsPerSchema, ok := SchemaquantitylimitsMap["maxNumberOfFieldsPerSchema"].(float64); ok {
		MaxNumberOfFieldsPerSchemaInt := int(MaxNumberOfFieldsPerSchema)
		o.MaxNumberOfFieldsPerSchema = &MaxNumberOfFieldsPerSchemaInt
	}
	
	if MaxNumberOfFieldsPerOrg, ok := SchemaquantitylimitsMap["maxNumberOfFieldsPerOrg"].(float64); ok {
		MaxNumberOfFieldsPerOrgInt := int(MaxNumberOfFieldsPerOrg)
		o.MaxNumberOfFieldsPerOrg = &MaxNumberOfFieldsPerOrgInt
	}
	
	if SelfUri, ok := SchemaquantitylimitsMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Schemaquantitylimits) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
