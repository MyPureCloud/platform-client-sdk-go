package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcobrowseeventtopiccobrowsemediaparticipant
type Conversationcobrowseeventtopiccobrowsemediaparticipant struct { 
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
	User *Conversationcobrowseeventtopicurireference `json:"user,omitempty"`


	// Queue
	Queue *Conversationcobrowseeventtopicurireference `json:"queue,omitempty"`


	// Team
	Team *Conversationcobrowseeventtopicurireference `json:"team,omitempty"`


	// Attributes
	Attributes *map[string]string `json:"attributes,omitempty"`


	// ErrorInfo
	ErrorInfo *Conversationcobrowseeventtopicerrorbody `json:"errorInfo,omitempty"`


	// Script
	Script *Conversationcobrowseeventtopicurireference `json:"script,omitempty"`


	// WrapupTimeoutMs
	WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`


	// WrapupSkipped
	WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`


	// AlertingTimeoutMs
	AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`


	// Provider
	Provider *string `json:"provider,omitempty"`


	// ExternalContact
	ExternalContact *Conversationcobrowseeventtopicurireference `json:"externalContact,omitempty"`


	// ExternalOrganization
	ExternalOrganization *Conversationcobrowseeventtopicurireference `json:"externalOrganization,omitempty"`


	// Wrapup
	Wrapup *Conversationcobrowseeventtopicwrapup `json:"wrapup,omitempty"`


	// ConversationRoutingData
	ConversationRoutingData *Conversationcobrowseeventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`


	// Peer
	Peer *string `json:"peer,omitempty"`


	// ScreenRecordingState
	ScreenRecordingState *string `json:"screenRecordingState,omitempty"`


	// FlaggedReason
	FlaggedReason *string `json:"flaggedReason,omitempty"`


	// JourneyContext
	JourneyContext *Conversationcobrowseeventtopicjourneycontext `json:"journeyContext,omitempty"`


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

func (o *Conversationcobrowseeventtopiccobrowsemediaparticipant) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcobrowseeventtopiccobrowsemediaparticipant
	
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
	
	ProviderEventTime := new(string)
	if o.ProviderEventTime != nil {
		
		*ProviderEventTime = timeutil.Strftime(o.ProviderEventTime, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		
		InitialState *string `json:"initialState,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		
		WrapupRequired *bool `json:"wrapupRequired,omitempty"`
		
		WrapupPrompt *string `json:"wrapupPrompt,omitempty"`
		
		User *Conversationcobrowseeventtopicurireference `json:"user,omitempty"`
		
		Queue *Conversationcobrowseeventtopicurireference `json:"queue,omitempty"`
		
		Team *Conversationcobrowseeventtopicurireference `json:"team,omitempty"`
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		
		ErrorInfo *Conversationcobrowseeventtopicerrorbody `json:"errorInfo,omitempty"`
		
		Script *Conversationcobrowseeventtopicurireference `json:"script,omitempty"`
		
		WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`
		
		WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`
		
		AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ExternalContact *Conversationcobrowseeventtopicurireference `json:"externalContact,omitempty"`
		
		ExternalOrganization *Conversationcobrowseeventtopicurireference `json:"externalOrganization,omitempty"`
		
		Wrapup *Conversationcobrowseeventtopicwrapup `json:"wrapup,omitempty"`
		
		ConversationRoutingData *Conversationcobrowseeventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`
		
		Peer *string `json:"peer,omitempty"`
		
		ScreenRecordingState *string `json:"screenRecordingState,omitempty"`
		
		FlaggedReason *string `json:"flaggedReason,omitempty"`
		
		JourneyContext *Conversationcobrowseeventtopicjourneycontext `json:"journeyContext,omitempty"`
		
		StartAcwTime *string `json:"startAcwTime,omitempty"`
		
		EndAcwTime *string `json:"endAcwTime,omitempty"`
		
		CobrowseSessionId *string `json:"cobrowseSessionId,omitempty"`
		
		CobrowseRole *string `json:"cobrowseRole,omitempty"`
		
		ViewerUrl *string `json:"viewerUrl,omitempty"`
		
		ProviderEventTime *string `json:"providerEventTime,omitempty"`
		
		Controlling *[]string `json:"controlling,omitempty"`
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
		
		CobrowseSessionId: o.CobrowseSessionId,
		
		CobrowseRole: o.CobrowseRole,
		
		ViewerUrl: o.ViewerUrl,
		
		ProviderEventTime: ProviderEventTime,
		
		Controlling: o.Controlling,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationcobrowseeventtopiccobrowsemediaparticipant) UnmarshalJSON(b []byte) error {
	var ConversationcobrowseeventtopiccobrowsemediaparticipantMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcobrowseeventtopiccobrowsemediaparticipantMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Address, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["address"].(string); ok {
		o.Address = &Address
	}
    
	if startTimeString, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["startTime"].(string); ok {
		StartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startTimeString)
		o.StartTime = &StartTime
	}
	
	if connectedTimeString, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if endTimeString, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["endTime"].(string); ok {
		EndTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endTimeString)
		o.EndTime = &EndTime
	}
	
	if startHoldTimeString, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["startHoldTime"].(string); ok {
		StartHoldTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startHoldTimeString)
		o.StartHoldTime = &StartHoldTime
	}
	
	if Purpose, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["purpose"].(string); ok {
		o.Purpose = &Purpose
	}
    
	if State, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["state"].(string); ok {
		o.State = &State
	}
    
	if InitialState, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["initialState"].(string); ok {
		o.InitialState = &InitialState
	}
    
	if Direction, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if DisconnectType, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    
	if Held, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["held"].(bool); ok {
		o.Held = &Held
	}
    
	if WrapupRequired, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["wrapupRequired"].(bool); ok {
		o.WrapupRequired = &WrapupRequired
	}
    
	if WrapupPrompt, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["wrapupPrompt"].(string); ok {
		o.WrapupPrompt = &WrapupPrompt
	}
    
	if User, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Queue, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Team, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["team"].(map[string]interface{}); ok {
		TeamString, _ := json.Marshal(Team)
		json.Unmarshal(TeamString, &o.Team)
	}
	
	if Attributes, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if ErrorInfo, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	
	if Script, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["script"].(map[string]interface{}); ok {
		ScriptString, _ := json.Marshal(Script)
		json.Unmarshal(ScriptString, &o.Script)
	}
	
	if WrapupTimeoutMs, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["wrapupTimeoutMs"].(float64); ok {
		WrapupTimeoutMsInt := int(WrapupTimeoutMs)
		o.WrapupTimeoutMs = &WrapupTimeoutMsInt
	}
	
	if WrapupSkipped, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["wrapupSkipped"].(bool); ok {
		o.WrapupSkipped = &WrapupSkipped
	}
    
	if AlertingTimeoutMs, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["alertingTimeoutMs"].(float64); ok {
		AlertingTimeoutMsInt := int(AlertingTimeoutMs)
		o.AlertingTimeoutMs = &AlertingTimeoutMsInt
	}
	
	if Provider, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["provider"].(string); ok {
		o.Provider = &Provider
	}
    
	if ExternalContact, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["externalContact"].(map[string]interface{}); ok {
		ExternalContactString, _ := json.Marshal(ExternalContact)
		json.Unmarshal(ExternalContactString, &o.ExternalContact)
	}
	
	if ExternalOrganization, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["externalOrganization"].(map[string]interface{}); ok {
		ExternalOrganizationString, _ := json.Marshal(ExternalOrganization)
		json.Unmarshal(ExternalOrganizationString, &o.ExternalOrganization)
	}
	
	if Wrapup, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if ConversationRoutingData, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["conversationRoutingData"].(map[string]interface{}); ok {
		ConversationRoutingDataString, _ := json.Marshal(ConversationRoutingData)
		json.Unmarshal(ConversationRoutingDataString, &o.ConversationRoutingData)
	}
	
	if Peer, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["peer"].(string); ok {
		o.Peer = &Peer
	}
    
	if ScreenRecordingState, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["screenRecordingState"].(string); ok {
		o.ScreenRecordingState = &ScreenRecordingState
	}
    
	if FlaggedReason, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["flaggedReason"].(string); ok {
		o.FlaggedReason = &FlaggedReason
	}
    
	if JourneyContext, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["journeyContext"].(map[string]interface{}); ok {
		JourneyContextString, _ := json.Marshal(JourneyContext)
		json.Unmarshal(JourneyContextString, &o.JourneyContext)
	}
	
	if startAcwTimeString, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["startAcwTime"].(string); ok {
		StartAcwTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startAcwTimeString)
		o.StartAcwTime = &StartAcwTime
	}
	
	if endAcwTimeString, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["endAcwTime"].(string); ok {
		EndAcwTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endAcwTimeString)
		o.EndAcwTime = &EndAcwTime
	}
	
	if CobrowseSessionId, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["cobrowseSessionId"].(string); ok {
		o.CobrowseSessionId = &CobrowseSessionId
	}
    
	if CobrowseRole, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["cobrowseRole"].(string); ok {
		o.CobrowseRole = &CobrowseRole
	}
    
	if ViewerUrl, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["viewerUrl"].(string); ok {
		o.ViewerUrl = &ViewerUrl
	}
    
	if providerEventTimeString, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["providerEventTime"].(string); ok {
		ProviderEventTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", providerEventTimeString)
		o.ProviderEventTime = &ProviderEventTime
	}
	
	if Controlling, ok := ConversationcobrowseeventtopiccobrowsemediaparticipantMap["controlling"].([]interface{}); ok {
		ControllingString, _ := json.Marshal(Controlling)
		json.Unmarshal(ControllingString, &o.Controlling)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcobrowseeventtopiccobrowsemediaparticipant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
