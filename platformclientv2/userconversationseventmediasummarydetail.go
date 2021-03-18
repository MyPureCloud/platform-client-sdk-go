package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Userconversationseventmediasummarydetail
type Userconversationseventmediasummarydetail struct { 
	// Active
	Active *int `json:"active,omitempty"`


	// Acw
	Acw *int `json:"acw,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userconversationseventmediasummarydetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
