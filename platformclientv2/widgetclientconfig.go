package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Widgetclientconfig
type Widgetclientconfig struct { 
	// V1
	V1 *Widgetclientconfigv1 `json:"v1,omitempty"`


	// V2
	V2 *Widgetclientconfigv2 `json:"v2,omitempty"`


	// V1Http
	V1Http *Widgetclientconfigv1http `json:"v1-http,omitempty"`


	// ThirdParty
	ThirdParty *Widgetclientconfigthirdparty `json:"third-party,omitempty"`

}

func (u *Widgetclientconfig) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Widgetclientconfig

	

	return json.Marshal(&struct { 
		V1 *Widgetclientconfigv1 `json:"v1,omitempty"`
		
		V2 *Widgetclientconfigv2 `json:"v2,omitempty"`
		
		V1Http *Widgetclientconfigv1http `json:"v1-http,omitempty"`
		
		ThirdParty *Widgetclientconfigthirdparty `json:"third-party,omitempty"`
		*Alias
	}{ 
		V1: u.V1,
		
		V2: u.V2,
		
		V1Http: u.V1Http,
		
		ThirdParty: u.ThirdParty,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Widgetclientconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
