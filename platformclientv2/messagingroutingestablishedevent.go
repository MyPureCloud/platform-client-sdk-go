package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagingroutingestablishedevent
type Messagingroutingestablishedevent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EventId - A unique (V4 UUID) eventId for this event
	EventId *string `json:"eventId,omitempty"`

	// EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`

	// ConversationId - A unique Id (V4 UUID) identifying this conversation
	ConversationId *string `json:"conversationId,omitempty"`

	// CommunicationId - A unique Id (V4 UUID) identifying this communication.
	CommunicationId *string `json:"communicationId,omitempty"`

	// QueueId - The id of the queue that is routing this conversation.
	QueueId *string `json:"queueId,omitempty"`

	// SkillIds - The unique identifiers for the skills that should be used to determine the destination for the conversation.
	SkillIds *[]string `json:"skillIds,omitempty"`

	// LanguageId - The unique identifier for the language that should be used to determine the destination for the conversation.
	LanguageId *string `json:"languageId,omitempty"`

	// Label - An optional label that categorizes the conversation. Max-utilization settings can be configured at a per-label level.
	Label *string `json:"label,omitempty"`

	// InitialConfiguration - Metadata about this communication.
	InitialConfiguration *Messaginginitialconfiguration `json:"initialConfiguration,omitempty"`

	// SourceConfiguration - Metadata about the source of this communication's interaction.
	SourceConfiguration *Sourceconfiguration `json:"sourceConfiguration,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Messagingroutingestablishedevent) SetField(field string, fieldValue interface{}) {
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

func (o Messagingroutingestablishedevent) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "EventDateTime", }
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
	type Alias Messagingroutingestablishedevent
	
	EventDateTime := new(string)
	if o.EventDateTime != nil {
		
		*EventDateTime = timeutil.Strftime(o.EventDateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventDateTime = nil
	}
	
	return json.Marshal(&struct { 
		EventId *string `json:"eventId,omitempty"`
		
		EventDateTime *string `json:"eventDateTime,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		CommunicationId *string `json:"communicationId,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		SkillIds *[]string `json:"skillIds,omitempty"`
		
		LanguageId *string `json:"languageId,omitempty"`
		
		Label *string `json:"label,omitempty"`
		
		InitialConfiguration *Messaginginitialconfiguration `json:"initialConfiguration,omitempty"`
		
		SourceConfiguration *Sourceconfiguration `json:"sourceConfiguration,omitempty"`
		Alias
	}{ 
		EventId: o.EventId,
		
		EventDateTime: EventDateTime,
		
		ConversationId: o.ConversationId,
		
		CommunicationId: o.CommunicationId,
		
		QueueId: o.QueueId,
		
		SkillIds: o.SkillIds,
		
		LanguageId: o.LanguageId,
		
		Label: o.Label,
		
		InitialConfiguration: o.InitialConfiguration,
		
		SourceConfiguration: o.SourceConfiguration,
		Alias:    (Alias)(o),
	})
}

func (o *Messagingroutingestablishedevent) UnmarshalJSON(b []byte) error {
	var MessagingroutingestablishedeventMap map[string]interface{}
	err := json.Unmarshal(b, &MessagingroutingestablishedeventMap)
	if err != nil {
		return err
	}
	
	if EventId, ok := MessagingroutingestablishedeventMap["eventId"].(string); ok {
		o.EventId = &EventId
	}
    
	if eventDateTimeString, ok := MessagingroutingestablishedeventMap["eventDateTime"].(string); ok {
		EventDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventDateTimeString)
		o.EventDateTime = &EventDateTime
	}
	
	if ConversationId, ok := MessagingroutingestablishedeventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if CommunicationId, ok := MessagingroutingestablishedeventMap["communicationId"].(string); ok {
		o.CommunicationId = &CommunicationId
	}
    
	if QueueId, ok := MessagingroutingestablishedeventMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if SkillIds, ok := MessagingroutingestablishedeventMap["skillIds"].([]interface{}); ok {
		SkillIdsString, _ := json.Marshal(SkillIds)
		json.Unmarshal(SkillIdsString, &o.SkillIds)
	}
	
	if LanguageId, ok := MessagingroutingestablishedeventMap["languageId"].(string); ok {
		o.LanguageId = &LanguageId
	}
    
	if Label, ok := MessagingroutingestablishedeventMap["label"].(string); ok {
		o.Label = &Label
	}
    
	if InitialConfiguration, ok := MessagingroutingestablishedeventMap["initialConfiguration"].(map[string]interface{}); ok {
		InitialConfigurationString, _ := json.Marshal(InitialConfiguration)
		json.Unmarshal(InitialConfigurationString, &o.InitialConfiguration)
	}
	
	if SourceConfiguration, ok := MessagingroutingestablishedeventMap["sourceConfiguration"].(map[string]interface{}); ok {
		SourceConfigurationString, _ := json.Marshal(SourceConfiguration)
		json.Unmarshal(SourceConfigurationString, &o.SourceConfiguration)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Messagingroutingestablishedevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
