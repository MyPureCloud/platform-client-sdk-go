package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmintradaydataupdatetopicintradaymetric
type Wfmintradaydataupdatetopicintradaymetric struct { 
	// Category
	Category *string `json:"category,omitempty"`


	// Version
	Version *string `json:"version,omitempty"`

}

func (u *Wfmintradaydataupdatetopicintradaymetric) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmintradaydataupdatetopicintradaymetric

	

	return json.Marshal(&struct { 
		Category *string `json:"category,omitempty"`
		
		Version *string `json:"version,omitempty"`
		*Alias
	}{ 
		Category: u.Category,
		
		Version: u.Version,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmintradaydataupdatetopicintradaymetric) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
