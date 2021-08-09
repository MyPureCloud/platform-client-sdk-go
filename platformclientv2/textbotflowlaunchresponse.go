package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotflowlaunchresponse - Information related to a successful launch of a bot flow. The ID will be used in subsequent turn requests of the bot flow.
type Textbotflowlaunchresponse struct { 
	// Id - The session ID of the bot flow, used to send to subsequent turn requests
	Id *string `json:"id,omitempty"`

}

// String returns a JSON representation of the model
func (o *Textbotflowlaunchresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
