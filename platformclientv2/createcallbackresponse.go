package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createcallbackresponse
type Createcallbackresponse struct { 
	// Conversation - The conversation associated with the callback
	Conversation *Domainentityref `json:"conversation,omitempty"`


	// CallbackIdentifiers - The list of communication identifiers for the callback participants
	CallbackIdentifiers *[]Callbackidentifier `json:"callbackIdentifiers,omitempty"`

}

func (u *Createcallbackresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createcallbackresponse

	

	return json.Marshal(&struct { 
		Conversation *Domainentityref `json:"conversation,omitempty"`
		
		CallbackIdentifiers *[]Callbackidentifier `json:"callbackIdentifiers,omitempty"`
		*Alias
	}{ 
		Conversation: u.Conversation,
		
		CallbackIdentifiers: u.CallbackIdentifiers,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Createcallbackresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
