package platformclientv2
import (
	"time"
	"encoding/json"
)

// Channel
type Channel struct { 
	// ConnectUri
	ConnectUri *string `json:"connectUri,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// Expires - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	Expires *time.Time `json:"expires,omitempty"`

}

// String returns a JSON representation of the model
func (o *Channel) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
