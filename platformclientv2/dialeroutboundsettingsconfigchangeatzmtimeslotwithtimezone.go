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

func (o *Dialeroutboundsettingsconfigchangeatzmtimeslotwithtimezone) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialeroutboundsettingsconfigchangeatzmtimeslotwithtimezone
	
	return json.Marshal(&struct { 
		EarliestCallableTime *string `json:"earliestCallableTime,omitempty"`
		
		LatestCallableTime *string `json:"latestCallableTime,omitempty"`
		
		TimeZoneId *string `json:"timeZoneId,omitempty"`
		*Alias
	}{ 
		EarliestCallableTime: o.EarliestCallableTime,
		
		LatestCallableTime: o.LatestCallableTime,
		
		TimeZoneId: o.TimeZoneId,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialeroutboundsettingsconfigchangeatzmtimeslotwithtimezone) UnmarshalJSON(b []byte) error {
	var DialeroutboundsettingsconfigchangeatzmtimeslotwithtimezoneMap map[string]interface{}
	err := json.Unmarshal(b, &DialeroutboundsettingsconfigchangeatzmtimeslotwithtimezoneMap)
	if err != nil {
		return err
	}
	
	if EarliestCallableTime, ok := DialeroutboundsettingsconfigchangeatzmtimeslotwithtimezoneMap["earliestCallableTime"].(string); ok {
		o.EarliestCallableTime = &EarliestCallableTime
	}
	
	if LatestCallableTime, ok := DialeroutboundsettingsconfigchangeatzmtimeslotwithtimezoneMap["latestCallableTime"].(string); ok {
		o.LatestCallableTime = &LatestCallableTime
	}
	
	if TimeZoneId, ok := DialeroutboundsettingsconfigchangeatzmtimeslotwithtimezoneMap["timeZoneId"].(string); ok {
		o.TimeZoneId = &TimeZoneId
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialeroutboundsettingsconfigchangeatzmtimeslotwithtimezone) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
