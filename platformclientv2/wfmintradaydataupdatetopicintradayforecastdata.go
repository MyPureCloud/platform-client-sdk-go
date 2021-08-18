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

func (u *Wfmintradaydataupdatetopicintradayforecastdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmintradaydataupdatetopicintradayforecastdata

	

	return json.Marshal(&struct { 
		Offered *float32 `json:"offered,omitempty"`
		
		AverageTalkTimeSeconds *float32 `json:"averageTalkTimeSeconds,omitempty"`
		
		AverageAfterCallWorkSeconds *float32 `json:"averageAfterCallWorkSeconds,omitempty"`
		*Alias
	}{ 
		Offered: u.Offered,
		
		AverageTalkTimeSeconds: u.AverageTalkTimeSeconds,
		
		AverageAfterCallWorkSeconds: u.AverageAfterCallWorkSeconds,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradayforecastdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
