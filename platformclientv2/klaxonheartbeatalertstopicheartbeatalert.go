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

func (u *Klaxonheartbeatalertstopicheartbeatalert) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Klaxonheartbeatalertstopicheartbeatalert

	
	StartDate := new(string)
	if u.StartDate != nil {
		
		*StartDate = timeutil.Strftime(u.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if u.EndDate != nil {
		
		*EndDate = timeutil.Strftime(u.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: u.Id,
		
		Name: u.Name,
		
		SenderId: u.SenderId,
		
		HeartBeatTimeoutInMinutes: u.HeartBeatTimeoutInMinutes,
		
		RuleId: u.RuleId,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		NotificationUsers: u.NotificationUsers,
		
		AlertTypes: u.AlertTypes,
		
		RuleType: u.RuleType,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Klaxonheartbeatalertstopicheartbeatalert) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
