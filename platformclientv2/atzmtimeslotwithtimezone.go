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

func (o *Atzmtimeslotwithtimezone) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Atzmtimeslotwithtimezone
	
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

func (o *Atzmtimeslotwithtimezone) UnmarshalJSON(b []byte) error {
	var AtzmtimeslotwithtimezoneMap map[string]interface{}
	err := json.Unmarshal(b, &AtzmtimeslotwithtimezoneMap)
	if err != nil {
		return err
	}
	
	if EarliestCallableTime, ok := AtzmtimeslotwithtimezoneMap["earliestCallableTime"].(string); ok {
		o.EarliestCallableTime = &EarliestCallableTime
	}
    
	if LatestCallableTime, ok := AtzmtimeslotwithtimezoneMap["latestCallableTime"].(string); ok {
		o.LatestCallableTime = &LatestCallableTime
	}
    
	if TimeZoneId, ok := AtzmtimeslotwithtimezoneMap["timeZoneId"].(string); ok {
		o.TimeZoneId = &TimeZoneId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Atzmtimeslotwithtimezone) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
