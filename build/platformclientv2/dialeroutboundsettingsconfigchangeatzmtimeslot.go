package platformclientv2
import (
	"encoding/json"
)

// Dialeroutboundsettingsconfigchangeatzmtimeslot
type Dialeroutboundsettingsconfigchangeatzmtimeslot struct { 
	// EarliestCallableTime
	EarliestCallableTime *string `json:"earliestCallableTime,omitempty"`


	// LatestCallableTime
	LatestCallableTime *string `json:"latestCallableTime,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialeroutboundsettingsconfigchangeatzmtimeslot) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
