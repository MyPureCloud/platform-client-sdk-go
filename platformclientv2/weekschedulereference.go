package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Weekschedulereference
type Weekschedulereference struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`


	// WeekDate - First day of this week schedule in yyyy-MM-dd format
	WeekDate *string `json:"weekDate,omitempty"`

}

func (o *Weekschedulereference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Weekschedulereference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		WeekDate *string `json:"weekDate,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		SelfUri: o.SelfUri,
		
		WeekDate: o.WeekDate,
		Alias:    (*Alias)(o),
	})
}

func (o *Weekschedulereference) UnmarshalJSON(b []byte) error {
	var WeekschedulereferenceMap map[string]interface{}
	err := json.Unmarshal(b, &WeekschedulereferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WeekschedulereferenceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if SelfUri, ok := WeekschedulereferenceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if WeekDate, ok := WeekschedulereferenceMap["weekDate"].(string); ok {
		o.WeekDate = &WeekDate
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Weekschedulereference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
