package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Domainnetworkcommandresponse
type Domainnetworkcommandresponse struct { 
	// CorrelationId
	CorrelationId *string `json:"correlationId,omitempty"`


	// CommandName
	CommandName *string `json:"commandName,omitempty"`


	// Acknowledged
	Acknowledged *bool `json:"acknowledged,omitempty"`


	// ErrorInfo
	ErrorInfo **Errordetails `json:"errorInfo,omitempty"`

}

func (u *Domainnetworkcommandresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Domainnetworkcommandresponse

	

	return json.Marshal(&struct { 
		CorrelationId *string `json:"correlationId,omitempty"`
		
		CommandName *string `json:"commandName,omitempty"`
		
		Acknowledged *bool `json:"acknowledged,omitempty"`
		
		ErrorInfo **Errordetails `json:"errorInfo,omitempty"`
		*Alias
	}{ 
		CorrelationId: u.CorrelationId,
		
		CommandName: u.CommandName,
		
		Acknowledged: u.Acknowledged,
		
		ErrorInfo: u.ErrorInfo,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Domainnetworkcommandresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
