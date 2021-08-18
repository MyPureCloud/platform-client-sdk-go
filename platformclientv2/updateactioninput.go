package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Updateactioninput
type Updateactioninput struct { 
	// Category - Category of action, Can be up to 256 characters long
	Category *string `json:"category,omitempty"`


	// Name - Name of action, Can be up to 256 characters long
	Name *string `json:"name,omitempty"`


	// Config - Configuration to support request and response processing
	Config *Actionconfig `json:"config,omitempty"`


	// Version - Version of this action
	Version *int `json:"version,omitempty"`

}

func (u *Updateactioninput) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Updateactioninput

	

	return json.Marshal(&struct { 
		Category *string `json:"category,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Config *Actionconfig `json:"config,omitempty"`
		
		Version *int `json:"version,omitempty"`
		*Alias
	}{ 
		Category: u.Category,
		
		Name: u.Name,
		
		Config: u.Config,
		
		Version: u.Version,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Updateactioninput) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
