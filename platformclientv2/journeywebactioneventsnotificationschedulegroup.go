package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Journeywebactioneventsnotificationschedulegroup
type Journeywebactioneventsnotificationschedulegroup struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Journeywebactioneventsnotificationschedulegroup) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Journeywebactioneventsnotificationschedulegroup
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Journeywebactioneventsnotificationschedulegroup) UnmarshalJSON(b []byte) error {
	var JourneywebactioneventsnotificationschedulegroupMap map[string]interface{}
	err := json.Unmarshal(b, &JourneywebactioneventsnotificationschedulegroupMap)
	if err != nil {
		return err
	}
	
	if Id, ok := JourneywebactioneventsnotificationschedulegroupMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if SelfUri, ok := JourneywebactioneventsnotificationschedulegroupMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Journeywebactioneventsnotificationschedulegroup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
