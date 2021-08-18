package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Validateworkplanresponse
type Validateworkplanresponse struct { 
	// WorkPlan - The work plan reference associated with this response
	WorkPlan *Workplanreference `json:"workPlan,omitempty"`


	// Valid - Whether the work plan is valid or not
	Valid *bool `json:"valid,omitempty"`


	// Messages - Validation messages for this work plan
	Messages *Validateworkplanmessages `json:"messages,omitempty"`

}

func (u *Validateworkplanresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Validateworkplanresponse

	

	return json.Marshal(&struct { 
		WorkPlan *Workplanreference `json:"workPlan,omitempty"`
		
		Valid *bool `json:"valid,omitempty"`
		
		Messages *Validateworkplanmessages `json:"messages,omitempty"`
		*Alias
	}{ 
		WorkPlan: u.WorkPlan,
		
		Valid: u.Valid,
		
		Messages: u.Messages,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Validateworkplanresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
