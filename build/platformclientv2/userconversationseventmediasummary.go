package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Userconversationseventmediasummary
type Userconversationseventmediasummary struct { 
	// ContactCenter
	ContactCenter *Userconversationseventmediasummarydetail `json:"contactCenter,omitempty"`


	// Enterprise
	Enterprise *Userconversationseventmediasummarydetail `json:"enterprise,omitempty"`

}

// String returns a JSON representation of the model
func (o *Userconversationseventmediasummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
