package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeoffrequestquerybody
type Timeoffrequestquerybody struct { 
	// UserIds - The set of user ids to filter time off requests
	UserIds *[]string `json:"userIds,omitempty"`


	// Statuses - The set of statuses to filter time off requests
	Statuses *[]string `json:"statuses,omitempty"`


	// DateRange - The inclusive range of dates to filter time off requests
	DateRange *Daterange `json:"dateRange,omitempty"`

}

func (u *Timeoffrequestquerybody) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Timeoffrequestquerybody

	

	return json.Marshal(&struct { 
		UserIds *[]string `json:"userIds,omitempty"`
		
		Statuses *[]string `json:"statuses,omitempty"`
		
		DateRange *Daterange `json:"dateRange,omitempty"`
		*Alias
	}{ 
		UserIds: u.UserIds,
		
		Statuses: u.Statuses,
		
		DateRange: u.DateRange,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Timeoffrequestquerybody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
