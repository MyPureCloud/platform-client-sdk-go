package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Dialeroutboundsettingsconfigchangeatzmtimeslot) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialeroutboundsettingsconfigchangeatzmtimeslot

	

	return json.Marshal(&struct { 
		EarliestCallableTime *string `json:"earliestCallableTime,omitempty"`
		
		LatestCallableTime *string `json:"latestCallableTime,omitempty"`
		*Alias
	}{ 
		EarliestCallableTime: u.EarliestCallableTime,
		
		LatestCallableTime: u.LatestCallableTime,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialeroutboundsettingsconfigchangeatzmtimeslot) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
