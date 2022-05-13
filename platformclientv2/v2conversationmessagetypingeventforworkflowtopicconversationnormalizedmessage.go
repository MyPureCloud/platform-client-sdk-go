package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessage
type V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessage struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Channel
	Channel *V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannel `json:"channel,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// Text
	Text *string `json:"text,omitempty"`


	// Content
	Content *[]V2conversationmessagetypingeventforworkflowtopicconversationmessagecontent `json:"content,omitempty"`


	// Events
	Events *[]V2conversationmessagetypingeventforworkflowtopicconversationmessageevent `json:"events,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// Reasons
	Reasons *[]V2conversationmessagetypingeventforworkflowtopicconversationreason `json:"reasons,omitempty"`


	// OriginatingEntity
	OriginatingEntity *string `json:"originatingEntity,omitempty"`


	// IsFinalReceipt
	IsFinalReceipt *bool `json:"isFinalReceipt,omitempty"`


	// Direction
	Direction *string `json:"direction,omitempty"`


	// Metadata
	Metadata *map[string]string `json:"metadata,omitempty"`

}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessage
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Channel *V2conversationmessagetypingeventforworkflowtopicconversationmessagingchannel `json:"channel,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Content *[]V2conversationmessagetypingeventforworkflowtopicconversationmessagecontent `json:"content,omitempty"`
		
		Events *[]V2conversationmessagetypingeventforworkflowtopicconversationmessageevent `json:"events,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Reasons *[]V2conversationmessagetypingeventforworkflowtopicconversationreason `json:"reasons,omitempty"`
		
		OriginatingEntity *string `json:"originatingEntity,omitempty"`
		
		IsFinalReceipt *bool `json:"isFinalReceipt,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		Metadata *map[string]string `json:"metadata,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Channel: o.Channel,
		
		VarType: o.VarType,
		
		Text: o.Text,
		
		Content: o.Content,
		
		Events: o.Events,
		
		Status: o.Status,
		
		Reasons: o.Reasons,
		
		OriginatingEntity: o.OriginatingEntity,
		
		IsFinalReceipt: o.IsFinalReceipt,
		
		Direction: o.Direction,
		
		Metadata: o.Metadata,
		Alias:    (*Alias)(o),
	})
}

func (o *V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessage) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap)
	if err != nil {
		return err
	}
	
	if Id, ok := V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Channel, ok := V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap["channel"].(map[string]interface{}); ok {
		ChannelString, _ := json.Marshal(Channel)
		json.Unmarshal(ChannelString, &o.Channel)
	}
	
	if VarType, ok := V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Text, ok := V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if Content, ok := V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap["content"].([]interface{}); ok {
		ContentString, _ := json.Marshal(Content)
		json.Unmarshal(ContentString, &o.Content)
	}
	
	if Events, ok := V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap["events"].([]interface{}); ok {
		EventsString, _ := json.Marshal(Events)
		json.Unmarshal(EventsString, &o.Events)
	}
	
	if Status, ok := V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if Reasons, ok := V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap["reasons"].([]interface{}); ok {
		ReasonsString, _ := json.Marshal(Reasons)
		json.Unmarshal(ReasonsString, &o.Reasons)
	}
	
	if OriginatingEntity, ok := V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap["originatingEntity"].(string); ok {
		o.OriginatingEntity = &OriginatingEntity
	}
    
	if IsFinalReceipt, ok := V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap["isFinalReceipt"].(bool); ok {
		o.IsFinalReceipt = &IsFinalReceipt
	}
    
	if Direction, ok := V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if Metadata, ok := V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessageMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforworkflowtopicconversationnormalizedmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
