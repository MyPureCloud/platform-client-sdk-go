package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Userroutingstatususerroutingstatus
type Userroutingstatususerroutingstatus struct { 
	// RoutingStatus
	RoutingStatus *Userroutingstatusroutingstatus `json:"routingStatus,omitempty"`


	// ErrorInfo
	ErrorInfo *Userroutingstatuserrorinfo `json:"errorInfo,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userroutingstatususerroutingstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
