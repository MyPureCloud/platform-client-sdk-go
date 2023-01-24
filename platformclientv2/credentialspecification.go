package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Credentialspecification - Specifies the requirements for a credential that can be provided for configuring an integration
type Credentialspecification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Required - Indicates if the credential must be provided in order for the integration configuration to be valid.
	Required *bool `json:"required,omitempty"`

	// Title - Title describing the usage for this credential.
	Title *string `json:"title,omitempty"`

	// CredentialTypes - List of acceptable credential types that can be provided for this credential.
	CredentialTypes *[]string `json:"credentialTypes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Credentialspecification) SetField(field string, fieldValue interface{}) {
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

func (o Credentialspecification) MarshalJSON() ([]byte, error) {
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
	type Alias Credentialspecification
	
	return json.Marshal(&struct { 
		Required *bool `json:"required,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		CredentialTypes *[]string `json:"credentialTypes,omitempty"`
		Alias
	}{ 
		Required: o.Required,
		
		Title: o.Title,
		
		CredentialTypes: o.CredentialTypes,
		Alias:    (Alias)(o),
	})
}

func (o *Credentialspecification) UnmarshalJSON(b []byte) error {
	var CredentialspecificationMap map[string]interface{}
	err := json.Unmarshal(b, &CredentialspecificationMap)
	if err != nil {
		return err
	}
	
	if Required, ok := CredentialspecificationMap["required"].(bool); ok {
		o.Required = &Required
	}
    
	if Title, ok := CredentialspecificationMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if CredentialTypes, ok := CredentialspecificationMap["credentialTypes"].([]interface{}); ok {
		CredentialTypesString, _ := json.Marshal(CredentialTypes)
		json.Unmarshal(CredentialTypesString, &o.CredentialTypes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Credentialspecification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
