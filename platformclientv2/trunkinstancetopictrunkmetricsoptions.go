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

func (u *Trunkinstancetopictrunkmetricsoptions) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkinstancetopictrunkmetricsoptions

	
	OptionStateTime := new(string)
	if u.OptionStateTime != nil {
		
		*OptionStateTime = timeutil.Strftime(u.OptionStateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		ProxyAddress: u.ProxyAddress,
		
		OptionState: u.OptionState,
		
		OptionStateTime: OptionStateTime,
		
		ErrorInfo: u.ErrorInfo,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Trunkinstancetopictrunkmetricsoptions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
