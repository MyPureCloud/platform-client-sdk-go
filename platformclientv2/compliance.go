package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Compliance
type Compliance struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// StopSettings - List of configurations for 'StopSettings' compliance
	StopSettings *[]Stopsettings `json:"stopSettings,omitempty"`

	// OptInSettings - List of configurations for 'OptInSettings' compliance
	OptInSettings *[]Optinsettings `json:"optInSettings,omitempty"`

	// HelpSettings - List of configurations for 'HelpSettings' compliance
	HelpSettings *[]Helpsettings `json:"helpSettings,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Compliance) SetField(field string, fieldValue interface{}) {
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

func (o Compliance) MarshalJSON() ([]byte, error) {
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
	type Alias Compliance
	
	return json.Marshal(&struct { 
		StopSettings *[]Stopsettings `json:"stopSettings,omitempty"`
		
		OptInSettings *[]Optinsettings `json:"optInSettings,omitempty"`
		
		HelpSettings *[]Helpsettings `json:"helpSettings,omitempty"`
		Alias
	}{ 
		StopSettings: o.StopSettings,
		
		OptInSettings: o.OptInSettings,
		
		HelpSettings: o.HelpSettings,
		Alias:    (Alias)(o),
	})
}

func (o *Compliance) UnmarshalJSON(b []byte) error {
	var ComplianceMap map[string]interface{}
	err := json.Unmarshal(b, &ComplianceMap)
	if err != nil {
		return err
	}
	
	if StopSettings, ok := ComplianceMap["stopSettings"].([]interface{}); ok {
		StopSettingsString, _ := json.Marshal(StopSettings)
		json.Unmarshal(StopSettingsString, &o.StopSettings)
	}
	
	if OptInSettings, ok := ComplianceMap["optInSettings"].([]interface{}); ok {
		OptInSettingsString, _ := json.Marshal(OptInSettings)
		json.Unmarshal(OptInSettingsString, &o.OptInSettings)
	}
	
	if HelpSettings, ok := ComplianceMap["helpSettings"].([]interface{}); ok {
		HelpSettingsString, _ := json.Marshal(HelpSettings)
		json.Unmarshal(HelpSettingsString, &o.HelpSettings)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Compliance) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
