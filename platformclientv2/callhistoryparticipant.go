package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Callhistoryparticipant
type Callhistoryparticipant struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The unique participant ID.
	Id *string `json:"id,omitempty"`

	// Name - The display friendly name of the participant.
	Name *string `json:"name,omitempty"`

	// Address - The participant address.
	Address *string `json:"address,omitempty"`

	// StartTime - The time when this participant first joined the conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartTime *time.Time `json:"startTime,omitempty"`

	// EndTime - The time when this participant went disconnected for this media (eg: video disconnected time). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndTime *time.Time `json:"endTime,omitempty"`

	// Purpose - The participant's purpose.  Values can be: 'agent', 'user', 'customer', 'external', 'acd', 'ivr
	Purpose *string `json:"purpose,omitempty"`

	// Direction - The participant's direction.  Values can be: 'inbound' or 'outbound'
	Direction *string `json:"direction,omitempty"`

	// Ani - The call ANI.
	Ani *string `json:"ani,omitempty"`

	// Dnis - The call DNIS.
	Dnis *string `json:"dnis,omitempty"`

	// User - The PureCloud user for this participant.
	User *User `json:"user,omitempty"`

	// Queue - The PureCloud queue for this participant.
	Queue *Queue `json:"queue,omitempty"`

	// Group - The group involved in the group ring call.
	Group *Group `json:"group,omitempty"`

	// DisconnectType - The reason the participant was disconnected from the conversation.
	DisconnectType *string `json:"disconnectType,omitempty"`

	// ExternalContact - The PureCloud external contact
	ExternalContact *Externalcontact `json:"externalContact,omitempty"`

	// ExternalOrganization - The PureCloud external organization
	ExternalOrganization *Externalorganization `json:"externalOrganization,omitempty"`

	// DidInteract - Indicates whether the contact ever connected
	DidInteract *bool `json:"didInteract,omitempty"`

	// SipResponseCodes - Indicates SIP Response codes associated with the participant
	SipResponseCodes *[]int `json:"sipResponseCodes,omitempty"`

	// FlaggedReason - The reason specifying why participant flagged the conversation.
	FlaggedReason *string `json:"flaggedReason,omitempty"`

	// OutboundCampaign - The outbound campaign associated with the participant
	OutboundCampaign *Campaign `json:"outboundCampaign,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Callhistoryparticipant) SetField(field string, fieldValue interface{}) {
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

func (o Callhistoryparticipant) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartTime","EndTime", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if contains(dateTimeFields, fieldName) {
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
	type Alias Callhistoryparticipant
	
	StartTime := new(string)
	if o.StartTime != nil {
		
		*StartTime = timeutil.Strftime(o.StartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartTime = nil
	}
	
	EndTime := new(string)
	if o.EndTime != nil {
		
		*EndTime = timeutil.Strftime(o.EndTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndTime = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Address *string `json:"address,omitempty"`
		
		StartTime *string `json:"startTime,omitempty"`
		
		EndTime *string `json:"endTime,omitempty"`
		
		Purpose *string `json:"purpose,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		Ani *string `json:"ani,omitempty"`
		
		Dnis *string `json:"dnis,omitempty"`
		
		User *User `json:"user,omitempty"`
		
		Queue *Queue `json:"queue,omitempty"`
		
		Group *Group `json:"group,omitempty"`
		
		DisconnectType *string `json:"disconnectType,omitempty"`
		
		ExternalContact *Externalcontact `json:"externalContact,omitempty"`
		
		ExternalOrganization *Externalorganization `json:"externalOrganization,omitempty"`
		
		DidInteract *bool `json:"didInteract,omitempty"`
		
		SipResponseCodes *[]int `json:"sipResponseCodes,omitempty"`
		
		FlaggedReason *string `json:"flaggedReason,omitempty"`
		
		OutboundCampaign *Campaign `json:"outboundCampaign,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Address: o.Address,
		
		StartTime: StartTime,
		
		EndTime: EndTime,
		
		Purpose: o.Purpose,
		
		Direction: o.Direction,
		
		Ani: o.Ani,
		
		Dnis: o.Dnis,
		
		User: o.User,
		
		Queue: o.Queue,
		
		Group: o.Group,
		
		DisconnectType: o.DisconnectType,
		
		ExternalContact: o.ExternalContact,
		
		ExternalOrganization: o.ExternalOrganization,
		
		DidInteract: o.DidInteract,
		
		SipResponseCodes: o.SipResponseCodes,
		
		FlaggedReason: o.FlaggedReason,
		
		OutboundCampaign: o.OutboundCampaign,
		Alias:    (Alias)(o),
	})
}

func (o *Callhistoryparticipant) UnmarshalJSON(b []byte) error {
	var CallhistoryparticipantMap map[string]interface{}
	err := json.Unmarshal(b, &CallhistoryparticipantMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CallhistoryparticipantMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CallhistoryparticipantMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Address, ok := CallhistoryparticipantMap["address"].(string); ok {
		o.Address = &Address
	}
    
	if startTimeString, ok := CallhistoryparticipantMap["startTime"].(string); ok {
		StartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startTimeString)
		o.StartTime = &StartTime
	}
	
	if endTimeString, ok := CallhistoryparticipantMap["endTime"].(string); ok {
		EndTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endTimeString)
		o.EndTime = &EndTime
	}
	
	if Purpose, ok := CallhistoryparticipantMap["purpose"].(string); ok {
		o.Purpose = &Purpose
	}
    
	if Direction, ok := CallhistoryparticipantMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if Ani, ok := CallhistoryparticipantMap["ani"].(string); ok {
		o.Ani = &Ani
	}
    
	if Dnis, ok := CallhistoryparticipantMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
    
	if User, ok := CallhistoryparticipantMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Queue, ok := CallhistoryparticipantMap["queue"].(map[string]interface{}); ok {
		QueueString, _ := json.Marshal(Queue)
		json.Unmarshal(QueueString, &o.Queue)
	}
	
	if Group, ok := CallhistoryparticipantMap["group"].(map[string]interface{}); ok {
		GroupString, _ := json.Marshal(Group)
		json.Unmarshal(GroupString, &o.Group)
	}
	
	if DisconnectType, ok := CallhistoryparticipantMap["disconnectType"].(string); ok {
		o.DisconnectType = &DisconnectType
	}
    
	if ExternalContact, ok := CallhistoryparticipantMap["externalContact"].(map[string]interface{}); ok {
		ExternalContactString, _ := json.Marshal(ExternalContact)
		json.Unmarshal(ExternalContactString, &o.ExternalContact)
	}
	
	if ExternalOrganization, ok := CallhistoryparticipantMap["externalOrganization"].(map[string]interface{}); ok {
		ExternalOrganizationString, _ := json.Marshal(ExternalOrganization)
		json.Unmarshal(ExternalOrganizationString, &o.ExternalOrganization)
	}
	
	if DidInteract, ok := CallhistoryparticipantMap["didInteract"].(bool); ok {
		o.DidInteract = &DidInteract
	}
    
	if SipResponseCodes, ok := CallhistoryparticipantMap["sipResponseCodes"].([]interface{}); ok {
		SipResponseCodesString, _ := json.Marshal(SipResponseCodes)
		json.Unmarshal(SipResponseCodesString, &o.SipResponseCodes)
	}
	
	if FlaggedReason, ok := CallhistoryparticipantMap["flaggedReason"].(string); ok {
		o.FlaggedReason = &FlaggedReason
	}
    
	if OutboundCampaign, ok := CallhistoryparticipantMap["outboundCampaign"].(map[string]interface{}); ok {
		OutboundCampaignString, _ := json.Marshal(OutboundCampaign)
		json.Unmarshal(OutboundCampaignString, &o.OutboundCampaign)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Callhistoryparticipant) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
