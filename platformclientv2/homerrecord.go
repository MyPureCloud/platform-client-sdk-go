package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Homerrecord
type Homerrecord struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Homerrecord) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Homerrecord) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "Date", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Homerrecord
	
	Date := new(string)
	if o.Date != nil {
		
		*Date = timeutil.Strftime(o.Date, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Date: Date,
		
		MilliTs: o.MilliTs,
		
		MicroTs: o.MicroTs,
		
		Method: o.Method,
		
		ReplyReason: o.ReplyReason,
		
		Ruri: o.Ruri,
		
		RuriUser: o.RuriUser,
		
		RuriDomain: o.RuriDomain,
		
		FromUser: o.FromUser,
		
		FromDomain: o.FromDomain,
		
		FromTag: o.FromTag,
		
		ToUser: o.ToUser,
		
		ToDomain: o.ToDomain,
		
		ToTag: o.ToTag,
		
		PidUser: o.PidUser,
		
		ContactUser: o.ContactUser,
		
		AuthUser: o.AuthUser,
		
		Callid: o.Callid,
		
		CallidAleg: o.CallidAleg,
		
		Via1: o.Via1,
		
		Via1Branch: o.Via1Branch,
		
		Cseq: o.Cseq,
		
		Diversion: o.Diversion,
		
		Reason: o.Reason,
		
		ContentType: o.ContentType,
		
		Auth: o.Auth,
		
		UserAgent: o.UserAgent,
		
		SourceIp: o.SourceIp,
		
		SourcePort: o.SourcePort,
		
		DestinationIp: o.DestinationIp,
		
		DestinationPort: o.DestinationPort,
		
		ContactIp: o.ContactIp,
		
		ContactPort: o.ContactPort,
		
		OriginatorIp: o.OriginatorIp,
		
		OriginatorPort: o.OriginatorPort,
		
		CorrelationId: o.CorrelationId,
		
		Proto: o.Proto,
		
		Family: o.Family,
		
		RtpStat: o.RtpStat,
		
		VarType: o.VarType,
		
		Node: o.Node,
		
		Trans: o.Trans,
		
		Dbnode: o.Dbnode,
		
		Msg: o.Msg,
		
		SourceAlias: o.SourceAlias,
		
		DestinationAlias: o.DestinationAlias,
		
		ConversationId: o.ConversationId,
		
		ParticipantId: o.ParticipantId,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Homerrecord) UnmarshalJSON(b []byte) error {
	var HomerrecordMap map[string]interface{}
	err := json.Unmarshal(b, &HomerrecordMap)
	if err != nil {
		return err
	}
	
	if Id, ok := HomerrecordMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := HomerrecordMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateString, ok := HomerrecordMap["date"].(string); ok {
		Date, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateString)
		o.Date = &Date
	}
	
	if MilliTs, ok := HomerrecordMap["milliTs"].(string); ok {
		o.MilliTs = &MilliTs
	}
    
	if MicroTs, ok := HomerrecordMap["microTs"].(string); ok {
		o.MicroTs = &MicroTs
	}
    
	if Method, ok := HomerrecordMap["method"].(string); ok {
		o.Method = &Method
	}
    
	if ReplyReason, ok := HomerrecordMap["replyReason"].(string); ok {
		o.ReplyReason = &ReplyReason
	}
    
	if Ruri, ok := HomerrecordMap["ruri"].(string); ok {
		o.Ruri = &Ruri
	}
    
	if RuriUser, ok := HomerrecordMap["ruriUser"].(string); ok {
		o.RuriUser = &RuriUser
	}
    
	if RuriDomain, ok := HomerrecordMap["ruriDomain"].(string); ok {
		o.RuriDomain = &RuriDomain
	}
    
	if FromUser, ok := HomerrecordMap["fromUser"].(string); ok {
		o.FromUser = &FromUser
	}
    
	if FromDomain, ok := HomerrecordMap["fromDomain"].(string); ok {
		o.FromDomain = &FromDomain
	}
    
	if FromTag, ok := HomerrecordMap["fromTag"].(string); ok {
		o.FromTag = &FromTag
	}
    
	if ToUser, ok := HomerrecordMap["toUser"].(string); ok {
		o.ToUser = &ToUser
	}
    
	if ToDomain, ok := HomerrecordMap["toDomain"].(string); ok {
		o.ToDomain = &ToDomain
	}
    
	if ToTag, ok := HomerrecordMap["toTag"].(string); ok {
		o.ToTag = &ToTag
	}
    
	if PidUser, ok := HomerrecordMap["pidUser"].(string); ok {
		o.PidUser = &PidUser
	}
    
	if ContactUser, ok := HomerrecordMap["contactUser"].(string); ok {
		o.ContactUser = &ContactUser
	}
    
	if AuthUser, ok := HomerrecordMap["authUser"].(string); ok {
		o.AuthUser = &AuthUser
	}
    
	if Callid, ok := HomerrecordMap["callid"].(string); ok {
		o.Callid = &Callid
	}
    
	if CallidAleg, ok := HomerrecordMap["callidAleg"].(string); ok {
		o.CallidAleg = &CallidAleg
	}
    
	if Via1, ok := HomerrecordMap["via1"].(string); ok {
		o.Via1 = &Via1
	}
    
	if Via1Branch, ok := HomerrecordMap["via1Branch"].(string); ok {
		o.Via1Branch = &Via1Branch
	}
    
	if Cseq, ok := HomerrecordMap["cseq"].(string); ok {
		o.Cseq = &Cseq
	}
    
	if Diversion, ok := HomerrecordMap["diversion"].(string); ok {
		o.Diversion = &Diversion
	}
    
	if Reason, ok := HomerrecordMap["reason"].(string); ok {
		o.Reason = &Reason
	}
    
	if ContentType, ok := HomerrecordMap["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if Auth, ok := HomerrecordMap["auth"].(string); ok {
		o.Auth = &Auth
	}
    
	if UserAgent, ok := HomerrecordMap["userAgent"].(string); ok {
		o.UserAgent = &UserAgent
	}
    
	if SourceIp, ok := HomerrecordMap["sourceIp"].(string); ok {
		o.SourceIp = &SourceIp
	}
    
	if SourcePort, ok := HomerrecordMap["sourcePort"].(string); ok {
		o.SourcePort = &SourcePort
	}
    
	if DestinationIp, ok := HomerrecordMap["destinationIp"].(string); ok {
		o.DestinationIp = &DestinationIp
	}
    
	if DestinationPort, ok := HomerrecordMap["destinationPort"].(string); ok {
		o.DestinationPort = &DestinationPort
	}
    
	if ContactIp, ok := HomerrecordMap["contactIp"].(string); ok {
		o.ContactIp = &ContactIp
	}
    
	if ContactPort, ok := HomerrecordMap["contactPort"].(string); ok {
		o.ContactPort = &ContactPort
	}
    
	if OriginatorIp, ok := HomerrecordMap["originatorIp"].(string); ok {
		o.OriginatorIp = &OriginatorIp
	}
    
	if OriginatorPort, ok := HomerrecordMap["originatorPort"].(string); ok {
		o.OriginatorPort = &OriginatorPort
	}
    
	if CorrelationId, ok := HomerrecordMap["correlationId"].(string); ok {
		o.CorrelationId = &CorrelationId
	}
    
	if Proto, ok := HomerrecordMap["proto"].(string); ok {
		o.Proto = &Proto
	}
    
	if Family, ok := HomerrecordMap["family"].(string); ok {
		o.Family = &Family
	}
    
	if RtpStat, ok := HomerrecordMap["rtpStat"].(string); ok {
		o.RtpStat = &RtpStat
	}
    
	if VarType, ok := HomerrecordMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Node, ok := HomerrecordMap["node"].(string); ok {
		o.Node = &Node
	}
    
	if Trans, ok := HomerrecordMap["trans"].(string); ok {
		o.Trans = &Trans
	}
    
	if Dbnode, ok := HomerrecordMap["dbnode"].(string); ok {
		o.Dbnode = &Dbnode
	}
    
	if Msg, ok := HomerrecordMap["msg"].(string); ok {
		o.Msg = &Msg
	}
    
	if SourceAlias, ok := HomerrecordMap["sourceAlias"].(string); ok {
		o.SourceAlias = &SourceAlias
	}
    
	if DestinationAlias, ok := HomerrecordMap["destinationAlias"].(string); ok {
		o.DestinationAlias = &DestinationAlias
	}
    
	if ConversationId, ok := HomerrecordMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if ParticipantId, ok := HomerrecordMap["participantId"].(string); ok {
		o.ParticipantId = &ParticipantId
	}
    
	if SelfUri, ok := HomerrecordMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Homerrecord) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
