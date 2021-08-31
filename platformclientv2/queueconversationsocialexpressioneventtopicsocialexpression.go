package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationsocialexpressioneventtopicsocialexpression
type Queueconversationsocialexpressioneventtopicsocialexpression struct { 
	// State
	State *string `json:"state,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// SocialMediaId
	SocialMediaId *string `json:"socialMediaId,omitempty"`


	// SocialMediaHub
	SocialMediaHub *string `json:"socialMediaHub,omitempty"`


	// SocialUserName
	SocialUserName *string `json:"socialUserName,omitempty"`


	// PreviewText
	PreviewText *string `json:"previewText,omitempty"`


	// RecordingId
	RecordingId *string `json:"recordingId,omitempty"`


	// Held
	Held *bool `json:"held,omitempty"`


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


	// Wrapup
	Wrapup *Queueconversationsocialexpressioneventtopicwrapup `json:"wrapup,omitempty"`


	// AfterCallWork
	AfterCallWork *Queueconversationsocialexpressioneventtopicaftercallwork `json:"afterCallWork,omitempty"`


	// AfterCallWorkRequired
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (o *Queueconversationsocialexpressioneventtopicsocialexpression) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationsocialexpressioneventtopicsocialexpression
	
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
		State *string `json:"state,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		SocialMediaId *string `json:"socialMediaId,omitempty"`
		
		SocialMediaHub *string `json:"socialMediaHub,omitempty"`
		
		SocialUserName *string `json:"socialUserName,omitempty"`
		
		PreviewText *string `json:"previewText,omitempty"`
		
		RecordingId *string `json:"recordingId,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ScriptId *string `json:"scriptId,omitempty"`
		
		PeerId *string `json:"peerId,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		StartHoldTime *string `json:"startHoldTime,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		DisconnectedTime *string `json:"disconnectedTime,omitempty"`
		
		Wrapup *Queueconversationsocialexpressioneventtopicwrapup `json:"wrapup,omitempty"`
		
		AfterCallWork *Queueconversationsocialexpressioneventtopicaftercallwork `json:"afterCallWork,omitempty"`
		
		AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		State: o.State,
		
		Id: o.Id,
		
		SocialMediaId: o.SocialMediaId,
		
		SocialMediaHub: o.SocialMediaHub,
		
		SocialUserName: o.SocialUserName,
		
		PreviewText: o.PreviewText,
		
		RecordingId: o.RecordingId,
		
		Held: o.Held,
		
		Provider: o.Provider,
		
		ScriptId: o.ScriptId,
		
		PeerId: o.PeerId,
		
		DisconnectType: o.DisconnectType,
		
		StartHoldTime: StartHoldTime,
		
		ConnectedTime: ConnectedTime,
		
		DisconnectedTime: DisconnectedTime,
		
		Wrapup: o.Wrapup,
		
		AfterCallWork: o.AfterCallWork,
		
		AfterCallWorkRequired: o.AfterCallWorkRequired,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationsocialexpressioneventtopicsocialexpression) UnmarshalJSON(b []byte) error {
	var QueueconversationsocialexpressioneventtopicsocialexpressionMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationsocialexpressioneventtopicsocialexpressionMap)
	if err != nil {
		return err
	}
	
	if State, ok := QueueconversationsocialexpressioneventtopicsocialexpressionMap["state"].(string); ok {
		o.State = &State
	}
	
	if Id, ok := QueueconversationsocialexpressioneventtopicsocialexpressionMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if SocialMediaId, ok := QueueconversationsocialexpressioneventtopicsocialexpressionMap["socialMediaId"].(string); ok {
		o.SocialMediaId = &SocialMediaId
	}
	
	if SocialMediaHub, ok := QueueconversationsocialexpressioneventtopicsocialexpressionMap["socialMediaHub"].(string); ok {
		o.SocialMediaHub = &SocialMediaHub
	}
	
	if SocialUserName, ok := QueueconversationsocialexpressioneventtopicsocialexpressionMap["socialUserName"].(string); ok {
		o.SocialUserName = &SocialUserName
	}
	
	if PreviewText, ok := QueueconversationsocialexpressioneventtopicsocialexpressionMap["previewText"].(string); ok {
		o.PreviewText = &PreviewText
	}
	
	if RecordingId, ok := QueueconversationsocialexpressioneventtopicsocialexpressionMap["recordingId"].(string); ok {
		o.RecordingId = &RecordingId
	}
	
	if Held, ok := QueueconversationsocialexpressioneventtopicsocialexpressionMap["held"].(bool); ok {
		o.Held = &Held
	}
	
	if Provider, ok := QueueconversationsocialexpressioneventtopicsocialexpressionMap["provider"].(string); ok {
		o.Provider = &Provider
	}
	
	if ScriptId, ok := QueueconversationsocialexpressioneventtopicsocialexpressionMap["scriptId"].(string); ok {
		o.ScriptId = &ScriptId
	}
	
	if PeerId, ok := QueueconversationsocialexpressioneventtopicsocialexpressionMap["peerId"].(string); ok {
		o.PeerId = &PeerId
	}
	
	if DisconnectType, ok := QueueconversationsocialexpressioneventtopicsocialexpressionMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
	
	if startHoldTimeString, ok := QueueconversationsocialexpressioneventtopicsocialexpressionMap["startHoldTime"].(string); ok {
		StartHoldTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startHoldTimeString)
		o.StartHoldTime = &StartHoldTime
	}
	
	if connectedTimeString, ok := QueueconversationsocialexpressioneventtopicsocialexpressionMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if disconnectedTimeString, ok := QueueconversationsocialexpressioneventtopicsocialexpressionMap["disconnectedTime"].(string); ok {
		DisconnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", disconnectedTimeString)
		o.DisconnectedTime = &DisconnectedTime
	}
	
	if Wrapup, ok := QueueconversationsocialexpressioneventtopicsocialexpressionMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if AfterCallWork, ok := QueueconversationsocialexpressioneventtopicsocialexpressionMap["afterCallWork"].(map[string]interface{}); ok {
		AfterCallWorkString, _ := json.Marshal(AfterCallWork)
		json.Unmarshal(AfterCallWorkString, &o.AfterCallWork)
	}
	
	if AfterCallWorkRequired, ok := QueueconversationsocialexpressioneventtopicsocialexpressionMap["afterCallWorkRequired"].(bool); ok {
		o.AfterCallWorkRequired = &AfterCallWorkRequired
	}
	
	if AdditionalProperties, ok := QueueconversationsocialexpressioneventtopicsocialexpressionMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicsocialexpression) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
