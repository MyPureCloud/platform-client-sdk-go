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

func (o *Trunkinstancetopictrunkmetricsregisters) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkinstancetopictrunkmetricsregisters
	
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
		
		ErrorInfo *Trunkinstancetopictrunkerrorinfo `json:"errorInfo,omitempty"`
		*Alias
	}{ 
		ProxyAddress: o.ProxyAddress,
		
		RegisterState: o.RegisterState,
		
		RegisterStateTime: RegisterStateTime,
		
		ErrorInfo: o.ErrorInfo,
		Alias:    (*Alias)(o),
	})
}

func (o *Trunkinstancetopictrunkmetricsregisters) UnmarshalJSON(b []byte) error {
	var TrunkinstancetopictrunkmetricsregistersMap map[string]interface{}
	err := json.Unmarshal(b, &TrunkinstancetopictrunkmetricsregistersMap)
	if err != nil {
		return err
	}
	
	if ProxyAddress, ok := TrunkinstancetopictrunkmetricsregistersMap["proxyAddress"].(string); ok {
		o.ProxyAddress = &ProxyAddress
	}
	
	if RegisterState, ok := TrunkinstancetopictrunkmetricsregistersMap["registerState"].(bool); ok {
		o.RegisterState = &RegisterState
	}
	
	if registerStateTimeString, ok := TrunkinstancetopictrunkmetricsregistersMap["registerStateTime"].(string); ok {
		RegisterStateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", registerStateTimeString)
		o.RegisterStateTime = &RegisterStateTime
	}
	
	if ErrorInfo, ok := TrunkinstancetopictrunkmetricsregistersMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trunkinstancetopictrunkmetricsregisters) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
