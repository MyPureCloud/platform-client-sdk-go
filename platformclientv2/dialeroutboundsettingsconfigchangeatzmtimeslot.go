package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialeroutboundsettingsconfigchangeatzmtimeslot - The time interval to place outbound calls
type Dialeroutboundsettingsconfigchangeatzmtimeslot struct { 
	// EarliestCallableTime - The earliest time to dial a contact
	EarliestCallableTime *string `json:"earliestCallableTime,omitempty"`


	// LatestCallableTime - The latest time to dial a contact
	LatestCallableTime *string `json:"latestCallableTime,omitempty"`

}

func (o *Dialeroutboundsettingsconfigchangeatzmtimeslot) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialeroutboundsettingsconfigchangeatzmtimeslot
	
	return json.Marshal(&struct { 
		EarliestCallableTime *string `json:"earliestCallableTime,omitempty"`
		
		LatestCallableTime *string `json:"latestCallableTime,omitempty"`
		*Alias
	}{ 
		EarliestCallableTime: o.EarliestCallableTime,
		
		LatestCallableTime: o.LatestCallableTime,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialeroutboundsettingsconfigchangeatzmtimeslot) UnmarshalJSON(b []byte) error {
	var DialeroutboundsettingsconfigchangeatzmtimeslotMap map[string]interface{}
	err := json.Unmarshal(b, &DialeroutboundsettingsconfigchangeatzmtimeslotMap)
	if err != nil {
		return err
	}
	
	if EarliestCallableTime, ok := DialeroutboundsettingsconfigchangeatzmtimeslotMap["earliestCallableTime"].(string); ok {
		o.EarliestCallableTime = &EarliestCallableTime
	}
	
	if LatestCallableTime, ok := DialeroutboundsettingsconfigchangeatzmtimeslotMap["latestCallableTime"].(string); ok {
		o.LatestCallableTime = &LatestCallableTime
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialeroutboundsettingsconfigchangeatzmtimeslot) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
