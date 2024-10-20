package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationsummarytopicvirtualagentsconversationsummaryevent
type Conversationsummarytopicvirtualagentsconversationsummaryevent struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`

	// Participants
	Participants *[]Conversationsummarytopicvirtualagentsconversationsummaryparticipant `json:"participants,omitempty"`

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
	Summary *Conversationsummarytopicvirtualagentsconversationsummary `json:"summary,omitempty"`

	// Headline
	Headline *Conversationsummarytopicvirtualagentsconversationheadline `json:"headline,omitempty"`

	// Reason
	Reason *Conversationsummarytopicvirtualagentsconversationreason `json:"reason,omitempty"`

	// Resolution
	Resolution *Conversationsummarytopicvirtualagentsconversationresolution `json:"resolution,omitempty"`

	// WrapUpCodes
	WrapUpCodes *[]Conversationsummarytopicvirtualagentsconversationwrapupcode `json:"wrapUpCodes,omitempty"`

	// TriggerSource
	TriggerSource *Conversationsummarytopicvirtualagentstriggersource `json:"triggerSource,omitempty"`

	// LastEditedBy
	LastEditedBy *Conversationsummarytopicvirtualagentsconversationsummaryparticipant `json:"lastEditedBy,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversationsummarytopicvirtualagentsconversationsummaryevent) SetField(field string, fieldValue interface{}) {
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

func (o Conversationsummarytopicvirtualagentsconversationsummaryevent) MarshalJSON() ([]byte, error) {
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
	type Alias Conversationsummarytopicvirtualagentsconversationsummaryevent
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	return json.Marshal(&struct { 
		ConversationId *string `json:"conversationId,omitempty"`
		
		Participants *[]Conversationsummarytopicvirtualagentsconversationsummaryparticipant `json:"participants,omitempty"`
		
		CommunicationIds *[]string `json:"communicationIds,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		MessageType *string `json:"messageType,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		SummaryId *string `json:"summaryId,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		Summary *Conversationsummarytopicvirtualagentsconversationsummary `json:"summary,omitempty"`
		
		Headline *Conversationsummarytopicvirtualagentsconversationheadline `json:"headline,omitempty"`
		
		Reason *Conversationsummarytopicvirtualagentsconversationreason `json:"reason,omitempty"`
		
		Resolution *Conversationsummarytopicvirtualagentsconversationresolution `json:"resolution,omitempty"`
		
		WrapUpCodes *[]Conversationsummarytopicvirtualagentsconversationwrapupcode `json:"wrapUpCodes,omitempty"`
		
		TriggerSource *Conversationsummarytopicvirtualagentstriggersource `json:"triggerSource,omitempty"`
		
		LastEditedBy *Conversationsummarytopicvirtualagentsconversationsummaryparticipant `json:"lastEditedBy,omitempty"`
		Alias
	}{ 
		ConversationId: o.ConversationId,
		
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
		
		WrapUpCodes: o.WrapUpCodes,
		
		TriggerSource: o.TriggerSource,
		
		LastEditedBy: o.LastEditedBy,
		Alias:    (Alias)(o),
	})
}

func (o *Conversationsummarytopicvirtualagentsconversationsummaryevent) UnmarshalJSON(b []byte) error {
	var ConversationsummarytopicvirtualagentsconversationsummaryeventMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationsummarytopicvirtualagentsconversationsummaryeventMap)
	if err != nil {
		return err
	}
	
	if ConversationId, ok := ConversationsummarytopicvirtualagentsconversationsummaryeventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if Participants, ok := ConversationsummarytopicvirtualagentsconversationsummaryeventMap["participants"].([]interface{}); ok {
		ParticipantsString, _ := json.Marshal(Participants)
		json.Unmarshal(ParticipantsString, &o.Participants)
	}
	
	if CommunicationIds, ok := ConversationsummarytopicvirtualagentsconversationsummaryeventMap["communicationIds"].([]interface{}); ok {
		CommunicationIdsString, _ := json.Marshal(CommunicationIds)
		json.Unmarshal(CommunicationIdsString, &o.CommunicationIds)
	}
	
	if createdDateString, ok := ConversationsummarytopicvirtualagentsconversationsummaryeventMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if MessageType, ok := ConversationsummarytopicvirtualagentsconversationsummaryeventMap["messageType"].(string); ok {
		o.MessageType = &MessageType
	}
    
	if MediaType, ok := ConversationsummarytopicvirtualagentsconversationsummaryeventMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if SummaryId, ok := ConversationsummarytopicvirtualagentsconversationsummaryeventMap["summaryId"].(string); ok {
		o.SummaryId = &SummaryId
	}
    
	if Language, ok := ConversationsummarytopicvirtualagentsconversationsummaryeventMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if Summary, ok := ConversationsummarytopicvirtualagentsconversationsummaryeventMap["summary"].(map[string]interface{}); ok {
		SummaryString, _ := json.Marshal(Summary)
		json.Unmarshal(SummaryString, &o.Summary)
	}
	
	if Headline, ok := ConversationsummarytopicvirtualagentsconversationsummaryeventMap["headline"].(map[string]interface{}); ok {
		HeadlineString, _ := json.Marshal(Headline)
		json.Unmarshal(HeadlineString, &o.Headline)
	}
	
	if Reason, ok := ConversationsummarytopicvirtualagentsconversationsummaryeventMap["reason"].(map[string]interface{}); ok {
		ReasonString, _ := json.Marshal(Reason)
		json.Unmarshal(ReasonString, &o.Reason)
	}
	
	if Resolution, ok := ConversationsummarytopicvirtualagentsconversationsummaryeventMap["resolution"].(map[string]interface{}); ok {
		ResolutionString, _ := json.Marshal(Resolution)
		json.Unmarshal(ResolutionString, &o.Resolution)
	}
	
	if WrapUpCodes, ok := ConversationsummarytopicvirtualagentsconversationsummaryeventMap["wrapUpCodes"].([]interface{}); ok {
		WrapUpCodesString, _ := json.Marshal(WrapUpCodes)
		json.Unmarshal(WrapUpCodesString, &o.WrapUpCodes)
	}
	
	if TriggerSource, ok := ConversationsummarytopicvirtualagentsconversationsummaryeventMap["triggerSource"].(map[string]interface{}); ok {
		TriggerSourceString, _ := json.Marshal(TriggerSource)
		json.Unmarshal(TriggerSourceString, &o.TriggerSource)
	}
	
	if LastEditedBy, ok := ConversationsummarytopicvirtualagentsconversationsummaryeventMap["lastEditedBy"].(map[string]interface{}); ok {
		LastEditedByString, _ := json.Marshal(LastEditedBy)
		json.Unmarshal(LastEditedByString, &o.LastEditedBy)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationsummarytopicvirtualagentsconversationsummaryevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
