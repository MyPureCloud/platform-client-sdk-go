package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Gcuccteamscalendareventsgcuccteamscalendareventv2
type Gcuccteamscalendareventsgcuccteamscalendareventv2 struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ExternalUserId
	ExternalUserId *string `json:"externalUserId,omitempty"`

	// GenesysUserId
	GenesysUserId *string `json:"genesysUserId,omitempty"`

	// Trigger
	Trigger *string `json:"trigger,omitempty"`

	// EventType
	EventType *string `json:"eventType,omitempty"`

	// EventTime
	EventTime *time.Time `json:"eventTime,omitempty"`

	// ResponseStatus
	ResponseStatus *string `json:"responseStatus,omitempty"`

	// MeetingId
	MeetingId *string `json:"meetingId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Gcuccteamscalendareventsgcuccteamscalendareventv2) SetField(field string, fieldValue interface{}) {
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

func (o Gcuccteamscalendareventsgcuccteamscalendareventv2) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "EventTime", }
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
	type Alias Gcuccteamscalendareventsgcuccteamscalendareventv2
	
	EventTime := new(string)
	if o.EventTime != nil {
		
		*EventTime = timeutil.Strftime(o.EventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventTime = nil
	}
	
	return json.Marshal(&struct { 
		ExternalUserId *string `json:"externalUserId,omitempty"`
		
		GenesysUserId *string `json:"genesysUserId,omitempty"`
		
		Trigger *string `json:"trigger,omitempty"`
		
		EventType *string `json:"eventType,omitempty"`
		
		EventTime *string `json:"eventTime,omitempty"`
		
		ResponseStatus *string `json:"responseStatus,omitempty"`
		
		MeetingId *string `json:"meetingId,omitempty"`
		Alias
	}{ 
		ExternalUserId: o.ExternalUserId,
		
		GenesysUserId: o.GenesysUserId,
		
		Trigger: o.Trigger,
		
		EventType: o.EventType,
		
		EventTime: EventTime,
		
		ResponseStatus: o.ResponseStatus,
		
		MeetingId: o.MeetingId,
		Alias:    (Alias)(o),
	})
}

func (o *Gcuccteamscalendareventsgcuccteamscalendareventv2) UnmarshalJSON(b []byte) error {
	var Gcuccteamscalendareventsgcuccteamscalendareventv2Map map[string]interface{}
	err := json.Unmarshal(b, &Gcuccteamscalendareventsgcuccteamscalendareventv2Map)
	if err != nil {
		return err
	}
	
	if ExternalUserId, ok := Gcuccteamscalendareventsgcuccteamscalendareventv2Map["externalUserId"].(string); ok {
		o.ExternalUserId = &ExternalUserId
	}
    
	if GenesysUserId, ok := Gcuccteamscalendareventsgcuccteamscalendareventv2Map["genesysUserId"].(string); ok {
		o.GenesysUserId = &GenesysUserId
	}
    
	if Trigger, ok := Gcuccteamscalendareventsgcuccteamscalendareventv2Map["trigger"].(string); ok {
		o.Trigger = &Trigger
	}
    
	if EventType, ok := Gcuccteamscalendareventsgcuccteamscalendareventv2Map["eventType"].(string); ok {
		o.EventType = &EventType
	}
    
	if eventTimeString, ok := Gcuccteamscalendareventsgcuccteamscalendareventv2Map["eventTime"].(string); ok {
		EventTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventTimeString)
		o.EventTime = &EventTime
	}
	
	if ResponseStatus, ok := Gcuccteamscalendareventsgcuccteamscalendareventv2Map["responseStatus"].(string); ok {
		o.ResponseStatus = &ResponseStatus
	}
    
	if MeetingId, ok := Gcuccteamscalendareventsgcuccteamscalendareventv2Map["meetingId"].(string); ok {
		o.MeetingId = &MeetingId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Gcuccteamscalendareventsgcuccteamscalendareventv2) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
