package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Userconversationseventmediasummarydetail) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userconversationseventmediasummarydetail

	

	return json.Marshal(&struct { 
		Active *int `json:"active,omitempty"`
		
		Acw *int `json:"acw,omitempty"`
		*Alias
	}{ 
		Active: u.Active,
		
		Acw: u.Acw,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userconversationseventmediasummarydetail) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
