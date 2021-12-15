package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcalleventtopicjourneyactionmap - Details about the action map from the Journey System which triggered this action
type Queueconversationcalleventtopicjourneyactionmap struct { 
	// Id - The ID of the actionMap in the Journey System which triggered this action
	Id *string `json:"id,omitempty"`


	// Version - The version number of the actionMap in the Journey System at the time this action was triggered
	Version *int `json:"version,omitempty"`

}

func (o *Queueconversationcalleventtopicjourneyactionmap) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationcalleventtopicjourneyactionmap
	
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

func (o *Queueconversationcalleventtopicjourneyactionmap) UnmarshalJSON(b []byte) error {
	var QueueconversationcalleventtopicjourneyactionmapMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationcalleventtopicjourneyactionmapMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationcalleventtopicjourneyactionmapMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Version, ok := QueueconversationcalleventtopicjourneyactionmapMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationcalleventtopicjourneyactionmap) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
