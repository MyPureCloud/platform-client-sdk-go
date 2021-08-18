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

func (u *Agentactivitychangedtopicagentactivity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Agentactivitychangedtopicagentactivity

	
	DateActiveQueuesChanged := new(string)
	if u.DateActiveQueuesChanged != nil {
		
		*DateActiveQueuesChanged = timeutil.Strftime(u.DateActiveQueuesChanged, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: u.Id,
		
		RoutingStatus: u.RoutingStatus,
		
		Presence: u.Presence,
		
		OutOfOffice: u.OutOfOffice,
		
		ActiveQueueIds: u.ActiveQueueIds,
		
		DateActiveQueuesChanged: DateActiveQueuesChanged,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Agentactivitychangedtopicagentactivity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
