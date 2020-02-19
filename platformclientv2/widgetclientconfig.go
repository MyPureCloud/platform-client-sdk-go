package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Widgetclientconfig) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
