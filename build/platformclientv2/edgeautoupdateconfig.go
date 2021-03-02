package platformclientv2
import (
	"time"
	"encoding/json"
)

// Edgeautoupdateconfig
type Edgeautoupdateconfig struct { 
	// TimeZone
	TimeZone *string `json:"timeZone,omitempty"`


	// Rrule
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
