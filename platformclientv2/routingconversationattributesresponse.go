package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Routingconversationattributesresponse
type Routingconversationattributesresponse struct { 
	// Priority - Current priority value on in-queue conversation. Range:[-25000000, 25000000]
	Priority *int `json:"priority,omitempty"`


	// Skills - Current routing skills on in-queue conversation
	Skills *[]Routingskill `json:"skills,omitempty"`


	// Language - Current language on in-queue conversation
	Language *Language `json:"language,omitempty"`

}

func (u *Routingconversationattributesresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Routingconversationattributesresponse

	

	return json.Marshal(&struct { 
		Priority *int `json:"priority,omitempty"`
		
		Skills *[]Routingskill `json:"skills,omitempty"`
		
		Language *Language `json:"language,omitempty"`
		*Alias
	}{ 
		Priority: u.Priority,
		
		Skills: u.Skills,
		
		Language: u.Language,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Routingconversationattributesresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
