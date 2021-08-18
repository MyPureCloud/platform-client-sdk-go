package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmusernotification
type Wfmusernotification struct { 
	// Id - The immutable globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// MutableGroupId - The group ID of the notification (mutable, may change  on update)
	MutableGroupId *string `json:"mutableGroupId,omitempty"`


	// Timestamp - The timestamp for this notification. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Timestamp *time.Time `json:"timestamp,omitempty"`


	// VarType - The type of this notification
	VarType *string `json:"type,omitempty"`


	// ShiftTrade - A shift trade notification.  Only set if type == ShiftTrade
	ShiftTrade *Shifttradenotification `json:"shiftTrade,omitempty"`


	// TimeOffRequest - A time off request notification.  Only set if type == TimeOffRequest
	TimeOffRequest *Timeoffrequestnotification `json:"timeOffRequest,omitempty"`


	// MarkedAsRead - Whether this notification has been marked \"read\"
	MarkedAsRead *bool `json:"markedAsRead,omitempty"`


	// AgentNotification - Whether this notification is for an agent
	AgentNotification *bool `json:"agentNotification,omitempty"`


	// OtherNotificationIdsInGroup - Other notification IDs in group.  This field is only populated in real-time notifications
	OtherNotificationIdsInGroup *[]string `json:"otherNotificationIdsInGroup,omitempty"`

}

func (u *Wfmusernotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmusernotification

	
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
		
		ShiftTrade *Shifttradenotification `json:"shiftTrade,omitempty"`
		
		TimeOffRequest *Timeoffrequestnotification `json:"timeOffRequest,omitempty"`
		
		MarkedAsRead *bool `json:"markedAsRead,omitempty"`
		
		AgentNotification *bool `json:"agentNotification,omitempty"`
		
		OtherNotificationIdsInGroup *[]string `json:"otherNotificationIdsInGroup,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		MutableGroupId: u.MutableGroupId,
		
		Timestamp: Timestamp,
		
		VarType: u.VarType,
		
		ShiftTrade: u.ShiftTrade,
		
		TimeOffRequest: u.TimeOffRequest,
		
		MarkedAsRead: u.MarkedAsRead,
		
		AgentNotification: u.AgentNotification,
		
		OtherNotificationIdsInGroup: u.OtherNotificationIdsInGroup,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmusernotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
