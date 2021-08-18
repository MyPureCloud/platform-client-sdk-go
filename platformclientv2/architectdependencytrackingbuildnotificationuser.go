package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectdependencytrackingbuildnotificationuser
type Architectdependencytrackingbuildnotificationuser struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// HomeOrg
	HomeOrg *Architectdependencytrackingbuildnotificationhomeorganization `json:"homeOrg,omitempty"`

}

func (u *Architectdependencytrackingbuildnotificationuser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectdependencytrackingbuildnotificationuser

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		HomeOrg *Architectdependencytrackingbuildnotificationhomeorganization `json:"homeOrg,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		HomeOrg: u.HomeOrg,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Architectdependencytrackingbuildnotificationuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
