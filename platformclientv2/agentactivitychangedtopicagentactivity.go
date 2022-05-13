package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentactivitychangedtopicagentactivity
type Agentactivitychangedtopicagentactivity struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// RoutingStatus
	RoutingStatus *Agentactivitychangedtopicroutingstatus `json:"routingStatus,omitempty"`


	// Presence
	Presence *Agentactivitychangedtopicpresence `json:"presence,omitempty"`


	// OutOfOffice
	OutOfOffice *Agentactivitychangedtopicoutofoffice `json:"outOfOffice,omitempty"`


	// ActiveQueueIds
	ActiveQueueIds *[]string `json:"activeQueueIds,omitempty"`


	// DateActiveQueuesChanged
	DateActiveQueuesChanged *time.Time `json:"dateActiveQueuesChanged,omitempty"`

}

func (o *Agentactivitychangedtopicagentactivity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Agentactivitychangedtopicagentactivity
	
	DateActiveQueuesChanged := new(string)
	if o.DateActiveQueuesChanged != nil {
		
		*DateActiveQueuesChanged = timeutil.Strftime(o.DateActiveQueuesChanged, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateActiveQueuesChanged = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		RoutingStatus *Agentactivitychangedtopicroutingstatus `json:"routingStatus,omitempty"`
		
		Presence *Agentactivitychangedtopicpresence `json:"presence,omitempty"`
		
		OutOfOffice *Agentactivitychangedtopicoutofoffice `json:"outOfOffice,omitempty"`
		
		ActiveQueueIds *[]string `json:"activeQueueIds,omitempty"`
		
		DateActiveQueuesChanged *string `json:"dateActiveQueuesChanged,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		RoutingStatus: o.RoutingStatus,
		
		Presence: o.Presence,
		
		OutOfOffice: o.OutOfOffice,
		
		ActiveQueueIds: o.ActiveQueueIds,
		
		DateActiveQueuesChanged: DateActiveQueuesChanged,
		Alias:    (*Alias)(o),
	})
}

func (o *Agentactivitychangedtopicagentactivity) UnmarshalJSON(b []byte) error {
	var AgentactivitychangedtopicagentactivityMap map[string]interface{}
	err := json.Unmarshal(b, &AgentactivitychangedtopicagentactivityMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AgentactivitychangedtopicagentactivityMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if RoutingStatus, ok := AgentactivitychangedtopicagentactivityMap["routingStatus"].(map[string]interface{}); ok {
		RoutingStatusString, _ := json.Marshal(RoutingStatus)
		json.Unmarshal(RoutingStatusString, &o.RoutingStatus)
	}
	
	if Presence, ok := AgentactivitychangedtopicagentactivityMap["presence"].(map[string]interface{}); ok {
		PresenceString, _ := json.Marshal(Presence)
		json.Unmarshal(PresenceString, &o.Presence)
	}
	
	if OutOfOffice, ok := AgentactivitychangedtopicagentactivityMap["outOfOffice"].(map[string]interface{}); ok {
		OutOfOfficeString, _ := json.Marshal(OutOfOffice)
		json.Unmarshal(OutOfOfficeString, &o.OutOfOffice)
	}
	
	if ActiveQueueIds, ok := AgentactivitychangedtopicagentactivityMap["activeQueueIds"].([]interface{}); ok {
		ActiveQueueIdsString, _ := json.Marshal(ActiveQueueIds)
		json.Unmarshal(ActiveQueueIdsString, &o.ActiveQueueIds)
	}
	
	if dateActiveQueuesChangedString, ok := AgentactivitychangedtopicagentactivityMap["dateActiveQueuesChanged"].(string); ok {
		DateActiveQueuesChanged, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateActiveQueuesChangedString)
		o.DateActiveQueuesChanged = &DateActiveQueuesChanged
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Agentactivitychangedtopicagentactivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
