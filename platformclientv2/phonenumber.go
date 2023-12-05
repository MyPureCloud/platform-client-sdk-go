package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Phonenumber
type Phonenumber struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Display - The displayed form of the phone number string. Users should input the phone number in this field, but it will be altered by the API on write. If the phone number can be read as E164, the value will be replaced with international formatted-version of the number. If the number cannot be read as E164, the value will be preserved as-is. In both cases, the provided input string will be copied to the userInput field.
	Display *string `json:"display,omitempty"`

	// Extension - An optional extension for the provided phone number.
	Extension *int `json:"extension,omitempty"`

	// AcceptsSMS - Whether this phone number can accept SMS messages.
	AcceptsSMS *bool `json:"acceptsSMS,omitempty"`

	// NormalizationCountryCode - The country code that will be used for E164 conversion of a provided phone number. If the country code is omitted from the provided phone number, the country code provided in this field will be used during the E164 conversion attempt. If this field is left empty, the default country code for any provided phone number that does not explicitly include a country code is assumed to be +1 (North America).
	NormalizationCountryCode *string `json:"normalizationCountryCode,omitempty"`

	// UserInput - The user-inputted phone number string that was provided to the display field on write. This field is not user-writeable and will always be set by the system.
	UserInput *string `json:"userInput,omitempty"`

	// E164 - The E164-formatted form of the provided phone number. This field is not user-writeable and will only be set when the provided phone number could be read as E164.
	E164 *string `json:"e164,omitempty"`

	// CountryCode - The detected country code from the provided phone number. This field is not user-writeable and will only be set when the provided phone number could be read as E164.
	CountryCode *string `json:"countryCode,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Phonenumber) SetField(field string, fieldValue interface{}) {
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

func (o Phonenumber) MarshalJSON() ([]byte, error) {
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
	type Alias Phonenumber
	
	return json.Marshal(&struct { 
		Display *string `json:"display,omitempty"`
		
		Extension *int `json:"extension,omitempty"`
		
		AcceptsSMS *bool `json:"acceptsSMS,omitempty"`
		
		NormalizationCountryCode *string `json:"normalizationCountryCode,omitempty"`
		
		UserInput *string `json:"userInput,omitempty"`
		
		E164 *string `json:"e164,omitempty"`
		
		CountryCode *string `json:"countryCode,omitempty"`
		Alias
	}{ 
		Display: o.Display,
		
		Extension: o.Extension,
		
		AcceptsSMS: o.AcceptsSMS,
		
		NormalizationCountryCode: o.NormalizationCountryCode,
		
		UserInput: o.UserInput,
		
		E164: o.E164,
		
		CountryCode: o.CountryCode,
		Alias:    (Alias)(o),
	})
}

func (o *Phonenumber) UnmarshalJSON(b []byte) error {
	var PhonenumberMap map[string]interface{}
	err := json.Unmarshal(b, &PhonenumberMap)
	if err != nil {
		return err
	}
	
	if Display, ok := PhonenumberMap["display"].(string); ok {
		o.Display = &Display
	}
    
	if Extension, ok := PhonenumberMap["extension"].(float64); ok {
		ExtensionInt := int(Extension)
		o.Extension = &ExtensionInt
	}
	
	if AcceptsSMS, ok := PhonenumberMap["acceptsSMS"].(bool); ok {
		o.AcceptsSMS = &AcceptsSMS
	}
    
	if NormalizationCountryCode, ok := PhonenumberMap["normalizationCountryCode"].(string); ok {
		o.NormalizationCountryCode = &NormalizationCountryCode
	}
    
	if UserInput, ok := PhonenumberMap["userInput"].(string); ok {
		o.UserInput = &UserInput
	}
    
	if E164, ok := PhonenumberMap["e164"].(string); ok {
		o.E164 = &E164
	}
    
	if CountryCode, ok := PhonenumberMap["countryCode"].(string); ok {
		o.CountryCode = &CountryCode
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Phonenumber) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
