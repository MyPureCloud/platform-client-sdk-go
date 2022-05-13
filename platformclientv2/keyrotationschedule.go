package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Keyrotationschedule
type Keyrotationschedule struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Period - Value to set schedule to
	Period *string `json:"period,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Keyrotationschedule) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Keyrotationschedule
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Period *string `json:"period,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Period: o.Period,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Keyrotationschedule) UnmarshalJSON(b []byte) error {
	var KeyrotationscheduleMap map[string]interface{}
	err := json.Unmarshal(b, &KeyrotationscheduleMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KeyrotationscheduleMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := KeyrotationscheduleMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Period, ok := KeyrotationscheduleMap["period"].(string); ok {
		o.Period = &Period
	}
    
	if SelfUri, ok := KeyrotationscheduleMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Keyrotationschedule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
