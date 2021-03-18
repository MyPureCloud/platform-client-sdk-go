package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Audituser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
