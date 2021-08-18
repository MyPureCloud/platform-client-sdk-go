package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Audituser
type Audituser struct { 
	// Id - The ID (UUID) of the user who initiated the action of this AuditMessage.
	Id *string `json:"id,omitempty"`


	// Name - The full username of the user who initiated the action of this AuditMessage.
	Name *string `json:"name,omitempty"`


	// Display - The display name of the user who initiated the action of this AuditMessage.
	Display *string `json:"display,omitempty"`

}

func (u *Audituser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Audituser

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Display *string `json:"display,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Display: u.Display,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Audituser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
