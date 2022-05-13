package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Urlresponse
type Urlresponse struct { 
	// Url
	Url *string `json:"url,omitempty"`

}

func (o *Urlresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Urlresponse
	
	return json.Marshal(&struct { 
		Url *string `json:"url,omitempty"`
		*Alias
	}{ 
		Url: o.Url,
		Alias:    (*Alias)(o),
	})
}

func (o *Urlresponse) UnmarshalJSON(b []byte) error {
	var UrlresponseMap map[string]interface{}
	err := json.Unmarshal(b, &UrlresponseMap)
	if err != nil {
		return err
	}
	
	if Url, ok := UrlresponseMap["url"].(string); ok {
		o.Url = &Url
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Urlresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
