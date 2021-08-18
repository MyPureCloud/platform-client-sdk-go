package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Businessunit
type Businessunit struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Settings - Settings for this business unit
	Settings *Businessunitsettings `json:"settings,omitempty"`


	// Division - The division to which this entity belongs.
	Division *Divisionreference `json:"division,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Businessunit) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Businessunit

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Settings *Businessunitsettings `json:"settings,omitempty"`
		
		Division *Divisionreference `json:"division,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Settings: u.Settings,
		
		Division: u.Division,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Businessunit) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
