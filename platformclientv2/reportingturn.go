package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Reportingturn
type Reportingturn struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Reportingturn) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Reportingturn) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

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
		Alias
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
		Alias:    (Alias)(o),
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
