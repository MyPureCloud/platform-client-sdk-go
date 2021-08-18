package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Userconversationseventmediasummary) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userconversationseventmediasummary

	

	return json.Marshal(&struct { 
		ContactCenter *Userconversationseventmediasummarydetail `json:"contactCenter,omitempty"`
		
		Enterprise *Userconversationseventmediasummarydetail `json:"enterprise,omitempty"`
		*Alias
	}{ 
		ContactCenter: u.ContactCenter,
		
		Enterprise: u.Enterprise,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userconversationseventmediasummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
