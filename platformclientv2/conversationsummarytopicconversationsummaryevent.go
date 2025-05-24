package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationsummarytopicconversationsummaryevent
type Conversationsummarytopicconversationsummaryevent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`

	// QueueId
	QueueId *string `json:"queueId,omitempty"`

	// Participants
	Participants *[]Conversationsummarytopicconversationsummaryparticipant `json:"participants,omitempty"`

	// CommunicationIds
	CommunicationIds *[]string `json:"communicationIds,omitempty"`

	// CreatedDate
	CreatedDate *time.Time `json:"createdDate,omitempty"`

	// MessageType
	MessageType *string `json:"messageType,omitempty"`

	// MediaType
	MediaType *string `json:"mediaType,omitempty"`

	// SummaryId
	SummaryId *string `json:"summaryId,omitempty"`

	// Language
	Language *string `json:"language,omitempty"`

	// Summary
	Summary *Conversationsummarytopicconversationsummary `json:"summary,omitempty"`

	// Headline
	Headline *Conversationsummarytopicconversationheadline `json:"headline,omitempty"`

	// Reason
	Reason *Conversationsummarytopicconversationreason `json:"reason,omitempty"`

	// Resolution
	Resolution *Conversationsummarytopicconversationresolution `json:"resolution,omitempty"`

	// FollowupActions
	FollowupActions *[]Conversationsummarytopicconversationfollowupaction `json:"followupActions,omitempty"`

	// ExtractedEntities
	ExtractedEntities *[]Conversationsummarytopicsummaryextractedcustomentity `json:"extractedEntities,omitempty"`

	// WrapUpCodes
	WrapUpCodes *[]Conversationsummarytopicconversationwrapupcode `json:"wrapUpCodes,omitempty"`

	// TriggerSource
	TriggerSource *Conversationsummarytopictriggersource `json:"triggerSource,omitempty"`

	// LastEditedBy
	LastEditedBy *Conversationsummarytopicconversationsummaryparticipant `json:"lastEditedBy,omitempty"`

	// ErrorType
	ErrorType *string `json:"errorType,omitempty"`

	// DurationMs
	DurationMs *int `json:"durationMs,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationsummarytopicconversationsummaryevent) SetField(field string, fieldValue interface{}) {
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

func (o Conversationsummarytopicconversationsummaryevent) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CreatedDate", }
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
	type Alias Conversationsummarytopicconversationsummaryevent
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	return json.Marshal(&struct { 
		ConversationId *string `json:"conversationId,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		Participants *[]Conversationsummarytopicconversationsummaryparticipant `json:"participants,omitempty"`
		
		CommunicationIds *[]string `json:"communicationIds,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		MessageType *string `json:"messageType,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		SummaryId *string `json:"summaryId,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		Summary *Conversationsummarytopicconversationsummary `json:"summary,omitempty"`
		
		Headline *Conversationsummarytopicconversationheadline `json:"headline,omitempty"`
		
		Reason *Conversationsummarytopicconversationreason `json:"reason,omitempty"`
		
		Resolution *Conversationsummarytopicconversationresolution `json:"resolution,omitempty"`
		
		FollowupActions *[]Conversationsummarytopicconversationfollowupaction `json:"followupActions,omitempty"`
		
		ExtractedEntities *[]Conversationsummarytopicsummaryextractedcustomentity `json:"extractedEntities,omitempty"`
		
		WrapUpCodes *[]Conversationsummarytopicconversationwrapupcode `json:"wrapUpCodes,omitempty"`
		
		TriggerSource *Conversationsummarytopictriggersource `json:"triggerSource,omitempty"`
		
		LastEditedBy *Conversationsummarytopicconversationsummaryparticipant `json:"lastEditedBy,omitempty"`
		
		ErrorType *string `json:"errorType,omitempty"`
		
		DurationMs *int `json:"durationMs,omitempty"`
		Alias
	}{ 
		ConversationId: o.ConversationId,
		
		QueueId: o.QueueId,
		
		Participants: o.Participants,
		
		CommunicationIds: o.CommunicationIds,
		
		CreatedDate: CreatedDate,
		
		MessageType: o.MessageType,
		
		MediaType: o.MediaType,
		
		SummaryId: o.SummaryId,
		
		Language: o.Language,
		
		Summary: o.Summary,
		
		Headline: o.Headline,
		
		Reason: o.Reason,
		
		Resolution: o.Resolution,
		
		FollowupActions: o.FollowupActions,
		
		ExtractedEntities: o.ExtractedEntities,
		
		WrapUpCodes: o.WrapUpCodes,
		
		TriggerSource: o.TriggerSource,
		
		LastEditedBy: o.LastEditedBy,
		
		ErrorType: o.ErrorType,
		
		DurationMs: o.DurationMs,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationsummarytopicconversationsummaryevent) UnmarshalJSON(b []byte) error {
	var ConversationsummarytopicconversationsummaryeventMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationsummarytopicconversationsummaryeventMap)
	if err != nil {
		return err
	}
	
	if ConversationId, ok := ConversationsummarytopicconversationsummaryeventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if QueueId, ok := ConversationsummarytopicconversationsummaryeventMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if Participants, ok := ConversationsummarytopicconversationsummaryeventMap["participants"].([]interface{}); ok {
		ParticipantsString, _ := json.Marshal(Participants)
		json.Unmarshal(ParticipantsString, &o.Participants)
	}
	
	if CommunicationIds, ok := ConversationsummarytopicconversationsummaryeventMap["communicationIds"].([]interface{}); ok {
		CommunicationIdsString, _ := json.Marshal(CommunicationIds)
		json.Unmarshal(CommunicationIdsString, &o.CommunicationIds)
	}
	
	if createdDateString, ok := ConversationsummarytopicconversationsummaryeventMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if MessageType, ok := ConversationsummarytopicconversationsummaryeventMap["messageType"].(string); ok {
		o.MessageType = &MessageType
	}
    
	if MediaType, ok := ConversationsummarytopicconversationsummaryeventMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if SummaryId, ok := ConversationsummarytopicconversationsummaryeventMap["summaryId"].(string); ok {
		o.SummaryId = &SummaryId
	}
    
	if Language, ok := ConversationsummarytopicconversationsummaryeventMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if Summary, ok := ConversationsummarytopicconversationsummaryeventMap["summary"].(map[string]interface{}); ok {
		SummaryString, _ := json.Marshal(Summary)
		json.Unmarshal(SummaryString, &o.Summary)
	}
	
	if Headline, ok := ConversationsummarytopicconversationsummaryeventMap["headline"].(map[string]interface{}); ok {
		HeadlineString, _ := json.Marshal(Headline)
		json.Unmarshal(HeadlineString, &o.Headline)
	}
	
	if Reason, ok := ConversationsummarytopicconversationsummaryeventMap["reason"].(map[string]interface{}); ok {
		ReasonString, _ := json.Marshal(Reason)
		json.Unmarshal(ReasonString, &o.Reason)
	}
	
	if Resolution, ok := ConversationsummarytopicconversationsummaryeventMap["resolution"].(map[string]interface{}); ok {
		ResolutionString, _ := json.Marshal(Resolution)
		json.Unmarshal(ResolutionString, &o.Resolution)
	}
	
	if FollowupActions, ok := ConversationsummarytopicconversationsummaryeventMap["followupActions"].([]interface{}); ok {
		FollowupActionsString, _ := json.Marshal(FollowupActions)
		json.Unmarshal(FollowupActionsString, &o.FollowupActions)
	}
	
	if ExtractedEntities, ok := ConversationsummarytopicconversationsummaryeventMap["extractedEntities"].([]interface{}); ok {
		ExtractedEntitiesString, _ := json.Marshal(ExtractedEntities)
		json.Unmarshal(ExtractedEntitiesString, &o.ExtractedEntities)
	}
	
	if WrapUpCodes, ok := ConversationsummarytopicconversationsummaryeventMap["wrapUpCodes"].([]interface{}); ok {
		WrapUpCodesString, _ := json.Marshal(WrapUpCodes)
		json.Unmarshal(WrapUpCodesString, &o.WrapUpCodes)
	}
	
	if TriggerSource, ok := ConversationsummarytopicconversationsummaryeventMap["triggerSource"].(map[string]interface{}); ok {
		TriggerSourceString, _ := json.Marshal(TriggerSource)
		json.Unmarshal(TriggerSourceString, &o.TriggerSource)
	}
	
	if LastEditedBy, ok := ConversationsummarytopicconversationsummaryeventMap["lastEditedBy"].(map[string]interface{}); ok {
		LastEditedByString, _ := json.Marshal(LastEditedBy)
		json.Unmarshal(LastEditedByString, &o.LastEditedBy)
	}
	
	if ErrorType, ok := ConversationsummarytopicconversationsummaryeventMap["errorType"].(string); ok {
		o.ErrorType = &ErrorType
	}
    
	if DurationMs, ok := ConversationsummarytopicconversationsummaryeventMap["durationMs"].(float64); ok {
		DurationMsInt := int(DurationMs)
		o.DurationMs = &DurationMsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationsummarytopicconversationsummaryevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
