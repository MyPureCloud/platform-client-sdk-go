package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentbodyvideo
type Documentbodyvideo struct { 
	// Url - The URL for the video.
	Url *string `json:"url,omitempty"`

}

func (o *Documentbodyvideo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Documentbodyvideo
	
	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		*Alias
	}{ 
		Url: o.Url,
		Alias:    (*Alias)(o),
	})
}

func (o *Documentbodyvideo) UnmarshalJSON(b []byte) error {
	var DocumentbodyvideoMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentbodyvideoMap)
	if err != nil {
		return err
	}
	
	if Url, ok := DocumentbodyvideoMap["url"].(string); ok {
		o.Url = &Url
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Documentbodyvideo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
