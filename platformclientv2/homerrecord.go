package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Homerrecord
type Homerrecord struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Date - metadata associated to the SIP calls. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Date *time.Time `json:"date,omitempty"`


	// MilliTs - metadata associated to the SIP calls
	MilliTs *string `json:"milliTs,omitempty"`


	// MicroTs - metadata associated to the SIP calls
	MicroTs *string `json:"microTs,omitempty"`


	// Method - metadata associated to the SIP calls
	Method *string `json:"method,omitempty"`


	// ReplyReason - metadata associated to the SIP calls
	ReplyReason *string `json:"replyReason,omitempty"`


	// Ruri - metadata associated to the SIP calls
	Ruri *string `json:"ruri,omitempty"`


	// RuriUser - metadata associated to the SIP calls
	RuriUser *string `json:"ruriUser,omitempty"`


	// RuriDomain - metadata associated to the SIP calls
	RuriDomain *string `json:"ruriDomain,omitempty"`


	// FromUser - metadata associated to the SIP calls
	FromUser *string `json:"fromUser,omitempty"`


	// FromDomain - metadata associated to the SIP calls
	FromDomain *string `json:"fromDomain,omitempty"`


	// FromTag - metadata associated to the SIP calls
	FromTag *string `json:"fromTag,omitempty"`


	// ToUser - metadata associated to the SIP calls
	ToUser *string `json:"toUser,omitempty"`


	// ToDomain - metadata associated to the SIP calls
	ToDomain *string `json:"toDomain,omitempty"`


	// ToTag - metadata associated to the SIP calls
	ToTag *string `json:"toTag,omitempty"`


	// PidUser - metadata associated to the SIP calls
	PidUser *string `json:"pidUser,omitempty"`


	// ContactUser - metadata associated to the SIP calls
	ContactUser *string `json:"contactUser,omitempty"`


	// AuthUser - metadata associated to the SIP calls
	AuthUser *string `json:"authUser,omitempty"`


	// Callid - metadata associated to the SIP calls
	Callid *string `json:"callid,omitempty"`


	// CallidAleg - metadata associated to the SIP calls
	CallidAleg *string `json:"callidAleg,omitempty"`


	// Via1 - metadata associated to the SIP calls
	Via1 *string `json:"via1,omitempty"`


	// Via1Branch - metadata associated to the SIP calls
	Via1Branch *string `json:"via1Branch,omitempty"`


	// Cseq - metadata associated to the SIP calls
	Cseq *string `json:"cseq,omitempty"`


	// Diversion - metadata associated to the SIP calls
	Diversion *string `json:"diversion,omitempty"`


	// Reason - metadata associated to the SIP calls
	Reason *string `json:"reason,omitempty"`


	// ContentType - metadata associated to the SIP calls
	ContentType *string `json:"contentType,omitempty"`


	// Auth - metadata associated to the SIP calls
	Auth *string `json:"auth,omitempty"`


	// UserAgent - metadata associated to the SIP calls
	UserAgent *string `json:"userAgent,omitempty"`


	// SourceIp - metadata associated to the SIP calls
	SourceIp *string `json:"sourceIp,omitempty"`


	// SourcePort - metadata associated to the SIP calls
	SourcePort *string `json:"sourcePort,omitempty"`


	// DestinationIp - metadata associated to the SIP calls
	DestinationIp *string `json:"destinationIp,omitempty"`


	// DestinationPort - metadata associated to the SIP calls
	DestinationPort *string `json:"destinationPort,omitempty"`


	// ContactIp - metadata associated to the SIP calls
	ContactIp *string `json:"contactIp,omitempty"`


	// ContactPort - metadata associated to the SIP calls
	ContactPort *string `json:"contactPort,omitempty"`


	// OriginatorIp - metadata associated to the SIP calls
	OriginatorIp *string `json:"originatorIp,omitempty"`


	// OriginatorPort - metadata associated to the SIP calls
	OriginatorPort *string `json:"originatorPort,omitempty"`


	// CorrelationId - metadata associated to the SIP calls
	CorrelationId *string `json:"correlationId,omitempty"`


	// Proto - metadata associated to the SIP calls
	Proto *string `json:"proto,omitempty"`


	// Family - metadata associated to the SIP calls
	Family *string `json:"family,omitempty"`


	// RtpStat - metadata associated to the SIP calls
	RtpStat *string `json:"rtpStat,omitempty"`


	// VarType - metadata associated to the SIP calls
	VarType *string `json:"type,omitempty"`


	// Node - metadata associated to the SIP calls
	Node *string `json:"node,omitempty"`


	// Trans - metadata associated to the SIP calls
	Trans *string `json:"trans,omitempty"`


	// Dbnode - metadata associated to the SIP calls
	Dbnode *string `json:"dbnode,omitempty"`


	// Msg - metadata associated to the SIP calls
	Msg *string `json:"msg,omitempty"`


	// SourceAlias - metadata associated to the SIP calls
	SourceAlias *string `json:"sourceAlias,omitempty"`


	// DestinationAlias - metadata associated to the SIP calls
	DestinationAlias *string `json:"destinationAlias,omitempty"`


	// ConversationId - metadata associated to the SIP calls
	ConversationId *string `json:"conversationId,omitempty"`


	// ParticipantId - metadata associated to the SIP calls
	ParticipantId *string `json:"participantId,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Homerrecord) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Homerrecord

	
	Date := new(string)
	if u.Date != nil {
		
		*Date = timeutil.Strftime(u.Date, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Date = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Date *string `json:"date,omitempty"`
		
		MilliTs *string `json:"milliTs,omitempty"`
		
		MicroTs *string `json:"microTs,omitempty"`
		
		Method *string `json:"method,omitempty"`
		
		ReplyReason *string `json:"replyReason,omitempty"`
		
		Ruri *string `json:"ruri,omitempty"`
		
		RuriUser *string `json:"ruriUser,omitempty"`
		
		RuriDomain *string `json:"ruriDomain,omitempty"`
		
		FromUser *string `json:"fromUser,omitempty"`
		
		FromDomain *string `json:"fromDomain,omitempty"`
		
		FromTag *string `json:"fromTag,omitempty"`
		
		ToUser *string `json:"toUser,omitempty"`
		
		ToDomain *string `json:"toDomain,omitempty"`
		
		ToTag *string `json:"toTag,omitempty"`
		
		PidUser *string `json:"pidUser,omitempty"`
		
		ContactUser *string `json:"contactUser,omitempty"`
		
		AuthUser *string `json:"authUser,omitempty"`
		
		Callid *string `json:"callid,omitempty"`
		
		CallidAleg *string `json:"callidAleg,omitempty"`
		
		Via1 *string `json:"via1,omitempty"`
		
		Via1Branch *string `json:"via1Branch,omitempty"`
		
		Cseq *string `json:"cseq,omitempty"`
		
		Diversion *string `json:"diversion,omitempty"`
		
		Reason *string `json:"reason,omitempty"`
		
		ContentType *string `json:"contentType,omitempty"`
		
		Auth *string `json:"auth,omitempty"`
		
		UserAgent *string `json:"userAgent,omitempty"`
		
		SourceIp *string `json:"sourceIp,omitempty"`
		
		SourcePort *string `json:"sourcePort,omitempty"`
		
		DestinationIp *string `json:"destinationIp,omitempty"`
		
		DestinationPort *string `json:"destinationPort,omitempty"`
		
		ContactIp *string `json:"contactIp,omitempty"`
		
		ContactPort *string `json:"contactPort,omitempty"`
		
		OriginatorIp *string `json:"originatorIp,omitempty"`
		
		OriginatorPort *string `json:"originatorPort,omitempty"`
		
		CorrelationId *string `json:"correlationId,omitempty"`
		
		Proto *string `json:"proto,omitempty"`
		
		Family *string `json:"family,omitempty"`
		
		RtpStat *string `json:"rtpStat,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Node *string `json:"node,omitempty"`
		
		Trans *string `json:"trans,omitempty"`
		
		Dbnode *string `json:"dbnode,omitempty"`
		
		Msg *string `json:"msg,omitempty"`
		
		SourceAlias *string `json:"sourceAlias,omitempty"`
		
		DestinationAlias *string `json:"destinationAlias,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		ParticipantId *string `json:"participantId,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Date: Date,
		
		MilliTs: u.MilliTs,
		
		MicroTs: u.MicroTs,
		
		Method: u.Method,
		
		ReplyReason: u.ReplyReason,
		
		Ruri: u.Ruri,
		
		RuriUser: u.RuriUser,
		
		RuriDomain: u.RuriDomain,
		
		FromUser: u.FromUser,
		
		FromDomain: u.FromDomain,
		
		FromTag: u.FromTag,
		
		ToUser: u.ToUser,
		
		ToDomain: u.ToDomain,
		
		ToTag: u.ToTag,
		
		PidUser: u.PidUser,
		
		ContactUser: u.ContactUser,
		
		AuthUser: u.AuthUser,
		
		Callid: u.Callid,
		
		CallidAleg: u.CallidAleg,
		
		Via1: u.Via1,
		
		Via1Branch: u.Via1Branch,
		
		Cseq: u.Cseq,
		
		Diversion: u.Diversion,
		
		Reason: u.Reason,
		
		ContentType: u.ContentType,
		
		Auth: u.Auth,
		
		UserAgent: u.UserAgent,
		
		SourceIp: u.SourceIp,
		
		SourcePort: u.SourcePort,
		
		DestinationIp: u.DestinationIp,
		
		DestinationPort: u.DestinationPort,
		
		ContactIp: u.ContactIp,
		
		ContactPort: u.ContactPort,
		
		OriginatorIp: u.OriginatorIp,
		
		OriginatorPort: u.OriginatorPort,
		
		CorrelationId: u.CorrelationId,
		
		Proto: u.Proto,
		
		Family: u.Family,
		
		RtpStat: u.RtpStat,
		
		VarType: u.VarType,
		
		Node: u.Node,
		
		Trans: u.Trans,
		
		Dbnode: u.Dbnode,
		
		Msg: u.Msg,
		
		SourceAlias: u.SourceAlias,
		
		DestinationAlias: u.DestinationAlias,
		
		ConversationId: u.ConversationId,
		
		ParticipantId: u.ParticipantId,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Homerrecord) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
