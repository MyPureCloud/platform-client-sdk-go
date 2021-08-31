package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkmetricsoptions
type Trunkmetricsoptions struct { 
	// ProxyAddress - Server proxy address that this options array element represents.
	ProxyAddress *string `json:"proxyAddress,omitempty"`


	// OptionState
	OptionState *bool `json:"optionState,omitempty"`


	// OptionStateTime - ISO 8601 format UTC absolute date & time of the last change of the option state.
	OptionStateTime *time.Time `json:"optionStateTime,omitempty"`


	// ErrorInfo
	ErrorInfo *Trunkerrorinfo `json:"errorInfo,omitempty"`

}

func (o *Trunkmetricsoptions) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkmetricsoptions
	
	OptionStateTime := new(string)
	if o.OptionStateTime != nil {
		
		*OptionStateTime = timeutil.Strftime(o.OptionStateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		OptionStateTime = nil
	}
	
	return json.Marshal(&struct { 
		ProxyAddress *string `json:"proxyAddress,omitempty"`
		
		OptionState *bool `json:"optionState,omitempty"`
		
		OptionStateTime *string `json:"optionStateTime,omitempty"`
		
		ErrorInfo *Trunkerrorinfo `json:"errorInfo,omitempty"`
		*Alias
	}{ 
		ProxyAddress: o.ProxyAddress,
		
		OptionState: o.OptionState,
		
		OptionStateTime: OptionStateTime,
		
		ErrorInfo: o.ErrorInfo,
		Alias:    (*Alias)(o),
	})
}

func (o *Trunkmetricsoptions) UnmarshalJSON(b []byte) error {
	var TrunkmetricsoptionsMap map[string]interface{}
	err := json.Unmarshal(b, &TrunkmetricsoptionsMap)
	if err != nil {
		return err
	}
	
	if ProxyAddress, ok := TrunkmetricsoptionsMap["proxyAddress"].(string); ok {
		o.ProxyAddress = &ProxyAddress
	}
	
	if OptionState, ok := TrunkmetricsoptionsMap["optionState"].(bool); ok {
		o.OptionState = &OptionState
	}
	
	if optionStateTimeString, ok := TrunkmetricsoptionsMap["optionStateTime"].(string); ok {
		OptionStateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", optionStateTimeString)
		o.OptionStateTime = &OptionStateTime
	}
	
	if ErrorInfo, ok := TrunkmetricsoptionsMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trunkmetricsoptions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
