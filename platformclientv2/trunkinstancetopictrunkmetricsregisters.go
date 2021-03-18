package platformclientv2
import (
	"time"
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

// String returns a JSON representation of the model
func (o *Trunkinstancetopictrunkmetricsregisters) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
