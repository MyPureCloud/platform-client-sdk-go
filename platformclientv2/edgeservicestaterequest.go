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

func (u *Edgeservicestaterequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Edgeservicestaterequest

	

	return json.Marshal(&struct { 
		InService *bool `json:"inService,omitempty"`
		
		CallDrainingWaitTimeSeconds *int `json:"callDrainingWaitTimeSeconds,omitempty"`
		*Alias
	}{ 
		InService: u.InService,
		
		CallDrainingWaitTimeSeconds: u.CallDrainingWaitTimeSeconds,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Edgeservicestaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
