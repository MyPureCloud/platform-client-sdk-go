package platformclientv2
import (
	"encoding/json"
)

// Wfmtimezone - Workforce management time zone
type Wfmtimezone struct { 
	// Id - The Olson format time zone ID (see https://en.wikipedia.org/wiki/Tz_database)
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmtimezone) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
