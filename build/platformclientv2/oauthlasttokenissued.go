package platformclientv2
import (
	"time"
	"encoding/json"
)

// Oauthlasttokenissued
type Oauthlasttokenissued struct { 
	// DateIssued - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateIssued *time.Time `json:"dateIssued,omitempty"`

}

// String returns a JSON representation of the model
func (o *Oauthlasttokenissued) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
