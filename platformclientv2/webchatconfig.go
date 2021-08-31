package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webchatconfig
type Webchatconfig struct { 
	// WebChatSkin - css class to be applied to the web chat widget.
	WebChatSkin *string `json:"webChatSkin,omitempty"`

}

func (o *Webchatconfig) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webchatconfig
	
	return json.Marshal(&struct { 
		WebChatSkin *string `json:"webChatSkin,omitempty"`
		*Alias
	}{ 
		WebChatSkin: o.WebChatSkin,
		Alias:    (*Alias)(o),
	})
}

func (o *Webchatconfig) UnmarshalJSON(b []byte) error {
	var WebchatconfigMap map[string]interface{}
	err := json.Unmarshal(b, &WebchatconfigMap)
	if err != nil {
		return err
	}
	
	if WebChatSkin, ok := WebchatconfigMap["webChatSkin"].(string); ok {
		o.WebChatSkin = &WebChatSkin
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Webchatconfig) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
