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


	// SessionEndDetails - The details related to end of bot flow session.
	SessionEndDetails *Sessionenddetails `json:"sessionEndDetails,omitempty"`


	// Conversation - The conversation details, across potentially multiple Bot Flow sessions.
	Conversation *Addressableentityref `json:"conversation,omitempty"`

}

func (o *Reportingturn) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reportingturn
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		
		SessionEndDetails *Sessionenddetails `json:"sessionEndDetails,omitempty"`
		
		Conversation *Addressableentityref `json:"conversation,omitempty"`
		*Alias
	}{ 
		UserInput: o.UserInput,
		
		BotPrompts: o.BotPrompts,
		
		SessionId: o.SessionId,
		
		AskAction: o.AskAction,
		
		Intent: o.Intent,
		
		Knowledge: o.Knowledge,
		
		DateCreated: DateCreated,
		
		AskActionResult: o.AskActionResult,
		
		SessionEndDetails: o.SessionEndDetails,
		
		Conversation: o.Conversation,
		Alias:    (*Alias)(o),
	})
}

func (o *Reportingturn) UnmarshalJSON(b []byte) error {
	var ReportingturnMap map[string]interface{}
	err := json.Unmarshal(b, &ReportingturnMap)
	if err != nil {
		return err
	}
	
	if UserInput, ok := ReportingturnMap["userInput"].(string); ok {
		o.UserInput = &UserInput
	}
    
	if BotPrompts, ok := ReportingturnMap["botPrompts"].([]interface{}); ok {
		BotPromptsString, _ := json.Marshal(BotPrompts)
		json.Unmarshal(BotPromptsString, &o.BotPrompts)
	}
	
	if SessionId, ok := ReportingturnMap["sessionId"].(string); ok {
		o.SessionId = &SessionId
	}
    
	if AskAction, ok := ReportingturnMap["askAction"].(map[string]interface{}); ok {
		AskActionString, _ := json.Marshal(AskAction)
		json.Unmarshal(AskActionString, &o.AskAction)
	}
	
	if Intent, ok := ReportingturnMap["intent"].(map[string]interface{}); ok {
		IntentString, _ := json.Marshal(Intent)
		json.Unmarshal(IntentString, &o.Intent)
	}
	
	if Knowledge, ok := ReportingturnMap["knowledge"].(map[string]interface{}); ok {
		KnowledgeString, _ := json.Marshal(Knowledge)
		json.Unmarshal(KnowledgeString, &o.Knowledge)
	}
	
	if dateCreatedString, ok := ReportingturnMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if AskActionResult, ok := ReportingturnMap["askActionResult"].(string); ok {
		o.AskActionResult = &AskActionResult
	}
    
	if SessionEndDetails, ok := ReportingturnMap["sessionEndDetails"].(map[string]interface{}); ok {
		SessionEndDetailsString, _ := json.Marshal(SessionEndDetails)
		json.Unmarshal(SessionEndDetailsString, &o.SessionEndDetails)
	}
	
	if Conversation, ok := ReportingturnMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Reportingturn) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
