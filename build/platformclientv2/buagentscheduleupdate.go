package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Buagentscheduleupdate
type Buagentscheduleupdate struct { 
	// VarType - The type of update
	VarType *string `json:"type,omitempty"`


	// ShiftStartDates - The start date for the affected shifts
	ShiftStartDates *[]time.Time `json:"shiftStartDates,omitempty"`

}

// String returns a JSON representation of the model
func (o *Buagentscheduleupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
