package platformclientv2
import (
	"time"
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Trunkinstancetopictrunkmetricsoptions) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
