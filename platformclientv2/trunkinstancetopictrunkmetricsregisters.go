package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkinstancetopictrunkmetricsregisters
type Trunkinstancetopictrunkmetricsregisters struct { 
	// ProxyAddress
	ProxyAddress *string `json:"proxyAddress,omitempty"`


	// RegisterState
	RegisterState *bool `json:"registerState,omitempty"`


	// RegisterStateTime
	RegisterStateTime *time.Time `json:"registerStateTime,omitempty"`


	// ErrorInfo
	ErrorInfo *Trunkinstancetopictrunkerrorinfo `json:"errorInfo,omitempty"`

}

func (u *Trunkinstancetopictrunkmetricsregisters) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkinstancetopictrunkmetricsregisters

	
	RegisterStateTime := new(string)
	if u.RegisterStateTime != nil {
		
		*RegisterStateTime = timeutil.Strftime(u.RegisterStateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		RegisterStateTime = nil
	}
	

	return json.Marshal(&struct { 
		ProxyAddress *string `json:"proxyAddress,omitempty"`
		
		RegisterState *bool `json:"registerState,omitempty"`
		
		RegisterStateTime *string `json:"registerStateTime,omitempty"`
		
		ErrorInfo *Trunkinstancetopictrunkerrorinfo `json:"errorInfo,omitempty"`
		*Alias
	}{ 
		ProxyAddress: u.ProxyAddress,
		
		RegisterState: u.RegisterState,
		
		RegisterStateTime: RegisterStateTime,
		
		ErrorInfo: u.ErrorInfo,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Trunkinstancetopictrunkmetricsregisters) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
