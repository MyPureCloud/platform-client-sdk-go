package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contentactions - User actions available on the content. All actions are optional and all actions are executed simultaneously.
type Contentactions struct { 
	// Url - A URL of a web page to direct the user to.
	Url *string `json:"url,omitempty"`


	// UrlTarget - The target window in which to open the URL. If empty will open a blank page or tab.
	UrlTarget *string `json:"urlTarget,omitempty"`


	// Textback - Text to be returned as the payload from a ButtonResponse when a button is clicked. The textback and title are a combination which will have to be unique across each card and carousel in order to determine which button was clicked in that card or carousel.
	Textback *string `json:"textback,omitempty"`

}

func (o *Contentactions) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contentactions
	
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

func (o *Contentactions) UnmarshalJSON(b []byte) error {
	var ContentactionsMap map[string]interface{}
	err := json.Unmarshal(b, &ContentactionsMap)
	if err != nil {
		return err
	}
	
	if Url, ok := ContentactionsMap["url"].(string); ok {
		o.Url = &Url
	}
	
	if UrlTarget, ok := ContentactionsMap["urlTarget"].(string); ok {
		o.UrlTarget = &UrlTarget
	}
	
	if Textback, ok := ContentactionsMap["textback"].(string); ok {
		o.Textback = &Textback
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contentactions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
