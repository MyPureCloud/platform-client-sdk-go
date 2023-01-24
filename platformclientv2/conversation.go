package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversation
type Conversation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// ExternalTag - The external tag associated with the conversation.
	ExternalTag *string `json:"externalTag,omitempty"`

	// StartTime - The time when the conversation started. This will be the time when the first participant joined the conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartTime *time.Time `json:"startTime,omitempty"`

	// EndTime - The time when the conversation ended. This will be the time when the last participant left the conversation, or null when the conversation is still active. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndTime *time.Time `json:"endTime,omitempty"`

	// Address - The address of the conversation as seen from an external participant. For phone calls this will be the DNIS for inbound calls and the ANI for outbound calls. For other media types this will be the address of the destination participant for inbound and the address of the initiating participant for outbound.
	Address *string `json:"address,omitempty"`

	// Participants - The list of all participants in the conversation.
	Participants *[]Participant `json:"participants,omitempty"`

	// ConversationIds - A list of conversations to merge into this conversation to create a conference. This field is null except when being used to create a conference.
	ConversationIds *[]string `json:"conversationIds,omitempty"`

	// MaxParticipants - If this is a conference conversation, then this field indicates the maximum number of participants allowed to participant in the conference.
	MaxParticipants *int `json:"maxParticipants,omitempty"`

	// RecordingState - On update, 'paused' initiates a secure pause, 'active' resumes any paused recordings; otherwise indicates state of conversation recording.
	RecordingState *string `json:"recordingState,omitempty"`

	// State - The conversation's state
	State *string `json:"state,omitempty"`

	// Divisions - Identifiers of divisions associated with this conversation
	Divisions *[]Conversationdivisionmembership `json:"divisions,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Conversation) SetField(field string, fieldValue interface{}) {
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

func (o Conversation) MarshalJSON() ([]byte, error) {
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
	type Alias Conversation
	
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
		
		ExternalTag *string `json:"externalTag,omitempty"`
		
		StartTime *string `json:"startTime,omitempty"`
		
		EndTime *string `json:"endTime,omitempty"`
		
		Address *string `json:"address,omitempty"`
		
		Participants *[]Participant `json:"participants,omitempty"`
		
		ConversationIds *[]string `json:"conversationIds,omitempty"`
		
		MaxParticipants *int `json:"maxParticipants,omitempty"`
		
		RecordingState *string `json:"recordingState,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Divisions *[]Conversationdivisionmembership `json:"divisions,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ExternalTag: o.ExternalTag,
		
		StartTime: StartTime,
		
		EndTime: EndTime,
		
		Address: o.Address,
		
		Participants: o.Participants,
		
		ConversationIds: o.ConversationIds,
		
		MaxParticipants: o.MaxParticipants,
		
		RecordingState: o.RecordingState,
		
		State: o.State,
		
		Divisions: o.Divisions,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Conversation) UnmarshalJSON(b []byte) error {
	var ConversationMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ConversationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ConversationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ExternalTag, ok := ConversationMap["externalTag"].(string); ok {
		o.ExternalTag = &ExternalTag
	}
    
	if startTimeString, ok := ConversationMap["startTime"].(string); ok {
		StartTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startTimeString)
		o.StartTime = &StartTime
	}
	
	if endTimeString, ok := ConversationMap["endTime"].(string); ok {
		EndTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endTimeString)
		o.EndTime = &EndTime
	}
	
	if Address, ok := ConversationMap["address"].(string); ok {
		o.Address = &Address
	}
    
	if Participants, ok := ConversationMap["participants"].([]interface{}); ok {
		ParticipantsString, _ := json.Marshal(Participants)
		json.Unmarshal(ParticipantsString, &o.Participants)
	}
	
	if ConversationIds, ok := ConversationMap["conversationIds"].([]interface{}); ok {
		ConversationIdsString, _ := json.Marshal(ConversationIds)
		json.Unmarshal(ConversationIdsString, &o.ConversationIds)
	}
	
	if MaxParticipants, ok := ConversationMap["maxParticipants"].(float64); ok {
		MaxParticipantsInt := int(MaxParticipants)
		o.MaxParticipants = &MaxParticipantsInt
	}
	
	if RecordingState, ok := ConversationMap["recordingState"].(string); ok {
		o.RecordingState = &RecordingState
	}
    
	if State, ok := ConversationMap["state"].(string); ok {
		o.State = &State
	}
    
	if Divisions, ok := ConversationMap["divisions"].([]interface{}); ok {
		DivisionsString, _ := json.Marshal(Divisions)
		json.Unmarshal(DivisionsString, &o.Divisions)
	}
	
	if SelfUri, ok := ConversationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
