package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Mediasummary
type Mediasummary struct { 
	// ContactCenter
	ContactCenter *Mediasummarydetail `json:"contactCenter,omitempty"`


	// Enterprise
	Enterprise *Mediasummarydetail `json:"enterprise,omitempty"`

}

func (u *Mediasummary) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Mediasummary

	

	return json.Marshal(&struct { 
		ContactCenter *Mediasummarydetail `json:"contactCenter,omitempty"`
		
		Enterprise *Mediasummarydetail `json:"enterprise,omitempty"`
		*Alias
	}{ 
		ContactCenter: u.ContactCenter,
		
		Enterprise: u.Enterprise,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Mediasummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
