package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalcontactscontactchangedtopicphonenumber
type Externalcontactscontactchangedtopicphonenumber struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Display
	Display *string `json:"display,omitempty"`

	// Extension
	Extension *int `json:"extension,omitempty"`

	// AcceptsSMS
	AcceptsSMS *bool `json:"acceptsSMS,omitempty"`

	// UserInput
	UserInput *string `json:"userInput,omitempty"`

	// E164
	E164 *string `json:"e164,omitempty"`

	// CountryCode
	CountryCode *string `json:"countryCode,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Externalcontactscontactchangedtopicphonenumber) SetField(field string, fieldValue interface{}) {
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

func (o Externalcontactscontactchangedtopicphonenumber) MarshalJSON() ([]byte, error) {
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
	type Alias Externalcontactscontactchangedtopicphonenumber
	
	return json.Marshal(&struct { 
		Display *string `json:"display,omitempty"`
		
		Extension *int `json:"extension,omitempty"`
		
		AcceptsSMS *bool `json:"acceptsSMS,omitempty"`
		
		UserInput *string `json:"userInput,omitempty"`
		
		E164 *string `json:"e164,omitempty"`
		
		CountryCode *string `json:"countryCode,omitempty"`
		Alias
	}{ 
		Display: o.Display,
		
		Extension: o.Extension,
		
		AcceptsSMS: o.AcceptsSMS,
		
		UserInput: o.UserInput,
		
		E164: o.E164,
		
		CountryCode: o.CountryCode,
		Alias:    (Alias)(o),
	})
}

func (o *Externalcontactscontactchangedtopicphonenumber) UnmarshalJSON(b []byte) error {
	var ExternalcontactscontactchangedtopicphonenumberMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalcontactscontactchangedtopicphonenumberMap)
	if err != nil {
		return err
	}
	
	if Display, ok := ExternalcontactscontactchangedtopicphonenumberMap["display"].(string); ok {
		o.Display = &Display
	}
    
	if Extension, ok := ExternalcontactscontactchangedtopicphonenumberMap["extension"].(float64); ok {
		ExtensionInt := int(Extension)
		o.Extension = &ExtensionInt
	}
	
	if AcceptsSMS, ok := ExternalcontactscontactchangedtopicphonenumberMap["acceptsSMS"].(bool); ok {
		o.AcceptsSMS = &AcceptsSMS
	}
    
	if UserInput, ok := ExternalcontactscontactchangedtopicphonenumberMap["userInput"].(string); ok {
		o.UserInput = &UserInput
	}
    
	if E164, ok := ExternalcontactscontactchangedtopicphonenumberMap["e164"].(string); ok {
		o.E164 = &E164
	}
    
	if CountryCode, ok := ExternalcontactscontactchangedtopicphonenumberMap["countryCode"].(string); ok {
		o.CountryCode = &CountryCode
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalcontactscontactchangedtopicphonenumber) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
