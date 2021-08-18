package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmusernotificationtopicwfmusernotification
type Wfmusernotificationtopicwfmusernotification struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// MutableGroupId
	MutableGroupId *string `json:"mutableGroupId,omitempty"`


	// Timestamp
	Timestamp *time.Time `json:"timestamp,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// ShiftTrade
	ShiftTrade *Wfmusernotificationtopicshifttradenotification `json:"shiftTrade,omitempty"`


	// TimeOffRequest
	TimeOffRequest *Wfmusernotificationtopictimeoffrequestnotification `json:"timeOffRequest,omitempty"`


	// AgentNotification
	AgentNotification *bool `json:"agentNotification,omitempty"`


	// OtherNotificationIdsInGroup
	OtherNotificationIdsInGroup *[]string `json:"otherNotificationIdsInGroup,omitempty"`


	// MarkedAsRead
	MarkedAsRead *bool `json:"markedAsRead,omitempty"`

}

func (u *Wfmusernotificationtopicwfmusernotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmusernotificationtopicwfmusernotification

	
	Timestamp := new(string)
	if u.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(u.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Timestamp = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		MutableGroupId *string `json:"mutableGroupId,omitempty"`
		
		Timestamp *string `json:"timestamp,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		ShiftTrade *Wfmusernotificationtopicshifttradenotification `json:"shiftTrade,omitempty"`
		
		TimeOffRequest *Wfmusernotificationtopictimeoffrequestnotification `json:"timeOffRequest,omitempty"`
		
		AgentNotification *bool `json:"agentNotification,omitempty"`
		
		OtherNotificationIdsInGroup *[]string `json:"otherNotificationIdsInGroup,omitempty"`
		
		MarkedAsRead *bool `json:"markedAsRead,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		MutableGroupId: u.MutableGroupId,
		
		Timestamp: Timestamp,
		
		VarType: u.VarType,
		
		ShiftTrade: u.ShiftTrade,
		
		TimeOffRequest: u.TimeOffRequest,
		
		AgentNotification: u.AgentNotification,
		
		OtherNotificationIdsInGroup: u.OtherNotificationIdsInGroup,
		
		MarkedAsRead: u.MarkedAsRead,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmusernotificationtopicwfmusernotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
