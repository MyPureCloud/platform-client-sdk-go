package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuscheduletopicbuschedulenotification
type Wfmbuscheduletopicbuschedulenotification struct { 
	// Status
	Status *string `json:"status,omitempty"`


	// OperationId
	OperationId *string `json:"operationId,omitempty"`


	// EventType
	EventType *string `json:"eventType,omitempty"`


	// Result
	Result *Wfmbuscheduletopicbuschedulemetadata `json:"result,omitempty"`

}

func (u *Wfmbuscheduletopicbuschedulenotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Wfmbuscheduletopicbuschedulenotification

	

	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		OperationId *string `json:"operationId,omitempty"`
		
		EventType *string `json:"eventType,omitempty"`
		
		Result *Wfmbuscheduletopicbuschedulemetadata `json:"result,omitempty"`
		*Alias
	}{ 
		Status: u.Status,
		
		OperationId: u.OperationId,
		
		EventType: u.EventType,
		
		Result: u.Result,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Wfmbuscheduletopicbuschedulenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
