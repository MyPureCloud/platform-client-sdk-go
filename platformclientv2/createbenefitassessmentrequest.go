package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createbenefitassessmentrequest
type Createbenefitassessmentrequest struct { 
	// QueueIds - The list of queue ids that are to be assessed for Predictive Routing benefit.
	QueueIds *[]string `json:"queueIds,omitempty"`

}

func (u *Createbenefitassessmentrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createbenefitassessmentrequest

	

	return json.Marshal(&struct { 
		QueueIds *[]string `json:"queueIds,omitempty"`
		*Alias
	}{ 
		QueueIds: u.QueueIds,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Createbenefitassessmentrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
