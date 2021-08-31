package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationvideoeventtopicsocialexpression
type Queueconversationvideoeventtopicsocialexpression struct { 
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
	Wrapup *Queueconversationvideoeventtopicwrapup `json:"wrapup,omitempty"`


	// AfterCallWork
	AfterCallWork *Queueconversationvideoeventtopicaftercallwork `json:"afterCallWork,omitempty"`


	// AfterCallWorkRequired
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (o *Queueconversationvideoeventtopicsocialexpression) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationvideoeventtopicsocialexpression
	
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
		
		Wrapup *Queueconversationvideoeventtopicwrapup `json:"wrapup,omitempty"`
		
		AfterCallWork *Queueconversationvideoeventtopicaftercallwork `json:"afterCallWork,omitempty"`
		
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

func (o *Queueconversationvideoeventtopicsocialexpression) UnmarshalJSON(b []byte) error {
	var QueueconversationvideoeventtopicsocialexpressionMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationvideoeventtopicsocialexpressionMap)
	if err != nil {
		return err
	}
	
	if State, ok := QueueconversationvideoeventtopicsocialexpressionMap["state"].(string); ok {
		o.State = &State
	}
	
	if Id, ok := QueueconversationvideoeventtopicsocialexpressionMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if SocialMediaId, ok := QueueconversationvideoeventtopicsocialexpressionMap["socialMediaId"].(string); ok {
		o.SocialMediaId = &SocialMediaId
	}
	
	if SocialMediaHub, ok := QueueconversationvideoeventtopicsocialexpressionMap["socialMediaHub"].(string); ok {
		o.SocialMediaHub = &SocialMediaHub
	}
	
	if SocialUserName, ok := QueueconversationvideoeventtopicsocialexpressionMap["socialUserName"].(string); ok {
		o.SocialUserName = &SocialUserName
	}
	
	if PreviewText, ok := QueueconversationvideoeventtopicsocialexpressionMap["previewText"].(string); ok {
		o.PreviewText = &PreviewText
	}
	
	if RecordingId, ok := QueueconversationvideoeventtopicsocialexpressionMap["recordingId"].(string); ok {
		o.RecordingId = &RecordingId
	}
	
	if Held, ok := QueueconversationvideoeventtopicsocialexpressionMap["held"].(bool); ok {
		o.Held = &Held
	}
	
	if Provider, ok := QueueconversationvideoeventtopicsocialexpressionMap["provider"].(string); ok {
		o.Provider = &Provider
	}
	
	if ScriptId, ok := QueueconversationvideoeventtopicsocialexpressionMap["scriptId"].(string); ok {
		o.ScriptId = &ScriptId
	}
	
	if PeerId, ok := QueueconversationvideoeventtopicsocialexpressionMap["peerId"].(string); ok {
		o.PeerId = &PeerId
	}
	
	if DisconnectType, ok := QueueconversationvideoeventtopicsocialexpressionMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
	
	if startHoldTimeString, ok := QueueconversationvideoeventtopicsocialexpressionMap["startHoldTime"].(string); ok {
		StartHoldTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startHoldTimeString)
		o.StartHoldTime = &StartHoldTime
	}
	
	if connectedTimeString, ok := QueueconversationvideoeventtopicsocialexpressionMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if disconnectedTimeString, ok := QueueconversationvideoeventtopicsocialexpressionMap["disconnectedTime"].(string); ok {
		DisconnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", disconnectedTimeString)
		o.DisconnectedTime = &DisconnectedTime
	}
	
	if Wrapup, ok := QueueconversationvideoeventtopicsocialexpressionMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if AfterCallWork, ok := QueueconversationvideoeventtopicsocialexpressionMap["afterCallWork"].(map[string]interface{}); ok {
		AfterCallWorkString, _ := json.Marshal(AfterCallWork)
		json.Unmarshal(AfterCallWorkString, &o.AfterCallWork)
	}
	
	if AfterCallWorkRequired, ok := QueueconversationvideoeventtopicsocialexpressionMap["afterCallWorkRequired"].(bool); ok {
		o.AfterCallWorkRequired = &AfterCallWorkRequired
	}
	
	if AdditionalProperties, ok := QueueconversationvideoeventtopicsocialexpressionMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicsocialexpression) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
