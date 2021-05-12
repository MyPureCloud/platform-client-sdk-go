package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Patchactionproperties
type Patchactionproperties struct { 
	// WebchatPrompt - Prompt message shown to user, used for webchat type action.
	WebchatPrompt *string `json:"webchatPrompt,omitempty"`


	// WebchatTitleText - Title shown to the user, used for webchat type action.
	WebchatTitleText *string `json:"webchatTitleText,omitempty"`


	// WebchatAcceptText - Accept button text shown to user, used for webchat type action.
	WebchatAcceptText *string `json:"webchatAcceptText,omitempty"`


	// WebchatDeclineText - Decline button text shown to user, used for webchat type action.
	WebchatDeclineText *string `json:"webchatDeclineText,omitempty"`


	// WebchatSurvey - Survey provided to the user, used for webchat type action.
	WebchatSurvey *Patchactionsurvey `json:"webchatSurvey,omitempty"`

}

// String returns a JSON representation of the model
func (o *Patchactionproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
