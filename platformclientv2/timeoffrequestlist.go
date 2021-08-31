package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Timeoffrequestlist
type Timeoffrequestlist struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// TimeOffRequests
	TimeOffRequests *[]Timeoffrequestresponse `json:"timeOffRequests,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Timeoffrequestlist) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Timeoffrequestlist
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		TimeOffRequests *[]Timeoffrequestresponse `json:"timeOffRequests,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		TimeOffRequests: o.TimeOffRequests,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Timeoffrequestlist) UnmarshalJSON(b []byte) error {
	var TimeoffrequestlistMap map[string]interface{}
	err := json.Unmarshal(b, &TimeoffrequestlistMap)
	if err != nil {
		return err
	}
	
	if Id, ok := TimeoffrequestlistMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := TimeoffrequestlistMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if TimeOffRequests, ok := TimeoffrequestlistMap["timeOffRequests"].([]interface{}); ok {
		TimeOffRequestsString, _ := json.Marshal(TimeOffRequests)
		json.Unmarshal(TimeOffRequestsString, &o.TimeOffRequests)
	}
	
	if SelfUri, ok := TimeoffrequestlistMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Timeoffrequestlist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
