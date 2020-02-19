package platformclientv2
import (
	"encoding/json"
)

// Faxsummary
type Faxsummary struct { 
	// ReadCount
	ReadCount *int32 `json:"readCount,omitempty"`


	// UnreadCount
	UnreadCount *int32 `json:"unreadCount,omitempty"`


	// TotalCount
	TotalCount *int32 `json:"totalCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Faxsummary) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
