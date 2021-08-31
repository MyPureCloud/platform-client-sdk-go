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

func (o *Trunkmetricstopictrunkmetricscalls) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkmetricstopictrunkmetricscalls
	
	return json.Marshal(&struct { 
		InboundCallCount *int `json:"inboundCallCount,omitempty"`
		
		OutboundCallCount *int `json:"outboundCallCount,omitempty"`
		*Alias
	}{ 
		InboundCallCount: o.InboundCallCount,
		
		OutboundCallCount: o.OutboundCallCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Trunkmetricstopictrunkmetricscalls) UnmarshalJSON(b []byte) error {
	var TrunkmetricstopictrunkmetricscallsMap map[string]interface{}
	err := json.Unmarshal(b, &TrunkmetricstopictrunkmetricscallsMap)
	if err != nil {
		return err
	}
	
	if InboundCallCount, ok := TrunkmetricstopictrunkmetricscallsMap["inboundCallCount"].(float64); ok {
		InboundCallCountInt := int(InboundCallCount)
		o.InboundCallCount = &InboundCallCountInt
	}
	
	if OutboundCallCount, ok := TrunkmetricstopictrunkmetricscallsMap["outboundCallCount"].(float64); ok {
		OutboundCallCountInt := int(OutboundCallCount)
		o.OutboundCallCount = &OutboundCallCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trunkmetricstopictrunkmetricscalls) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
