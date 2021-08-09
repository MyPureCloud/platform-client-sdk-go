package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Routingconversationattributesresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
