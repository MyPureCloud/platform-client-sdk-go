package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Deleteretention
type Deleteretention struct { 
	// Days
	Days *int `json:"days,omitempty"`

}

// String returns a JSON representation of the model
func (o *Deleteretention) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
