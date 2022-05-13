package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebeventsnotificationactionmap
type Journeywebeventsnotificationactionmap struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// DisplayName
	DisplayName *string `json:"displayName,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`

}

func (o *Journeywebeventsnotificationactionmap) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeywebeventsnotificationactionmap
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		
		Version *int `json:"version,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		SelfUri: o.SelfUri,
		
		DisplayName: o.DisplayName,
		
		Version: o.Version,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeywebeventsnotificationactionmap) UnmarshalJSON(b []byte) error {
	var JourneywebeventsnotificationactionmapMap map[string]interface{}
	err := json.Unmarshal(b, &JourneywebeventsnotificationactionmapMap)
	if err != nil {
		return err
	}
	
	if Id, ok := JourneywebeventsnotificationactionmapMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if SelfUri, ok := JourneywebeventsnotificationactionmapMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if DisplayName, ok := JourneywebeventsnotificationactionmapMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
    
	if Version, ok := JourneywebeventsnotificationactionmapMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeywebeventsnotificationactionmap) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
