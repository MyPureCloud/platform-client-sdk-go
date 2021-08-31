package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationsocialexpressioneventtopicscreenshare
type Queueconversationsocialexpressioneventtopicscreenshare struct { 
	// State
	State *string `json:"state,omitempty"`


	// Self
	Self *Queueconversationsocialexpressioneventtopicaddress `json:"self,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// Context
	Context *string `json:"context,omitempty"`


	// Sharing
	Sharing *bool `json:"sharing,omitempty"`


	// Provider
	Provider *string `json:"provider,omitempty"`


	// ScriptId
	ScriptId *string `json:"scriptId,omitempty"`


	// PeerId
	PeerId *string `json:"peerId,omitempty"`


	// DisconnectType
	DisconnectType *string `json:"disconnectType,omitempty"`


	// ConnectedTime
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// DisconnectedTime
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`


	// Wrapup
	Wrapup *Queueconversationsocialexpressioneventtopicwrapup `json:"wrapup,omitempty"`


	// AfterCallWork
	AfterCallWork *Queueconversationsocialexpressioneventtopicaftercallwork `json:"afterCallWork,omitempty"`


	// AfterCallWorkRequired
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (o *Queueconversationsocialexpressioneventtopicscreenshare) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationsocialexpressioneventtopicscreenshare
	
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
		State *string `json:"state,omitempty"`
		
		Self *Queueconversationsocialexpressioneventtopicaddress `json:"self,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Context *string `json:"context,omitempty"`
		
		Sharing *bool `json:"sharing,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ScriptId *string `json:"scriptId,omitempty"`
		
		PeerId *string `json:"peerId,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		DisconnectedTime *string `json:"disconnectedTime,omitempty"`
		
		Wrapup *Queueconversationsocialexpressioneventtopicwrapup `json:"wrapup,omitempty"`
		
		AfterCallWork *Queueconversationsocialexpressioneventtopicaftercallwork `json:"afterCallWork,omitempty"`
		
		AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		State: o.State,
		
		Self: o.Self,
		
		Id: o.Id,
		
		Context: o.Context,
		
		Sharing: o.Sharing,
		
		Provider: o.Provider,
		
		ScriptId: o.ScriptId,
		
		PeerId: o.PeerId,
		
		DisconnectType: o.DisconnectType,
		
		ConnectedTime: ConnectedTime,
		
		DisconnectedTime: DisconnectedTime,
		
		Wrapup: o.Wrapup,
		
		AfterCallWork: o.AfterCallWork,
		
		AfterCallWorkRequired: o.AfterCallWorkRequired,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationsocialexpressioneventtopicscreenshare) UnmarshalJSON(b []byte) error {
	var QueueconversationsocialexpressioneventtopicscreenshareMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationsocialexpressioneventtopicscreenshareMap)
	if err != nil {
		return err
	}
	
	if State, ok := QueueconversationsocialexpressioneventtopicscreenshareMap["state"].(string); ok {
		o.State = &State
	}
	
	if Self, ok := QueueconversationsocialexpressioneventtopicscreenshareMap["self"].(map[string]interface{}); ok {
		SelfString, _ := json.Marshal(Self)
		json.Unmarshal(SelfString, &o.Self)
	}
	
	if Id, ok := QueueconversationsocialexpressioneventtopicscreenshareMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Context, ok := QueueconversationsocialexpressioneventtopicscreenshareMap["context"].(string); ok {
		o.Context = &Context
	}
	
	if Sharing, ok := QueueconversationsocialexpressioneventtopicscreenshareMap["sharing"].(bool); ok {
		o.Sharing = &Sharing
	}
	
	if Provider, ok := QueueconversationsocialexpressioneventtopicscreenshareMap["provider"].(string); ok {
		o.Provider = &Provider
	}
	
	if ScriptId, ok := QueueconversationsocialexpressioneventtopicscreenshareMap["scriptId"].(string); ok {
		o.ScriptId = &ScriptId
	}
	
	if PeerId, ok := QueueconversationsocialexpressioneventtopicscreenshareMap["peerId"].(string); ok {
		o.PeerId = &PeerId
	}
	
	if DisconnectType, ok := QueueconversationsocialexpressioneventtopicscreenshareMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
	
	if connectedTimeString, ok := QueueconversationsocialexpressioneventtopicscreenshareMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if disconnectedTimeString, ok := QueueconversationsocialexpressioneventtopicscreenshareMap["disconnectedTime"].(string); ok {
		DisconnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", disconnectedTimeString)
		o.DisconnectedTime = &DisconnectedTime
	}
	
	if Wrapup, ok := QueueconversationsocialexpressioneventtopicscreenshareMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if AfterCallWork, ok := QueueconversationsocialexpressioneventtopicscreenshareMap["afterCallWork"].(map[string]interface{}); ok {
		AfterCallWorkString, _ := json.Marshal(AfterCallWork)
		json.Unmarshal(AfterCallWorkString, &o.AfterCallWork)
	}
	
	if AfterCallWorkRequired, ok := QueueconversationsocialexpressioneventtopicscreenshareMap["afterCallWorkRequired"].(bool); ok {
		o.AfterCallWorkRequired = &AfterCallWorkRequired
	}
	
	if AdditionalProperties, ok := QueueconversationsocialexpressioneventtopicscreenshareMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicscreenshare) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
