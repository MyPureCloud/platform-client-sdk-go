package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Mergeoperation
type Mergeoperation struct { 
	// SourceContact - The source contact for the merge operation
	SourceContact *Addressableentityref `json:"sourceContact,omitempty"`


	// TargetContact - The target contact for the merge operation
	TargetContact *Addressableentityref `json:"targetContact,omitempty"`


	// ResultingContact - The contact created as a result of the merge operation
	ResultingContact *Addressableentityref `json:"resultingContact,omitempty"`

}

func (u *Mergeoperation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Mergeoperation

	

	return json.Marshal(&struct { 
		SourceContact *Addressableentityref `json:"sourceContact,omitempty"`
		
		TargetContact *Addressableentityref `json:"targetContact,omitempty"`
		
		ResultingContact *Addressableentityref `json:"resultingContact,omitempty"`
		*Alias
	}{ 
		SourceContact: u.SourceContact,
		
		TargetContact: u.TargetContact,
		
		ResultingContact: u.ResultingContact,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Mergeoperation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
