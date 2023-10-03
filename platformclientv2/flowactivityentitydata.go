package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowactivityentitydata
type Flowactivityentitydata struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ActivityDate - The time at which the activity was observed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ActivityDate *time.Time `json:"activityDate,omitempty"`

	// Metric - Activity metric
	Metric *string `json:"metric,omitempty"`

	// ActiveRouting - Active routing method
	ActiveRouting *string `json:"activeRouting,omitempty"`

	// AddressFrom - The address that initiated an action
	AddressFrom *string `json:"addressFrom,omitempty"`

	// AddressTo - The address receiving an action
	AddressTo *string `json:"addressTo,omitempty"`

	// Ani - Automatic Number Identification (caller's number)
	Ani *string `json:"ani,omitempty"`

	// ConversationId - Unique identifier for the conversation
	ConversationId *string `json:"conversationId,omitempty"`

	// ConvertedFrom - Session media type that was converted from in case of a media type conversion
	ConvertedFrom *string `json:"convertedFrom,omitempty"`

	// ConvertedTo - Session media type that was converted to in case of a media type conversion
	ConvertedTo *string `json:"convertedTo,omitempty"`

	// Direction - The direction of the communication
	Direction *string `json:"direction,omitempty"`

	// Dnis - Dialed number identification service (number dialed by the calling party)
	Dnis *string `json:"dnis,omitempty"`

	// FlowId - The unique identifier of this flow
	FlowId *string `json:"flowId,omitempty"`

	// FlowType - The type of this flow
	FlowType *string `json:"flowType,omitempty"`

	// MediaType - The session media type
	MediaType *string `json:"mediaType,omitempty"`

	// ParticipantName - A human readable name identifying the participant
	ParticipantName *string `json:"participantName,omitempty"`

	// QueueId - Queue identifier
	QueueId *string `json:"queueId,omitempty"`

	// RequestedLanguageId - Unique identifier for the language requested for an interaction
	RequestedLanguageId *string `json:"requestedLanguageId,omitempty"`

	// RequestedRoutingSkillIds - Unique identifier(s) for skill(s) requested for an interaction
	RequestedRoutingSkillIds *[]string `json:"requestedRoutingSkillIds,omitempty"`

	// RequestedRoutings - Routing type(s) for requested/attempted routing methods.
	RequestedRoutings *[]string `json:"requestedRoutings,omitempty"`

	// RoutingPriority - Routing priority for the current interaction
	RoutingPriority *int `json:"routingPriority,omitempty"`

	// SessionId - The unique identifier of this session
	SessionId *string `json:"sessionId,omitempty"`

	// TeamId - The team ID the user is a member of
	TeamId *string `json:"teamId,omitempty"`

	// UsedRouting - Complete routing method
	UsedRouting *string `json:"usedRouting,omitempty"`

	// UserId - Unique identifier for the user
	UserId *string `json:"userId,omitempty"`

	// ScoredAgents - Scored agents
	ScoredAgents *[]Flowactivityscoredagent `json:"scoredAgents,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Flowactivityentitydata) SetField(field string, fieldValue interface{}) {
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

func (o Flowactivityentitydata) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ActivityDate", }
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
	type Alias Flowactivityentitydata
	
	ActivityDate := new(string)
	if o.ActivityDate != nil {
		
		*ActivityDate = timeutil.Strftime(o.ActivityDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ActivityDate = nil
	}
	
	return json.Marshal(&struct { 
		ActivityDate *string `json:"activityDate,omitempty"`
		
		Metric *string `json:"metric,omitempty"`
		
		ActiveRouting *string `json:"activeRouting,omitempty"`
		
		AddressFrom *string `json:"addressFrom,omitempty"`
		
		AddressTo *string `json:"addressTo,omitempty"`
		
		Ani *string `json:"ani,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		ConvertedFrom *string `json:"convertedFrom,omitempty"`
		
		ConvertedTo *string `json:"convertedTo,omitempty"`
		
		Direction *string `json:"direction,omitempty"`
		
		Dnis *string `json:"dnis,omitempty"`
		
		FlowId *string `json:"flowId,omitempty"`
		
		FlowType *string `json:"flowType,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		ParticipantName *string `json:"participantName,omitempty"`
		
		QueueId *string `json:"queueId,omitempty"`
		
		RequestedLanguageId *string `json:"requestedLanguageId,omitempty"`
		
		RequestedRoutingSkillIds *[]string `json:"requestedRoutingSkillIds,omitempty"`
		
		RequestedRoutings *[]string `json:"requestedRoutings,omitempty"`
		
		RoutingPriority *int `json:"routingPriority,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		TeamId *string `json:"teamId,omitempty"`
		
		UsedRouting *string `json:"usedRouting,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		ScoredAgents *[]Flowactivityscoredagent `json:"scoredAgents,omitempty"`
		Alias
	}{ 
		ActivityDate: ActivityDate,
		
		Metric: o.Metric,
		
		ActiveRouting: o.ActiveRouting,
		
		AddressFrom: o.AddressFrom,
		
		AddressTo: o.AddressTo,
		
		Ani: o.Ani,
		
		ConversationId: o.ConversationId,
		
		ConvertedFrom: o.ConvertedFrom,
		
		ConvertedTo: o.ConvertedTo,
		
		Direction: o.Direction,
		
		Dnis: o.Dnis,
		
		FlowId: o.FlowId,
		
		FlowType: o.FlowType,
		
		MediaType: o.MediaType,
		
		ParticipantName: o.ParticipantName,
		
		QueueId: o.QueueId,
		
		RequestedLanguageId: o.RequestedLanguageId,
		
		RequestedRoutingSkillIds: o.RequestedRoutingSkillIds,
		
		RequestedRoutings: o.RequestedRoutings,
		
		RoutingPriority: o.RoutingPriority,
		
		SessionId: o.SessionId,
		
		TeamId: o.TeamId,
		
		UsedRouting: o.UsedRouting,
		
		UserId: o.UserId,
		
		ScoredAgents: o.ScoredAgents,
		Alias:    (Alias)(o),
	})
}

func (o *Flowactivityentitydata) UnmarshalJSON(b []byte) error {
	var FlowactivityentitydataMap map[string]interface{}
	err := json.Unmarshal(b, &FlowactivityentitydataMap)
	if err != nil {
		return err
	}
	
	if activityDateString, ok := FlowactivityentitydataMap["activityDate"].(string); ok {
		ActivityDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", activityDateString)
		o.ActivityDate = &ActivityDate
	}
	
	if Metric, ok := FlowactivityentitydataMap["metric"].(string); ok {
		o.Metric = &Metric
	}
    
	if ActiveRouting, ok := FlowactivityentitydataMap["activeRouting"].(string); ok {
		o.ActiveRouting = &ActiveRouting
	}
    
	if AddressFrom, ok := FlowactivityentitydataMap["addressFrom"].(string); ok {
		o.AddressFrom = &AddressFrom
	}
    
	if AddressTo, ok := FlowactivityentitydataMap["addressTo"].(string); ok {
		o.AddressTo = &AddressTo
	}
    
	if Ani, ok := FlowactivityentitydataMap["ani"].(string); ok {
		o.Ani = &Ani
	}
    
	if ConversationId, ok := FlowactivityentitydataMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if ConvertedFrom, ok := FlowactivityentitydataMap["convertedFrom"].(string); ok {
		o.ConvertedFrom = &ConvertedFrom
	}
    
	if ConvertedTo, ok := FlowactivityentitydataMap["convertedTo"].(string); ok {
		o.ConvertedTo = &ConvertedTo
	}
    
	if Direction, ok := FlowactivityentitydataMap["direction"].(string); ok {
		o.Direction = &Direction
	}
    
	if Dnis, ok := FlowactivityentitydataMap["dnis"].(string); ok {
		o.Dnis = &Dnis
	}
    
	if FlowId, ok := FlowactivityentitydataMap["flowId"].(string); ok {
		o.FlowId = &FlowId
	}
    
	if FlowType, ok := FlowactivityentitydataMap["flowType"].(string); ok {
		o.FlowType = &FlowType
	}
    
	if MediaType, ok := FlowactivityentitydataMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if ParticipantName, ok := FlowactivityentitydataMap["participantName"].(string); ok {
		o.ParticipantName = &ParticipantName
	}
    
	if QueueId, ok := FlowactivityentitydataMap["queueId"].(string); ok {
		o.QueueId = &QueueId
	}
    
	if RequestedLanguageId, ok := FlowactivityentitydataMap["requestedLanguageId"].(string); ok {
		o.RequestedLanguageId = &RequestedLanguageId
	}
    
	if RequestedRoutingSkillIds, ok := FlowactivityentitydataMap["requestedRoutingSkillIds"].([]interface{}); ok {
		RequestedRoutingSkillIdsString, _ := json.Marshal(RequestedRoutingSkillIds)
		json.Unmarshal(RequestedRoutingSkillIdsString, &o.RequestedRoutingSkillIds)
	}
	
	if RequestedRoutings, ok := FlowactivityentitydataMap["requestedRoutings"].([]interface{}); ok {
		RequestedRoutingsString, _ := json.Marshal(RequestedRoutings)
		json.Unmarshal(RequestedRoutingsString, &o.RequestedRoutings)
	}
	
	if RoutingPriority, ok := FlowactivityentitydataMap["routingPriority"].(float64); ok {
		RoutingPriorityInt := int(RoutingPriority)
		o.RoutingPriority = &RoutingPriorityInt
	}
	
	if SessionId, ok := FlowactivityentitydataMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if TeamId, ok := FlowactivityentitydataMap["teamId"].(string); ok {
		o.TeamId = &TeamId
	}
    
	if UsedRouting, ok := FlowactivityentitydataMap["usedRouting"].(string); ok {
		o.UsedRouting = &UsedRouting
	}
    
	if UserId, ok := FlowactivityentitydataMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if ScoredAgents, ok := FlowactivityentitydataMap["scoredAgents"].([]interface{}); ok {
		ScoredAgentsString, _ := json.Marshal(ScoredAgents)
		json.Unmarshal(ScoredAgentsString, &o.ScoredAgents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Flowactivityentitydata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
