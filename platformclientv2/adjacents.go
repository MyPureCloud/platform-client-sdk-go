package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Adjacents
type Adjacents struct { 
	// Superiors
	Superiors *[]User `json:"superiors,omitempty"`


	// Siblings
	Siblings *[]User `json:"siblings,omitempty"`


	// DirectReports
	DirectReports *[]User `json:"directReports,omitempty"`

}

func (u *Adjacents) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Adjacents

	

	return json.Marshal(&struct { 
		Superiors *[]User `json:"superiors,omitempty"`
		
		Siblings *[]User `json:"siblings,omitempty"`
		
		DirectReports *[]User `json:"directReports,omitempty"`
		*Alias
	}{ 
		Superiors: u.Superiors,
		
		Siblings: u.Siblings,
		
		DirectReports: u.DirectReports,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Adjacents) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
