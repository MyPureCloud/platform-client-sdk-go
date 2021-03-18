package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdatenotification
type Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdatenotification struct { 
	// Status
	Status *string `json:"status,omitempty"`


	// OperationId
	OperationId *string `json:"operationId,omitempty"`


	// Result
	Result *Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdateresultlisting `json:"result,omitempty"`


	// WeekDate
	WeekDate *Wfmbulkshifttradestateupdatenotificationtopiclocaldate `json:"weekDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmbulkshifttradestateupdatenotificationtopicbulkshifttradestateupdatenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
