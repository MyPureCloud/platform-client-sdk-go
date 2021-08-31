package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationvideoeventtopicjourneyactionmap
type Queueconversationvideoeventtopicjourneyactionmap struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`

}

func (o *Queueconversationvideoeventtopicjourneyactionmap) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationvideoeventtopicjourneyactionmap
	
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

func (o *Queueconversationvideoeventtopicjourneyactionmap) UnmarshalJSON(b []byte) error {
	var QueueconversationvideoeventtopicjourneyactionmapMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationvideoeventtopicjourneyactionmapMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationvideoeventtopicjourneyactionmapMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Version, ok := QueueconversationvideoeventtopicjourneyactionmapMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationvideoeventtopicjourneyactionmap) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
