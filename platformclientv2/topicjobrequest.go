package platformclientv2
import (
	"encoding/json"
)

// Topicjobrequest
type Topicjobrequest struct { 
	// TopicIds - The ids of the topics used for this job
	TopicIds *[]string `json:"topicIds,omitempty"`

}

// String returns a JSON representation of the model
func (o *Topicjobrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
