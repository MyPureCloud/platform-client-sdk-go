package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Routingconversationattributesrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Routingconversationattributesrequest

	

	return json.Marshal(&struct { 
		Priority *int `json:"priority,omitempty"`
		
		SkillIds *[]string `json:"skillIds,omitempty"`
		
		LanguageId *string `json:"languageId,omitempty"`
		*Alias
	}{ 
		Priority: u.Priority,
		
		SkillIds: u.SkillIds,
		
		LanguageId: u.LanguageId,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Routingconversationattributesrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
