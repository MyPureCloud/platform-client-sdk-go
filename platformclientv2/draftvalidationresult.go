package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Draftvalidationresult - Validation results
type Draftvalidationresult struct { 
	// Valid - Indicates if configuration is valid
	Valid *bool `json:"valid,omitempty"`


	// Errors - List of errors causing validation failure
	Errors *[]Errorbody `json:"errors,omitempty"`

}

func (u *Draftvalidationresult) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Draftvalidationresult

	

	return json.Marshal(&struct { 
		Valid *bool `json:"valid,omitempty"`
		
		Errors *[]Errorbody `json:"errors,omitempty"`
		*Alias
	}{ 
		Valid: u.Valid,
		
		Errors: u.Errors,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Draftvalidationresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
