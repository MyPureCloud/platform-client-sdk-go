package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Topicjobrequest
type Topicjobrequest struct { 
	// TopicIds - The ids of the topics used for this job
	TopicIds *[]string `json:"topicIds,omitempty"`

}

func (u *Topicjobrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Topicjobrequest

	

	return json.Marshal(&struct { 
		TopicIds *[]string `json:"topicIds,omitempty"`
		*Alias
	}{ 
		TopicIds: u.TopicIds,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Topicjobrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
