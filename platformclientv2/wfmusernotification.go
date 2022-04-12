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


	// AdherenceExplanation - An adherence explanation notification.  Only set if type == AdherenceExplanation
	AdherenceExplanation *Adherenceexplanationnotification `json:"adherenceExplanation,omitempty"`


	// MarkedAsRead - Whether this notification has been marked \"read\"
	MarkedAsRead *bool `json:"markedAsRead,omitempty"`


	// AgentNotification - Whether this notification is for an agent
	AgentNotification *bool `json:"agentNotification,omitempty"`


	// OtherNotificationIdsInGroup - Other notification IDs in group.  This field is only populated in real-time notifications
	OtherNotificationIdsInGroup *[]string `json:"otherNotificationIdsInGroup,omitempty"`

}

func (o *Wfmusernotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmusernotification
	
	Timestamp := new(string)
	if o.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(o.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		
		AdherenceExplanation *Adherenceexplanationnotification `json:"adherenceExplanation,omitempty"`
		
		MarkedAsRead *bool `json:"markedAsRead,omitempty"`
		
		AgentNotification *bool `json:"agentNotification,omitempty"`
		
		OtherNotificationIdsInGroup *[]string `json:"otherNotificationIdsInGroup,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		MutableGroupId: o.MutableGroupId,
		
		Timestamp: Timestamp,
		
		VarType: o.VarType,
		
		ShiftTrade: o.ShiftTrade,
		
		TimeOffRequest: o.TimeOffRequest,
		
		AdherenceExplanation: o.AdherenceExplanation,
		
		MarkedAsRead: o.MarkedAsRead,
		
		AgentNotification: o.AgentNotification,
		
		OtherNotificationIdsInGroup: o.OtherNotificationIdsInGroup,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmusernotification) UnmarshalJSON(b []byte) error {
	var WfmusernotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmusernotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WfmusernotificationMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if MutableGroupId, ok := WfmusernotificationMap["mutableGroupId"].(string); ok {
		o.MutableGroupId = &MutableGroupId
	}
	
	if timestampString, ok := WfmusernotificationMap["timestamp"].(string); ok {
		Timestamp, _ := time.Parse("2006-01-02T15:04:05.999999Z", timestampString)
		o.Timestamp = &Timestamp
	}
	
	if VarType, ok := WfmusernotificationMap["type"].(string); ok {
		o.VarType = &VarType
	}
	
	if ShiftTrade, ok := WfmusernotificationMap["shiftTrade"].(map[string]interface{}); ok {
		ShiftTradeString, _ := json.Marshal(ShiftTrade)
		json.Unmarshal(ShiftTradeString, &o.ShiftTrade)
	}
	
	if TimeOffRequest, ok := WfmusernotificationMap["timeOffRequest"].(map[string]interface{}); ok {
		TimeOffRequestString, _ := json.Marshal(TimeOffRequest)
		json.Unmarshal(TimeOffRequestString, &o.TimeOffRequest)
	}
	
	if AdherenceExplanation, ok := WfmusernotificationMap["adherenceExplanation"].(map[string]interface{}); ok {
		AdherenceExplanationString, _ := json.Marshal(AdherenceExplanation)
		json.Unmarshal(AdherenceExplanationString, &o.AdherenceExplanation)
	}
	
	if MarkedAsRead, ok := WfmusernotificationMap["markedAsRead"].(bool); ok {
		o.MarkedAsRead = &MarkedAsRead
	}
	
	if AgentNotification, ok := WfmusernotificationMap["agentNotification"].(bool); ok {
		o.AgentNotification = &AgentNotification
	}
	
	if OtherNotificationIdsInGroup, ok := WfmusernotificationMap["otherNotificationIdsInGroup"].([]interface{}); ok {
		OtherNotificationIdsInGroupString, _ := json.Marshal(OtherNotificationIdsInGroup)
		json.Unmarshal(OtherNotificationIdsInGroupString, &o.OtherNotificationIdsInGroup)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmusernotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
