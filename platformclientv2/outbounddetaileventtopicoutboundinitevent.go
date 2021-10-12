package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Outbounddetaileventtopicoutboundinitevent
type Outbounddetaileventtopicoutboundinitevent struct { 
	// EventTime
	EventTime *int `json:"eventTime,omitempty"`


	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`


	// ParticipantId
	ParticipantId *string `json:"participantId,omitempty"`


	// SessionId
	SessionId *string `json:"sessionId,omitempty"`


	// MediaType
	MediaType *string `json:"mediaType,omitempty"`


	// Provider
	Provider *string `json:"provider,omitempty"`


	// Direction
	Direction *string `json:"direction,omitempty"`


	// Ani
	Ani *string `json:"ani,omitempty"`


	// Dnis
	Dnis *string `json:"dnis,omitempty"`


	// AddressTo
	AddressTo *string `json:"addressTo,omitempty"`


	// AddressFrom
	AddressFrom *string `json:"addressFrom,omitempty"`


	// Subject
	Subject *string `json:"subject,omitempty"`


	// MessageType
	MessageType *string `json:"messageType,omitempty"`


	// OutboundCampaignId
	OutboundCampaignId *string `json:"outboundCampaignId,omitempty"`


	// DivisionId
	DivisionId *string `json:"divisionId,omitempty"`


	// OutboundContactListId
	OutboundContactListId *string `json:"outboundContactListId,omitempty"`


	// OutboundContactId
	OutboundContactId *string `json:"outboundContactId,omitempty"`

}

func (o *Outbounddetaileventtopicoutboundinitevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Outbounddetaileventtopicoutboundinitevent
	
	return json.Marshal(&struct { 
		EventTime *int `json:"eventTime,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		ParticipantId *string `json:"participantId,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		Ani *string `json:"ani,omitempty"`
		
		Dnis *string `json:"dnis,omitempty"`
		
		AddressTo *string `json:"addressTo,omitempty"`
		
		AddressFrom *string `json:"addressFrom,omitempty"`
		
		Subject *string `json:"subject,omitempty"`
		
		MessageType *string `json:"messageType,omitempty"`
		
		OutboundCampaignId *string `json:"outboundCampaignId,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		OutboundContactListId *string `json:"outboundContactListId,omitempty"`
		
		OutboundContactId *string `json:"outboundContactId,omitempty"`
		*Alias
	}{ 
		EventTime: o.EventTime,
		
		ConversationId: o.ConversationId,
		
		ParticipantId: o.ParticipantId,
		
		SessionId: o.SessionId,
		
		MediaType: o.MediaType,
		
		Provider: o.Provider,
		
		Direction: o.Direction,
		
		Ani: o.Ani,
		
		Dnis: o.Dnis,
		
		AddressTo: o.AddressTo,
		
		AddressFrom: o.AddressFrom,
		
		Subject: o.Subject,
		
		MessageType: o.MessageType,
		
		OutboundCampaignId: o.OutboundCampaignId,
		
		DivisionId: o.DivisionId,
		
		OutboundContactListId: o.OutboundContactListId,
		
		OutboundContactId: o.OutboundContactId,
		Alias:    (*Alias)(o),
	})
}

func (o *Outbounddetaileventtopicoutboundinitevent) UnmarshalJSON(b []byte) error {
	var OutbounddetaileventtopicoutboundiniteventMap map[string]interface{}
	err := json.Unmarshal(b, &OutbounddetaileventtopicoutboundiniteventMap)
	if err != nil {
		return err
	}
	
	if EventTime, ok := OutbounddetaileventtopicoutboundiniteventMap["eventTime"].(float64); ok {
		EventTimeInt := int(EventTime)
		o.EventTime = &EventTimeInt
	}
	
	if ConversationId, ok := OutbounddetaileventtopicoutboundiniteventMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
	
	if ParticipantId, ok := OutbounddetaileventtopicoutboundiniteventMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
	
	if SessionId, ok := OutbounddetaileventtopicoutboundiniteventMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
	
	if MediaType, ok := OutbounddetaileventtopicoutboundiniteventMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
	
	if Provider, ok := OutbounddetaileventtopicoutboundiniteventMap["provider"].(string); ok {
		o.Provider = &Provider
	}
	
	if Direction, ok := OutbounddetaileventtopicoutboundiniteventMap["direction"].(string); ok {
		o.Direction = &Direction
	}
	
	if Ani, ok := OutbounddetaileventtopicoutboundiniteventMap["ani"].(string); ok {
		o.Ani = &Ani
	}
	
	if Dnis, ok := OutbounddetaileventtopicoutboundiniteventMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
	
	if AddressTo, ok := OutbounddetaileventtopicoutboundiniteventMap["addressTo"].(string); ok {
		o.AddressTo = &AddressTo
	}
	
	if AddressFrom, ok := OutbounddetaileventtopicoutboundiniteventMap["addressFrom"].(string); ok {
		o.AddressFrom = &AddressFrom
	}
	
	if Subject, ok := OutbounddetaileventtopicoutboundiniteventMap["subject"].(string); ok {
		o.Subject = &Subject
	}
	
	if MessageType, ok := OutbounddetaileventtopicoutboundiniteventMap["messageType"].(string); ok {
		o.MessageType = &MessageType
	}
	
	if OutboundCampaignId, ok := OutbounddetaileventtopicoutboundiniteventMap["outboundCampaignId"].(string); ok {
		o.OutboundCampaignId = &OutboundCampaignId
	}
	
	if DivisionId, ok := OutbounddetaileventtopicoutboundiniteventMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
	
	if OutboundContactListId, ok := OutbounddetaileventtopicoutboundiniteventMap["outboundContactListId"].(string); ok {
		o.OutboundContactListId = &OutboundContactListId
	}
	
	if OutboundContactId, ok := OutbounddetaileventtopicoutboundiniteventMap["outboundContactId"].(string); ok {
		o.OutboundContactId = &OutboundContactId
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Outbounddetaileventtopicoutboundinitevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
