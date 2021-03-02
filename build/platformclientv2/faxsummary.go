package platformclientv2
import (
	"encoding/json"
)

// Faxsummary
type Faxsummary struct { 
	// ReadCount
	ReadCount *int `json:"readCount,omitempty"`


	// UnreadCount
	UnreadCount *int `json:"unreadCount,omitempty"`


	// TotalCount
	TotalCount *int `json:"totalCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Faxsummary) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
