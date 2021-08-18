package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Klaxonheartbeatrulestopicheartbeatrule) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Klaxonheartbeatrulestopicheartbeatrule

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SenderId *string `json:"senderId,omitempty"`
		
		HeartBeatTimeoutInMinutes *float32 `json:"heartBeatTimeoutInMinutes,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		InAlarm *bool `json:"inAlarm,omitempty"`
		
		NotificationUsers *[]Klaxonheartbeatrulestopicnotificationuser `json:"notificationUsers,omitempty"`
		
		AlertTypes *[]string `json:"alertTypes,omitempty"`
		
		RuleType *string `json:"ruleType,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		SenderId: u.SenderId,
		
		HeartBeatTimeoutInMinutes: u.HeartBeatTimeoutInMinutes,
		
		Enabled: u.Enabled,
		
		InAlarm: u.InAlarm,
		
		NotificationUsers: u.NotificationUsers,
		
		AlertTypes: u.AlertTypes,
		
		RuleType: u.RuleType,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Klaxonheartbeatrulestopicheartbeatrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
