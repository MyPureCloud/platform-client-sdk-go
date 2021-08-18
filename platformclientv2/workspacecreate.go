package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workspacecreate
type Workspacecreate struct { 
	// Name - The workspace name
	Name *string `json:"name,omitempty"`


	// Bucket
	Bucket *string `json:"bucket,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`

}

func (u *Workspacecreate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workspacecreate

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Bucket *string `json:"bucket,omitempty"`
		
		Description *string `json:"description,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Bucket: u.Bucket,
		
		Description: u.Description,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Workspacecreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
