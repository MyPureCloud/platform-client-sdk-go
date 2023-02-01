package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Callhistoryconversation
type Callhistoryconversation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Participants - The list of participants involved in the conversation.
	Participants *[]Callhistoryparticipant `json:"participants,omitempty"`

	// Direction - The direction of the call relating to the current user
	Direction *string `json:"direction,omitempty"`

	// WentToVoicemail - Did the call end in the current user's voicemail
	WentToVoicemail *bool `json:"wentToVoicemail,omitempty"`

	// MissedCall - Did the user not answer this conversation
	MissedCall *bool `json:"missedCall,omitempty"`

	// StartTime - The time the user joined the conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartTime *time.Time `json:"startTime,omitempty"`

	// WasConference - Was this conversation a conference
	WasConference *bool `json:"wasConference,omitempty"`

	// WasCallback - Was this conversation a callback
	WasCallback *bool `json:"wasCallback,omitempty"`

	// HadScreenShare - Did this conversation have a screen share session
	HadScreenShare *bool `json:"hadScreenShare,omitempty"`

	// HadCobrowse - Did this conversation have a cobrowse session
	HadCobrowse *bool `json:"hadCobrowse,omitempty"`

	// WasOutboundCampaign - Was this conversation associated with an outbound campaign
	WasOutboundCampaign *bool `json:"wasOutboundCampaign,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Callhistoryconversation) SetField(field string, fieldValue interface{}) {
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

func (o Callhistoryconversation) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartTime", }
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
	type Alias Callhistoryconversation
	
	StartTime := new(string)
	if o.StartTime != nil {
		
		*StartTime = timeutil.Strftime(o.StartTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartTime = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Participants *[]Callhistoryparticipant `json:"participants,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		WentToVoicemail *bool `json:"wentToVoicemail,omitempty"`
		
		MissedCall *bool `json:"missedCall,omitempty"`
		
		StartTime *string `json:"startTime,omitempty"`
		
		WasConference *bool `json:"wasConference,omitempty"`
		
		WasCallback *bool `json:"wasCallback,omitempty"`
		
		HadScreenShare *bool `json:"hadScreenShare,omitempty"`
		
		HadCobrowse *bool `json:"hadCobrowse,omitempty"`
		
		WasOutboundCampaign *bool `json:"wasOutboundCampaign,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Participants: o.Participants,
		
		Direction: o.Direction,
		
		WentToVoicemail: o.WentToVoicemail,
		
		MissedCall: o.MissedCall,
		
		StartTime: StartTime,
		
		WasConference: o.WasConference,
		
		WasCallback: o.WasCallback,
		
		HadScreenShare: o.HadScreenShare,
		
		HadCobrowse: o.HadCobrowse,
		
		WasOutboundCampaign: o.WasOutboundCampaign,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Callhistoryconversation) UnmarshalJSON(b []byte) error {
	var CallhistoryconversationMap map[string]interface{}
	err := json.Unmarshal(b, &CallhistoryconversationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CallhistoryconversationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CallhistoryconversationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Participants, ok := CallhistoryconversationMap["participants"].([]interface{}); ok {
		ParticipantsString, _ := json.Marshal(Participants)
		json.Unmarshal(ParticipantsString, &o.Participants)
	}
	
	if Direction, ok := CallhistoryconversationMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if WentToVoicemail, ok := CallhistoryconversationMap["wentToVoicemail"].(bool); ok {
		o.WentToVoicemail = &WentToVoicemail
	}
    
	if MissedCall, ok := CallhistoryconversationMap["missedCall"].(bool); ok {
		o.MissedCall = &MissedCall
	}
    
	if startTimeString, ok := CallhistoryconversationMap["startTime"].(string); ok {
		StartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startTimeString)
		o.StartTime = &StartTime
	}
	
	if WasConference, ok := CallhistoryconversationMap["wasConference"].(bool); ok {
		o.WasConference = &WasConference
	}
    
	if WasCallback, ok := CallhistoryconversationMap["wasCallback"].(bool); ok {
		o.WasCallback = &WasCallback
	}
    
	if HadScreenShare, ok := CallhistoryconversationMap["hadScreenShare"].(bool); ok {
		o.HadScreenShare = &HadScreenShare
	}
    
	if HadCobrowse, ok := CallhistoryconversationMap["hadCobrowse"].(bool); ok {
		o.HadCobrowse = &HadCobrowse
	}
    
	if WasOutboundCampaign, ok := CallhistoryconversationMap["wasOutboundCampaign"].(bool); ok {
		o.WasOutboundCampaign = &WasOutboundCampaign
	}
    
	if SelfUri, ok := CallhistoryconversationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Callhistoryconversation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
