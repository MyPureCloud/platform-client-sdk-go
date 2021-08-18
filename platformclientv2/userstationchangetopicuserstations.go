package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userstationchangetopicuserstations
type Userstationchangetopicuserstations struct { 
	// AssociatedStation
	AssociatedStation *Userstationchangetopicuserstation `json:"associatedStation,omitempty"`

}

func (u *Userstationchangetopicuserstations) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userstationchangetopicuserstations

	

	return json.Marshal(&struct { 
		AssociatedStation *Userstationchangetopicuserstation `json:"associatedStation,omitempty"`
		*Alias
	}{ 
		AssociatedStation: u.AssociatedStation,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userstationchangetopicuserstations) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
