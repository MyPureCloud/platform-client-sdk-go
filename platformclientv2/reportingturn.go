package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Reportingturn
type Reportingturn struct { 
	// UserInput - The chosen user input associated with this reporting turn.
	UserInput *string `json:"userInput,omitempty"`


	// BotPrompts - The bot prompts associated with this reporting turn.
	BotPrompts *[]string `json:"botPrompts,omitempty"`


	// SessionId - The bot session ID that this reporting turn is grouped under.
	SessionId *string `json:"sessionId,omitempty"`


	// AskAction - The bot flow 'ask' action associated with this reporting turn (e.g. AskForIntent).
	AskAction *Reportingturnaction `json:"askAction,omitempty"`


	// Intent - The intent and associated slots detected during this reporting turn.
	Intent *Reportingturnintent `json:"intent,omitempty"`


	// Knowledge - The knowledge data captured during this reporting turn.
	Knowledge *Reportingturnknowledge `json:"knowledge,omitempty"`


	// DateCreated - Timestamp indicating when the original turn was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// AskActionResult - Result of the bot flow 'ask' action.
	AskActionResult *string `json:"askActionResult,omitempty"`


	// Conversation - The conversation details, across potentially multiple Cicero sessions.
	Conversation *Addressableentityref `json:"conversation,omitempty"`

}

func (u *Reportingturn) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reportingturn

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	

	return json.Marshal(&struct { 
		UserInput *string `json:"userInput,omitempty"`
		
		BotPrompts *[]string `json:"botPrompts,omitempty"`
		
		SessionId *string `json:"sessionId,omitempty"`
		
		AskAction *Reportingturnaction `json:"askAction,omitempty"`
		
		Intent *Reportingturnintent `json:"intent,omitempty"`
		
		Knowledge *Reportingturnknowledge `json:"knowledge,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		AskActionResult *string `json:"askActionResult,omitempty"`
		
		Conversation *Addressableentityref `json:"conversation,omitempty"`
		*Alias
	}{ 
		UserInput: u.UserInput,
		
		BotPrompts: u.BotPrompts,
		
		SessionId: u.SessionId,
		
		AskAction: u.AskAction,
		
		Intent: u.Intent,
		
		Knowledge: u.Knowledge,
		
		DateCreated: DateCreated,
		
		AskActionResult: u.AskActionResult,
		
		Conversation: u.Conversation,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Reportingturn) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
