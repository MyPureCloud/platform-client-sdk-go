package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Articlecontent
type Articlecontent struct { 
	// Body - Body of the article content.
	Body *Articlecontentbody `json:"body,omitempty"`

}

func (o *Articlecontent) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Articlecontent
	
	return json.Marshal(&struct { 
		Body *Articlecontentbody `json:"body,omitempty"`
		*Alias
	}{ 
		Body: o.Body,
		Alias:    (*Alias)(o),
	})
}

func (o *Articlecontent) UnmarshalJSON(b []byte) error {
	var ArticlecontentMap map[string]interface{}
	err := json.Unmarshal(b, &ArticlecontentMap)
	if err != nil {
		return err
	}
	
	if Body, ok := ArticlecontentMap["body"].(map[string]interface{}); ok {
		BodyString, _ := json.Marshal(Body)
		json.Unmarshal(BodyString, &o.Body)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Articlecontent) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
