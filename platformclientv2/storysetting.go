package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Storysetting
type Storysetting struct { 
	// Mention - Setting relating to Story Mentions
	Mention *Inboundonlysetting `json:"mention,omitempty"`


	// Reply - Setting relating to Story Replies
	Reply *Inboundonlysetting `json:"reply,omitempty"`

}

func (o *Storysetting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Storysetting
	
	return json.Marshal(&struct { 
		Mention *Inboundonlysetting `json:"mention,omitempty"`
		
		Reply *Inboundonlysetting `json:"reply,omitempty"`
		*Alias
	}{ 
		Mention: o.Mention,
		
		Reply: o.Reply,
		Alias:    (*Alias)(o),
	})
}

func (o *Storysetting) UnmarshalJSON(b []byte) error {
	var StorysettingMap map[string]interface{}
	err := json.Unmarshal(b, &StorysettingMap)
	if err != nil {
		return err
	}
	
	if Mention, ok := StorysettingMap["mention"].(map[string]interface{}); ok {
		MentionString, _ := json.Marshal(Mention)
		json.Unmarshal(MentionString, &o.Mention)
	}
	
	if Reply, ok := StorysettingMap["reply"].(map[string]interface{}); ok {
		ReplyString, _ := json.Marshal(Reply)
		json.Unmarshal(ReplyString, &o.Reply)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Storysetting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
