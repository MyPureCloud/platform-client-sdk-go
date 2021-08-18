package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Actionproperties
type Actionproperties struct { 
	// WebchatPrompt - Prompt message shown to user, used for webchat type action.
	WebchatPrompt *string `json:"webchatPrompt,omitempty"`


	// WebchatTitleText - Title shown to the user, used for webchat type action.
	WebchatTitleText *string `json:"webchatTitleText,omitempty"`


	// WebchatAcceptText - Accept button text shown to user, used for webchat type action.
	WebchatAcceptText *string `json:"webchatAcceptText,omitempty"`


	// WebchatDeclineText - Decline button text shown to user, used for webchat type action.
	WebchatDeclineText *string `json:"webchatDeclineText,omitempty"`


	// WebchatSurvey - Survey provided to the user, used for webchat type action.
	WebchatSurvey *Actionsurvey `json:"webchatSurvey,omitempty"`

}

func (u *Actionproperties) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actionproperties

	

	return json.Marshal(&struct { 
		WebchatPrompt *string `json:"webchatPrompt,omitempty"`
		
		WebchatTitleText *string `json:"webchatTitleText,omitempty"`
		
		WebchatAcceptText *string `json:"webchatAcceptText,omitempty"`
		
		WebchatDeclineText *string `json:"webchatDeclineText,omitempty"`
		
		WebchatSurvey *Actionsurvey `json:"webchatSurvey,omitempty"`
		*Alias
	}{ 
		WebchatPrompt: u.WebchatPrompt,
		
		WebchatTitleText: u.WebchatTitleText,
		
		WebchatAcceptText: u.WebchatAcceptText,
		
		WebchatDeclineText: u.WebchatDeclineText,
		
		WebchatSurvey: u.WebchatSurvey,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Actionproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
