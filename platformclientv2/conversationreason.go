package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationreason - Reasons for a failed message receipt.
type Conversationreason struct { 
	// Code - The reason code for the failed message receipt.
	Code *string `json:"code,omitempty"`


	// Message - Description of the reason for the failed message receipt.
	Message *string `json:"message,omitempty"`

}

// String returns a JSON representation of the model
func (o *Conversationreason) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
