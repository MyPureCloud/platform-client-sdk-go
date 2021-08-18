package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Auditqueryservice
type Auditqueryservice struct { 
	// Name - Name of the Service
	Name *string `json:"name,omitempty"`


	// Entities - List of Entities
	Entities *[]Auditqueryentity `json:"entities,omitempty"`

}

func (u *Auditqueryservice) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Auditqueryservice

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Entities *[]Auditqueryentity `json:"entities,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		Entities: u.Entities,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Auditqueryservice) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
