package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Userroutingstatususerroutingstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userroutingstatususerroutingstatus

	

	return json.Marshal(&struct { 
		RoutingStatus *Userroutingstatusroutingstatus `json:"routingStatus,omitempty"`
		
		ErrorInfo *Userroutingstatuserrorinfo `json:"errorInfo,omitempty"`
		*Alias
	}{ 
		RoutingStatus: u.RoutingStatus,
		
		ErrorInfo: u.ErrorInfo,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userroutingstatususerroutingstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
