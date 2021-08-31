package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgerebootparameters
type Edgerebootparameters struct { 
	// CallDrainingWaitTimeSeconds - The number of seconds to wait for call draining to complete before initiating the reboot. A value of 0 will prevent call draining and all calls will disconnect immediately.
	CallDrainingWaitTimeSeconds *int `json:"callDrainingWaitTimeSeconds,omitempty"`

}

func (o *Edgerebootparameters) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgerebootparameters
	
	return json.Marshal(&struct { 
		CallDrainingWaitTimeSeconds *int `json:"callDrainingWaitTimeSeconds,omitempty"`
		*Alias
	}{ 
		CallDrainingWaitTimeSeconds: o.CallDrainingWaitTimeSeconds,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgerebootparameters) UnmarshalJSON(b []byte) error {
	var EdgerebootparametersMap map[string]interface{}
	err := json.Unmarshal(b, &EdgerebootparametersMap)
	if err != nil {
		return err
	}
	
	if CallDrainingWaitTimeSeconds, ok := EdgerebootparametersMap["callDrainingWaitTimeSeconds"].(float64); ok {
		CallDrainingWaitTimeSecondsInt := int(CallDrainingWaitTimeSeconds)
		o.CallDrainingWaitTimeSeconds = &CallDrainingWaitTimeSecondsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Edgerebootparameters) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
