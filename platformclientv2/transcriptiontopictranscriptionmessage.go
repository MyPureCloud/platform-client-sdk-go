package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Transcriptiontopictranscriptionmessage
type Transcriptiontopictranscriptionmessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EventTime
	EventTime *time.Time `json:"eventTime,omitempty"`

	// OrganizationId
	OrganizationId *string `json:"organizationId,omitempty"`

	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`

	// CommunicationId
	CommunicationId *string `json:"communicationId,omitempty"`

	// SessionStartTimeMs
	SessionStartTimeMs *int `json:"sessionStartTimeMs,omitempty"`

	// TranscriptionStartTimeMs
	TranscriptionStartTimeMs *int `json:"transcriptionStartTimeMs,omitempty"`

	// Transcripts
	Transcripts *[]Transcriptiontopictranscriptresult `json:"transcripts,omitempty"`

	// Status
	Status *Transcriptiontopictranscriptionrequeststatus `json:"status,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Transcriptiontopictranscriptionmessage) SetField(field string, fieldValue interface{}) {
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

func (o Transcriptiontopictranscriptionmessage) MarshalJSON() ([]byte, error) {
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
	type Alias Transcriptiontopictranscriptionmessage
	
	EventTime := new(string)
	if o.EventTime != nil {
		
		*EventTime = timeutil.Strftime(o.EventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventTime = nil
	}
	
	return json.Marshal(&struct { 
		EventTime *string `json:"eventTime,omitempty"`
		
		OrganizationId *string `json:"organizationId,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		CommunicationId *string `json:"communicationId,omitempty"`
		
		SessionStartTimeMs *int `json:"sessionStartTimeMs,omitempty"`
		
		TranscriptionStartTimeMs *int `json:"transcriptionStartTimeMs,omitempty"`
		
		Transcripts *[]Transcriptiontopictranscriptresult `json:"transcripts,omitempty"`
		
		Status *Transcriptiontopictranscriptionrequeststatus `json:"status,omitempty"`
		Alias
	}{ 
		EventTime: EventTime,
		
		OrganizationId: o.OrganizationId,
		
		ConversationId: o.ConversationId,
		
		CommunicationId: o.CommunicationId,
		
		SessionStartTimeMs: o.SessionStartTimeMs,
		
		TranscriptionStartTimeMs: o.TranscriptionStartTimeMs,
		
		Transcripts: o.Transcripts,
		
		Status: o.Status,
		Alias:    (Alias)(o),
	})
}

func (o *Transcriptiontopictranscriptionmessage) UnmarshalJSON(b []byte) error {
	var TranscriptiontopictranscriptionmessageMap map[string]interface{}
	err := json.Unmarshal(b, &TranscriptiontopictranscriptionmessageMap)
	if err != nil {
		return err
	}
	
	if eventTimeString, ok := TranscriptiontopictranscriptionmessageMap["eventTime"].(string); ok {
		EventTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventTimeString)
		o.EventTime = &EventTime
	}
	
	if OrganizationId, ok := TranscriptiontopictranscriptionmessageMap["organizationId"].(string); ok {
		o.OrganizationId = &OrganizationId
	}
    
	if ConversationId, ok := TranscriptiontopictranscriptionmessageMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if CommunicationId, ok := TranscriptiontopictranscriptionmessageMap["communicationId"].(string); ok {
		o.CommunicationId = &CommunicationId
	}
    
	if SessionStartTimeMs, ok := TranscriptiontopictranscriptionmessageMap["sessionStartTimeMs"].(float64); ok {
		SessionStartTimeMsInt := int(SessionStartTimeMs)
		o.SessionStartTimeMs = &SessionStartTimeMsInt
	}
	
	if TranscriptionStartTimeMs, ok := TranscriptiontopictranscriptionmessageMap["transcriptionStartTimeMs"].(float64); ok {
		TranscriptionStartTimeMsInt := int(TranscriptionStartTimeMs)
		o.TranscriptionStartTimeMs = &TranscriptionStartTimeMsInt
	}
	
	if Transcripts, ok := TranscriptiontopictranscriptionmessageMap["transcripts"].([]interface{}); ok {
		TranscriptsString, _ := json.Marshal(Transcripts)
		json.Unmarshal(TranscriptsString, &o.Transcripts)
	}
	
	if Status, ok := TranscriptiontopictranscriptionmessageMap["status"].(map[string]interface{}); ok {
		StatusString, _ := json.Marshal(Status)
		json.Unmarshal(StatusString, &o.Status)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Transcriptiontopictranscriptionmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
