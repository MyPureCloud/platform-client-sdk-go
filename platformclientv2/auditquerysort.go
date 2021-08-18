package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Auditquerysort
type Auditquerysort struct { 
	// Name - Name of the property to sort.
	Name *string `json:"name,omitempty"`


	// SortOrder - Sort Order
	SortOrder *string `json:"sortOrder,omitempty"`

}

func (u *Auditquerysort) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Auditquerysort

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		SortOrder *string `json:"sortOrder,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		SortOrder: u.SortOrder,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Auditquerysort) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
