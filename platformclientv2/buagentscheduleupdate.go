package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Buagentscheduleupdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Buagentscheduleupdate

	

	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		ShiftStartDates *[]time.Time `json:"shiftStartDates,omitempty"`
		*Alias
	}{ 
		VarType: u.VarType,
		
		ShiftStartDates: u.ShiftStartDates,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Buagentscheduleupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
