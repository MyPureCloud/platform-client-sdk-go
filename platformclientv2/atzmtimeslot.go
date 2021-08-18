package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Atzmtimeslot
type Atzmtimeslot struct { 
	// EarliestCallableTime - The earliest time to dial a contact. Valid format is HH:mm
	EarliestCallableTime *string `json:"earliestCallableTime,omitempty"`


	// LatestCallableTime - The latest time to dial a contact. Valid format is HH:mm
	LatestCallableTime *string `json:"latestCallableTime,omitempty"`

}

func (u *Atzmtimeslot) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Atzmtimeslot

	

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
func (o *Atzmtimeslot) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
