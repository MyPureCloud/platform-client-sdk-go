package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Lexbotalias
type Lexbotalias struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Bot - The Lex bot this is an alias for
	Bot *Lexbot `json:"bot,omitempty"`


	// BotVersion - The version of the Lex bot this alias points at
	BotVersion *string `json:"botVersion,omitempty"`


	// Status - The status of the Lex bot alias
	Status *string `json:"status,omitempty"`


	// FailureReason - If the status is FAILED, Amazon Lex explains why it failed to build the bot
	FailureReason *string `json:"failureReason,omitempty"`


	// Language - The target language of the Lex bot
	Language *string `json:"language,omitempty"`


	// Intents - An array of Intents associated with this bot alias
	Intents *[]Lexintent `json:"intents,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Lexbotalias) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Lexbotalias
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Bot *Lexbot `json:"bot,omitempty"`
		
		BotVersion *string `json:"botVersion,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		FailureReason *string `json:"failureReason,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		Intents *[]Lexintent `json:"intents,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Bot: o.Bot,
		
		BotVersion: o.BotVersion,
		
		Status: o.Status,
		
		FailureReason: o.FailureReason,
		
		Language: o.Language,
		
		Intents: o.Intents,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Lexbotalias) UnmarshalJSON(b []byte) error {
	var LexbotaliasMap map[string]interface{}
	err := json.Unmarshal(b, &LexbotaliasMap)
	if err != nil {
		return err
	}
	
	if Id, ok := LexbotaliasMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := LexbotaliasMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Bot, ok := LexbotaliasMap["bot"].(map[string]interface{}); ok {
		BotString, _ := json.Marshal(Bot)
		json.Unmarshal(BotString, &o.Bot)
	}
	
	if BotVersion, ok := LexbotaliasMap["botVersion"].(string); ok {
		o.BotVersion = &BotVersion
	}
	
	if Status, ok := LexbotaliasMap["status"].(string); ok {
		o.Status = &Status
	}
	
	if FailureReason, ok := LexbotaliasMap["failureReason"].(string); ok {
		o.FailureReason = &FailureReason
	}
	
	if Language, ok := LexbotaliasMap["language"].(string); ok {
		o.Language = &Language
	}
	
	if Intents, ok := LexbotaliasMap["intents"].([]interface{}); ok {
		IntentsString, _ := json.Marshal(Intents)
		json.Unmarshal(IntentsString, &o.Intents)
	}
	
	if SelfUri, ok := LexbotaliasMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Lexbotalias) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
