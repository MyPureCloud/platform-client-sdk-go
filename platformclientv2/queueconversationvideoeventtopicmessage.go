package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationvideoeventtopicmessage
type Queueconversationvideoeventtopicmessage struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// State
	State *string `json:"state,omitempty"`


	// Held
	Held *bool `json:"held,omitempty"`


	// ErrorInfo
	ErrorInfo *Queueconversationvideoeventtopicerrordetails `json:"errorInfo,omitempty"`


	// Provider
	Provider *string `json:"provider,omitempty"`


	// ScriptId
	ScriptId *string `json:"scriptId,omitempty"`


	// PeerId
	PeerId *string `json:"peerId,omitempty"`


	// DisconnectType
	DisconnectType *string `json:"disconnectType,omitempty"`


	// StartHoldTime
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`


	// ConnectedTime
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// DisconnectedTime
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`


	// ToAddress
	ToAddress *Queueconversationvideoeventtopicaddress `json:"toAddress,omitempty"`


	// FromAddress
	FromAddress *Queueconversationvideoeventtopicaddress `json:"fromAddress,omitempty"`


	// Messages
	Messages *[]Queueconversationvideoeventtopicmessagedetails `json:"messages,omitempty"`


	// MessagesTranscriptUri
	MessagesTranscriptUri *string `json:"messagesTranscriptUri,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// RecipientCountry
	RecipientCountry *string `json:"recipientCountry,omitempty"`


	// RecipientType
	RecipientType *string `json:"recipientType,omitempty"`


	// JourneyContext
	JourneyContext *Queueconversationvideoeventtopicjourneycontext `json:"journeyContext,omitempty"`


	// Wrapup
	Wrapup *Queueconversationvideoeventtopicwrapup `json:"wrapup,omitempty"`


	// AfterCallWork
	AfterCallWork *Queueconversationvideoeventtopicaftercallwork `json:"afterCallWork,omitempty"`


	// AfterCallWorkRequired
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (o *Queueconversationvideoeventtopicmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationvideoeventtopicmessage
	
	StartHoldTime := new(string)
	if o.StartHoldTime != nil {
		
		*StartHoldTime = timeutil.Strftime(o.StartHoldTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartHoldTime = nil
	}
	
	ConnectedTime := new(string)
	if o.ConnectedTime != nil {
		
		*ConnectedTime = timeutil.Strftime(o.ConnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConnectedTime = nil
	}
	
	DisconnectedTime := new(string)
	if o.DisconnectedTime != nil {
		
		*DisconnectedTime = timeutil.Strftime(o.DisconnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DisconnectedTime = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		
		ErrorInfo *Queueconversationvideoeventtopicerrordetails `json:"errorInfo,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ScriptId *string `json:"scriptId,omitempty"`
		
		PeerId *string `json:"peerId,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		StartHoldTime *string `json:"startHoldTime,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		DisconnectedTime *string `json:"disconnectedTime,omitempty"`
		
		ToAddress *Queueconversationvideoeventtopicaddress `json:"toAddress,omitempty"`
		
		FromAddress *Queueconversationvideoeventtopicaddress `json:"fromAddress,omitempty"`
		
		Messages *[]Queueconversationvideoeventtopicmessagedetails `json:"messages,omitempty"`
		
		MessagesTranscriptUri *string `json:"messagesTranscriptUri,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		RecipientCountry *string `json:"recipientCountry,omitempty"`
		
		RecipientType *string `json:"recipientType,omitempty"`
		
		JourneyContext *Queueconversationvideoeventtopicjourneycontext `json:"journeyContext,omitempty"`
		
		Wrapup *Queueconversationvideoeventtopicwrapup `json:"wrapup,omitempty"`
		
		AfterCallWork *Queueconversationvideoeventtopicaftercallwork `json:"afterCallWork,omitempty"`
		
		AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		State: o.State,
		
		Held: o.Held,
		
		ErrorInfo: o.ErrorInfo,
		
		Provider: o.Provider,
		
		ScriptId: o.ScriptId,
		
		PeerId: o.PeerId,
		
		DisconnectType: o.DisconnectType,
		
		StartHoldTime: StartHoldTime,
		
		ConnectedTime: ConnectedTime,
		
		DisconnectedTime: DisconnectedTime,
		
		ToAddress: o.ToAddress,
		
		FromAddress: o.FromAddress,
		
		Messages: o.Messages,
		
		MessagesTranscriptUri: o.MessagesTranscriptUri,
		
		VarType: o.VarType,
		
		RecipientCountry: o.RecipientCountry,
		
		RecipientType: o.RecipientType,
		
		JourneyContext: o.JourneyContext,
		
		Wrapup: o.Wrapup,
		
		AfterCallWork: o.AfterCallWork,
		
		AfterCallWorkRequired: o.AfterCallWorkRequired,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationvideoeventtopicmessage) UnmarshalJSON(b []byte) error {
	var QueueconversationvideoeventtopicmessageMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationvideoeventtopicmessageMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationvideoeventtopicmessageMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if State, ok := QueueconversationvideoeventtopicmessageMap["state"].(string); ok {
		o.State = &State
	}
	
	if Held, ok := QueueconversationvideoeventtopicmessageMap["held"].(bool); ok {
		o.Held = &Held
	}
	
	if ErrorInfo, ok := QueueconversationvideoeventtopicmessageMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	
	if Provider, ok := QueueconversationvideoeventtopicmessageMap["provider"].(string); ok {
		o.Provider = &Provider
	}
	
	if ScriptId, ok := QueueconversationvideoeventtopicmessageMap["scriptId"].(string); ok {
		o.ScriptId = &ScriptId
	}
	
	if PeerId, ok := QueueconversationvideoeventtopicmessageMap["peerId"].(string); ok {
		o.PeerId = &PeerId
	}
	
	if DisconnectType, ok := QueueconversationvideoeventtopicmessageMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
	
	if startHoldTimeString, ok := QueueconversationvideoeventtopicmessageMap["startHoldTime"].(string); ok {
		StartHoldTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startHoldTimeString)
		o.StartHoldTime = &StartHoldTime
	}
	
	if connectedTimeString, ok := QueueconversationvideoeventtopicmessageMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if disconnectedTimeString, ok := QueueconversationvideoeventtopicmessageMap["disconnectedTime"].(string); ok {
		DisconnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", disconnectedTimeString)
		o.DisconnectedTime = &DisconnectedTime
	}
	
	if ToAddress, ok := QueueconversationvideoeventtopicmessageMap["toAddress"].(map[string]interface{}); ok {
		ToAddressString, _ := json.Marshal(ToAddress)
		json.Unmarshal(ToAddressString, &o.ToAddress)
	}
	
	if FromAddress, ok := QueueconversationvideoeventtopicmessageMap["fromAddress"].(map[string]interface{}); ok {
		FromAddressString, _ := json.Marshal(FromAddress)
		json.Unmarshal(FromAddressString, &o.FromAddress)
	}
	
	if Messages, ok := QueueconversationvideoeventtopicmessageMap["messages"].([]interface{}); ok {
		MessagesString, _ := json.Marshal(Messages)
		json.Unmarshal(MessagesString, &o.Messages)
	}
	
	if MessagesTranscriptUri, ok := QueueconversationvideoeventtopicmessageMap["messagesTranscriptUri"].(string); ok {
		o.MessagesTranscriptUri = &MessagesTranscriptUri
	}
	
	if VarType, ok := QueueconversationvideoeventtopicmessageMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if RecipientCountry, ok := QueueconversationvideoeventtopicmessageMap["recipientCountry"].(string); ok {
		o.RecipientCountry = &RecipientCountry
	}
	
	if RecipientType, ok := QueueconversationvideoeventtopicmessageMap["recipientType"].(string); ok {
		o.RecipientType = &RecipientType
	}
	
	if JourneyContext, ok := QueueconversationvideoeventtopicmessageMap["journeyContext"].(map[string]interface{}); ok {
		JourneyContextString, _ := json.Marshal(JourneyContext)
		json.Unmarshal(JourneyContextString, &o.JourneyContext)
	}
	
	if Wrapup, ok := QueueconversationvideoeventtopicmessageMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if AfterCallWork, ok := QueueconversationvideoeventtopicmessageMap["afterCallWork"].(map[string]interface{}); ok {
		AfterCallWorkString, _ := json.Marshal(AfterCallWork)
		json.Unmarshal(AfterCallWorkString, &o.AfterCallWork)
	}
	
	if AfterCallWorkRequired, ok := QueueconversationvideoeventtopicmessageMap["afterCallWorkRequired"].(bool); ok {
		o.AfterCallWorkRequired = &AfterCallWorkRequired
	}
	
	if AdditionalProperties, ok := QueueconversationvideoeventtopicmessageMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
