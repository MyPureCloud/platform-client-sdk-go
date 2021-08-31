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

func (o *Userroutingstatususerroutingstatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userroutingstatususerroutingstatus
	
	return json.Marshal(&struct { 
		RoutingStatus *Userroutingstatusroutingstatus `json:"routingStatus,omitempty"`
		
		ErrorInfo *Userroutingstatuserrorinfo `json:"errorInfo,omitempty"`
		*Alias
	}{ 
		RoutingStatus: o.RoutingStatus,
		
		ErrorInfo: o.ErrorInfo,
		Alias:    (*Alias)(o),
	})
}

func (o *Userroutingstatususerroutingstatus) UnmarshalJSON(b []byte) error {
	var UserroutingstatususerroutingstatusMap map[string]interface{}
	err := json.Unmarshal(b, &UserroutingstatususerroutingstatusMap)
	if err != nil {
		return err
	}
	
	if RoutingStatus, ok := UserroutingstatususerroutingstatusMap["routingStatus"].(map[string]interface{}); ok {
		RoutingStatusString, _ := json.Marshal(RoutingStatus)
		json.Unmarshal(RoutingStatusString, &o.RoutingStatus)
	}
	
	if ErrorInfo, ok := UserroutingstatususerroutingstatusMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userroutingstatususerroutingstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
