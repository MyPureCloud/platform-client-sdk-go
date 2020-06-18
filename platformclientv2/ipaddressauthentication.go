package platformclientv2
import (
	"encoding/json"
)

// Ipaddressauthentication
type Ipaddressauthentication struct { 
	// NetworkWhitelist
	NetworkWhitelist *[]string `json:"networkWhitelist,omitempty"`

}

// String returns a JSON representation of the model
func (o *Ipaddressauthentication) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
