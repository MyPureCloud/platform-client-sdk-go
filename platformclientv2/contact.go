package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contact
type Contact struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Address - Email address or phone number for this contact type
	Address *string `json:"address,omitempty"`

	// Display - Formatted version of the address property
	Display *string `json:"display,omitempty"`

	// MediaType
	MediaType *string `json:"mediaType,omitempty"`

	// VarType
	VarType *string `json:"type,omitempty"`

	// Extension - Use internal extension instead of address. Mutually exclusive with the address field.
	Extension *string `json:"extension,omitempty"`

	// CountryCode
	CountryCode *string `json:"countryCode,omitempty"`

	// Integration - Integration tag value if this number is associated with an external integration.
	Integration *string `json:"integration,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contact) SetField(field string, fieldValue interface{}) {
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

func (o Contact) MarshalJSON() ([]byte, error) {
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
	type Alias Contact
	
	return json.Marshal(&struct { 
		Address *string `json:"address,omitempty"`
		
		Display *string `json:"display,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Extension *string `json:"extension,omitempty"`
		
		CountryCode *string `json:"countryCode,omitempty"`
		
		Integration *string `json:"integration,omitempty"`
		Alias
	}{ 
		Address: o.Address,
		
		Display: o.Display,
		
		MediaType: o.MediaType,
		
		VarType: o.VarType,
		
		Extension: o.Extension,
		
		CountryCode: o.CountryCode,
		
		Integration: o.Integration,
		Alias:    (Alias)(o),
	})
}

func (o *Contact) UnmarshalJSON(b []byte) error {
	var ContactMap map[string]interface{}
	err := json.Unmarshal(b, &ContactMap)
	if err != nil {
		return err
	}
	
	if Address, ok := ContactMap["address"].(string); ok {
		o.Address = &Address
	}
    
	if Display, ok := ContactMap["display"].(string); ok {
		o.Display = &Display
	}
    
	if MediaType, ok := ContactMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if VarType, ok := ContactMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Extension, ok := ContactMap["extension"].(string); ok {
		o.Extension = &Extension
	}
    
	if CountryCode, ok := ContactMap["countryCode"].(string); ok {
		o.CountryCode = &CountryCode
	}
    
	if Integration, ok := ContactMap["integration"].(string); ok {
		o.Integration = &Integration
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contact) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
