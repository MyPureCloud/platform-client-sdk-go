package platformclientv2
import (
	"time"
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Klaxonheartbeatalertstopicheartbeatalert) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
