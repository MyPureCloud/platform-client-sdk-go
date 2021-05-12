package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Journeycontext
type Journeycontext struct { 
	// Customer - A subset of the Journey System's customer data at a point-in-time (for external linkage and internal usage/context)
	Customer *Journeycustomer `json:"customer,omitempty"`


	// CustomerSession - A subset of the Journey System's tracked customer session data at a point-in-time (for external linkage and internal usage/context)
	CustomerSession *Journeycustomersession `json:"customerSession,omitempty"`


	// TriggeringAction - A subset of the Journey System's action data relevant to a part of a conversation (for external linkage and internal usage/context)
	TriggeringAction *Journeyaction `json:"triggeringAction,omitempty"`

}

// String returns a JSON representation of the model
func (o *Journeycontext) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
