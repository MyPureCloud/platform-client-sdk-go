package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userbestpoints
type Userbestpoints struct { 
	// User - The requested user for the best points
	User *Userreference `json:"user,omitempty"`


	// BestPoints - List of best point for the requested user
	BestPoints *[]Userbestpointsitem `json:"bestPoints,omitempty"`

}

func (u *Userbestpoints) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userbestpoints

	

	return json.Marshal(&struct { 
		User *Userreference `json:"user,omitempty"`
		
		BestPoints *[]Userbestpointsitem `json:"bestPoints,omitempty"`
		*Alias
	}{ 
		User: u.User,
		
		BestPoints: u.BestPoints,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userbestpoints) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
