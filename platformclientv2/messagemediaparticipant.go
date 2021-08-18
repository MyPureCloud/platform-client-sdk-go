package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagemediaparticipant
type Messagemediaparticipant struct { 
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


	// ToAddress - Address for the participant on receiving side of the message conversation. If the address is a phone number, E.164 format is recommended.
	ToAddress *Address `json:"toAddress,omitempty"`


	// FromAddress - Address for the participant on the sending side of the message conversation. If the address is a phone number, E.164 format is recommended.
	FromAddress *Address `json:"fromAddress,omitempty"`


	// Messages - Message instance details on the communication.
	Messages *[]Messagedetails `json:"messages,omitempty"`


	// VarType - Indicates the type of message platform from which the message originated.
	VarType *string `json:"type,omitempty"`


	// RecipientCountry - Indicates the country where the recipient is associated in ISO 3166-1 alpha-2 format.
	RecipientCountry *string `json:"recipientCountry,omitempty"`


	// RecipientType - The type of the recipient. Eg: Provisioned phoneNumber is the recipient for sms message type.
	RecipientType *string `json:"recipientType,omitempty"`

}

func (u *Messagemediaparticipant) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messagemediaparticipant

	
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
		
		ToAddress *Address `json:"toAddress,omitempty"`
		
		FromAddress *Address `json:"fromAddress,omitempty"`
		
		Messages *[]Messagedetails `json:"messages,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		RecipientCountry *string `json:"recipientCountry,omitempty"`
		
		RecipientType *string `json:"recipientType,omitempty"`
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
		
		Peer: u.Peer,
		
		FlaggedReason: u.FlaggedReason,
		
		JourneyContext: u.JourneyContext,
		
		ConversationRoutingData: u.ConversationRoutingData,
		
		StartAcwTime: StartAcwTime,
		
		EndAcwTime: EndAcwTime,
		
		ToAddress: u.ToAddress,
		
		FromAddress: u.FromAddress,
		
		Messages: u.Messages,
		
		VarType: u.VarType,
		
		RecipientCountry: u.RecipientCountry,
		
		RecipientType: u.RecipientType,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Messagemediaparticipant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
