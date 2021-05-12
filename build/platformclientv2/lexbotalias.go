package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Lexbotalias) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
