package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationchateventtopicchatmediaparticipant
type Queueconversationchateventtopicchatmediaparticipant struct { 
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
	User *Queueconversationchateventtopicurireference `json:"user,omitempty"`


	// Queue
	Queue *Queueconversationchateventtopicurireference `json:"queue,omitempty"`


	// Team
	Team *Queueconversationchateventtopicurireference `json:"team,omitempty"`


	// Attributes
	Attributes *map[string]string `json:"attributes,omitempty"`


	// ErrorInfo
	ErrorInfo *Queueconversationchateventtopicerrorbody `json:"errorInfo,omitempty"`


	// Script
	Script *Queueconversationchateventtopicurireference `json:"script,omitempty"`


	// WrapupTimeoutMs
	WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`


	// WrapupSkipped
	WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`


	// AlertingTimeoutMs
	AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`


	// Provider
	Provider *string `json:"provider,omitempty"`


	// ExternalContact
	ExternalContact *Queueconversationchateventtopicurireference `json:"externalContact,omitempty"`


	// ExternalOrganization
	ExternalOrganization *Queueconversationchateventtopicurireference `json:"externalOrganization,omitempty"`


	// Wrapup
	Wrapup *Queueconversationchateventtopicwrapup `json:"wrapup,omitempty"`


	// ConversationRoutingData
	ConversationRoutingData *Queueconversationchateventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`


	// Peer
	Peer *string `json:"peer,omitempty"`


	// ScreenRecordingState
	ScreenRecordingState *string `json:"screenRecordingState,omitempty"`


	// FlaggedReason
	FlaggedReason *string `json:"flaggedReason,omitempty"`


	// JourneyContext
	JourneyContext *Queueconversationchateventtopicjourneycontext `json:"journeyContext,omitempty"`


	// StartAcwTime
	StartAcwTime *time.Time `json:"startAcwTime,omitempty"`


	// EndAcwTime
	EndAcwTime *time.Time `json:"endAcwTime,omitempty"`


	// RoomId
	RoomId *string `json:"roomId,omitempty"`


	// AvatarImageUrl
	AvatarImageUrl *string `json:"avatarImageUrl,omitempty"`

}

func (o *Queueconversationchateventtopicchatmediaparticipant) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationchateventtopicchatmediaparticipant
	
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
		
		Direction *string `json:"direction,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		Held *bool `json:"held,omitempty"`
		
		WrapupRequired *bool `json:"wrapupRequired,omitempty"`
		
		WrapupPrompt *string `json:"wrapupPrompt,omitempty"`
		
		User *Queueconversationchateventtopicurireference `json:"user,omitempty"`
		
		Queue *Queueconversationchateventtopicurireference `json:"queue,omitempty"`
		
		Team *Queueconversationchateventtopicurireference `json:"team,omitempty"`
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		
		ErrorInfo *Queueconversationchateventtopicerrorbody `json:"errorInfo,omitempty"`
		
		Script *Queueconversationchateventtopicurireference `json:"script,omitempty"`
		
		WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`
		
		WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`
		
		AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ExternalContact *Queueconversationchateventtopicurireference `json:"externalContact,omitempty"`
		
		ExternalOrganization *Queueconversationchateventtopicurireference `json:"externalOrganization,omitempty"`
		
		Wrapup *Queueconversationchateventtopicwrapup `json:"wrapup,omitempty"`
		
		ConversationRoutingData *Queueconversationchateventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`
		
		Peer *string `json:"peer,omitempty"`
		
		ScreenRecordingState *string `json:"screenRecordingState,omitempty"`
		
		FlaggedReason *string `json:"flaggedReason,omitempty"`
		
		JourneyContext *Queueconversationchateventtopicjourneycontext `json:"journeyContext,omitempty"`
		
		StartAcwTime *string `json:"startAcwTime,omitempty"`
		
		EndAcwTime *string `json:"endAcwTime,omitempty"`
		
		RoomId *string `json:"roomId,omitempty"`
		
		AvatarImageUrl *string `json:"avatarImageUrl,omitempty"`
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
		
		RoomId: o.RoomId,
		
		AvatarImageUrl: o.AvatarImageUrl,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationchateventtopicchatmediaparticipant) UnmarshalJSON(b []byte) error {
	var QueueconversationchateventtopicchatmediaparticipantMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationchateventtopicchatmediaparticipantMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationchateventtopicchatmediaparticipantMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := QueueconversationchateventtopicchatmediaparticipantMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Address, ok := QueueconversationchateventtopicchatmediaparticipantMap["address"].(string); ok {
		o.Address = &Address
	}
	
	if startTimeString, ok := QueueconversationchateventtopicchatmediaparticipantMap["startTime"].(string); ok {
		StartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startTimeString)
		o.StartTime = &StartTime
	}
	
	if connectedTimeString, ok := QueueconversationchateventtopicchatmediaparticipantMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if endTimeString, ok := QueueconversationchateventtopicchatmediaparticipantMap["endTime"].(string); ok {
		EndTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endTimeString)
		o.EndTime = &EndTime
	}
	
	if startHoldTimeString, ok := QueueconversationchateventtopicchatmediaparticipantMap["startHoldTime"].(string); ok {
		StartHoldTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startHoldTimeString)
		o.StartHoldTime = &StartHoldTime
	}
	
	if Purpose, ok := QueueconversationchateventtopicchatmediaparticipantMap["purpose"].(string); ok {
		o.Purpose = &Purpose
	}
	
	if State, ok := QueueconversationchateventtopicchatmediaparticipantMap["state"].(string); ok {
		o.State = &State
	}
	
	if Direction, ok := QueueconversationchateventtopicchatmediaparticipantMap["direction"].(string); ok {
		o.Direction = &Direction
	}
	
	if DisconnectType, ok := QueueconversationchateventtopicchatmediaparticipantMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
	
	if Held, ok := QueueconversationchateventtopicchatmediaparticipantMap["held"].(bool); ok {
		o.Held = &Held
	}
	
	if WrapupRequired, ok := QueueconversationchateventtopicchatmediaparticipantMap["wrapupRequired"].(bool); ok {
		o.WrapupRequired = &WrapupRequired
	}
	
	if WrapupPrompt, ok := QueueconversationchateventtopicchatmediaparticipantMap["wrapupPrompt"].(string); ok {
		o.WrapupPrompt = &WrapupPrompt
	}
	
	if User, ok := QueueconversationchateventtopicchatmediaparticipantMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Queue, ok := QueueconversationchateventtopicchatmediaparticipantMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Team, ok := QueueconversationchateventtopicchatmediaparticipantMap["team"].(map[string]interface{}); ok {
		TeamString, _ := json.Marshal(Team)
		json.Unmarshal(TeamString, &o.Team)
	}
	
	if Attributes, ok := QueueconversationchateventtopicchatmediaparticipantMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if ErrorInfo, ok := QueueconversationchateventtopicchatmediaparticipantMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	
	if Script, ok := QueueconversationchateventtopicchatmediaparticipantMap["script"].(map[string]interface{}); ok {
		ScriptString, _ := json.Marshal(Script)
		json.Unmarshal(ScriptString, &o.Script)
	}
	
	if WrapupTimeoutMs, ok := QueueconversationchateventtopicchatmediaparticipantMap["wrapupTimeoutMs"].(float64); ok {
		WrapupTimeoutMsInt := int(WrapupTimeoutMs)
		o.WrapupTimeoutMs = &WrapupTimeoutMsInt
	}
	
	if WrapupSkipped, ok := QueueconversationchateventtopicchatmediaparticipantMap["wrapupSkipped"].(bool); ok {
		o.WrapupSkipped = &WrapupSkipped
	}
	
	if AlertingTimeoutMs, ok := QueueconversationchateventtopicchatmediaparticipantMap["alertingTimeoutMs"].(float64); ok {
		AlertingTimeoutMsInt := int(AlertingTimeoutMs)
		o.AlertingTimeoutMs = &AlertingTimeoutMsInt
	}
	
	if Provider, ok := QueueconversationchateventtopicchatmediaparticipantMap["provider"].(string); ok {
		o.Provider = &Provider
	}
	
	if ExternalContact, ok := QueueconversationchateventtopicchatmediaparticipantMap["externalContact"].(map[string]interface{}); ok {
		ExternalContactString, _ := json.Marshal(ExternalContact)
		json.Unmarshal(ExternalContactString, &o.ExternalContact)
	}
	
	if ExternalOrganization, ok := QueueconversationchateventtopicchatmediaparticipantMap["externalOrganization"].(map[string]interface{}); ok {
		ExternalOrganizationString, _ := json.Marshal(ExternalOrganization)
		json.Unmarshal(ExternalOrganizationString, &o.ExternalOrganization)
	}
	
	if Wrapup, ok := QueueconversationchateventtopicchatmediaparticipantMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if ConversationRoutingData, ok := QueueconversationchateventtopicchatmediaparticipantMap["conversationRoutingData"].(map[string]interface{}); ok {
		ConversationRoutingDataString, _ := json.Marshal(ConversationRoutingData)
		json.Unmarshal(ConversationRoutingDataString, &o.ConversationRoutingData)
	}
	
	if Peer, ok := QueueconversationchateventtopicchatmediaparticipantMap["peer"].(string); ok {
		o.Peer = &Peer
	}
	
	if ScreenRecordingState, ok := QueueconversationchateventtopicchatmediaparticipantMap["screenRecordingState"].(string); ok {
		o.ScreenRecordingState = &ScreenRecordingState
	}
	
	if FlaggedReason, ok := QueueconversationchateventtopicchatmediaparticipantMap["flaggedReason"].(string); ok {
		o.FlaggedReason = &FlaggedReason
	}
	
	if JourneyContext, ok := QueueconversationchateventtopicchatmediaparticipantMap["journeyContext"].(map[string]interface{}); ok {
		JourneyContextString, _ := json.Marshal(JourneyContext)
		json.Unmarshal(JourneyContextString, &o.JourneyContext)
	}
	
	if startAcwTimeString, ok := QueueconversationchateventtopicchatmediaparticipantMap["startAcwTime"].(string); ok {
		StartAcwTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startAcwTimeString)
		o.StartAcwTime = &StartAcwTime
	}
	
	if endAcwTimeString, ok := QueueconversationchateventtopicchatmediaparticipantMap["endAcwTime"].(string); ok {
		EndAcwTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endAcwTimeString)
		o.EndAcwTime = &EndAcwTime
	}
	
	if RoomId, ok := QueueconversationchateventtopicchatmediaparticipantMap["roomId"].(string); ok {
		o.RoomId = &RoomId
	}
	
	if AvatarImageUrl, ok := QueueconversationchateventtopicchatmediaparticipantMap["avatarImageUrl"].(string); ok {
		o.AvatarImageUrl = &AvatarImageUrl
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationchateventtopicchatmediaparticipant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
