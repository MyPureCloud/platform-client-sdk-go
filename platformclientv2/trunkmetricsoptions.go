package platformclientv2
import (
	"time"
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Trunkmetricsoptions) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
