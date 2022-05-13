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

func (o *Atzmtimeslot) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Atzmtimeslot
	
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

func (o *Atzmtimeslot) UnmarshalJSON(b []byte) error {
	var AtzmtimeslotMap map[string]interface{}
	err := json.Unmarshal(b, &AtzmtimeslotMap)
	if err != nil {
		return err
	}
	
	if EarliestCallableTime, ok := AtzmtimeslotMap["earliestCallableTime"].(string); ok {
		o.EarliestCallableTime = &EarliestCallableTime
	}
    
	if LatestCallableTime, ok := AtzmtimeslotMap["latestCallableTime"].(string); ok {
		o.LatestCallableTime = &LatestCallableTime
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Atzmtimeslot) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
