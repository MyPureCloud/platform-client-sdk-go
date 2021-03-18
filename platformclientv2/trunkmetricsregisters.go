package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Trunkmetricsregisters
type Trunkmetricsregisters struct { 
	// ProxyAddress - Server proxy address that this registers array element represents.
	ProxyAddress *string `json:"proxyAddress,omitempty"`


	// RegisterState - True if last REGISTER message had positive response; false if error response or no response.
	RegisterState *bool `json:"registerState,omitempty"`


	// RegisterStateTime - ISO 8601 format UTC absolute date & time of the last change of the register state.
	RegisterStateTime *time.Time `json:"registerStateTime,omitempty"`


	// ErrorInfo
	ErrorInfo *Trunkerrorinfo `json:"errorInfo,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trunkmetricsregisters) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
