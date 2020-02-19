package platformclientv2
import (
	"encoding/json"
)

// Dialeroutboundsettingsconfigchangeatzmtimeslotwithtimezone
type Dialeroutboundsettingsconfigchangeatzmtimeslotwithtimezone struct { 
	// EarliestCallableTime
	EarliestCallableTime *string `json:"earliestCallableTime,omitempty"`


	// LatestCallableTime
	LatestCallableTime *string `json:"latestCallableTime,omitempty"`


	// TimeZoneId
	TimeZoneId *string `json:"timeZoneId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialeroutboundsettingsconfigchangeatzmtimeslotwithtimezone) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
