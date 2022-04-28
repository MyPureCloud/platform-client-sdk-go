package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userroutingstatususerroutingstatus
type Userroutingstatususerroutingstatus struct { 
	// Id - The unique identifier of the user.
	Id *string `json:"id,omitempty"`


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
		Id *string `json:"id,omitempty"`
		
		RoutingStatus *Userroutingstatusroutingstatus `json:"routingStatus,omitempty"`
		
		ErrorInfo *Userroutingstatuserrorinfo `json:"errorInfo,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
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
	
	if Id, ok := UserroutingstatususerroutingstatusMap["id"].(string); ok {
		o.Id = &Id
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
