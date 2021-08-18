package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotchannel - Channel information relevant to a bot flow.
type Textbotchannel struct { 
	// Name - The name of the channel.
	Name *string `json:"name,omitempty"`


	// InputModes - The input modes for the channel.
	InputModes *[]string `json:"inputModes,omitempty"`


	// OutputModes - The output modes for the channel.
	OutputModes *[]string `json:"outputModes,omitempty"`


	// UserAgent - Information about the end user agent calling the bot flow.
	UserAgent *Textbotuseragent `json:"userAgent,omitempty"`

}

func (u *Textbotchannel) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Textbotchannel

	

	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		InputModes *[]string `json:"inputModes,omitempty"`
		
		OutputModes *[]string `json:"outputModes,omitempty"`
		
		UserAgent *Textbotuseragent `json:"userAgent,omitempty"`
		*Alias
	}{ 
		Name: u.Name,
		
		InputModes: u.InputModes,
		
		OutputModes: u.OutputModes,
		
		UserAgent: u.UserAgent,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Textbotchannel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
