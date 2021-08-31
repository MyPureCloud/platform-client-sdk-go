package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkmetricsregisters
type Trunkmetricsregisters struct { 
	// ProxyAddress - Server proxy address that this registers array element represents.
	ProxyAddress *string `json:"proxyAddress,omitempty"`


	// RegisterState - True if last REGISTER message had positive response; false if error response or no response.
	RegisterState *bool `json:"registerState,omitempty"`


	// RegisterStateTime - ISO 8601 format UTC absolute date & time of the last change of the register state.
	RegisterStateTime *time.Time `json:"registerStateTime,omitempty"`


	// ErrorInfo
	ErrorInfo *Trunkerrorinfo `json:"errorInfo,omitempty"`

}

func (o *Trunkmetricsregisters) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkmetricsregisters
	
	RegisterStateTime := new(string)
	if o.RegisterStateTime != nil {
		
		*RegisterStateTime = timeutil.Strftime(o.RegisterStateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		RegisterStateTime = nil
	}
	
	return json.Marshal(&struct { 
		ProxyAddress *string `json:"proxyAddress,omitempty"`
		
		RegisterState *bool `json:"registerState,omitempty"`
		
		RegisterStateTime *string `json:"registerStateTime,omitempty"`
		
		ErrorInfo *Trunkerrorinfo `json:"errorInfo,omitempty"`
		*Alias
	}{ 
		ProxyAddress: o.ProxyAddress,
		
		RegisterState: o.RegisterState,
		
		RegisterStateTime: RegisterStateTime,
		
		ErrorInfo: o.ErrorInfo,
		Alias:    (*Alias)(o),
	})
}

func (o *Trunkmetricsregisters) UnmarshalJSON(b []byte) error {
	var TrunkmetricsregistersMap map[string]interface{}
	err := json.Unmarshal(b, &TrunkmetricsregistersMap)
	if err != nil {
		return err
	}
	
	if ProxyAddress, ok := TrunkmetricsregistersMap["proxyAddress"].(string); ok {
		o.ProxyAddress = &ProxyAddress
	}
	
	if RegisterState, ok := TrunkmetricsregistersMap["registerState"].(bool); ok {
		o.RegisterState = &RegisterState
	}
	
	if registerStateTimeString, ok := TrunkmetricsregistersMap["registerStateTime"].(string); ok {
		RegisterStateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", registerStateTimeString)
		o.RegisterStateTime = &RegisterStateTime
	}
	
	if ErrorInfo, ok := TrunkmetricsregistersMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trunkmetricsregisters) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
