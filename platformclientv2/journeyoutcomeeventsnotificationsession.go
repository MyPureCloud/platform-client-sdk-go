package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeyoutcomeeventsnotificationsession
type Journeyoutcomeeventsnotificationsession struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`

}

func (o *Journeyoutcomeeventsnotificationsession) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeyoutcomeeventsnotificationsession
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		SelfUri: o.SelfUri,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeyoutcomeeventsnotificationsession) UnmarshalJSON(b []byte) error {
	var JourneyoutcomeeventsnotificationsessionMap map[string]interface{}
	err := json.Unmarshal(b, &JourneyoutcomeeventsnotificationsessionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := JourneyoutcomeeventsnotificationsessionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if SelfUri, ok := JourneyoutcomeeventsnotificationsessionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if VarType, ok := JourneyoutcomeeventsnotificationsessionMap["type"].(string); ok {
		o.VarType = &VarType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeyoutcomeeventsnotificationsession) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
