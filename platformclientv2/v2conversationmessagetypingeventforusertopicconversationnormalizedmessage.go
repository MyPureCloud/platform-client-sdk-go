package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2conversationmessagetypingeventforusertopicconversationnormalizedmessage
type V2conversationmessagetypingeventforusertopicconversationnormalizedmessage struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Channel
	Channel *V2conversationmessagetypingeventforusertopicconversationmessagingchannel `json:"channel,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// Text
	Text *string `json:"text,omitempty"`


	// Content
	Content *[]V2conversationmessagetypingeventforusertopicconversationmessagecontent `json:"content,omitempty"`


	// Events
	Events *[]V2conversationmessagetypingeventforusertopicconversationmessageevent `json:"events,omitempty"`


	// Status
	Status *string `json:"status,omitempty"`


	// Reasons
	Reasons *[]V2conversationmessagetypingeventforusertopicconversationreason `json:"reasons,omitempty"`


	// OriginatingEntity
	OriginatingEntity *string `json:"originatingEntity,omitempty"`


	// IsFinalReceipt
	IsFinalReceipt *bool `json:"isFinalReceipt,omitempty"`


	// Direction
	Direction *string `json:"direction,omitempty"`


	// Metadata
	Metadata *map[string]string `json:"metadata,omitempty"`

}

func (o *V2conversationmessagetypingeventforusertopicconversationnormalizedmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2conversationmessagetypingeventforusertopicconversationnormalizedmessage
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Channel *V2conversationmessagetypingeventforusertopicconversationmessagingchannel `json:"channel,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Text *string `json:"text,omitempty"`
		
		Content *[]V2conversationmessagetypingeventforusertopicconversationmessagecontent `json:"content,omitempty"`
		
		Events *[]V2conversationmessagetypingeventforusertopicconversationmessageevent `json:"events,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		Reasons *[]V2conversationmessagetypingeventforusertopicconversationreason `json:"reasons,omitempty"`
		
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

func (o *V2conversationmessagetypingeventforusertopicconversationnormalizedmessage) UnmarshalJSON(b []byte) error {
	var V2conversationmessagetypingeventforusertopicconversationnormalizedmessageMap map[string]interface{}
	err := json.Unmarshal(b, &V2conversationmessagetypingeventforusertopicconversationnormalizedmessageMap)
	if err != nil {
		return err
	}
	
	if Id, ok := V2conversationmessagetypingeventforusertopicconversationnormalizedmessageMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Channel, ok := V2conversationmessagetypingeventforusertopicconversationnormalizedmessageMap["channel"].(map[string]interface{}); ok {
		ChannelString, _ := json.Marshal(Channel)
		json.Unmarshal(ChannelString, &o.Channel)
	}
	
	if VarType, ok := V2conversationmessagetypingeventforusertopicconversationnormalizedmessageMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if Text, ok := V2conversationmessagetypingeventforusertopicconversationnormalizedmessageMap["text"].(string); ok {
		o.Text = &Text
	}
	
	if Content, ok := V2conversationmessagetypingeventforusertopicconversationnormalizedmessageMap["content"].([]interface{}); ok {
		ContentString, _ := json.Marshal(Content)
		json.Unmarshal(ContentString, &o.Content)
	}
	
	if Events, ok := V2conversationmessagetypingeventforusertopicconversationnormalizedmessageMap["events"].([]interface{}); ok {
		EventsString, _ := json.Marshal(Events)
		json.Unmarshal(EventsString, &o.Events)
	}
	
	if Status, ok := V2conversationmessagetypingeventforusertopicconversationnormalizedmessageMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if Reasons, ok := V2conversationmessagetypingeventforusertopicconversationnormalizedmessageMap["reasons"].([]interface{}); ok {
		ReasonsString, _ := json.Marshal(Reasons)
		json.Unmarshal(ReasonsString, &o.Reasons)
	}
	
	if OriginatingEntity, ok := V2conversationmessagetypingeventforusertopicconversationnormalizedmessageMap["originatingEntity"].(string); ok {
		o.OriginatingEntity = &OriginatingEntity
	}
	
	if IsFinalReceipt, ok := V2conversationmessagetypingeventforusertopicconversationnormalizedmessageMap["isFinalReceipt"].(bool); ok {
		o.IsFinalReceipt = &IsFinalReceipt
	}
	
	if Direction, ok := V2conversationmessagetypingeventforusertopicconversationnormalizedmessageMap["direction"].(string); ok {
		o.Direction = &Direction
	}
	
	if Metadata, ok := V2conversationmessagetypingeventforusertopicconversationnormalizedmessageMap["metadata"].(map[string]interface{}); ok {
		MetadataString, _ := json.Marshal(Metadata)
		json.Unmarshal(MetadataString, &o.Metadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2conversationmessagetypingeventforusertopicconversationnormalizedmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
