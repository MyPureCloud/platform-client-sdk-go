package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createbusinessunitrequest
type Createbusinessunitrequest struct { 
	// Name - The name of the business unit
	Name *string `json:"name,omitempty"`


	// DivisionId - The ID of the division to which the business unit should be added
	DivisionId *string `json:"divisionId,omitempty"`


	// Settings - Configuration for the business unit
	Settings *Createbusinessunitsettings `json:"settings,omitempty"`

}

func (u *Createbusinessunitrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createbusinessunitrequest

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		Settings *Createbusinessunitsettings `json:"settings,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		DivisionId: u.DivisionId,
		
		Settings: u.Settings,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Createbusinessunitrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
