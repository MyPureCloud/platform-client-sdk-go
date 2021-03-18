package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Searchshifttradesrequest
type Searchshifttradesrequest struct { 
	// ReceivingScheduleId - The ID of the schedule for which to search for available shift trades
	ReceivingScheduleId *string `json:"receivingScheduleId,omitempty"`


	// ReceivingShiftIds - The IDs of shifts that the receiving user would potentially be willing to trade. If empty, only returns one sided trades (pick up a shift)
	ReceivingShiftIds *[]string `json:"receivingShiftIds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Searchshifttradesrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
