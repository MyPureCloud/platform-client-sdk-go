package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationemaileventtopicjourneyactionmap
type Queueconversationemaileventtopicjourneyactionmap struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`

}

func (o *Queueconversationemaileventtopicjourneyactionmap) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationemaileventtopicjourneyactionmap
	
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

func (o *Queueconversationemaileventtopicjourneyactionmap) UnmarshalJSON(b []byte) error {
	var QueueconversationemaileventtopicjourneyactionmapMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationemaileventtopicjourneyactionmapMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationemaileventtopicjourneyactionmapMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Version, ok := QueueconversationemaileventtopicjourneyactionmapMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationemaileventtopicjourneyactionmap) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
