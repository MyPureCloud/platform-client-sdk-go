package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Phonechangetopicphonestatus
type Phonechangetopicphonestatus struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// OperationalStatus
	OperationalStatus *string `json:"operationalStatus,omitempty"`


	// Edge
	Edge *Phonechangetopicedgereference `json:"edge,omitempty"`


	// Provision
	Provision *Phonechangetopicprovisioninfo `json:"provision,omitempty"`


	// LineStatuses
	LineStatuses *[]Phonechangetopiclinestatus `json:"lineStatuses,omitempty"`

}

func (u *Phonechangetopicphonestatus) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Phonechangetopicphonestatus

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		OperationalStatus *string `json:"operationalStatus,omitempty"`
		
		Edge *Phonechangetopicedgereference `json:"edge,omitempty"`
		
		Provision *Phonechangetopicprovisioninfo `json:"provision,omitempty"`
		
		LineStatuses *[]Phonechangetopiclinestatus `json:"lineStatuses,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		OperationalStatus: u.OperationalStatus,
		
		Edge: u.Edge,
		
		Provision: u.Provision,
		
		LineStatuses: u.LineStatuses,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Phonechangetopicphonestatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
