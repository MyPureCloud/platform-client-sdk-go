package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createverifierrequest
type Createverifierrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Algorithm - The hashing algorithm for the TOTP verifier.
	Algorithm *string `json:"algorithm,omitempty"`

	// Digits - The number of digits in the TOTP code. Must be between 6 and 12.
	Digits *int `json:"digits,omitempty"`

	// Enabled - Indicates whether this verifier will be enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Name - The name of the verifier. Maximum length is 100 characters.
	Name *string `json:"name,omitempty"`

	// Period - The time period in seconds for the TOTP code.
	Period *int `json:"period,omitempty"`

	// SecretSize - The size of the shared secret in bytes. Must be between 10 and 64.
	SecretSize *int `json:"secretSize,omitempty"`

	// VarDefault - Indicates whether this will be the default verifier.
	VarDefault *bool `json:"default,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createverifierrequest) SetField(field string, fieldValue interface{}) {
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

func (o Createverifierrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Createverifierrequest
	
	return json.Marshal(&struct { 
		Algorithm *string `json:"algorithm,omitempty"`
		
		Digits *int `json:"digits,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Period *int `json:"period,omitempty"`
		
		SecretSize *int `json:"secretSize,omitempty"`
		
		VarDefault *bool `json:"default,omitempty"`
		Alias
	}{ 
		Algorithm: o.Algorithm,
		
		Digits: o.Digits,
		
		Enabled: o.Enabled,
		
		Name: o.Name,
		
		Period: o.Period,
		
		SecretSize: o.SecretSize,
		
		VarDefault: o.VarDefault,
		Alias:    (Alias)(o),
	})
}

func (o *Createverifierrequest) UnmarshalJSON(b []byte) error {
	var CreateverifierrequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreateverifierrequestMap)
	if err != nil {
		return err
	}
	
	if Algorithm, ok := CreateverifierrequestMap["algorithm"].(string); ok {
		o.Algorithm = &Algorithm
	}
    
	if Digits, ok := CreateverifierrequestMap["digits"].(float64); ok {
		DigitsInt := int(Digits)
		o.Digits = &DigitsInt
	}
	
	if Enabled, ok := CreateverifierrequestMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    
	if Name, ok := CreateverifierrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Period, ok := CreateverifierrequestMap["period"].(float64); ok {
		PeriodInt := int(Period)
		o.Period = &PeriodInt
	}
	
	if SecretSize, ok := CreateverifierrequestMap["secretSize"].(float64); ok {
		SecretSizeInt := int(SecretSize)
		o.SecretSize = &SecretSizeInt
	}
	
	if VarDefault, ok := CreateverifierrequestMap["default"].(bool); ok {
		o.VarDefault = &VarDefault
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Createverifierrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
