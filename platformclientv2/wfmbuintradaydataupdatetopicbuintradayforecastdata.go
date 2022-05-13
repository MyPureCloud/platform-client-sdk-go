package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuintradaydataupdatetopicbuintradayforecastdata
type Wfmbuintradaydataupdatetopicbuintradayforecastdata struct { 
	// Offered
	Offered *float32 `json:"offered,omitempty"`


	// AverageHandleTimeSeconds
	AverageHandleTimeSeconds *float32 `json:"averageHandleTimeSeconds,omitempty"`

}

func (o *Wfmbuintradaydataupdatetopicbuintradayforecastdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuintradaydataupdatetopicbuintradayforecastdata
	
	return json.Marshal(&struct { 
		Offered *float32 `json:"offered,omitempty"`
		
		AverageHandleTimeSeconds *float32 `json:"averageHandleTimeSeconds,omitempty"`
		*Alias
	}{ 
		Offered: o.Offered,
		
		AverageHandleTimeSeconds: o.AverageHandleTimeSeconds,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmbuintradaydataupdatetopicbuintradayforecastdata) UnmarshalJSON(b []byte) error {
	var WfmbuintradaydataupdatetopicbuintradayforecastdataMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbuintradaydataupdatetopicbuintradayforecastdataMap)
	if err != nil {
		return err
	}
	
	if Offered, ok := WfmbuintradaydataupdatetopicbuintradayforecastdataMap["offered"].(float64); ok {
		OfferedFloat32 := float32(Offered)
		o.Offered = &OfferedFloat32
	}
    
	if AverageHandleTimeSeconds, ok := WfmbuintradaydataupdatetopicbuintradayforecastdataMap["averageHandleTimeSeconds"].(float64); ok {
		AverageHandleTimeSecondsFloat32 := float32(AverageHandleTimeSeconds)
		o.AverageHandleTimeSeconds = &AverageHandleTimeSecondsFloat32
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbuintradaydataupdatetopicbuintradayforecastdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
