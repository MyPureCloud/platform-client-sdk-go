package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Publishform
type Publishform struct { 
	// Published - Is this form published
	Published *bool `json:"published,omitempty"`


	// Id - Unique Id for this version of this form
	Id *string `json:"id,omitempty"`

}

func (o *Publishform) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Publishform
	
	return json.Marshal(&struct { 
		Published *bool `json:"published,omitempty"`
		
		Id *string `json:"id,omitempty"`
		*Alias
	}{ 
		Published: o.Published,
		
		Id: o.Id,
		Alias:    (*Alias)(o),
	})
}

func (o *Publishform) UnmarshalJSON(b []byte) error {
	var PublishformMap map[string]interface{}
	err := json.Unmarshal(b, &PublishformMap)
	if err != nil {
		return err
	}
	
	if Published, ok := PublishformMap["published"].(bool); ok {
		o.Published = &Published
	}
	
	if Id, ok := PublishformMap["id"].(string); ok {
		o.Id = &Id
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Publishform) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
