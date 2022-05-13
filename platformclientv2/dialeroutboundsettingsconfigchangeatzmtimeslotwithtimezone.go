package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialeroutboundsettingsconfigchangeatzmtimeslotwithtimezone - The time interval to place outbound calls
type Dialeroutboundsettingsconfigchangeatzmtimeslotwithtimezone struct { 
	// TimeZoneId - The time zone to use for contacts that cannot be mapped
	TimeZoneId *string `json:"timeZoneId,omitempty"`


	// EarliestCallableTime - The earliest time to dial a contact
	EarliestCallableTime *string `json:"earliestCallableTime,omitempty"`


	// LatestCallableTime - The latest time to dial a contact
	LatestCallableTime *string `json:"latestCallableTime,omitempty"`

}

func (o *Dialeroutboundsettingsconfigchangeatzmtimeslotwithtimezone) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialeroutboundsettingsconfigchangeatzmtimeslotwithtimezone
	
	return json.Marshal(&struct { 
		TimeZoneId *string `json:"timeZoneId,omitempty"`
		
		EarliestCallableTime *string `json:"earliestCallableTime,omitempty"`
		
		LatestCallableTime *string `json:"latestCallableTime,omitempty"`
		*Alias
	}{ 
		TimeZoneId: o.TimeZoneId,
		
		EarliestCallableTime: o.EarliestCallableTime,
		
		LatestCallableTime: o.LatestCallableTime,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialeroutboundsettingsconfigchangeatzmtimeslotwithtimezone) UnmarshalJSON(b []byte) error {
	var DialeroutboundsettingsconfigchangeatzmtimeslotwithtimezoneMap map[string]interface{}
	err := json.Unmarshal(b, &DialeroutboundsettingsconfigchangeatzmtimeslotwithtimezoneMap)
	if err != nil {
		return err
	}
	
	if TimeZoneId, ok := DialeroutboundsettingsconfigchangeatzmtimeslotwithtimezoneMap["timeZoneId"].(string); ok {
		o.TimeZoneId = &TimeZoneId
	}
    
	if EarliestCallableTime, ok := DialeroutboundsettingsconfigchangeatzmtimeslotwithtimezoneMap["earliestCallableTime"].(string); ok {
		o.EarliestCallableTime = &EarliestCallableTime
	}
    
	if LatestCallableTime, ok := DialeroutboundsettingsconfigchangeatzmtimeslotwithtimezoneMap["latestCallableTime"].(string); ok {
		o.LatestCallableTime = &LatestCallableTime
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Dialeroutboundsettingsconfigchangeatzmtimeslotwithtimezone) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
