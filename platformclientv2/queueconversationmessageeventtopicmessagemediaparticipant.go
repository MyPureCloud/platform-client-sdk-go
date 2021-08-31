package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationmessageeventtopicmessagemediaparticipant
type Queueconversationmessageeventtopicmessagemediaparticipant struct { 
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
	User *Queueconversationmessageeventtopicurireference `json:"user,omitempty"`


	// Queue
	Queue *Queueconversationmessageeventtopicurireference `json:"queue,omitempty"`


	// Team
	Team *Queueconversationmessageeventtopicurireference `json:"team,omitempty"`


	// Attributes
	Attributes *map[string]string `json:"attributes,omitempty"`


	// ErrorInfo
	ErrorInfo *Queueconversationmessageeventtopicerrorbody `json:"errorInfo,omitempty"`


	// Script
	Script *Queueconversationmessageeventtopicurireference `json:"script,omitempty"`


	// WrapupTimeoutMs
	WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`


	// WrapupSkipped
	WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`


	// AlertingTimeoutMs
	AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`


	// Provider
	Provider *string `json:"provider,omitempty"`


	// ExternalContact
	ExternalContact *Queueconversationmessageeventtopicurireference `json:"externalContact,omitempty"`


	// ExternalOrganization
	ExternalOrganization *Queueconversationmessageeventtopicurireference `json:"externalOrganization,omitempty"`


	// Wrapup
	Wrapup *Queueconversationmessageeventtopicwrapup `json:"wrapup,omitempty"`


	// ConversationRoutingData
	ConversationRoutingData *Queueconversationmessageeventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`


	// Peer
	Peer *string `json:"peer,omitempty"`


	// ScreenRecordingState
	ScreenRecordingState *string `json:"screenRecordingState,omitempty"`


	// FlaggedReason
	FlaggedReason *string `json:"flaggedReason,omitempty"`


	// JourneyContext
	JourneyContext *Queueconversationmessageeventtopicjourneycontext `json:"journeyContext,omitempty"`


	// StartAcwTime
	StartAcwTime *time.Time `json:"startAcwTime,omitempty"`


	// EndAcwTime
	EndAcwTime *time.Time `json:"endAcwTime,omitempty"`


	// Messages
	Messages *[]Queueconversationmessageeventtopicmessagedetails `json:"messages,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// RecipientCountry
	RecipientCountry *string `json:"recipientCountry,omitempty"`


	// RecipientType
	RecipientType *string `json:"recipientType,omitempty"`

}

func (o *Queueconversationmessageeventtopicmessagemediaparticipant) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationmessageeventtopicmessagemediaparticipant
	
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
		
		User *Queueconversationmessageeventtopicurireference `json:"user,omitempty"`
		
		Queue *Queueconversationmessageeventtopicurireference `json:"queue,omitempty"`
		
		Team *Queueconversationmessageeventtopicurireference `json:"team,omitempty"`
		
		Attributes *map[string]string `json:"attributes,omitempty"`
		
		ErrorInfo *Queueconversationmessageeventtopicerrorbody `json:"errorInfo,omitempty"`
		
		Script *Queueconversationmessageeventtopicurireference `json:"script,omitempty"`
		
		WrapupTimeoutMs *int `json:"wrapupTimeoutMs,omitempty"`
		
		WrapupSkipped *bool `json:"wrapupSkipped,omitempty"`
		
		AlertingTimeoutMs *int `json:"alertingTimeoutMs,omitempty"`
		
		Provider *string `json:"provider,omitempty"`
		
		ExternalContact *Queueconversationmessageeventtopicurireference `json:"externalContact,omitempty"`
		
		ExternalOrganization *Queueconversationmessageeventtopicurireference `json:"externalOrganization,omitempty"`
		
		Wrapup *Queueconversationmessageeventtopicwrapup `json:"wrapup,omitempty"`
		
		ConversationRoutingData *Queueconversationmessageeventtopicconversationroutingdata `json:"conversationRoutingData,omitempty"`
		
		Peer *string `json:"peer,omitempty"`
		
		ScreenRecordingState *string `json:"screenRecordingState,omitempty"`
		
		FlaggedReason *string `json:"flaggedReason,omitempty"`
		
		JourneyContext *Queueconversationmessageeventtopicjourneycontext `json:"journeyContext,omitempty"`
		
		StartAcwTime *string `json:"startAcwTime,omitempty"`
		
		EndAcwTime *string `json:"endAcwTime,omitempty"`
		
		Messages *[]Queueconversationmessageeventtopicmessagedetails `json:"messages,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		RecipientCountry *string `json:"recipientCountry,omitempty"`
		
		RecipientType *string `json:"recipientType,omitempty"`
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
		
		Messages: o.Messages,
		
		VarType: o.VarType,
		
		RecipientCountry: o.RecipientCountry,
		
		RecipientType: o.RecipientType,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationmessageeventtopicmessagemediaparticipant) UnmarshalJSON(b []byte) error {
	var QueueconversationmessageeventtopicmessagemediaparticipantMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationmessageeventtopicmessagemediaparticipantMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Address, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["address"].(string); ok {
		o.Address = &Address
	}
	
	if startTimeString, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["startTime"].(string); ok {
		StartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startTimeString)
		o.StartTime = &StartTime
	}
	
	if connectedTimeString, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["connectedTime"].(string); ok {
		ConnectedTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", connectedTimeString)
		o.ConnectedTime = &ConnectedTime
	}
	
	if endTimeString, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["endTime"].(string); ok {
		EndTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endTimeString)
		o.EndTime = &EndTime
	}
	
	if startHoldTimeString, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["startHoldTime"].(string); ok {
		StartHoldTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startHoldTimeString)
		o.StartHoldTime = &StartHoldTime
	}
	
	if Purpose, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["purpose"].(string); ok {
		o.Purpose = &Purpose
	}
	
	if State, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["state"].(string); ok {
		o.State = &State
	}
	
	if Direction, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["direction"].(string); ok {
		o.Direction = &Direction
	}
	
	if DisconnectType, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
	
	if Held, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["held"].(bool); ok {
		o.Held = &Held
	}
	
	if WrapupRequired, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["wrapupRequired"].(bool); ok {
		o.WrapupRequired = &WrapupRequired
	}
	
	if WrapupPrompt, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["wrapupPrompt"].(string); ok {
		o.WrapupPrompt = &WrapupPrompt
	}
	
	if User, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Queue, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Team, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["team"].(map[string]interface{}); ok {
		TeamString, _ := json.Marshal(Team)
		json.Unmarshal(TeamString, &o.Team)
	}
	
	if Attributes, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["attributes"].(map[string]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if ErrorInfo, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	
	if Script, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["script"].(map[string]interface{}); ok {
		ScriptString, _ := json.Marshal(Script)
		json.Unmarshal(ScriptString, &o.Script)
	}
	
	if WrapupTimeoutMs, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["wrapupTimeoutMs"].(float64); ok {
		WrapupTimeoutMsInt := int(WrapupTimeoutMs)
		o.WrapupTimeoutMs = &WrapupTimeoutMsInt
	}
	
	if WrapupSkipped, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["wrapupSkipped"].(bool); ok {
		o.WrapupSkipped = &WrapupSkipped
	}
	
	if AlertingTimeoutMs, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["alertingTimeoutMs"].(float64); ok {
		AlertingTimeoutMsInt := int(AlertingTimeoutMs)
		o.AlertingTimeoutMs = &AlertingTimeoutMsInt
	}
	
	if Provider, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["provider"].(string); ok {
		o.Provider = &Provider
	}
	
	if ExternalContact, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["externalContact"].(map[string]interface{}); ok {
		ExternalContactString, _ := json.Marshal(ExternalContact)
		json.Unmarshal(ExternalContactString, &o.ExternalContact)
	}
	
	if ExternalOrganization, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["externalOrganization"].(map[string]interface{}); ok {
		ExternalOrganizationString, _ := json.Marshal(ExternalOrganization)
		json.Unmarshal(ExternalOrganizationString, &o.ExternalOrganization)
	}
	
	if Wrapup, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["wrapup"].(map[string]interface{}); ok {
		WrapupString, _ := json.Marshal(Wrapup)
		json.Unmarshal(WrapupString, &o.Wrapup)
	}
	
	if ConversationRoutingData, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["conversationRoutingData"].(map[string]interface{}); ok {
		ConversationRoutingDataString, _ := json.Marshal(ConversationRoutingData)
		json.Unmarshal(ConversationRoutingDataString, &o.ConversationRoutingData)
	}
	
	if Peer, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["peer"].(string); ok {
		o.Peer = &Peer
	}
	
	if ScreenRecordingState, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["screenRecordingState"].(string); ok {
		o.ScreenRecordingState = &ScreenRecordingState
	}
	
	if FlaggedReason, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["flaggedReason"].(string); ok {
		o.FlaggedReason = &FlaggedReason
	}
	
	if JourneyContext, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["journeyContext"].(map[string]interface{}); ok {
		JourneyContextString, _ := json.Marshal(JourneyContext)
		json.Unmarshal(JourneyContextString, &o.JourneyContext)
	}
	
	if startAcwTimeString, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["startAcwTime"].(string); ok {
		StartAcwTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startAcwTimeString)
		o.StartAcwTime = &StartAcwTime
	}
	
	if endAcwTimeString, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["endAcwTime"].(string); ok {
		EndAcwTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endAcwTimeString)
		o.EndAcwTime = &EndAcwTime
	}
	
	if Messages, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["messages"].([]interface{}); ok {
		MessagesString, _ := json.Marshal(Messages)
		json.Unmarshal(MessagesString, &o.Messages)
	}
	
	if VarType, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if RecipientCountry, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["recipientCountry"].(string); ok {
		o.RecipientCountry = &RecipientCountry
	}
	
	if RecipientType, ok := QueueconversationmessageeventtopicmessagemediaparticipantMap["recipientType"].(string); ok {
		o.RecipientType = &RecipientType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationmessageeventtopicmessagemediaparticipant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
