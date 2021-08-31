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

func (o *Klaxonheartbeatrulestopicheartbeatrule) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Name: o.Name,
		
		SenderId: o.SenderId,
		
		HeartBeatTimeoutInMinutes: o.HeartBeatTimeoutInMinutes,
		
		Enabled: o.Enabled,
		
		InAlarm: o.InAlarm,
		
		NotificationUsers: o.NotificationUsers,
		
		AlertTypes: o.AlertTypes,
		
		RuleType: o.RuleType,
		Alias:    (*Alias)(o),
	})
}

func (o *Klaxonheartbeatrulestopicheartbeatrule) UnmarshalJSON(b []byte) error {
	var KlaxonheartbeatrulestopicheartbeatruleMap map[string]interface{}
	err := json.Unmarshal(b, &KlaxonheartbeatrulestopicheartbeatruleMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KlaxonheartbeatrulestopicheartbeatruleMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := KlaxonheartbeatrulestopicheartbeatruleMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if SenderId, ok := KlaxonheartbeatrulestopicheartbeatruleMap["senderId"].(string); ok {
		o.SenderId = &SenderId
	}
	
	if HeartBeatTimeoutInMinutes, ok := KlaxonheartbeatrulestopicheartbeatruleMap["heartBeatTimeoutInMinutes"].(float64); ok {
		HeartBeatTimeoutInMinutesFloat32 := float32(HeartBeatTimeoutInMinutes)
		o.HeartBeatTimeoutInMinutes = &HeartBeatTimeoutInMinutesFloat32
	}
	
	if Enabled, ok := KlaxonheartbeatrulestopicheartbeatruleMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
	
	if InAlarm, ok := KlaxonheartbeatrulestopicheartbeatruleMap["inAlarm"].(bool); ok {
		o.InAlarm = &InAlarm
	}
	
	if NotificationUsers, ok := KlaxonheartbeatrulestopicheartbeatruleMap["notificationUsers"].([]interface{}); ok {
		NotificationUsersString, _ := json.Marshal(NotificationUsers)
		json.Unmarshal(NotificationUsersString, &o.NotificationUsers)
	}
	
	if AlertTypes, ok := KlaxonheartbeatrulestopicheartbeatruleMap["alertTypes"].([]interface{}); ok {
		AlertTypesString, _ := json.Marshal(AlertTypes)
		json.Unmarshal(AlertTypesString, &o.AlertTypes)
	}
	
	if RuleType, ok := KlaxonheartbeatrulestopicheartbeatruleMap["ruleType"].(string); ok {
		o.RuleType = &RuleType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Klaxonheartbeatrulestopicheartbeatrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
