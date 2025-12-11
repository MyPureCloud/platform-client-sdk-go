package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Activationtrigger
type Activationtrigger struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// TriggerType - Trigger type that activated this checklist.
	TriggerType *string `json:"triggerType,omitempty"`

	// TriggerDate - Date when the checklist was triggered. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	TriggerDate *time.Time `json:"triggerDate,omitempty"`

	// IntentId - The intent ID if checklist was triggered by an intent.
	IntentId *string `json:"intentId,omitempty"`

	// IntentName - The intent name if checklist was triggered by an intent.
	IntentName *string `json:"intentName,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Activationtrigger) SetField(field string, fieldValue interface{}) {
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

func (o Activationtrigger) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "TriggerDate", }
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
	type Alias Activationtrigger
	
	TriggerDate := new(string)
	if o.TriggerDate != nil {
		
		*TriggerDate = timeutil.Strftime(o.TriggerDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		TriggerDate = nil
	}
	
	return json.Marshal(&struct { 
		TriggerType *string `json:"triggerType,omitempty"`
		
		TriggerDate *string `json:"triggerDate,omitempty"`
		
		IntentId *string `json:"intentId,omitempty"`
		
		IntentName *string `json:"intentName,omitempty"`
		Alias
	}{ 
		TriggerType: o.TriggerType,
		
		TriggerDate: TriggerDate,
		
		IntentId: o.IntentId,
		
		IntentName: o.IntentName,
		Alias:    (Alias)(o),
	})
}

func (o *Activationtrigger) UnmarshalJSON(b []byte) error {
	var ActivationtriggerMap map[string]interface{}
	err := json.Unmarshal(b, &ActivationtriggerMap)
	if err != nil {
		return err
	}
	
	if TriggerType, ok := ActivationtriggerMap["triggerType"].(string); ok {
		o.TriggerType = &TriggerType
	}
    
	if triggerDateString, ok := ActivationtriggerMap["triggerDate"].(string); ok {
		TriggerDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", triggerDateString)
		o.TriggerDate = &TriggerDate
	}
	
	if IntentId, ok := ActivationtriggerMap["intentId"].(string); ok {
		o.IntentId = &IntentId
	}
    
	if IntentName, ok := ActivationtriggerMap["intentName"].(string); ok {
		o.IntentName = &IntentName
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Activationtrigger) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
