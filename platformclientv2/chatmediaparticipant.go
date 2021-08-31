package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Chatmediaparticipant
type Chatmediaparticipant struct { 
	// Id - The unique participant ID.
	Id *string `json:"id,omitempty"`


	// Name - The display friendly name of the participant.
	Name *string `json:"name,omitempty"`


	// Address - The participant address.
	Address *string `json:"address,omitempty"`


	// StartTime - The time when this participant first joined the conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartTime *time.Time `json:"startTime,omitempty"`


	// ConnectedTime - The time when this participant went connected for this media (eg: video connected time). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ConnectedTime *time.Time `json:"connectedTime,omitempty"`


	// EndTime - The time when this participant went disconnected for this media (eg: video disconnected time). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndTime *time.Time `json:"endTime,omitempty"`


	// StartHoldTime - The time when this participant's hold started. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartHoldTime *time.Time `json:"startHoldTime,omitempty"`


	// Purpose - The participant's purpose.  Values can be: 'agent', 'user', 'customer', 'external', 'acd', 'ivr
	Purpose *string `json:"purpose,omitempty"`


	// State - The participant's state.  Values can be: 'alerting', 'connected', 'disconnected', 'dialing', 'contacting
	State *string `json:"state,omitempty"`


	// Direction - The participant's direction.  Values can be: 'inbound' or 'outbound'
	Direction *string `json:"direction,omitempty"`


	// DisconnectType - The reason the participant was disconnected from the conversation.
	DisconnectType *string `json:"disconnectType,omitempty"`


	// Held - Value is true when the participant is on hold.
	Held *bool `json:"held,omitempty"`


	// WrapupRequired - Value is true when the participant requires wrap-up.
	WrapupRequired *bool `json:"wrapupRequired,omitempty"`


	// WrapupPrompt - The wrap-up prompt indicating the type of wrap-up to be performed.
	WrapupPrompt *string `json:"wrapupPrompt,omitempty"`


	// User - The PureCloud user for this participant.
	User *Domainentityref `json:"user,omitempty"`


	// Queue - The PureCloud queue for this participant.
	Queue *Domainentityref `json:"queue,omitempty"`


	// Team - The PureCloud team for this participant.
	Team *Domainentityref `json:"team,omitempty"`


	// Attributes - A list of ad-hoc attributes for the participant.
	Attributes *map[string]string `json:"attributes,omitempty"`


	// ErrorInfo - If the conversation ends in error, contains additional error details.
	ErrorInfo *Errorinfo `json:"errorInfo,omitempty"`


	// Script - The Engage script that should be used by this participant.
	Script *Domainentityref `json:"script,omitempty"`


	// WrapupTimeoutMs - The amount of time the participant has to complete wrap-up.
	WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`


	// WrapupSkipped - Value is true when the participant has skipped wrap-up.
	WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`


	// AlertingTimeoutMs - Specifies how long the agent has to answer an interaction before being marked as not responding.
	AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`


	// Provider - The source provider for the communication.
	Provider *string `json:"provider,omitempty"`


	// ExternalContact - If this participant represents an external contact, then this will be the reference for the external contact.
	ExternalContact *Domainentityref `json:"externalContact,omitempty"`


	// ExternalOrganization - If this participant represents an external org, then this will be the reference for the external org.
	ExternalOrganization *Domainentityref `json:"externalOrganization,omitempty"`


	// Wrapup - Wrapup for this participant, if it has been applied.
	Wrapup *Wrapup `json:"wrapup,omitempty"`


	// Peer - The peer communication corresponding to a matching leg for this communication.
	Peer *string `json:"peer,omitempty"`


	// FlaggedReason - The reason specifying why participant flagged the conversation.
	FlaggedReason *string `json:"flaggedReason,omitempty"`


	// JourneyContext - Journey System data/context that is applicable to this communication.  When used for historical purposes, the context should be immutable.  When null, there is no applicable Journey System context.
	JourneyContext *Journeycontext `json:"journeyContext,omitempty"`


	// ConversationRoutingData - Information on how a communication should be routed to an agent.
	ConversationRoutingData *Conversationroutingdata `json:"conversationRoutingData,omitempty"`


	// StartAcwTime - The timestamp when this participant started after-call work. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartAcwTime *time.Time `json:"startAcwTime,omitempty"`


	// EndAcwTime - The timestamp when this participant ended after-call work. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndAcwTime *time.Time `json:"endAcwTime,omitempty"`


	// RoomId - The ID of the chat room.
	RoomId *string `json:"roomId,omitempty"`


	// AvatarImageUrl - If available, the URI to the avatar image of this communication.
	AvatarImageUrl *string `json:"avatarImageUrl,omitempty"`

}

func (o *Chatmediaparticipant) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Chatmediaparticipant
	
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
		
		User *Domainentityref `json:"user,omitempty"`
		
		Queue *Domainentityref `json:"queue,omitempty"`
		
		Team *Domainentityref `json:"team,omitempty"`
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		
		ErrorInfo *Errorinfo `json:"errorInfo,omitempty"`
		
		Script *Domainentityref `json:"script,omitempty"`
		
		WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`
		
		WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`
		
		AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ExternalContact *Domainentityref `json:"externalContact,omitempty"`
		
		ExternalOrganization *Domainentityref `json:"externalOrganization,omitempty"`
		
		Wrapup *Wrapup `json:"wrapup,omitempty"`
		
		Peer *string `json:"peer,omitempty"`
		
		FlaggedReason *string `json:"flaggedReason,omitempty"`
		
		JourneyContext *Journeycontext `json:"journeyContext,omitempty"`
		
		ConversationRoutingData *Conversationroutingdata `json:"conversationRoutingData,omitempty"`
		
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
		
		Peer: o.Peer,
		
		FlaggedReason: o.FlaggedReason,
		
		JourneyContext: o.JourneyContext,
		
		ConversationRoutingData: o.ConversationRoutingData,
		
		StartAcwTime: StartAcwTime,
		
		EndAcwTime: EndAcwTime,
		
		RoomId: o.RoomId,
		
		AvatarImageUrl: o.AvatarImageUrl,
		Alias:    (*Alias)(o),
	})
}

func (o *Chatmediaparticipant) UnmarshalJSON(b []byte) error {
	var ChatmediaparticipantMap map[string]interface{}
	err := json.Unmarshal(b, &ChatmediaparticipantMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ChatmediaparticipantMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := ChatmediaparticipantMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Address, ok := ChatmediaparticipantMap["address"].(string); ok {
		o.Address = &Address
	}
	
	if startTimeString, ok := ChatmediaparticipantMap["startTime"].(string); ok {
		StartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startTimeString)
		o.StartTime = &StartTime
	}
	
	if connectedTimeString, ok := ChatmediaparticipantMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if endTimeString, ok := ChatmediaparticipantMap["endTime"].(string); ok {
		EndTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endTimeString)
		o.EndTime = &EndTime
	}
	
	if startHoldTimeString, ok := ChatmediaparticipantMap["startHoldTime"].(string); ok {
		StartHoldTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startHoldTimeString)
		o.StartHoldTime = &StartHoldTime
	}
	
	if Purpose, ok := ChatmediaparticipantMap["purpose"].(string); ok {
		o.Purpose = &Purpose
	}
	
	if State, ok := ChatmediaparticipantMap["state"].(string); ok {
		o.State = &State
	}
	
	if Direction, ok := ChatmediaparticipantMap["direction"].(string); ok {
		o.Direction = &Direction
	}
	
	if DisconnectType, ok := ChatmediaparticipantMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
	
	if Held, ok := ChatmediaparticipantMap["held"].(bool); ok {
		o.Held = &Held
	}
	
	if WrapupRequired, ok := ChatmediaparticipantMap["wrapupRequired"].(bool); ok {
		o.WrapupRequired = &WrapupRequired
	}
	
	if WrapupPrompt, ok := ChatmediaparticipantMap["wrapupPrompt"].(string); ok {
		o.WrapupPrompt = &WrapupPrompt
	}
	
	if User, ok := ChatmediaparticipantMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Queue, ok := ChatmediaparticipantMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Team, ok := ChatmediaparticipantMap["team"].(map[string]interface{}); ok {
		TeamString, _ := json.Marshal(Team)
		json.Unmarshal(TeamString, &o.Team)
	}
	
	if Attributes, ok := ChatmediaparticipantMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if ErrorInfo, ok := ChatmediaparticipantMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	
	if Script, ok := ChatmediaparticipantMap["script"].(map[string]interface{}); ok {
		ScriptString, _ := json.Marshal(Script)
		json.Unmarshal(ScriptString, &o.Script)
	}
	
	if WrapupTimeoutMs, ok := ChatmediaparticipantMap["wrapupTimeoutMs"].(float64); ok {
		WrapupTimeoutMsInt := int(WrapupTimeoutMs)
		o.WrapupTimeoutMs = &WrapupTimeoutMsInt
	}
	
	if WrapupSkipped, ok := ChatmediaparticipantMap["wrapupSkipped"].(bool); ok {
		o.WrapupSkipped = &WrapupSkipped
	}
	
	if AlertingTimeoutMs, ok := ChatmediaparticipantMap["alertingTimeoutMs"].(float64); ok {
		AlertingTimeoutMsInt := int(AlertingTimeoutMs)
		o.AlertingTimeoutMs = &AlertingTimeoutMsInt
	}
	
	if Provider, ok := ChatmediaparticipantMap["provider"].(string); ok {
		o.Provider = &Provider
	}
	
	if ExternalContact, ok := ChatmediaparticipantMap["externalContact"].(map[string]interface{}); ok {
		ExternalContactString, _ := json.Marshal(ExternalContact)
		json.Unmarshal(ExternalContactString, &o.ExternalContact)
	}
	
	if ExternalOrganization, ok := ChatmediaparticipantMap["externalOrganization"].(map[string]interface{}); ok {
		ExternalOrganizationString, _ := json.Marshal(ExternalOrganization)
		json.Unmarshal(ExternalOrganizationString, &o.ExternalOrganization)
	}
	
	if Wrapup, ok := ChatmediaparticipantMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if Peer, ok := ChatmediaparticipantMap["peer"].(string); ok {
		o.Peer = &Peer
	}
	
	if FlaggedReason, ok := ChatmediaparticipantMap["flaggedReason"].(string); ok {
		o.FlaggedReason = &FlaggedReason
	}
	
	if JourneyContext, ok := ChatmediaparticipantMap["journeyContext"].(map[string]interface{}); ok {
		JourneyContextString, _ := json.Marshal(JourneyContext)
		json.Unmarshal(JourneyContextString, &o.JourneyContext)
	}
	
	if ConversationRoutingData, ok := ChatmediaparticipantMap["conversationRoutingData"].(map[string]interface{}); ok {
		ConversationRoutingDataString, _ := json.Marshal(ConversationRoutingData)
		json.Unmarshal(ConversationRoutingDataString, &o.ConversationRoutingData)
	}
	
	if startAcwTimeString, ok := ChatmediaparticipantMap["startAcwTime"].(string); ok {
		StartAcwTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startAcwTimeString)
		o.StartAcwTime = &StartAcwTime
	}
	
	if endAcwTimeString, ok := ChatmediaparticipantMap["endAcwTime"].(string); ok {
		EndAcwTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endAcwTimeString)
		o.EndAcwTime = &EndAcwTime
	}
	
	if RoomId, ok := ChatmediaparticipantMap["roomId"].(string); ok {
		o.RoomId = &RoomId
	}
	
	if AvatarImageUrl, ok := ChatmediaparticipantMap["avatarImageUrl"].(string); ok {
		o.AvatarImageUrl = &AvatarImageUrl
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Chatmediaparticipant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
