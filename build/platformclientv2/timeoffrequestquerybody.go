package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Timeoffrequestquerybody) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
