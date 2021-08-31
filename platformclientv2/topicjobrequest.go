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

func (o *Topicjobrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Topicjobrequest
	
	return json.Marshal(&struct { 
		TopicIds *[]string `json:"topicIds,omitempty"`
		*Alias
	}{ 
		TopicIds: o.TopicIds,
		Alias:    (*Alias)(o),
	})
}

func (o *Topicjobrequest) UnmarshalJSON(b []byte) error {
	var TopicjobrequestMap map[string]interface{}
	err := json.Unmarshal(b, &TopicjobrequestMap)
	if err != nil {
		return err
	}
	
	if TopicIds, ok := TopicjobrequestMap["topicIds"].([]interface{}); ok {
		TopicIdsString, _ := json.Marshal(TopicIds)
		json.Unmarshal(TopicIdsString, &o.TopicIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Topicjobrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
