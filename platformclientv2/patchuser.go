package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchuser
type Patchuser struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// AcdAutoAnswer - The value that denotes if acdAutoAnswer is set on the user
	AcdAutoAnswer *bool `json:"acdAutoAnswer,omitempty"`

}

func (u *Patchuser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchuser

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		AcdAutoAnswer *bool `json:"acdAutoAnswer,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		AcdAutoAnswer: u.AcdAutoAnswer,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Patchuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
