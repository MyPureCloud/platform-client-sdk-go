package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Searchshifttradesrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Searchshifttradesrequest

	

	return json.Marshal(&struct { 
		ReceivingScheduleId *string `json:"receivingScheduleId,omitempty"`
		
		ReceivingShiftIds *[]string `json:"receivingShiftIds,omitempty"`
		*Alias
	}{ 
		ReceivingScheduleId: u.ReceivingScheduleId,
		
		ReceivingShiftIds: u.ReceivingShiftIds,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Searchshifttradesrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
