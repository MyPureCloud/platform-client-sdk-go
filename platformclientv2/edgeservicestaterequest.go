package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Edgeservicestaterequest
type Edgeservicestaterequest struct { 
	// InService - A boolean that sets the Edge in-service or out-of-service.
	InService *bool `json:"inService,omitempty"`


	// CallDrainingWaitTimeSeconds - The number of seconds to wait for call draining to complete before initiating the reboot. A value of 0 will prevent call draining and all calls will disconnect immediately.
	CallDrainingWaitTimeSeconds *int `json:"callDrainingWaitTimeSeconds,omitempty"`

}

func (o *Edgeservicestaterequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgeservicestaterequest
	
	return json.Marshal(&struct { 
		InService *bool `json:"inService,omitempty"`
		
		CallDrainingWaitTimeSeconds *int `json:"callDrainingWaitTimeSeconds,omitempty"`
		*Alias
	}{ 
		InService: o.InService,
		
		CallDrainingWaitTimeSeconds: o.CallDrainingWaitTimeSeconds,
		Alias:    (*Alias)(o),
	})
}

func (o *Edgeservicestaterequest) UnmarshalJSON(b []byte) error {
	var EdgeservicestaterequestMap map[string]interface{}
	err := json.Unmarshal(b, &EdgeservicestaterequestMap)
	if err != nil {
		return err
	}
	
	if InService, ok := EdgeservicestaterequestMap["inService"].(bool); ok {
		o.InService = &InService
	}
    
	if CallDrainingWaitTimeSeconds, ok := EdgeservicestaterequestMap["callDrainingWaitTimeSeconds"].(float64); ok {
		CallDrainingWaitTimeSecondsInt := int(CallDrainingWaitTimeSeconds)
		o.CallDrainingWaitTimeSeconds = &CallDrainingWaitTimeSecondsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Edgeservicestaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
