package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Participantmetrics
type Participantmetrics struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// AgentDurationPercentage - Percentage of Agent duration in the conversation
	AgentDurationPercentage *float64 `json:"agentDurationPercentage,omitempty"`

	// CustomerDurationPercentage - Percentage of Customer duration in the conversation
	CustomerDurationPercentage *float64 `json:"customerDurationPercentage,omitempty"`

	// SilenceDurationPercentage - Percentage of Silence duration in the conversation
	SilenceDurationPercentage *float64 `json:"silenceDurationPercentage,omitempty"`

	// IvrDurationPercentage - Percentage of IVR duration in the conversation
	IvrDurationPercentage *float64 `json:"ivrDurationPercentage,omitempty"`

	// AcdDurationPercentage - Percentage of ACD duration in the conversation
	AcdDurationPercentage *float64 `json:"acdDurationPercentage,omitempty"`

	// OvertalkDurationPercentage - Percentage of Overtalk duration in the conversation
	OvertalkDurationPercentage *float64 `json:"overtalkDurationPercentage,omitempty"`

	// OtherDurationPercentage - Percentage of Other events duration in the conversation
	OtherDurationPercentage *float64 `json:"otherDurationPercentage,omitempty"`

	// OvertalkCount - Number of Overtalks in the conversation
	OvertalkCount *int `json:"overtalkCount,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Participantmetrics) SetField(field string, fieldValue interface{}) {
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

func (o Participantmetrics) MarshalJSON() ([]byte, error) {
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
	type Alias Participantmetrics
	
	return json.Marshal(&struct { 
		AgentDurationPercentage *float64 `json:"agentDurationPercentage,omitempty"`
		
		CustomerDurationPercentage *float64 `json:"customerDurationPercentage,omitempty"`
		
		SilenceDurationPercentage *float64 `json:"silenceDurationPercentage,omitempty"`
		
		IvrDurationPercentage *float64 `json:"ivrDurationPercentage,omitempty"`
		
		AcdDurationPercentage *float64 `json:"acdDurationPercentage,omitempty"`
		
		OvertalkDurationPercentage *float64 `json:"overtalkDurationPercentage,omitempty"`
		
		OtherDurationPercentage *float64 `json:"otherDurationPercentage,omitempty"`
		
		OvertalkCount *int `json:"overtalkCount,omitempty"`
		Alias
	}{ 
		AgentDurationPercentage: o.AgentDurationPercentage,
		
		CustomerDurationPercentage: o.CustomerDurationPercentage,
		
		SilenceDurationPercentage: o.SilenceDurationPercentage,
		
		IvrDurationPercentage: o.IvrDurationPercentage,
		
		AcdDurationPercentage: o.AcdDurationPercentage,
		
		OvertalkDurationPercentage: o.OvertalkDurationPercentage,
		
		OtherDurationPercentage: o.OtherDurationPercentage,
		
		OvertalkCount: o.OvertalkCount,
		Alias:    (Alias)(o),
	})
}

func (o *Participantmetrics) UnmarshalJSON(b []byte) error {
	var ParticipantmetricsMap map[string]interface{}
	err := json.Unmarshal(b, &ParticipantmetricsMap)
	if err != nil {
		return err
	}
	
	if AgentDurationPercentage, ok := ParticipantmetricsMap["agentDurationPercentage"].(float64); ok {
		o.AgentDurationPercentage = &AgentDurationPercentage
	}
    
	if CustomerDurationPercentage, ok := ParticipantmetricsMap["customerDurationPercentage"].(float64); ok {
		o.CustomerDurationPercentage = &CustomerDurationPercentage
	}
    
	if SilenceDurationPercentage, ok := ParticipantmetricsMap["silenceDurationPercentage"].(float64); ok {
		o.SilenceDurationPercentage = &SilenceDurationPercentage
	}
    
	if IvrDurationPercentage, ok := ParticipantmetricsMap["ivrDurationPercentage"].(float64); ok {
		o.IvrDurationPercentage = &IvrDurationPercentage
	}
    
	if AcdDurationPercentage, ok := ParticipantmetricsMap["acdDurationPercentage"].(float64); ok {
		o.AcdDurationPercentage = &AcdDurationPercentage
	}
    
	if OvertalkDurationPercentage, ok := ParticipantmetricsMap["overtalkDurationPercentage"].(float64); ok {
		o.OvertalkDurationPercentage = &OvertalkDurationPercentage
	}
    
	if OtherDurationPercentage, ok := ParticipantmetricsMap["otherDurationPercentage"].(float64); ok {
		o.OtherDurationPercentage = &OtherDurationPercentage
	}
    
	if OvertalkCount, ok := ParticipantmetricsMap["overtalkCount"].(float64); ok {
		OvertalkCountInt := int(OvertalkCount)
		o.OvertalkCount = &OvertalkCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Participantmetrics) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
