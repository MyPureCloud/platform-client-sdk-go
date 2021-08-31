package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkmetricscalls
type Trunkmetricscalls struct { 
	// InboundCallCount
	InboundCallCount *int `json:"inboundCallCount,omitempty"`


	// OutboundCallCount
	OutboundCallCount *int `json:"outboundCallCount,omitempty"`

}

func (o *Trunkmetricscalls) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkmetricscalls
	
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

func (o *Trunkmetricscalls) UnmarshalJSON(b []byte) error {
	var TrunkmetricscallsMap map[string]interface{}
	err := json.Unmarshal(b, &TrunkmetricscallsMap)
	if err != nil {
		return err
	}
	
	if InboundCallCount, ok := TrunkmetricscallsMap["inboundCallCount"].(float64); ok {
		InboundCallCountInt := int(InboundCallCount)
		o.InboundCallCount = &InboundCallCountInt
	}
	
	if OutboundCallCount, ok := TrunkmetricscallsMap["outboundCallCount"].(float64); ok {
		OutboundCallCountInt := int(OutboundCallCount)
		o.OutboundCallCount = &OutboundCallCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trunkmetricscalls) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
