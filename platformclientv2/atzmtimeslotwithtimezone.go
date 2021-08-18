package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Atzmtimeslotwithtimezone
type Atzmtimeslotwithtimezone struct { 
	// EarliestCallableTime - The earliest time to dial a contact. Valid format is HH:mm
	EarliestCallableTime *string `json:"earliestCallableTime,omitempty"`


	// LatestCallableTime - The latest time to dial a contact. Valid format is HH:mm
	LatestCallableTime *string `json:"latestCallableTime,omitempty"`


	// TimeZoneId - The time zone to use for contacts that cannot be mapped.
	TimeZoneId *string `json:"timeZoneId,omitempty"`

}

func (u *Atzmtimeslotwithtimezone) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Atzmtimeslotwithtimezone

	

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
func (o *Atzmtimeslotwithtimezone) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
