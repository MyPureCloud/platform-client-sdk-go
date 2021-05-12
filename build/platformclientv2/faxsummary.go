package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
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
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
