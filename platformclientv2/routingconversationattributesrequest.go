package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Routingconversationattributesrequest
type Routingconversationattributesrequest struct { 
	// Priority - Priority to be updated on in-queue conversation. Range:[-25000000, 25000000]
	Priority *int `json:"priority,omitempty"`


	// SkillIds - Skills to be updated on in-queue conversation.
	SkillIds *[]string `json:"skillIds,omitempty"`


	// LanguageId - Language required on the in-queue conversation.
	LanguageId *string `json:"languageId,omitempty"`

}

// String returns a JSON representation of the model
func (o *Routingconversationattributesrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
