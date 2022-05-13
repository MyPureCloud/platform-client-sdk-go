package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebactioneventsnotificationsession
type Journeywebactioneventsnotificationsession struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`

}

func (o *Journeywebactioneventsnotificationsession) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeywebactioneventsnotificationsession
	
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

func (o *Journeywebactioneventsnotificationsession) UnmarshalJSON(b []byte) error {
	var JourneywebactioneventsnotificationsessionMap map[string]interface{}
	err := json.Unmarshal(b, &JourneywebactioneventsnotificationsessionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := JourneywebactioneventsnotificationsessionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if SelfUri, ok := JourneywebactioneventsnotificationsessionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if VarType, ok := JourneywebactioneventsnotificationsessionMap["type"].(string); ok {
		o.VarType = &VarType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Journeywebactioneventsnotificationsession) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
