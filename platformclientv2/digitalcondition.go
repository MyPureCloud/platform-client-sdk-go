package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Digitalcondition
type Digitalcondition struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Inverted - If true, inverts the result of evaluating this condition. Default is false.
	Inverted *bool `json:"inverted,omitempty"`

	// ContactColumnConditionSettings - The settings for a 'contact list column' condition.
	ContactColumnConditionSettings *Contactcolumnconditionsettings `json:"contactColumnConditionSettings,omitempty"`

	// ContactAddressConditionSettings - The settings for a 'contact address' condition.
	ContactAddressConditionSettings *Contactaddressconditionsettings `json:"contactAddressConditionSettings,omitempty"`

	// ContactAddressTypeConditionSettings - The settings for a 'contact address type' condition.
	ContactAddressTypeConditionSettings *Contactaddresstypeconditionsettings `json:"contactAddressTypeConditionSettings,omitempty"`

	// LastAttemptByColumnConditionSettings - The settings for a 'last attempt by column' condition.
	LastAttemptByColumnConditionSettings *Lastattemptbycolumnconditionsettings `json:"lastAttemptByColumnConditionSettings,omitempty"`

	// LastAttemptOverallConditionSettings - The settings for a 'last attempt overall' condition.
	LastAttemptOverallConditionSettings *Lastattemptoverallconditionsettings `json:"lastAttemptOverallConditionSettings,omitempty"`

	// LastResultByColumnConditionSettings - The settings for a 'last result by column' condition.
	LastResultByColumnConditionSettings *Lastresultbycolumnconditionsettings `json:"lastResultByColumnConditionSettings,omitempty"`

	// LastResultOverallConditionSettings - The settings for a 'last result overall' condition.
	LastResultOverallConditionSettings *Lastresultoverallconditionsettings `json:"lastResultOverallConditionSettings,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Digitalcondition) SetField(field string, fieldValue interface{}) {
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

func (o Digitalcondition) MarshalJSON() ([]byte, error) {
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
	type Alias Digitalcondition
	
	return json.Marshal(&struct { 
		Inverted *bool `json:"inverted,omitempty"`
		
		ContactColumnConditionSettings *Contactcolumnconditionsettings `json:"contactColumnConditionSettings,omitempty"`
		
		ContactAddressConditionSettings *Contactaddressconditionsettings `json:"contactAddressConditionSettings,omitempty"`
		
		ContactAddressTypeConditionSettings *Contactaddresstypeconditionsettings `json:"contactAddressTypeConditionSettings,omitempty"`
		
		LastAttemptByColumnConditionSettings *Lastattemptbycolumnconditionsettings `json:"lastAttemptByColumnConditionSettings,omitempty"`
		
		LastAttemptOverallConditionSettings *Lastattemptoverallconditionsettings `json:"lastAttemptOverallConditionSettings,omitempty"`
		
		LastResultByColumnConditionSettings *Lastresultbycolumnconditionsettings `json:"lastResultByColumnConditionSettings,omitempty"`
		
		LastResultOverallConditionSettings *Lastresultoverallconditionsettings `json:"lastResultOverallConditionSettings,omitempty"`
		Alias
	}{ 
		Inverted: o.Inverted,
		
		ContactColumnConditionSettings: o.ContactColumnConditionSettings,
		
		ContactAddressConditionSettings: o.ContactAddressConditionSettings,
		
		ContactAddressTypeConditionSettings: o.ContactAddressTypeConditionSettings,
		
		LastAttemptByColumnConditionSettings: o.LastAttemptByColumnConditionSettings,
		
		LastAttemptOverallConditionSettings: o.LastAttemptOverallConditionSettings,
		
		LastResultByColumnConditionSettings: o.LastResultByColumnConditionSettings,
		
		LastResultOverallConditionSettings: o.LastResultOverallConditionSettings,
		Alias:    (Alias)(o),
	})
}

func (o *Digitalcondition) UnmarshalJSON(b []byte) error {
	var DigitalconditionMap map[string]interface{}
	err := json.Unmarshal(b, &DigitalconditionMap)
	if err != nil {
		return err
	}
	
	if Inverted, ok := DigitalconditionMap["inverted"].(bool); ok {
		o.Inverted = &Inverted
	}
    
	if ContactColumnConditionSettings, ok := DigitalconditionMap["contactColumnConditionSettings"].(map[string]interface{}); ok {
		ContactColumnConditionSettingsString, _ := json.Marshal(ContactColumnConditionSettings)
		json.Unmarshal(ContactColumnConditionSettingsString, &o.ContactColumnConditionSettings)
	}
	
	if ContactAddressConditionSettings, ok := DigitalconditionMap["contactAddressConditionSettings"].(map[string]interface{}); ok {
		ContactAddressConditionSettingsString, _ := json.Marshal(ContactAddressConditionSettings)
		json.Unmarshal(ContactAddressConditionSettingsString, &o.ContactAddressConditionSettings)
	}
	
	if ContactAddressTypeConditionSettings, ok := DigitalconditionMap["contactAddressTypeConditionSettings"].(map[string]interface{}); ok {
		ContactAddressTypeConditionSettingsString, _ := json.Marshal(ContactAddressTypeConditionSettings)
		json.Unmarshal(ContactAddressTypeConditionSettingsString, &o.ContactAddressTypeConditionSettings)
	}
	
	if LastAttemptByColumnConditionSettings, ok := DigitalconditionMap["lastAttemptByColumnConditionSettings"].(map[string]interface{}); ok {
		LastAttemptByColumnConditionSettingsString, _ := json.Marshal(LastAttemptByColumnConditionSettings)
		json.Unmarshal(LastAttemptByColumnConditionSettingsString, &o.LastAttemptByColumnConditionSettings)
	}
	
	if LastAttemptOverallConditionSettings, ok := DigitalconditionMap["lastAttemptOverallConditionSettings"].(map[string]interface{}); ok {
		LastAttemptOverallConditionSettingsString, _ := json.Marshal(LastAttemptOverallConditionSettings)
		json.Unmarshal(LastAttemptOverallConditionSettingsString, &o.LastAttemptOverallConditionSettings)
	}
	
	if LastResultByColumnConditionSettings, ok := DigitalconditionMap["lastResultByColumnConditionSettings"].(map[string]interface{}); ok {
		LastResultByColumnConditionSettingsString, _ := json.Marshal(LastResultByColumnConditionSettings)
		json.Unmarshal(LastResultByColumnConditionSettingsString, &o.LastResultByColumnConditionSettings)
	}
	
	if LastResultOverallConditionSettings, ok := DigitalconditionMap["lastResultOverallConditionSettings"].(map[string]interface{}); ok {
		LastResultOverallConditionSettingsString, _ := json.Marshal(LastResultOverallConditionSettings)
		json.Unmarshal(LastResultOverallConditionSettingsString, &o.LastResultOverallConditionSettings)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Digitalcondition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
