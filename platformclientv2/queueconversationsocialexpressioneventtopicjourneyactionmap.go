package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationsocialexpressioneventtopicjourneyactionmap
type Queueconversationsocialexpressioneventtopicjourneyactionmap struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`

}

func (o *Queueconversationsocialexpressioneventtopicjourneyactionmap) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationsocialexpressioneventtopicjourneyactionmap
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Version *int `json:"version,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Version: o.Version,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationsocialexpressioneventtopicjourneyactionmap) UnmarshalJSON(b []byte) error {
	var QueueconversationsocialexpressioneventtopicjourneyactionmapMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationsocialexpressioneventtopicjourneyactionmapMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationsocialexpressioneventtopicjourneyactionmapMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Version, ok := QueueconversationsocialexpressioneventtopicjourneyactionmapMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationsocialexpressioneventtopicjourneyactionmap) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
