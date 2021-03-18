package platformclientv2
import (
	"time"
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

// String returns a JSON representation of the model
func (o *Wfmusernotificationtopicwfmusernotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
