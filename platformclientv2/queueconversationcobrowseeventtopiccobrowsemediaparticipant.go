package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcobrowseeventtopiccobrowsemediaparticipant
type Queueconversationcobrowseeventtopiccobrowsemediaparticipant struct { 
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
	User *Queueconversationcobrowseeventtopicurireference `json:"user,omitempty"`


	// Queue
	Queue *Queueconversationcobrowseeventtopicurireference `json:"queue,omitempty"`


	// Team
	Team *Queueconversationcobrowseeventtopicurireference `json:"team,omitempty"`


	// Attributes
	Attributes *map[string]string `json:"attributes,omitempty"`


	// ErrorInfo
	ErrorInfo *Queueconversationcobrowseeventtopicerrorbody `json:"errorInfo,omitempty"`


	// Script
	Script *Queueconversationcobrowseeventtopicurireference `json:"script,omitempty"`


	// WrapupTimeoutMs
	WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`


	// WrapupSkipped
	WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`


	// AlertingTimeoutMs
	AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`


	// Provider
	Provider *string `json:"provider,omitempty"`


	// ExternalContact
	ExternalContact *Queueconversationcobrowseeventtopicurireference `json:"externalContact,omitempty"`


	// ExternalOrganization
	ExternalOrganization *Queueconversationcobrowseeventtopicurireference `json:"externalOrganization,omitempty"`


	// Wrapup
	Wrapup *Queueconversationcobrowseeventtopicwrapup `json:"wrapup,omitempty"`


	// ConversationRoutingData
	ConversationRoutingData *Queueconversationcobrowseeventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`


	// Peer
	Peer *string `json:"peer,omitempty"`


	// ScreenRecordingState
	ScreenRecordingState *string `json:"screenRecordingState,omitempty"`


	// FlaggedReason
	FlaggedReason *string `json:"flaggedReason,omitempty"`


	// JourneyContext
	JourneyContext *Queueconversationcobrowseeventtopicjourneycontext `json:"journeyContext,omitempty"`


	// StartAcwTime
	StartAcwTime *time.Time `json:"startAcwTime,omitempty"`


	// EndAcwTime
	EndAcwTime *time.Time `json:"endAcwTime,omitempty"`


	// CobrowseSessionId
	CobrowseSessionId *string `json:"cobrowseSessionId,omitempty"`


	// CobrowseRole
	CobrowseRole *string `json:"cobrowseRole,omitempty"`


	// ViewerUrl
	ViewerUrl *string `json:"viewerUrl,omitempty"`


	// ProviderEventTime
	ProviderEventTime *time.Time `json:"providerEventTime,omitempty"`


	// Controlling
	Controlling *[]string `json:"controlling,omitempty"`

}

func (u *Queueconversationcobrowseeventtopiccobrowsemediaparticipant) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationcobrowseeventtopiccobrowsemediaparticipant

	
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
	
	ProviderEventTime := new(string)
	if u.ProviderEventTime != nil {
		
		*ProviderEventTime = timeutil.Strftime(u.ProviderEventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ProviderEventTime = nil
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
		
		User *Queueconversationcobrowseeventtopicurireference `json:"user,omitempty"`
		
		Queue *Queueconversationcobrowseeventtopicurireference `json:"queue,omitempty"`
		
		Team *Queueconversationcobrowseeventtopicurireference `json:"team,omitempty"`
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		
		ErrorInfo *Queueconversationcobrowseeventtopicerrorbody `json:"errorInfo,omitempty"`
		
		Script *Queueconversationcobrowseeventtopicurireference `json:"script,omitempty"`
		
		WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`
		
		WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`
		
		AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ExternalContact *Queueconversationcobrowseeventtopicurireference `json:"externalContact,omitempty"`
		
		ExternalOrganization *Queueconversationcobrowseeventtopicurireference `json:"externalOrganization,omitempty"`
		
		Wrapup *Queueconversationcobrowseeventtopicwrapup `json:"wrapup,omitempty"`
		
		ConversationRoutingData *Queueconversationcobrowseeventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`
		
		Peer *string `json:"peer,omitempty"`
		
		ScreenRecordingState *string `json:"screenRecordingState,omitempty"`
		
		FlaggedReason *string `json:"flaggedReason,omitempty"`
		
		JourneyContext *Queueconversationcobrowseeventtopicjourneycontext `json:"journeyContext,omitempty"`
		
		StartAcwTime *string `json:"startAcwTime,omitempty"`
		
		EndAcwTime *string `json:"endAcwTime,omitempty"`
		
		CobrowseSessionId *string `json:"cobrowseSessionId,omitempty"`
		
		CobrowseRole *string `json:"cobrowseRole,omitempty"`
		
		ViewerUrl *string `json:"viewerUrl,omitempty"`
		
		ProviderEventTime *string `json:"providerEventTime,omitempty"`
		
		Controlling *[]string `json:"controlling,omitempty"`
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
		
		CobrowseSessionId: u.CobrowseSessionId,
		
		CobrowseRole: u.CobrowseRole,
		
		ViewerUrl: u.ViewerUrl,
		
		ProviderEventTime: ProviderEventTime,
		
		Controlling: u.Controlling,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Queueconversationcobrowseeventtopiccobrowsemediaparticipant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
