package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Passwordrequirements
type Passwordrequirements struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// MinimumLength
	MinimumLength *int `json:"minimumLength,omitempty"`

	// MinimumDigits
	MinimumDigits *int `json:"minimumDigits,omitempty"`

	// MinimumLetters
	MinimumLetters *int `json:"minimumLetters,omitempty"`

	// MinimumUpper
	MinimumUpper *int `json:"minimumUpper,omitempty"`

	// MinimumLower
	MinimumLower *int `json:"minimumLower,omitempty"`

	// MinimumSpecials
	MinimumSpecials *int `json:"minimumSpecials,omitempty"`

	// MinimumAgeSeconds
	MinimumAgeSeconds *int `json:"minimumAgeSeconds,omitempty"`

	// ExpirationDays
	ExpirationDays *int `json:"expirationDays,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Passwordrequirements) SetField(field string, fieldValue interface{}) {
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

func (o Passwordrequirements) MarshalJSON() ([]byte, error) {
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
	type Alias Passwordrequirements
	
	return json.Marshal(&struct { 
		MinimumLength *int `json:"minimumLength,omitempty"`
		
		MinimumDigits *int `json:"minimumDigits,omitempty"`
		
		MinimumLetters *int `json:"minimumLetters,omitempty"`
		
		MinimumUpper *int `json:"minimumUpper,omitempty"`
		
		MinimumLower *int `json:"minimumLower,omitempty"`
		
		MinimumSpecials *int `json:"minimumSpecials,omitempty"`
		
		MinimumAgeSeconds *int `json:"minimumAgeSeconds,omitempty"`
		
		ExpirationDays *int `json:"expirationDays,omitempty"`
		Alias
	}{ 
		MinimumLength: o.MinimumLength,
		
		MinimumDigits: o.MinimumDigits,
		
		MinimumLetters: o.MinimumLetters,
		
		MinimumUpper: o.MinimumUpper,
		
		MinimumLower: o.MinimumLower,
		
		MinimumSpecials: o.MinimumSpecials,
		
		MinimumAgeSeconds: o.MinimumAgeSeconds,
		
		ExpirationDays: o.ExpirationDays,
		Alias:    (Alias)(o),
	})
}

func (o *Passwordrequirements) UnmarshalJSON(b []byte) error {
	var PasswordrequirementsMap map[string]interface{}
	err := json.Unmarshal(b, &PasswordrequirementsMap)
	if err != nil {
		return err
	}
	
	if MinimumLength, ok := PasswordrequirementsMap["minimumLength"].(float64); ok {
		MinimumLengthInt := int(MinimumLength)
		o.MinimumLength = &MinimumLengthInt
	}
	
	if MinimumDigits, ok := PasswordrequirementsMap["minimumDigits"].(float64); ok {
		MinimumDigitsInt := int(MinimumDigits)
		o.MinimumDigits = &MinimumDigitsInt
	}
	
	if MinimumLetters, ok := PasswordrequirementsMap["minimumLetters"].(float64); ok {
		MinimumLettersInt := int(MinimumLetters)
		o.MinimumLetters = &MinimumLettersInt
	}
	
	if MinimumUpper, ok := PasswordrequirementsMap["minimumUpper"].(float64); ok {
		MinimumUpperInt := int(MinimumUpper)
		o.MinimumUpper = &MinimumUpperInt
	}
	
	if MinimumLower, ok := PasswordrequirementsMap["minimumLower"].(float64); ok {
		MinimumLowerInt := int(MinimumLower)
		o.MinimumLower = &MinimumLowerInt
	}
	
	if MinimumSpecials, ok := PasswordrequirementsMap["minimumSpecials"].(float64); ok {
		MinimumSpecialsInt := int(MinimumSpecials)
		o.MinimumSpecials = &MinimumSpecialsInt
	}
	
	if MinimumAgeSeconds, ok := PasswordrequirementsMap["minimumAgeSeconds"].(float64); ok {
		MinimumAgeSecondsInt := int(MinimumAgeSeconds)
		o.MinimumAgeSeconds = &MinimumAgeSecondsInt
	}
	
	if ExpirationDays, ok := PasswordrequirementsMap["expirationDays"].(float64); ok {
		ExpirationDaysInt := int(ExpirationDays)
		o.ExpirationDays = &ExpirationDaysInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Passwordrequirements) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
