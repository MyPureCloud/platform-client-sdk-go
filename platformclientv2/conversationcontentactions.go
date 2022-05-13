package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationcontentactions - User actions available on the content. All actions are optional and all actions are executed simultaneously.
type Conversationcontentactions struct { 
	// Url - A URL of a web page to direct the user to.
	Url *string `json:"url,omitempty"`


	// UrlTarget - The target window in which to open the URL. If empty will open a blank page or tab.
	UrlTarget *string `json:"urlTarget,omitempty"`


	// Textback - Text to be sent back in reply when the item is selected.
	Textback *string `json:"textback,omitempty"`

}

func (o *Conversationcontentactions) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationcontentactions
	
	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		
		UrlTarget *string `json:"urlTarget,omitempty"`
		
		Textback *string `json:"textback,omitempty"`
		*Alias
	}{ 
		Url: o.Url,
		
		UrlTarget: o.UrlTarget,
		
		Textback: o.Textback,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationcontentactions) UnmarshalJSON(b []byte) error {
	var ConversationcontentactionsMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationcontentactionsMap)
	if err != nil {
		return err
	}
	
	if Url, ok := ConversationcontentactionsMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if UrlTarget, ok := ConversationcontentactionsMap["urlTarget"].(string); ok {
		o.UrlTarget = &UrlTarget
	}
    
	if Textback, ok := ConversationcontentactionsMap["textback"].(string); ok {
		o.Textback = &Textback
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationcontentactions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
