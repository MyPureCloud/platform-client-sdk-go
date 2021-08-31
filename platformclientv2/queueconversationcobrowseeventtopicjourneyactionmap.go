package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationcobrowseeventtopicjourneyactionmap
type Queueconversationcobrowseeventtopicjourneyactionmap struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`

}

func (o *Queueconversationcobrowseeventtopicjourneyactionmap) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationcobrowseeventtopicjourneyactionmap
	
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

func (o *Queueconversationcobrowseeventtopicjourneyactionmap) UnmarshalJSON(b []byte) error {
	var QueueconversationcobrowseeventtopicjourneyactionmapMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationcobrowseeventtopicjourneyactionmapMap)
	if err != nil {
		return err
	}
	
	if Id, ok := QueueconversationcobrowseeventtopicjourneyactionmapMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Version, ok := QueueconversationcobrowseeventtopicjourneyactionmapMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationcobrowseeventtopicjourneyactionmap) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
