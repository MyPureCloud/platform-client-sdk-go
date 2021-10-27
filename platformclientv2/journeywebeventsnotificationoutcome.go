package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebeventsnotificationoutcome
type Journeywebeventsnotificationoutcome struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// DisplayName
	DisplayName *string `json:"displayName,omitempty"`


	// Version
	Version *string `json:"version,omitempty"`

}

func (o *Journeywebeventsnotificationoutcome) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeywebeventsnotificationoutcome
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		
		Version *string `json:"version,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		SelfUri: o.SelfUri,
		
		DisplayName: o.DisplayName,
		
		Version: o.Version,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeywebeventsnotificationoutcome) UnmarshalJSON(b []byte) error {
	var JourneywebeventsnotificationoutcomeMap map[string]interface{}
	err := json.Unmarshal(b, &JourneywebeventsnotificationoutcomeMap)
	if err != nil {
		return err
	}
	
	if Id, ok := JourneywebeventsnotificationoutcomeMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if SelfUri, ok := JourneywebeventsnotificationoutcomeMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	
	if DisplayName, ok := JourneywebeventsnotificationoutcomeMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
	
	if Version, ok := JourneywebeventsnotificationoutcomeMap["version"].(string); ok {
		o.Version = &Version
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeywebeventsnotificationoutcome) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
