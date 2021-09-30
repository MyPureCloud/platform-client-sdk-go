package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationeventtopicmessage
type Queueconversationeventtopicmessage struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// State
	State *string `json:"state,omitempty"`


	// Held
	Held *bool `json:"held,omitempty"`


	// ErrorInfo
	ErrorInfo *Queueconversationeventtopicerrordetails `json:"errorInfo,omitempty"`


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
	ToAddress *Queueconversationeventtopicaddress `json:"toAddress,omitempty"`


	// FromAddress
	FromAddress *Queueconversationeventtopicaddress `json:"fromAddress,omitempty"`


	// Messages
	Messages *[]Queueconversationeventtopicmessagedetails `json:"messages,omitempty"`


	// MessagesTranscriptUri
	MessagesTranscriptUri *string `json:"messagesTranscriptUri,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// RecipientCountry
	RecipientCountry *string `json:"recipientCountry,omitempty"`


	// RecipientType
	RecipientType *string `json:"recipientType,omitempty"`


	// JourneyContext
	JourneyContext *Queueconversationeventtopicjourneycontext `json:"journeyContext,omitempty"`


	// Wrapup
	Wrapup *Queueconversationeventtopicwrapup `json:"wrapup,omitempty"`


	// AfterCallWork
	AfterCallWork *Queueconversationeventtopicaftercallwork `json:"afterCallWork,omitempty"`


	// AfterCallWorkRequired
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`


	// AgentAssistantId
	AgentAssistantId *string `json:"agentAssistantId,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (o *Queueconversationeventtopicmessage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationeventtopicmessage
	
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
		
		ErrorInfo *Queueconversationeventtopicerrordetails `json:"errorInfo,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ScriptId *string `json:"scriptId,omitempty"`
		
		PeerId *string `json:"peerId,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		StartHoldTime *string `json:"startHoldTime,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		DisconnectedTime *string `json:"disconnectedTime,omitempty"`
		
		ToAddress *Queueconversationeventtopicaddress `json:"toAddress,omitempty"`
		
		FromAddress *Queueconversationeventtopicaddress `json:"fromAddress,omitempty"`
		
		Messages *[]Queueconversationeventtopicmessagedetails `json:"messages,omitempty"`
		
		MessagesTranscriptUri *string `json:"messagesTranscriptUri,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		RecipientCountry *string `json:"recipientCountry,omitempty"`
		
		RecipientType *string `json:"recipientType,omitempty"`
		
		JourneyContext *Queueconversationeventtopicjourneycontext `json:"journeyContext,omitempty"`
		
		Wrapup *Queueconversationeventtopicwrapup `json:"wrapup,omitempty"`
		
		AfterCallWork *Queueconversationeventtopicaftercallwork `json:"afterCallWork,omitempty"`
		
		AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
		
		AgentAssistantId *string `json:"agentAssistantId,omitempty"`
		
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
		
		AgentAssistantId: o.AgentAssistantId,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationeventtopicmessage) UnmarshalJSON(b []byte) error {
	var QueueconversationeventtopicmessageMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationeventtopicmessageMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationeventtopicmessageMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if State, ok := QueueconversationeventtopicmessageMap["state"].(string); ok {
		o.State = &State
	}
	
	if Held, ok := QueueconversationeventtopicmessageMap["held"].(bool); ok {
		o.Held = &Held
	}
	
	if ErrorInfo, ok := QueueconversationeventtopicmessageMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	
	if Provider, ok := QueueconversationeventtopicmessageMap["provider"].(string); ok {
		o.Provider = &Provider
	}
	
	if ScriptId, ok := QueueconversationeventtopicmessageMap["scriptId"].(string); ok {
		o.ScriptId = &ScriptId
	}
	
	if PeerId, ok := QueueconversationeventtopicmessageMap["peerId"].(string); ok {
		o.PeerId = &PeerId
	}
	
	if DisconnectType, ok := QueueconversationeventtopicmessageMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
	
	if startHoldTimeString, ok := QueueconversationeventtopicmessageMap["startHoldTime"].(string); ok {
		StartHoldTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startHoldTimeString)
		o.StartHoldTime = &StartHoldTime
	}
	
	if connectedTimeString, ok := QueueconversationeventtopicmessageMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if disconnectedTimeString, ok := QueueconversationeventtopicmessageMap["disconnectedTime"].(string); ok {
		DisconnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", disconnectedTimeString)
		o.DisconnectedTime = &DisconnectedTime
	}
	
	if ToAddress, ok := QueueconversationeventtopicmessageMap["toAddress"].(map[string]interface{}); ok {
		ToAddressString, _ := json.Marshal(ToAddress)
		json.Unmarshal(ToAddressString, &o.ToAddress)
	}
	
	if FromAddress, ok := QueueconversationeventtopicmessageMap["fromAddress"].(map[string]interface{}); ok {
		FromAddressString, _ := json.Marshal(FromAddress)
		json.Unmarshal(FromAddressString, &o.FromAddress)
	}
	
	if Messages, ok := QueueconversationeventtopicmessageMap["messages"].([]interface{}); ok {
		MessagesString, _ := json.Marshal(Messages)
		json.Unmarshal(MessagesString, &o.Messages)
	}
	
	if MessagesTranscriptUri, ok := QueueconversationeventtopicmessageMap["messagesTranscriptUri"].(string); ok {
		o.MessagesTranscriptUri = &MessagesTranscriptUri
	}
	
	if VarType, ok := QueueconversationeventtopicmessageMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if RecipientCountry, ok := QueueconversationeventtopicmessageMap["recipientCountry"].(string); ok {
		o.RecipientCountry = &RecipientCountry
	}
	
	if RecipientType, ok := QueueconversationeventtopicmessageMap["recipientType"].(string); ok {
		o.RecipientType = &RecipientType
	}
	
	if JourneyContext, ok := QueueconversationeventtopicmessageMap["journeyContext"].(map[string]interface{}); ok {
		JourneyContextString, _ := json.Marshal(JourneyContext)
		json.Unmarshal(JourneyContextString, &o.JourneyContext)
	}
	
	if Wrapup, ok := QueueconversationeventtopicmessageMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if AfterCallWork, ok := QueueconversationeventtopicmessageMap["afterCallWork"].(map[string]interface{}); ok {
		AfterCallWorkString, _ := json.Marshal(AfterCallWork)
		json.Unmarshal(AfterCallWorkString, &o.AfterCallWork)
	}
	
	if AfterCallWorkRequired, ok := QueueconversationeventtopicmessageMap["afterCallWorkRequired"].(bool); ok {
		o.AfterCallWorkRequired = &AfterCallWorkRequired
	}
	
	if AgentAssistantId, ok := QueueconversationeventtopicmessageMap["agentAssistantId"].(string); ok {
		o.AgentAssistantId = &AgentAssistantId
	}
	
	if AdditionalProperties, ok := QueueconversationeventtopicmessageMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopicmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
