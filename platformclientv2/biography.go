package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Biography
type Biography struct { 
	// Biography - Personal detailed description
	Biography *string `json:"biography,omitempty"`


	// Interests
	Interests *[]string `json:"interests,omitempty"`


	// Hobbies
	Hobbies *[]string `json:"hobbies,omitempty"`


	// Spouse
	Spouse *string `json:"spouse,omitempty"`


	// Education - User education details
	Education *[]Education `json:"education,omitempty"`

}

func (u *Biography) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Biography

	

	return json.Marshal(&struct { 
		Biography *string `json:"biography,omitempty"`
		
		Interests *[]string `json:"interests,omitempty"`
		
		Hobbies *[]string `json:"hobbies,omitempty"`
		
		Spouse *string `json:"spouse,omitempty"`
		
		Education *[]Education `json:"education,omitempty"`
		*Alias
	}{ 
		Biography: u.Biography,
		
		Interests: u.Interests,
		
		Hobbies: u.Hobbies,
		
		Spouse: u.Spouse,
		
		Education: u.Education,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Biography) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
