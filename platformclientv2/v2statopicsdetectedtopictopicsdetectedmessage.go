package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2statopicsdetectedtopictopicsdetectedmessage
type V2statopicsdetectedtopictopicsdetectedmessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`

	// CommunicationId
	CommunicationId *string `json:"communicationId,omitempty"`

	// RecordingId
	RecordingId *string `json:"recordingId,omitempty"`

	// TranscriptId
	TranscriptId *string `json:"transcriptId,omitempty"`

	// MediaType
	MediaType *string `json:"mediaType,omitempty"`

	// Topics
	Topics *[]V2statopicsdetectedtopictopicdetected `json:"topics,omitempty"`

	// Participants
	Participants *[]V2statopicsdetectedtopicparticipant `json:"participants,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2statopicsdetectedtopictopicsdetectedmessage) SetField(field string, fieldValue interface{}) {
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

func (o V2statopicsdetectedtopictopicsdetectedmessage) MarshalJSON() ([]byte, error) {
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
	type Alias V2statopicsdetectedtopictopicsdetectedmessage
	
	return json.Marshal(&struct { 
		ConversationId *string `json:"conversationId,omitempty"`
		
		CommunicationId *string `json:"communicationId,omitempty"`
		
		RecordingId *string `json:"recordingId,omitempty"`
		
		TranscriptId *string `json:"transcriptId,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		Topics *[]V2statopicsdetectedtopictopicdetected `json:"topics,omitempty"`
		
		Participants *[]V2statopicsdetectedtopicparticipant `json:"participants,omitempty"`
		Alias
	}{ 
		ConversationId: o.ConversationId,
		
		CommunicationId: o.CommunicationId,
		
		RecordingId: o.RecordingId,
		
		TranscriptId: o.TranscriptId,
		
		MediaType: o.MediaType,
		
		Topics: o.Topics,
		
		Participants: o.Participants,
		Alias:    (Alias)(o),
	})
}

func (o *V2statopicsdetectedtopictopicsdetectedmessage) UnmarshalJSON(b []byte) error {
	var V2statopicsdetectedtopictopicsdetectedmessageMap map[string]interface{}
	err := json.Unmarshal(b, &V2statopicsdetectedtopictopicsdetectedmessageMap)
	if err != nil {
		return err
	}
	
	if ConversationId, ok := V2statopicsdetectedtopictopicsdetectedmessageMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if CommunicationId, ok := V2statopicsdetectedtopictopicsdetectedmessageMap["communicationId"].(string); ok {
		o.CommunicationId = &CommunicationId
	}
    
	if RecordingId, ok := V2statopicsdetectedtopictopicsdetectedmessageMap["recordingId"].(string); ok {
		o.RecordingId = &RecordingId
	}
    
	if TranscriptId, ok := V2statopicsdetectedtopictopicsdetectedmessageMap["transcriptId"].(string); ok {
		o.TranscriptId = &TranscriptId
	}
    
	if MediaType, ok := V2statopicsdetectedtopictopicsdetectedmessageMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if Topics, ok := V2statopicsdetectedtopictopicsdetectedmessageMap["topics"].([]interface{}); ok {
		TopicsString, _ := json.Marshal(Topics)
		json.Unmarshal(TopicsString, &o.Topics)
	}
	
	if Participants, ok := V2statopicsdetectedtopictopicsdetectedmessageMap["participants"].([]interface{}); ok {
		ParticipantsString, _ := json.Marshal(Participants)
		json.Unmarshal(ParticipantsString, &o.Participants)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2statopicsdetectedtopictopicsdetectedmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
