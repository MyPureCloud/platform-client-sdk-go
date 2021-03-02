package platformclientv2
import (
	"encoding/json"
)

// Lineuserid - Channel-specific User ID for Line accounts
type Lineuserid struct { 
	// UserId - The unique channel-specific userId for the user
	UserId *string `json:"userId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Lineuserid) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
