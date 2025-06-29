package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2queueobservationobservation
type V2queueobservationobservation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ObservationDate - The timestamp when the observation started
	ObservationDate *time.Time `json:"observationDate,omitempty"`

	// ConversationId - Unique identifier of the conversation to which this observation belongs
	ConversationId *string `json:"conversationId,omitempty"`

	// SessionId - Unique identifier of the user session associated with this observation
	SessionId *string `json:"sessionId,omitempty"`

	// RequestedRoutingSkillIds - Unique identifiers for skills requested for an interaction
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

	// TeamId - The team Id the user is a member of
	TeamId *string `json:"teamId,omitempty"`

	// ScoredAgents - Scored agents for this conversation
	ScoredAgents *[]V2queueobservationscoredagent `json:"scoredAgents,omitempty"`

	// RequestedRoutings - All routing types for requested/attempted routing methods.
	RequestedRoutings *[]string `json:"requestedRoutings,omitempty"`

	// UsedRouting - Complete routing method
	UsedRouting *string `json:"usedRouting,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2queueobservationobservation) SetField(field string, fieldValue interface{}) {
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

func (o V2queueobservationobservation) MarshalJSON() ([]byte, error) {
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
	type Alias V2queueobservationobservation
	
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
		
		ScoredAgents *[]V2queueobservationscoredagent `json:"scoredAgents,omitempty"`
		
		RequestedRoutings *[]string `json:"requestedRoutings,omitempty"`
		
		UsedRouting *string `json:"usedRouting,omitempty"`
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
		
		ScoredAgents: o.ScoredAgents,
		
		RequestedRoutings: o.RequestedRoutings,
		
		UsedRouting: o.UsedRouting,
		Alias:    (Alias)(o),
	})
}

func (o *V2queueobservationobservation) UnmarshalJSON(b []byte) error {
	var V2queueobservationobservationMap map[string]interface{}
	err := json.Unmarshal(b, &V2queueobservationobservationMap)
	if err != nil {
		return err
	}
	
	if observationDateString, ok := V2queueobservationobservationMap["observationDate"].(string); ok {
		ObservationDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", observationDateString)
		o.ObservationDate = &ObservationDate
	}
	
	if ConversationId, ok := V2queueobservationobservationMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if SessionId, ok := V2queueobservationobservationMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if RequestedRoutingSkillIds, ok := V2queueobservationobservationMap["requestedRoutingSkillIds"].([]interface{}); ok {
		RequestedRoutingSkillIdsString, _ := json.Marshal(RequestedRoutingSkillIds)
		json.Unmarshal(RequestedRoutingSkillIdsString, &o.RequestedRoutingSkillIds)
	}
	
	if RequestedLanguageId, ok := V2queueobservationobservationMap["requestedLanguageId"].(string); ok {
		o.RequestedLanguageId = &RequestedLanguageId
	}
    
	if RoutingPriority, ok := V2queueobservationobservationMap["routingPriority"].(float64); ok {
		RoutingPriorityInt := int(RoutingPriority)
		o.RoutingPriority = &RoutingPriorityInt
	}
	
	if ParticipantName, ok := V2queueobservationobservationMap["participantName"].(string); ok {
		o.ParticipantName = &ParticipantName
	}
    
	if UserId, ok := V2queueobservationobservationMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if Direction, ok := V2queueobservationobservationMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if ConvertedFrom, ok := V2queueobservationobservationMap["convertedFrom"].(string); ok {
		o.ConvertedFrom = &ConvertedFrom
	}
    
	if ConvertedTo, ok := V2queueobservationobservationMap["convertedTo"].(string); ok {
		o.ConvertedTo = &ConvertedTo
	}
    
	if AddressFrom, ok := V2queueobservationobservationMap["addressFrom"].(string); ok {
		o.AddressFrom = &AddressFrom
	}
    
	if AddressTo, ok := V2queueobservationobservationMap["addressTo"].(string); ok {
		o.AddressTo = &AddressTo
	}
    
	if Ani, ok := V2queueobservationobservationMap["ani"].(string); ok {
		o.Ani = &Ani
	}
    
	if Dnis, ok := V2queueobservationobservationMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
    
	if TeamId, ok := V2queueobservationobservationMap["teamId"].(string); ok {
		o.TeamId = &TeamId
	}
    
	if ScoredAgents, ok := V2queueobservationobservationMap["scoredAgents"].([]interface{}); ok {
		ScoredAgentsString, _ := json.Marshal(ScoredAgents)
		json.Unmarshal(ScoredAgentsString, &o.ScoredAgents)
	}
	
	if RequestedRoutings, ok := V2queueobservationobservationMap["requestedRoutings"].([]interface{}); ok {
		RequestedRoutingsString, _ := json.Marshal(RequestedRoutings)
		json.Unmarshal(RequestedRoutingsString, &o.RequestedRoutings)
	}
	
	if UsedRouting, ok := V2queueobservationobservationMap["usedRouting"].(string); ok {
		o.UsedRouting = &UsedRouting
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V2queueobservationobservation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
