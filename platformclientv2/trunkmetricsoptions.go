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

func (u *Trunkmetricsoptions) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Trunkmetricsoptions

	
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
		
		ErrorInfo *Trunkerrorinfo `json:"errorInfo,omitempty"`
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
func (o *Trunkmetricsoptions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
