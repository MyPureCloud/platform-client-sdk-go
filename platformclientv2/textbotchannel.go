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

func (o *Textbotchannel) MarshalJSON() ([]byte, error) {
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
		Name: o.Name,
		
		InputModes: o.InputModes,
		
		OutputModes: o.OutputModes,
		
		UserAgent: o.UserAgent,
		Alias:    (*Alias)(o),
	})
}

func (o *Textbotchannel) UnmarshalJSON(b []byte) error {
	var TextbotchannelMap map[string]interface{}
	err := json.Unmarshal(b, &TextbotchannelMap)
	if err != nil {
		return err
	}
	
	if Name, ok := TextbotchannelMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if InputModes, ok := TextbotchannelMap["inputModes"].([]interface{}); ok {
		InputModesString, _ := json.Marshal(InputModes)
		json.Unmarshal(InputModesString, &o.InputModes)
	}
	
	if OutputModes, ok := TextbotchannelMap["outputModes"].([]interface{}); ok {
		OutputModesString, _ := json.Marshal(OutputModes)
		json.Unmarshal(OutputModesString, &o.OutputModes)
	}
	
	if UserAgent, ok := TextbotchannelMap["userAgent"].(map[string]interface{}); ok {
		UserAgentString, _ := json.Marshal(UserAgent)
		json.Unmarshal(UserAgentString, &o.UserAgent)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Textbotchannel) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
