package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
