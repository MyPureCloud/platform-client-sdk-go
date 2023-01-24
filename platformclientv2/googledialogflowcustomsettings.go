package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Googledialogflowcustomsettings
type Googledialogflowcustomsettings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Environment - If set this environment will be used to initiate the dialogflow bot, otherwise the default configuration will be used.  See https://cloud.google.com/dialogflow/docs/agents-versions
	Environment *string `json:"environment,omitempty"`

	// EventName - If set this eventName will be used to initiate the dialogflow bot rather than language processing on the input text.  See https://cloud.google.com/dialogflow/es/docs/events-overview
	EventName *string `json:"eventName,omitempty"`

	// WebhookQueryParameters - Parameters passed to the fulfillment webhook of the bot (if any).
	WebhookQueryParameters *map[string]string `json:"webhookQueryParameters,omitempty"`

	// EventInputParameters - Parameters passed to the event input of the bot.
	EventInputParameters *map[string]string `json:"eventInputParameters,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Googledialogflowcustomsettings) SetField(field string, fieldValue interface{}) {
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

func (o Googledialogflowcustomsettings) MarshalJSON() ([]byte, error) {
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
	type Alias Googledialogflowcustomsettings
	
	return json.Marshal(&struct { 
		Environment *string `json:"environment,omitempty"`
		
		EventName *string `json:"eventName,omitempty"`
		
		WebhookQueryParameters *map[string]string `json:"webhookQueryParameters,omitempty"`
		
		EventInputParameters *map[string]string `json:"eventInputParameters,omitempty"`
		Alias
	}{ 
		Environment: o.Environment,
		
		EventName: o.EventName,
		
		WebhookQueryParameters: o.WebhookQueryParameters,
		
		EventInputParameters: o.EventInputParameters,
		Alias:    (Alias)(o),
	})
}

func (o *Googledialogflowcustomsettings) UnmarshalJSON(b []byte) error {
	var GoogledialogflowcustomsettingsMap map[string]interface{}
	err := json.Unmarshal(b, &GoogledialogflowcustomsettingsMap)
	if err != nil {
		return err
	}
	
	if Environment, ok := GoogledialogflowcustomsettingsMap["environment"].(string); ok {
		o.Environment = &Environment
	}
    
	if EventName, ok := GoogledialogflowcustomsettingsMap["eventName"].(string); ok {
		o.EventName = &EventName
	}
    
	if WebhookQueryParameters, ok := GoogledialogflowcustomsettingsMap["webhookQueryParameters"].(map[string]interface{}); ok {
		WebhookQueryParametersString, _ := json.Marshal(WebhookQueryParameters)
		json.Unmarshal(WebhookQueryParametersString, &o.WebhookQueryParameters)
	}
	
	if EventInputParameters, ok := GoogledialogflowcustomsettingsMap["eventInputParameters"].(map[string]interface{}); ok {
		EventInputParametersString, _ := json.Marshal(EventInputParameters)
		json.Unmarshal(EventInputParametersString, &o.EventInputParameters)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Googledialogflowcustomsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
