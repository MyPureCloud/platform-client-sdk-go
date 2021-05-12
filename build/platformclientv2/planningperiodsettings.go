package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Planningperiodsettings
type Planningperiodsettings struct { 
	// WeekCount - Planning period length in weeks
	WeekCount *int `json:"weekCount,omitempty"`


	// StartDate - Start date of the planning period in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	StartDate *time.Time `json:"startDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Planningperiodsettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
