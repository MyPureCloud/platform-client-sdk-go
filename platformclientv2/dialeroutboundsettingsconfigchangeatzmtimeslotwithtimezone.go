package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (u *Dialeroutboundsettingsconfigchangeatzmtimeslotwithtimezone) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialeroutboundsettingsconfigchangeatzmtimeslotwithtimezone

	

	return json.Marshal(&struct { 
		EarliestCallableTime *string `json:"earliestCallableTime,omitempty"`
		
		LatestCallableTime *string `json:"latestCallableTime,omitempty"`
		
		TimeZoneId *string `json:"timeZoneId,omitempty"`
		*Alias
	}{ 
		EarliestCallableTime: u.EarliestCallableTime,
		
		LatestCallableTime: u.LatestCallableTime,
		
		TimeZoneId: u.TimeZoneId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialeroutboundsettingsconfigchangeatzmtimeslotwithtimezone) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
