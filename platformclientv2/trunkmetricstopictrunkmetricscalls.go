package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkmetricstopictrunkmetricscalls
type Trunkmetricstopictrunkmetricscalls struct { 
	// InboundCallCount
	InboundCallCount *int `json:"inboundCallCount,omitempty"`


	// OutboundCallCount
	OutboundCallCount *int `json:"outboundCallCount,omitempty"`

}

func (u *Trunkmetricstopictrunkmetricscalls) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkmetricstopictrunkmetricscalls

	

	return json.Marshal(&struct { 
		InboundCallCount *int `json:"inboundCallCount,omitempty"`
		
		OutboundCallCount *int `json:"outboundCallCount,omitempty"`
		*Alias
	}{ 
		InboundCallCount: u.InboundCallCount,
		
		OutboundCallCount: u.OutboundCallCount,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Trunkmetricstopictrunkmetricscalls) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
