package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationscreenshareeventtopicjourneyactionmap
type Queueconversationscreenshareeventtopicjourneyactionmap struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`

}

func (o *Queueconversationscreenshareeventtopicjourneyactionmap) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationscreenshareeventtopicjourneyactionmap
	
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

func (o *Queueconversationscreenshareeventtopicjourneyactionmap) UnmarshalJSON(b []byte) error {
	var QueueconversationscreenshareeventtopicjourneyactionmapMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationscreenshareeventtopicjourneyactionmapMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationscreenshareeventtopicjourneyactionmapMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Version, ok := QueueconversationscreenshareeventtopicjourneyactionmapMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationscreenshareeventtopicjourneyactionmap) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
