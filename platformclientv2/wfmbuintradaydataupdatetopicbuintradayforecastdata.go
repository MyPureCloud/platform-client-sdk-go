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

func (u *Wfmbuintradaydataupdatetopicbuintradayforecastdata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuintradaydataupdatetopicbuintradayforecastdata

	

	return json.Marshal(&struct { 
		Offered *float32 `json:"offered,omitempty"`
		
		AverageHandleTimeSeconds *float32 `json:"averageHandleTimeSeconds,omitempty"`
		*Alias
	}{ 
		Offered: u.Offered,
		
		AverageHandleTimeSeconds: u.AverageHandleTimeSeconds,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmbuintradaydataupdatetopicbuintradayforecastdata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
