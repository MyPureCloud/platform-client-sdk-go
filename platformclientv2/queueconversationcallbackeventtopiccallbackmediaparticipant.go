package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcallbackeventtopiccallbackmediaparticipant
type Queueconversationcallbackeventtopiccallbackmediaparticipant struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Address
	Address *string `json:"address,omitempty"`


	// StartTime
	StartTime *time.Time `json:"startTime,omitempty"`


	// ConnectedTime
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// EndTime
	EndTime *time.Time `json:"endTime,omitempty"`


	// StartHoldTime
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`


	// Purpose
	Purpose *string `json:"purpose,omitempty"`


	// State
	State *string `json:"state,omitempty"`


	// Direction
	Direction *string `json:"direction,omitempty"`


	// DisconnectType
	DisconnectType *string `json:"disconnectType,omitempty"`


	// Held
	Held *bool `json:"held,omitempty"`


	// WrapupRequired
	WrapupRequired *bool `json:"wrapupRequired,omitempty"`


	// WrapupPrompt
	WrapupPrompt *string `json:"wrapupPrompt,omitempty"`


	// User
	User *Queueconversationcallbackeventtopicurireference `json:"user,omitempty"`


	// Queue
	Queue *Queueconversationcallbackeventtopicurireference `json:"queue,omitempty"`


	// Team
	Team *Queueconversationcallbackeventtopicurireference `json:"team,omitempty"`


	// Attributes
	Attributes *map[string]string `json:"attributes,omitempty"`


	// ErrorInfo
	ErrorInfo *Queueconversationcallbackeventtopicerrorbody `json:"errorInfo,omitempty"`


	// Script
	Script *Queueconversationcallbackeventtopicurireference `json:"script,omitempty"`


	// WrapupTimeoutMs
	WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`


	// WrapupSkipped
	WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`


	// AlertingTimeoutMs
	AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`


	// Provider
	Provider *string `json:"provider,omitempty"`


	// ExternalContact
	ExternalContact *Queueconversationcallbackeventtopicurireference `json:"externalContact,omitempty"`


	// ExternalOrganization
	ExternalOrganization *Queueconversationcallbackeventtopicurireference `json:"externalOrganization,omitempty"`


	// Wrapup
	Wrapup *Queueconversationcallbackeventtopicwrapup `json:"wrapup,omitempty"`


	// ConversationRoutingData
	ConversationRoutingData *Queueconversationcallbackeventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`


	// Peer
	Peer *string `json:"peer,omitempty"`


	// ScreenRecordingState
	ScreenRecordingState *string `json:"screenRecordingState,omitempty"`


	// FlaggedReason
	FlaggedReason *string `json:"flaggedReason,omitempty"`


	// JourneyContext
	JourneyContext *Queueconversationcallbackeventtopicjourneycontext `json:"journeyContext,omitempty"`


	// StartAcwTime
	StartAcwTime *time.Time `json:"startAcwTime,omitempty"`


	// EndAcwTime
	EndAcwTime *time.Time `json:"endAcwTime,omitempty"`


	// OutboundPreview
	OutboundPreview *Queueconversationcallbackeventtopicdialerpreview `json:"outboundPreview,omitempty"`


	// Voicemail
	Voicemail *Queueconversationcallbackeventtopicvoicemail `json:"voicemail,omitempty"`


	// CallbackNumbers
	CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`


	// CallbackUserName
	CallbackUserName *string `json:"callbackUserName,omitempty"`


	// SkipEnabled
	SkipEnabled *bool `json:"skipEnabled,omitempty"`


	// ExternalCampaign
	ExternalCampaign *bool `json:"externalCampaign,omitempty"`


	// TimeoutSeconds
	TimeoutSeconds *int `json:"timeoutSeconds,omitempty"`


	// CallbackScheduledTime
	CallbackScheduledTime *time.Time `json:"callbackScheduledTime,omitempty"`


	// AutomatedCallbackConfigId
	AutomatedCallbackConfigId *string `json:"automatedCallbackConfigId,omitempty"`

}

func (u *Queueconversationcallbackeventtopiccallbackmediaparticipant) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationcallbackeventtopiccallbackmediaparticipant

	
	StartTime := new(string)
	if u.StartTime != nil {
		
		*StartTime = timeutil.Strftime(u.StartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartTime = nil
	}
	
	ConnectedTime := new(string)
	if u.ConnectedTime != nil {
		
		*ConnectedTime = timeutil.Strftime(u.ConnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConnectedTime = nil
	}
	
	EndTime := new(string)
	if u.EndTime != nil {
		
		*EndTime = timeutil.Strftime(u.EndTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndTime = nil
	}
	
	StartHoldTime := new(string)
	if u.StartHoldTime != nil {
		
		*StartHoldTime = timeutil.Strftime(u.StartHoldTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartHoldTime = nil
	}
	
	StartAcwTime := new(string)
	if u.StartAcwTime != nil {
		
		*StartAcwTime = timeutil.Strftime(u.StartAcwTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartAcwTime = nil
	}
	
	EndAcwTime := new(string)
	if u.EndAcwTime != nil {
		
		*EndAcwTime = timeutil.Strftime(u.EndAcwTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndAcwTime = nil
	}
	
	CallbackScheduledTime := new(string)
	if u.CallbackScheduledTime != nil {
		
		*CallbackScheduledTime = timeutil.Strftime(u.CallbackScheduledTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CallbackScheduledTime = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Address *string `json:"address,omitempty"`
		
		StartTime *string `json:"startTime,omitempty"`
		
		ConnectedTime *string `json:"connectedTime,omitempty"`
		
		EndTime *string `json:"endTime,omitempty"`
		
		StartHoldTime *string `json:"startHoldTime,omitempty"`
		
		Purpose *string `json:"purpose,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		
		WrapupRequired *bool `json:"wrapupRequired,omitempty"`
		
		WrapupPrompt *string `json:"wrapupPrompt,omitempty"`
		
		User *Queueconversationcallbackeventtopicurireference `json:"user,omitempty"`
		
		Queue *Queueconversationcallbackeventtopicurireference `json:"queue,omitempty"`
		
		Team *Queueconversationcallbackeventtopicurireference `json:"team,omitempty"`
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		
		ErrorInfo *Queueconversationcallbackeventtopicerrorbody `json:"errorInfo,omitempty"`
		
		Script *Queueconversationcallbackeventtopicurireference `json:"script,omitempty"`
		
		WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`
		
		WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`
		
		AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ExternalContact *Queueconversationcallbackeventtopicurireference `json:"externalContact,omitempty"`
		
		ExternalOrganization *Queueconversationcallbackeventtopicurireference `json:"externalOrganization,omitempty"`
		
		Wrapup *Queueconversationcallbackeventtopicwrapup `json:"wrapup,omitempty"`
		
		ConversationRoutingData *Queueconversationcallbackeventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`
		
		Peer *string `json:"peer,omitempty"`
		
		ScreenRecordingState *string `json:"screenRecordingState,omitempty"`
		
		FlaggedReason *string `json:"flaggedReason,omitempty"`
		
		JourneyContext *Queueconversationcallbackeventtopicjourneycontext `json:"journeyContext,omitempty"`
		
		StartAcwTime *string `json:"startAcwTime,omitempty"`
		
		EndAcwTime *string `json:"endAcwTime,omitempty"`
		
		OutboundPreview *Queueconversationcallbackeventtopicdialerpreview `json:"outboundPreview,omitempty"`
		
		Voicemail *Queueconversationcallbackeventtopicvoicemail `json:"voicemail,omitempty"`
		
		CallbackNumbers *[]string `json:"callbackNumbers,omitempty"`
		
		CallbackUserName *string `json:"callbackUserName,omitempty"`
		
		SkipEnabled *bool `json:"skipEnabled,omitempty"`
		
		ExternalCampaign *bool `json:"externalCampaign,omitempty"`
		
		TimeoutSeconds *int `json:"timeoutSeconds,omitempty"`
		
		CallbackScheduledTime *string `json:"callbackScheduledTime,omitempty"`
		
		AutomatedCallbackConfigId *string `json:"automatedCallbackConfigId,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Address: u.Address,
		
		StartTime: StartTime,
		
		ConnectedTime: ConnectedTime,
		
		EndTime: EndTime,
		
		StartHoldTime: StartHoldTime,
		
		Purpose: u.Purpose,
		
		State: u.State,
		
		Direction: u.Direction,
		
		DisconnectType: u.DisconnectType,
		
		Held: u.Held,
		
		WrapupRequired: u.WrapupRequired,
		
		WrapupPrompt: u.WrapupPrompt,
		
		User: u.User,
		
		Queue: u.Queue,
		
		Team: u.Team,
		
		Attributes: u.Attributes,
		
		ErrorInfo: u.ErrorInfo,
		
		Script: u.Script,
		
		WrapupTimeoutMs: u.WrapupTimeoutMs,
		
		WrapupSkipped: u.WrapupSkipped,
		
		AlertingTimeoutMs: u.AlertingTimeoutMs,
		
		Provider: u.Provider,
		
		ExternalContact: u.ExternalContact,
		
		ExternalOrganization: u.ExternalOrganization,
		
		Wrapup: u.Wrapup,
		
		ConversationRoutingData: u.ConversationRoutingData,
		
		Peer: u.Peer,
		
		ScreenRecordingState: u.ScreenRecordingState,
		
		FlaggedReason: u.FlaggedReason,
		
		JourneyContext: u.JourneyContext,
		
		StartAcwTime: StartAcwTime,
		
		EndAcwTime: EndAcwTime,
		
		OutboundPreview: u.OutboundPreview,
		
		Voicemail: u.Voicemail,
		
		CallbackNumbers: u.CallbackNumbers,
		
		CallbackUserName: u.CallbackUserName,
		
		SkipEnabled: u.SkipEnabled,
		
		ExternalCampaign: u.ExternalCampaign,
		
		TimeoutSeconds: u.TimeoutSeconds,
		
		CallbackScheduledTime: CallbackScheduledTime,
		
		AutomatedCallbackConfigId: u.AutomatedCallbackConfigId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Queueconversationcallbackeventtopiccallbackmediaparticipant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
