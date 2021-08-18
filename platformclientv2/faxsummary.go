package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Faxsummary) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Faxsummary

	

	return json.Marshal(&struct { 
		ReadCount *int `json:"readCount,omitempty"`
		
		UnreadCount *int `json:"unreadCount,omitempty"`
		
		TotalCount *int `json:"totalCount,omitempty"`
		*Alias
	}{ 
		ReadCount: u.ReadCount,
		
		UnreadCount: u.UnreadCount,
		
		TotalCount: u.TotalCount,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Faxsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
