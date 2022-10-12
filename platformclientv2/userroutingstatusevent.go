package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userroutingstatusevent
type Userroutingstatusevent struct { 
	// EventId - A unique (UUID) eventId for this event
	EventId *string `json:"eventId,omitempty"`


	// EventDateTime - A timestamp as epoch representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EventDateTime *time.Time `json:"eventDateTime,omitempty"`


	// AgentId - Unique identifier of the agent.
	AgentId *string `json:"agentId,omitempty"`


	// Status - The agent's current routing status.
	Status *string `json:"status,omitempty"`


	// SourceId - The agent's source platform Id.
	SourceId *string `json:"sourceId,omitempty"`

}

func (o *Userroutingstatusevent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userroutingstatusevent
	
	EventDateTime := new(string)
	if o.EventDateTime != nil {
		
		*EventDateTime = timeutil.Strftime(o.EventDateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EventDateTime = nil
	}
	
	return json.Marshal(&struct { 
		EventId *string `json:"eventId,omitempty"`
		
		EventDateTime *string `json:"eventDateTime,omitempty"`
		
		AgentId *string `json:"agentId,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		SourceId *string `json:"sourceId,omitempty"`
		*Alias
	}{ 
		EventId: o.EventId,
		
		EventDateTime: EventDateTime,
		
		AgentId: o.AgentId,
		
		Status: o.Status,
		
		SourceId: o.SourceId,
		Alias:    (*Alias)(o),
	})
}

func (o *Userroutingstatusevent) UnmarshalJSON(b []byte) error {
	var UserroutingstatuseventMap map[string]interface{}
	err := json.Unmarshal(b, &UserroutingstatuseventMap)
	if err != nil {
		return err
	}
	
	if EventId, ok := UserroutingstatuseventMap["eventId"].(string); ok {
		o.EventId = &EventId
	}
    
	if eventDateTimeString, ok := UserroutingstatuseventMap["eventDateTime"].(string); ok {
		EventDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", eventDateTimeString)
		o.EventDateTime = &EventDateTime
	}
	
	if AgentId, ok := UserroutingstatuseventMap["agentId"].(string); ok {
		o.AgentId = &AgentId
	}
    
	if Status, ok := UserroutingstatuseventMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if SourceId, ok := UserroutingstatuseventMap["sourceId"].(string); ok {
		o.SourceId = &SourceId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Userroutingstatusevent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
