package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Klaxonheartbeatalertstopicheartbeatalert
type Klaxonheartbeatalertstopicheartbeatalert struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// SenderId
	SenderId *string `json:"senderId,omitempty"`


	// HeartBeatTimeoutInMinutes
	HeartBeatTimeoutInMinutes *float32 `json:"heartBeatTimeoutInMinutes,omitempty"`


	// RuleId
	RuleId *string `json:"ruleId,omitempty"`


	// StartDate
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate
	EndDate *time.Time `json:"endDate,omitempty"`


	// NotificationUsers
	NotificationUsers *[]Klaxonheartbeatalertstopicnotificationuser `json:"notificationUsers,omitempty"`


	// AlertTypes
	AlertTypes *[]string `json:"alertTypes,omitempty"`


	// RuleType
	RuleType *string `json:"ruleType,omitempty"`

}

func (o *Klaxonheartbeatalertstopicheartbeatalert) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Klaxonheartbeatalertstopicheartbeatalert
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if o.EndDate != nil {
		
		*EndDate = timeutil.Strftime(o.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SenderId *string `json:"senderId,omitempty"`
		
		HeartBeatTimeoutInMinutes *float32 `json:"heartBeatTimeoutInMinutes,omitempty"`
		
		RuleId *string `json:"ruleId,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		NotificationUsers *[]Klaxonheartbeatalertstopicnotificationuser `json:"notificationUsers,omitempty"`
		
		AlertTypes *[]string `json:"alertTypes,omitempty"`
		
		RuleType *string `json:"ruleType,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SenderId: o.SenderId,
		
		HeartBeatTimeoutInMinutes: o.HeartBeatTimeoutInMinutes,
		
		RuleId: o.RuleId,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		NotificationUsers: o.NotificationUsers,
		
		AlertTypes: o.AlertTypes,
		
		RuleType: o.RuleType,
		Alias:    (*Alias)(o),
	})
}

func (o *Klaxonheartbeatalertstopicheartbeatalert) UnmarshalJSON(b []byte) error {
	var KlaxonheartbeatalertstopicheartbeatalertMap map[string]interface{}
	err := json.Unmarshal(b, &KlaxonheartbeatalertstopicheartbeatalertMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KlaxonheartbeatalertstopicheartbeatalertMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := KlaxonheartbeatalertstopicheartbeatalertMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if SenderId, ok := KlaxonheartbeatalertstopicheartbeatalertMap["senderId"].(string); ok {
		o.SenderId = &SenderId
	}
	
	if HeartBeatTimeoutInMinutes, ok := KlaxonheartbeatalertstopicheartbeatalertMap["heartBeatTimeoutInMinutes"].(float64); ok {
		HeartBeatTimeoutInMinutesFloat32 := float32(HeartBeatTimeoutInMinutes)
		o.HeartBeatTimeoutInMinutes = &HeartBeatTimeoutInMinutesFloat32
	}
	
	if RuleId, ok := KlaxonheartbeatalertstopicheartbeatalertMap["ruleId"].(string); ok {
		o.RuleId = &RuleId
	}
	
	if startDateString, ok := KlaxonheartbeatalertstopicheartbeatalertMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := KlaxonheartbeatalertstopicheartbeatalertMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if NotificationUsers, ok := KlaxonheartbeatalertstopicheartbeatalertMap["notificationUsers"].([]interface{}); ok {
		NotificationUsersString, _ := json.Marshal(NotificationUsers)
		json.Unmarshal(NotificationUsersString, &o.NotificationUsers)
	}
	
	if AlertTypes, ok := KlaxonheartbeatalertstopicheartbeatalertMap["alertTypes"].([]interface{}); ok {
		AlertTypesString, _ := json.Marshal(AlertTypes)
		json.Unmarshal(AlertTypesString, &o.AlertTypes)
	}
	
	if RuleType, ok := KlaxonheartbeatalertstopicheartbeatalertMap["ruleType"].(string); ok {
		o.RuleType = &RuleType
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Klaxonheartbeatalertstopicheartbeatalert) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
