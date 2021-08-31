package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgemetricstopicedgemetricsubsystemmedia
type Edgemetricstopicedgemetricsubsystemmedia struct { 
	// ProcessName
	ProcessName *string `json:"processName,omitempty"`


	// DelayMs
	DelayMs *int `json:"delayMs,omitempty"`

}

func (o *Edgemetricstopicedgemetricsubsystemmedia) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgemetricstopicedgemetricsubsystemmedia
	
	return json.Marshal(&struct { 
		ProcessName *string `json:"processName,omitempty"`
		
		DelayMs *int `json:"delayMs,omitempty"`
		*Alias
	}{ 
		ProcessName: o.ProcessName,
		
		DelayMs: o.DelayMs,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgemetricstopicedgemetricsubsystemmedia) UnmarshalJSON(b []byte) error {
	var EdgemetricstopicedgemetricsubsystemmediaMap map[string]interface{}
	err := json.Unmarshal(b, &EdgemetricstopicedgemetricsubsystemmediaMap)
	if err != nil {
		return err
	}
	
	if ProcessName, ok := EdgemetricstopicedgemetricsubsystemmediaMap["processName"].(string); ok {
		o.ProcessName = &ProcessName
	}
	
	if DelayMs, ok := EdgemetricstopicedgemetricsubsystemmediaMap["delayMs"].(float64); ok {
		DelayMsInt := int(DelayMs)
		o.DelayMs = &DelayMsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Edgemetricstopicedgemetricsubsystemmedia) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
