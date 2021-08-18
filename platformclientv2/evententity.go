package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Evententity
type Evententity struct { 
	// EntityType - Type of entity the event pertains to. e.g. integration
	EntityType *string `json:"entityType,omitempty"`


	// Id - ID of the entity the event pertains to.
	Id *string `json:"id,omitempty"`

}

func (u *Evententity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Evententity

	

	return json.Marshal(&struct { 
		EntityType *string `json:"entityType,omitempty"`
		
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		EntityType: u.EntityType,
		
		Id: u.Id,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Evententity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
