package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Updatemanagementunitrequest
type Updatemanagementunitrequest struct { 
	// Name - The new name of the management unit
	Name *string `json:"name,omitempty"`


	// DivisionId - The new division id for the management unit
	DivisionId *string `json:"divisionId,omitempty"`


	// Settings - Updated settings for the management unit
	Settings *Managementunitsettingsrequest `json:"settings,omitempty"`

}

func (u *Updatemanagementunitrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Updatemanagementunitrequest

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		Settings *Managementunitsettingsrequest `json:"settings,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		DivisionId: u.DivisionId,
		
		Settings: u.Settings,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Updatemanagementunitrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
