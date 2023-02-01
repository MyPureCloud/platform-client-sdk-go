package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Observationvalue
type Observationvalue struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ObservationDate - The time at which the observation occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ObservationDate *time.Time `json:"observationDate,omitempty"`

	// ConversationId - Unique identifier for the conversation
	ConversationId *string `json:"conversationId,omitempty"`

	// SessionId - The unique identifier of this session
	SessionId *string `json:"sessionId,omitempty"`

	// RequestedRoutingSkillIds - Unique identifier for a skill requested for an interaction
	RequestedRoutingSkillIds *[]string `json:"requestedRoutingSkillIds,omitempty"`

	// RequestedLanguageId - Unique identifier for the language requested for an interaction
	RequestedLanguageId *string `json:"requestedLanguageId,omitempty"`

	// RoutingPriority - Routing priority for the current interaction
	RoutingPriority *int `json:"routingPriority,omitempty"`

	// ParticipantName - A human readable name identifying the participant
	ParticipantName *string `json:"participantName,omitempty"`

	// UserId - Unique identifier for the user
	UserId *string `json:"userId,omitempty"`

	// Direction - The direction of the communication
	Direction *string `json:"direction,omitempty"`

	// ConvertedFrom - Session media type that was converted from in case of a media type conversion
	ConvertedFrom *string `json:"convertedFrom,omitempty"`

	// ConvertedTo - Session media type that was converted to in case of a media type conversion
	ConvertedTo *string `json:"convertedTo,omitempty"`

	// AddressFrom - The address that initiated an action
	AddressFrom *string `json:"addressFrom,omitempty"`

	// AddressTo - The address receiving an action
	AddressTo *string `json:"addressTo,omitempty"`

	// Ani - Automatic Number Identification (caller's number)
	Ani *string `json:"ani,omitempty"`

	// Dnis - Dialed number identification service (number dialed by the calling party)
	Dnis *string `json:"dnis,omitempty"`

	// TeamId - The team id the user is a member of
	TeamId *string `json:"teamId,omitempty"`

	// RequestedRoutings - All routing types for requested/attempted routing methods
	RequestedRoutings *[]string `json:"requestedRoutings,omitempty"`

	// UsedRouting - Complete routing method
	UsedRouting *string `json:"usedRouting,omitempty"`

	// ScoredAgents
	ScoredAgents *[]Analyticsscoredagent `json:"scoredAgents,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Observationvalue) SetField(field string, fieldValue interface{}) {
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

func (o Observationvalue) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ObservationDate", }
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
	type Alias Observationvalue
	
	ObservationDate := new(string)
	if o.ObservationDate != nil {
		
		*ObservationDate = timeutil.Strftime(o.ObservationDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ObservationDate = nil
	}
	
	return json.Marshal(&struct { 
		ObservationDate *string `json:"observationDate,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		RequestedRoutingSkillIds *[]string `json:"requestedRoutingSkillIds,omitempty"`
		
		RequestedLanguageId *string `json:"requestedLanguageId,omitempty"`
		
		RoutingPriority *int `json:"routingPriority,omitempty"`
		
		ParticipantName *string `json:"participantName,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		ConvertedFrom *string `json:"convertedFrom,omitempty"`
		
		ConvertedTo *string `json:"convertedTo,omitempty"`
		
		AddressFrom *string `json:"addressFrom,omitempty"`
		
		AddressTo *string `json:"addressTo,omitempty"`
		
		Ani *string `json:"ani,omitempty"`
		
		Dnis *string `json:"dnis,omitempty"`
		
		TeamId *string `json:"teamId,omitempty"`
		
		RequestedRoutings *[]string `json:"requestedRoutings,omitempty"`
		
		UsedRouting *string `json:"usedRouting,omitempty"`
		
		ScoredAgents *[]Analyticsscoredagent `json:"scoredAgents,omitempty"`
		Alias
	}{ 
		ObservationDate: ObservationDate,
		
		ConversationId: o.ConversationId,
		
		SessionId: o.SessionId,
		
		RequestedRoutingSkillIds: o.RequestedRoutingSkillIds,
		
		RequestedLanguageId: o.RequestedLanguageId,
		
		RoutingPriority: o.RoutingPriority,
		
		ParticipantName: o.ParticipantName,
		
		UserId: o.UserId,
		
		Direction: o.Direction,
		
		ConvertedFrom: o.ConvertedFrom,
		
		ConvertedTo: o.ConvertedTo,
		
		AddressFrom: o.AddressFrom,
		
		AddressTo: o.AddressTo,
		
		Ani: o.Ani,
		
		Dnis: o.Dnis,
		
		TeamId: o.TeamId,
		
		RequestedRoutings: o.RequestedRoutings,
		
		UsedRouting: o.UsedRouting,
		
		ScoredAgents: o.ScoredAgents,
		Alias:    (Alias)(o),
	})
}

func (o *Observationvalue) UnmarshalJSON(b []byte) error {
	var ObservationvalueMap map[string]interface{}
	err := json.Unmarshal(b, &ObservationvalueMap)
	if err != nil {
		return err
	}
	
	if observationDateString, ok := ObservationvalueMap["observationDate"].(string); ok {
		ObservationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", observationDateString)
		o.ObservationDate = &ObservationDate
	}
	
	if ConversationId, ok := ObservationvalueMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if SessionId, ok := ObservationvalueMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if RequestedRoutingSkillIds, ok := ObservationvalueMap["requestedRoutingSkillIds"].([]interface{}); ok {
		RequestedRoutingSkillIdsString, _ := json.Marshal(RequestedRoutingSkillIds)
		json.Unmarshal(RequestedRoutingSkillIdsString, &o.RequestedRoutingSkillIds)
	}
	
	if RequestedLanguageId, ok := ObservationvalueMap["requestedLanguageId"].(string); ok {
		o.RequestedLanguageId = &RequestedLanguageId
	}
    
	if RoutingPriority, ok := ObservationvalueMap["routingPriority"].(float64); ok {
		RoutingPriorityInt := int(RoutingPriority)
		o.RoutingPriority = &RoutingPriorityInt
	}
	
	if ParticipantName, ok := ObservationvalueMap["participantName"].(string); ok {
		o.ParticipantName = &ParticipantName
	}
    
	if UserId, ok := ObservationvalueMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if Direction, ok := ObservationvalueMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if ConvertedFrom, ok := ObservationvalueMap["convertedFrom"].(string); ok {
		o.ConvertedFrom = &ConvertedFrom
	}
    
	if ConvertedTo, ok := ObservationvalueMap["convertedTo"].(string); ok {
		o.ConvertedTo = &ConvertedTo
	}
    
	if AddressFrom, ok := ObservationvalueMap["addressFrom"].(string); ok {
		o.AddressFrom = &AddressFrom
	}
    
	if AddressTo, ok := ObservationvalueMap["addressTo"].(string); ok {
		o.AddressTo = &AddressTo
	}
    
	if Ani, ok := ObservationvalueMap["ani"].(string); ok {
		o.Ani = &Ani
	}
    
	if Dnis, ok := ObservationvalueMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
    
	if TeamId, ok := ObservationvalueMap["teamId"].(string); ok {
		o.TeamId = &TeamId
	}
    
	if RequestedRoutings, ok := ObservationvalueMap["requestedRoutings"].([]interface{}); ok {
		RequestedRoutingsString, _ := json.Marshal(RequestedRoutings)
		json.Unmarshal(RequestedRoutingsString, &o.RequestedRoutings)
	}
	
	if UsedRouting, ok := ObservationvalueMap["usedRouting"].(string); ok {
		o.UsedRouting = &UsedRouting
	}
    
	if ScoredAgents, ok := ObservationvalueMap["scoredAgents"].([]interface{}); ok {
		ScoredAgentsString, _ := json.Marshal(ScoredAgents)
		json.Unmarshal(ScoredAgentsString, &o.ScoredAgents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Observationvalue) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
