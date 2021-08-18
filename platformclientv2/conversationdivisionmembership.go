package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationdivisionmembership
type Conversationdivisionmembership struct { 
	// Division - A division the conversation belongs to.
	Division *Domainentityref `json:"division,omitempty"`


	// Entities - The entities on the conversation within the division. These are the users, queues, work flows, etc. that can be on conversations and and be assigned to different divisions.
	Entities *[]Domainentityref `json:"entities,omitempty"`

}

func (u *Conversationdivisionmembership) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationdivisionmembership

	

	return json.Marshal(&struct { 
		Division *Domainentityref `json:"division,omitempty"`
		
		Entities *[]Domainentityref `json:"entities,omitempty"`
		*Alias
	}{ 
		Division: u.Division,
		
		Entities: u.Entities,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Conversationdivisionmembership) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
