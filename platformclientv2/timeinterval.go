package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Timeinterval
type Timeinterval struct { 
	// Months
	Months *int `json:"months,omitempty"`


	// Weeks
	Weeks *int `json:"weeks,omitempty"`


	// Days
	Days *int `json:"days,omitempty"`


	// Hours
	Hours *int `json:"hours,omitempty"`

}

// String returns a JSON representation of the model
func (o *Timeinterval) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
