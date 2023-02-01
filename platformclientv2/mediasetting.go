package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Mediasetting
type Mediasetting struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EnableAutoAnswer - Indicates if auto-answer is enabled for the given media type or subtype (default is false).  Subtype settings take precedence over media type settings.
	EnableAutoAnswer *bool `json:"enableAutoAnswer,omitempty"`

	// AlertingTimeoutSeconds
	AlertingTimeoutSeconds *int `json:"alertingTimeoutSeconds,omitempty"`

	// ServiceLevel
	ServiceLevel *Servicelevel `json:"serviceLevel,omitempty"`

	// SubTypeSettings - Map of media subtype to media subtype specific settings.
	SubTypeSettings *map[string]Basemediasettings `json:"subTypeSettings,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Mediasetting) SetField(field string, fieldValue interface{}) {
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

func (o Mediasetting) MarshalJSON() ([]byte, error) {
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
	type Alias Mediasetting
	
	return json.Marshal(&struct { 
		EnableAutoAnswer *bool `json:"enableAutoAnswer,omitempty"`
		
		AlertingTimeoutSeconds *int `json:"alertingTimeoutSeconds,omitempty"`
		
		ServiceLevel *Servicelevel `json:"serviceLevel,omitempty"`
		
		SubTypeSettings *map[string]Basemediasettings `json:"subTypeSettings,omitempty"`
		Alias
	}{ 
		EnableAutoAnswer: o.EnableAutoAnswer,
		
		AlertingTimeoutSeconds: o.AlertingTimeoutSeconds,
		
		ServiceLevel: o.ServiceLevel,
		
		SubTypeSettings: o.SubTypeSettings,
		Alias:    (Alias)(o),
	})
}

func (o *Mediasetting) UnmarshalJSON(b []byte) error {
	var MediasettingMap map[string]interface{}
	err := json.Unmarshal(b, &MediasettingMap)
	if err != nil {
		return err
	}
	
	if EnableAutoAnswer, ok := MediasettingMap["enableAutoAnswer"].(bool); ok {
		o.EnableAutoAnswer = &EnableAutoAnswer
	}
    
	if AlertingTimeoutSeconds, ok := MediasettingMap["alertingTimeoutSeconds"].(float64); ok {
		AlertingTimeoutSecondsInt := int(AlertingTimeoutSeconds)
		o.AlertingTimeoutSeconds = &AlertingTimeoutSecondsInt
	}
	
	if ServiceLevel, ok := MediasettingMap["serviceLevel"].(map[string]interface{}); ok {
		ServiceLevelString, _ := json.Marshal(ServiceLevel)
		json.Unmarshal(ServiceLevelString, &o.ServiceLevel)
	}
	
	if SubTypeSettings, ok := MediasettingMap["subTypeSettings"].(map[string]interface{}); ok {
		SubTypeSettingsString, _ := json.Marshal(SubTypeSettings)
		json.Unmarshal(SubTypeSettingsString, &o.SubTypeSettings)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Mediasetting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
