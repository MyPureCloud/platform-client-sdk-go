package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userstationchangetopicuserstation
type Userstationchangetopicuserstation struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// AssociatedUser
	AssociatedUser *Userstationchangetopicuser `json:"associatedUser,omitempty"`

}

func (u *Userstationchangetopicuserstation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userstationchangetopicuserstation

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		AssociatedUser *Userstationchangetopicuser `json:"associatedUser,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		AssociatedUser: u.AssociatedUser,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userstationchangetopicuserstation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
