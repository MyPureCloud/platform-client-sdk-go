package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Buintradayforecastdata
type Buintradayforecastdata struct { 
	// Offered - The number of interactions routed into the queues in the selected planning groups for the given media type for an agent to answer
	Offered *float64 `json:"offered,omitempty"`


	// AverageHandleTimeSeconds - The average handle time in seconds an agent spent handling interactions
	AverageHandleTimeSeconds *float64 `json:"averageHandleTimeSeconds,omitempty"`

}

func (o *Buintradayforecastdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buintradayforecastdata
	
	return json.Marshal(&struct { 
		Offered *float64 `json:"offered,omitempty"`
		
		AverageHandleTimeSeconds *float64 `json:"averageHandleTimeSeconds,omitempty"`
		*Alias
	}{ 
		Offered: o.Offered,
		
		AverageHandleTimeSeconds: o.AverageHandleTimeSeconds,
		Alias:    (*Alias)(o),
	})
}

func (o *Buintradayforecastdata) UnmarshalJSON(b []byte) error {
	var BuintradayforecastdataMap map[string]interface{}
	err := json.Unmarshal(b, &BuintradayforecastdataMap)
	if err != nil {
		return err
	}
	
	if Offered, ok := BuintradayforecastdataMap["offered"].(float64); ok {
		o.Offered = &Offered
	}
	
	if AverageHandleTimeSeconds, ok := BuintradayforecastdataMap["averageHandleTimeSeconds"].(float64); ok {
		o.AverageHandleTimeSeconds = &AverageHandleTimeSeconds
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buintradayforecastdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
