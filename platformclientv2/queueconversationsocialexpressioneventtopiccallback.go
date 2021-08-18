package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationsocialexpressioneventtopiccallback
type Queueconversationsocialexpressioneventtopiccallback struct { 
	// State
	State *string `json:"state,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// Direction
	Direction *string `json:"direction,omitempty"`


	// Held
	Held *bool `json:"held,omitempty"`


	// DisconnectType
	DisconnectType *string `json:"disconnectType,omitempty"`


	// StartHoldTime
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`


	// DialerPreview
	DialerPreview *Queueconversationsocialexpressioneventtopicdialerpreview `json:"dialerPreview,omitempty"`


	// Voicemail
	Voicemail *Queueconversationsocialexpressioneventtopicvoicemail `json:"voicemail,omitempty"`


	// CallbackNumbers
	CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`


	// CallbackUserName
	CallbackUserName *string `json:"callbackUserName,omitempty"`


	// ScriptId
	ScriptId *string `json:"scriptId,omitempty"`


	// PeerId
	PeerId *string `json:"peerId,omitempty"`


	// ExternalCampaign
	ExternalCampaign *bool `json:"externalCampaign,omitempty"`


	// SkipEnabled
	SkipEnabled *bool `json:"skipEnabled,omitempty"`


	// Provider
	Provider *string `json:"provider,omitempty"`


	// TimeoutSeconds
	TimeoutSeconds *int `json:"timeoutSeconds,omitempty"`


	// ConnectedTime
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// DisconnectedTime
	DisconnectedTime *time.Time `json:"disconnectedTime,omitempty"`


	// CallbackScheduledTime
	CallbackScheduledTime *time.Time `json:"callbackScheduledTime,omitempty"`


	// AutomatedCallbackConfigId
	AutomatedCallbackConfigId *string `json:"automatedCallbackConfigId,omitempty"`


	// Wrapup
	Wrapup *Queueconversationsocialexpressioneventtopicwrapup `json:"wrapup,omitempty"`


	// AfterCallWork
	AfterCallWork *Queueconversationsocialexpressioneventtopicaftercallwork `json:"afterCallWork,omitempty"`


	// AfterCallWorkRequired
	AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`


	// CallerId
	CallerId *string `json:"callerId,omitempty"`


	// CallerIdName
	CallerIdName *string `json:"callerIdName,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (u *Queueconversationsocialexpressioneventtopiccallback) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationsocialexpressioneventtopiccallback

	
	StartHoldTime := new(string)
	if u.StartHoldTime != nil {
		
		*StartHoldTime = timeutil.Strftime(u.StartHoldTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartHoldTime = nil
	}
	
	ConnectedTime := new(string)
	if u.ConnectedTime != nil {
		
		*ConnectedTime = timeutil.Strftime(u.ConnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConnectedTime = nil
	}
	
	DisconnectedTime := new(string)
	if u.DisconnectedTime != nil {
		
		*DisconnectedTime = timeutil.Strftime(u.DisconnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DisconnectedTime = nil
	}
	
	CallbackScheduledTime := new(string)
	if u.CallbackScheduledTime != nil {
		
		*CallbackScheduledTime = timeutil.Strftime(u.CallbackScheduledTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CallbackScheduledTime = nil
	}
	

	return json.Marshal(&struct { 
		State *string `json:"state,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		StartHoldTime *string `json:"startHoldTime,omitempty"`
		
		DialerPreview *Queueconversationsocialexpressioneventtopicdialerpreview `json:"dialerPreview,omitempty"`
		
		Voicemail *Queueconversationsocialexpressioneventtopicvoicemail `json:"voicemail,omitempty"`
		
		CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`
		
		CallbackUserName *string `json:"callbackUserName,omitempty"`
		
		ScriptId *string `json:"scriptId,omitempty"`
		
		PeerId *string `json:"peerId,omitempty"`
		
		ExternalCampaign *bool `json:"externalCampaign,omitempty"`
		
		SkipEnabled *bool `json:"skipEnabled,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		TimeoutSeconds *int `json:"timeoutSeconds,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		DisconnectedTime *string `json:"disconnectedTime,omitempty"`
		
		CallbackScheduledTime *string `json:"callbackScheduledTime,omitempty"`
		
		AutomatedCallbackConfigId *string `json:"automatedCallbackConfigId,omitempty"`
		
		Wrapup *Queueconversationsocialexpressioneventtopicwrapup `json:"wrapup,omitempty"`
		
		AfterCallWork *Queueconversationsocialexpressioneventtopicaftercallwork `json:"afterCallWork,omitempty"`
		
		AfterCallWorkRequired *bool `json:"afterCallWorkRequired,omitempty"`
		
		CallerId *string `json:"callerId,omitempty"`
		
		CallerIdName *string `json:"callerIdName,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		State: u.State,
		
		Id: u.Id,
		
		Direction: u.Direction,
		
		Held: u.Held,
		
		DisconnectType: u.DisconnectType,
		
		StartHoldTime: StartHoldTime,
		
		DialerPreview: u.DialerPreview,
		
		Voicemail: u.Voicemail,
		
		CallbackNumbers: u.CallbackNumbers,
		
		CallbackUserName: u.CallbackUserName,
		
		ScriptId: u.ScriptId,
		
		PeerId: u.PeerId,
		
		ExternalCampaign: u.ExternalCampaign,
		
		SkipEnabled: u.SkipEnabled,
		
		Provider: u.Provider,
		
		TimeoutSeconds: u.TimeoutSeconds,
		
		ConnectedTime: ConnectedTime,
		
		DisconnectedTime: DisconnectedTime,
		
		CallbackScheduledTime: CallbackScheduledTime,
		
		AutomatedCallbackConfigId: u.AutomatedCallbackConfigId,
		
		Wrapup: u.Wrapup,
		
		AfterCallWork: u.AfterCallWork,
		
		AfterCallWorkRequired: u.AfterCallWorkRequired,
		
		CallerId: u.CallerId,
		
		CallerIdName: u.CallerIdName,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopiccallback) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
