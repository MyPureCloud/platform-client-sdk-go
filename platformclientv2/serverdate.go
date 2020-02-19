package platformclientv2
import (
	"time"
	"encoding/json"
)

// Serverdate
type Serverdate struct { 
	// CurrentDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	CurrentDate *time.Time `json:"currentDate,omitempty"`

}

// String returns a JSON representation of the model
func (o *Serverdate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
