package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Klaxonheartbeatrulestopicheartbeatrule
type Klaxonheartbeatrulestopicheartbeatrule struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// SenderId
	SenderId *string `json:"senderId,omitempty"`


	// HeartBeatTimeoutInMinutes
	HeartBeatTimeoutInMinutes *float32 `json:"heartBeatTimeoutInMinutes,omitempty"`


	// Enabled
	Enabled *bool `json:"enabled,omitempty"`


	// InAlarm
	InAlarm *bool `json:"inAlarm,omitempty"`


	// NotificationUsers
	NotificationUsers *[]Klaxonheartbeatrulestopicnotificationuser `json:"notificationUsers,omitempty"`


	// AlertTypes
	AlertTypes *[]string `json:"alertTypes,omitempty"`


	// RuleType
	RuleType *string `json:"ruleType,omitempty"`

}

// String returns a JSON representation of the model
func (o *Klaxonheartbeatrulestopicheartbeatrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
