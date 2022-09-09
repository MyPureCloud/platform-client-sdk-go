package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationscreenshareeventtopicscreensharemediaparticipant
type Conversationscreenshareeventtopicscreensharemediaparticipant struct { 
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


	// InitialState
	InitialState *string `json:"initialState,omitempty"`


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
	User *Conversationscreenshareeventtopicurireference `json:"user,omitempty"`


	// Queue
	Queue *Conversationscreenshareeventtopicurireference `json:"queue,omitempty"`


	// Team
	Team *Conversationscreenshareeventtopicurireference `json:"team,omitempty"`


	// Attributes
	Attributes *map[string]string `json:"attributes,omitempty"`


	// ErrorInfo
	ErrorInfo *Conversationscreenshareeventtopicerrorbody `json:"errorInfo,omitempty"`


	// Script
	Script *Conversationscreenshareeventtopicurireference `json:"script,omitempty"`


	// WrapupTimeoutMs
	WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`


	// WrapupSkipped
	WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`


	// AlertingTimeoutMs
	AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`


	// Provider
	Provider *string `json:"provider,omitempty"`


	// ExternalContact
	ExternalContact *Conversationscreenshareeventtopicurireference `json:"externalContact,omitempty"`


	// ExternalOrganization
	ExternalOrganization *Conversationscreenshareeventtopicurireference `json:"externalOrganization,omitempty"`


	// Wrapup
	Wrapup *Conversationscreenshareeventtopicwrapup `json:"wrapup,omitempty"`


	// ConversationRoutingData
	ConversationRoutingData *Conversationscreenshareeventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`


	// Peer
	Peer *string `json:"peer,omitempty"`


	// ScreenRecordingState
	ScreenRecordingState *string `json:"screenRecordingState,omitempty"`


	// FlaggedReason
	FlaggedReason *string `json:"flaggedReason,omitempty"`


	// JourneyContext
	JourneyContext *Conversationscreenshareeventtopicjourneycontext `json:"journeyContext,omitempty"`


	// StartAcwTime
	StartAcwTime *time.Time `json:"startAcwTime,omitempty"`


	// EndAcwTime
	EndAcwTime *time.Time `json:"endAcwTime,omitempty"`


	// Context
	Context *string `json:"context,omitempty"`


	// PeerCount
	PeerCount *int `json:"peerCount,omitempty"`


	// Sharing
	Sharing *bool `json:"sharing,omitempty"`

}

func (o *Conversationscreenshareeventtopicscreensharemediaparticipant) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationscreenshareeventtopicscreensharemediaparticipant
	
	StartTime := new(string)
	if o.StartTime != nil {
		
		*StartTime = timeutil.Strftime(o.StartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartTime = nil
	}
	
	ConnectedTime := new(string)
	if o.ConnectedTime != nil {
		
		*ConnectedTime = timeutil.Strftime(o.ConnectedTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ConnectedTime = nil
	}
	
	EndTime := new(string)
	if o.EndTime != nil {
		
		*EndTime = timeutil.Strftime(o.EndTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndTime = nil
	}
	
	StartHoldTime := new(string)
	if o.StartHoldTime != nil {
		
		*StartHoldTime = timeutil.Strftime(o.StartHoldTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartHoldTime = nil
	}
	
	StartAcwTime := new(string)
	if o.StartAcwTime != nil {
		
		*StartAcwTime = timeutil.Strftime(o.StartAcwTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartAcwTime = nil
	}
	
	EndAcwTime := new(string)
	if o.EndAcwTime != nil {
		
		*EndAcwTime = timeutil.Strftime(o.EndAcwTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndAcwTime = nil
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
		
		InitialState *string `json:"initialState,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		
		WrapupRequired *bool `json:"wrapupRequired,omitempty"`
		
		WrapupPrompt *string `json:"wrapupPrompt,omitempty"`
		
		User *Conversationscreenshareeventtopicurireference `json:"user,omitempty"`
		
		Queue *Conversationscreenshareeventtopicurireference `json:"queue,omitempty"`
		
		Team *Conversationscreenshareeventtopicurireference `json:"team,omitempty"`
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		
		ErrorInfo *Conversationscreenshareeventtopicerrorbody `json:"errorInfo,omitempty"`
		
		Script *Conversationscreenshareeventtopicurireference `json:"script,omitempty"`
		
		WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`
		
		WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`
		
		AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ExternalContact *Conversationscreenshareeventtopicurireference `json:"externalContact,omitempty"`
		
		ExternalOrganization *Conversationscreenshareeventtopicurireference `json:"externalOrganization,omitempty"`
		
		Wrapup *Conversationscreenshareeventtopicwrapup `json:"wrapup,omitempty"`
		
		ConversationRoutingData *Conversationscreenshareeventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`
		
		Peer *string `json:"peer,omitempty"`
		
		ScreenRecordingState *string `json:"screenRecordingState,omitempty"`
		
		FlaggedReason *string `json:"flaggedReason,omitempty"`
		
		JourneyContext *Conversationscreenshareeventtopicjourneycontext `json:"journeyContext,omitempty"`
		
		StartAcwTime *string `json:"startAcwTime,omitempty"`
		
		EndAcwTime *string `json:"endAcwTime,omitempty"`
		
		Context *string `json:"context,omitempty"`
		
		PeerCount *int `json:"peerCount,omitempty"`
		
		Sharing *bool `json:"sharing,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Address: o.Address,
		
		StartTime: StartTime,
		
		ConnectedTime: ConnectedTime,
		
		EndTime: EndTime,
		
		StartHoldTime: StartHoldTime,
		
		Purpose: o.Purpose,
		
		State: o.State,
		
		InitialState: o.InitialState,
		
		Direction: o.Direction,
		
		DisconnectType: o.DisconnectType,
		
		Held: o.Held,
		
		WrapupRequired: o.WrapupRequired,
		
		WrapupPrompt: o.WrapupPrompt,
		
		User: o.User,
		
		Queue: o.Queue,
		
		Team: o.Team,
		
		Attributes: o.Attributes,
		
		ErrorInfo: o.ErrorInfo,
		
		Script: o.Script,
		
		WrapupTimeoutMs: o.WrapupTimeoutMs,
		
		WrapupSkipped: o.WrapupSkipped,
		
		AlertingTimeoutMs: o.AlertingTimeoutMs,
		
		Provider: o.Provider,
		
		ExternalContact: o.ExternalContact,
		
		ExternalOrganization: o.ExternalOrganization,
		
		Wrapup: o.Wrapup,
		
		ConversationRoutingData: o.ConversationRoutingData,
		
		Peer: o.Peer,
		
		ScreenRecordingState: o.ScreenRecordingState,
		
		FlaggedReason: o.FlaggedReason,
		
		JourneyContext: o.JourneyContext,
		
		StartAcwTime: StartAcwTime,
		
		EndAcwTime: EndAcwTime,
		
		Context: o.Context,
		
		PeerCount: o.PeerCount,
		
		Sharing: o.Sharing,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationscreenshareeventtopicscreensharemediaparticipant) UnmarshalJSON(b []byte) error {
	var ConversationscreenshareeventtopicscreensharemediaparticipantMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationscreenshareeventtopicscreensharemediaparticipantMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Address, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["address"].(string); ok {
		o.Address = &Address
	}
    
	if startTimeString, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["startTime"].(string); ok {
		StartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startTimeString)
		o.StartTime = &StartTime
	}
	
	if connectedTimeString, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if endTimeString, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["endTime"].(string); ok {
		EndTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endTimeString)
		o.EndTime = &EndTime
	}
	
	if startHoldTimeString, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["startHoldTime"].(string); ok {
		StartHoldTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startHoldTimeString)
		o.StartHoldTime = &StartHoldTime
	}
	
	if Purpose, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["purpose"].(string); ok {
		o.Purpose = &Purpose
	}
    
	if State, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["state"].(string); ok {
		o.State = &State
	}
    
	if InitialState, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["initialState"].(string); ok {
		o.InitialState = &InitialState
	}
    
	if Direction, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if DisconnectType, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    
	if Held, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["held"].(bool); ok {
		o.Held = &Held
	}
    
	if WrapupRequired, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["wrapupRequired"].(bool); ok {
		o.WrapupRequired = &WrapupRequired
	}
    
	if WrapupPrompt, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["wrapupPrompt"].(string); ok {
		o.WrapupPrompt = &WrapupPrompt
	}
    
	if User, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Queue, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Team, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["team"].(map[string]interface{}); ok {
		TeamString, _ := json.Marshal(Team)
		json.Unmarshal(TeamString, &o.Team)
	}
	
	if Attributes, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if ErrorInfo, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	
	if Script, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["script"].(map[string]interface{}); ok {
		ScriptString, _ := json.Marshal(Script)
		json.Unmarshal(ScriptString, &o.Script)
	}
	
	if WrapupTimeoutMs, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["wrapupTimeoutMs"].(float64); ok {
		WrapupTimeoutMsInt := int(WrapupTimeoutMs)
		o.WrapupTimeoutMs = &WrapupTimeoutMsInt
	}
	
	if WrapupSkipped, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["wrapupSkipped"].(bool); ok {
		o.WrapupSkipped = &WrapupSkipped
	}
    
	if AlertingTimeoutMs, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["alertingTimeoutMs"].(float64); ok {
		AlertingTimeoutMsInt := int(AlertingTimeoutMs)
		o.AlertingTimeoutMs = &AlertingTimeoutMsInt
	}
	
	if Provider, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if ExternalContact, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["externalContact"].(map[string]interface{}); ok {
		ExternalContactString, _ := json.Marshal(ExternalContact)
		json.Unmarshal(ExternalContactString, &o.ExternalContact)
	}
	
	if ExternalOrganization, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["externalOrganization"].(map[string]interface{}); ok {
		ExternalOrganizationString, _ := json.Marshal(ExternalOrganization)
		json.Unmarshal(ExternalOrganizationString, &o.ExternalOrganization)
	}
	
	if Wrapup, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if ConversationRoutingData, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["conversationRoutingData"].(map[string]interface{}); ok {
		ConversationRoutingDataString, _ := json.Marshal(ConversationRoutingData)
		json.Unmarshal(ConversationRoutingDataString, &o.ConversationRoutingData)
	}
	
	if Peer, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["peer"].(string); ok {
		o.Peer = &Peer
	}
    
	if ScreenRecordingState, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["screenRecordingState"].(string); ok {
		o.ScreenRecordingState = &ScreenRecordingState
	}
    
	if FlaggedReason, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["flaggedReason"].(string); ok {
		o.FlaggedReason = &FlaggedReason
	}
    
	if JourneyContext, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["journeyContext"].(map[string]interface{}); ok {
		JourneyContextString, _ := json.Marshal(JourneyContext)
		json.Unmarshal(JourneyContextString, &o.JourneyContext)
	}
	
	if startAcwTimeString, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["startAcwTime"].(string); ok {
		StartAcwTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startAcwTimeString)
		o.StartAcwTime = &StartAcwTime
	}
	
	if endAcwTimeString, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["endAcwTime"].(string); ok {
		EndAcwTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endAcwTimeString)
		o.EndAcwTime = &EndAcwTime
	}
	
	if Context, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["context"].(string); ok {
		o.Context = &Context
	}
    
	if PeerCount, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["peerCount"].(float64); ok {
		PeerCountInt := int(PeerCount)
		o.PeerCount = &PeerCountInt
	}
	
	if Sharing, ok := ConversationscreenshareeventtopicscreensharemediaparticipantMap["sharing"].(bool); ok {
		o.Sharing = &Sharing
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationscreenshareeventtopicscreensharemediaparticipant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
