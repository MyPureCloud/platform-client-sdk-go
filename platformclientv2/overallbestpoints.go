package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Overallbestpoints
type Overallbestpoints struct { 
	// Division - The requested division
	Division *Division `json:"division,omitempty"`


	// BestPoints - List of gamification best point items
	BestPoints *[]Overallbestpointsitem `json:"bestPoints,omitempty"`

}

func (u *Overallbestpoints) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Overallbestpoints

	

	return json.Marshal(&struct { 
		Division *Division `json:"division,omitempty"`
		
		BestPoints *[]Overallbestpointsitem `json:"bestPoints,omitempty"`
		*Alias
	}{ 
		Division: u.Division,
		
		BestPoints: u.BestPoints,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Overallbestpoints) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
