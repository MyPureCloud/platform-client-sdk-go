package platformclientv2
import (
	"time"
	"encoding/json"
)

// Edgeautoupdateconfig
type Edgeautoupdateconfig struct { 
	// TimeZone - The timezone of the window in which any updates to the edges assigned to the site can be applied. The minimum size of the window is 2 hours.
	TimeZone *string `json:"timeZone,omitempty"`


	// Rrule - The recurrence rule for updating the Edges assigned to the site. The only supported frequencies are daily and weekly. Weekly frequencies require a day list with at least oneday specified. All other configurations are not supported.
	Rrule *string `json:"rrule,omitempty"`


	// Start - Date time is represented as an ISO-8601 string without a timezone. For example: yyyy-MM-ddTHH:mm:ss.SSS
	Start *time.Time `json:"start,omitempty"`


	// End - Date time is represented as an ISO-8601 string without a timezone. For example: yyyy-MM-ddTHH:mm:ss.SSS
	End *time.Time `json:"end,omitempty"`

}

// String returns a JSON representation of the model
func (o *Edgeautoupdateconfig) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
