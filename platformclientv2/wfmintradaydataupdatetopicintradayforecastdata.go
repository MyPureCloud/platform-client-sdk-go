package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmintradaydataupdatetopicintradayforecastdata
type Wfmintradaydataupdatetopicintradayforecastdata struct { 
	// Offered
	Offered *float32 `json:"offered,omitempty"`


	// AverageTalkTimeSeconds
	AverageTalkTimeSeconds *float32 `json:"averageTalkTimeSeconds,omitempty"`


	// AverageAfterCallWorkSeconds
	AverageAfterCallWorkSeconds *float32 `json:"averageAfterCallWorkSeconds,omitempty"`

}

func (o *Wfmintradaydataupdatetopicintradayforecastdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmintradaydataupdatetopicintradayforecastdata
	
	return json.Marshal(&struct { 
		Offered *float32 `json:"offered,omitempty"`
		
		AverageTalkTimeSeconds *float32 `json:"averageTalkTimeSeconds,omitempty"`
		
		AverageAfterCallWorkSeconds *float32 `json:"averageAfterCallWorkSeconds,omitempty"`
		*Alias
	}{ 
		Offered: o.Offered,
		
		AverageTalkTimeSeconds: o.AverageTalkTimeSeconds,
		
		AverageAfterCallWorkSeconds: o.AverageAfterCallWorkSeconds,
		Alias:    (*Alias)(o),
	})
}

func (o *Wfmintradaydataupdatetopicintradayforecastdata) UnmarshalJSON(b []byte) error {
	var WfmintradaydataupdatetopicintradayforecastdataMap map[string]interface{}
	err := json.Unmarshal(b, &WfmintradaydataupdatetopicintradayforecastdataMap)
	if err != nil {
		return err
	}
	
	if Offered, ok := WfmintradaydataupdatetopicintradayforecastdataMap["offered"].(float64); ok {
		OfferedFloat32 := float32(Offered)
		o.Offered = &OfferedFloat32
	}
    
	if AverageTalkTimeSeconds, ok := WfmintradaydataupdatetopicintradayforecastdataMap["averageTalkTimeSeconds"].(float64); ok {
		AverageTalkTimeSecondsFloat32 := float32(AverageTalkTimeSeconds)
		o.AverageTalkTimeSeconds = &AverageTalkTimeSecondsFloat32
	}
    
	if AverageAfterCallWorkSeconds, ok := WfmintradaydataupdatetopicintradayforecastdataMap["averageAfterCallWorkSeconds"].(float64); ok {
		AverageAfterCallWorkSecondsFloat32 := float32(AverageAfterCallWorkSeconds)
		o.AverageAfterCallWorkSeconds = &AverageAfterCallWorkSecondsFloat32
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradayforecastdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
