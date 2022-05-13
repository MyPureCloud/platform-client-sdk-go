package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkinstancetopictrunkmetricsoptions
type Trunkinstancetopictrunkmetricsoptions struct { 
	// ProxyAddress
	ProxyAddress *string `json:"proxyAddress,omitempty"`


	// OptionState
	OptionState *bool `json:"optionState,omitempty"`


	// OptionStateTime
	OptionStateTime *time.Time `json:"optionStateTime,omitempty"`


	// ErrorInfo
	ErrorInfo *Trunkinstancetopictrunkerrorinfo `json:"errorInfo,omitempty"`

}

func (o *Trunkinstancetopictrunkmetricsoptions) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkinstancetopictrunkmetricsoptions
	
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
		
		ErrorInfo *Trunkinstancetopictrunkerrorinfo `json:"errorInfo,omitempty"`
		*Alias
	}{ 
		ProxyAddress: o.ProxyAddress,
		
		OptionState: o.OptionState,
		
		OptionStateTime: OptionStateTime,
		
		ErrorInfo: o.ErrorInfo,
		Alias:    (*Alias)(o),
	})
}

func (o *Trunkinstancetopictrunkmetricsoptions) UnmarshalJSON(b []byte) error {
	var TrunkinstancetopictrunkmetricsoptionsMap map[string]interface{}
	err := json.Unmarshal(b, &TrunkinstancetopictrunkmetricsoptionsMap)
	if err != nil {
		return err
	}
	
	if ProxyAddress, ok := TrunkinstancetopictrunkmetricsoptionsMap["proxyAddress"].(string); ok {
		o.ProxyAddress = &ProxyAddress
	}
    
	if OptionState, ok := TrunkinstancetopictrunkmetricsoptionsMap["optionState"].(bool); ok {
		o.OptionState = &OptionState
	}
    
	if optionStateTimeString, ok := TrunkinstancetopictrunkmetricsoptionsMap["optionStateTime"].(string); ok {
		OptionStateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", optionStateTimeString)
		o.OptionStateTime = &OptionStateTime
	}
	
	if ErrorInfo, ok := TrunkinstancetopictrunkmetricsoptionsMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Trunkinstancetopictrunkmetricsoptions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
