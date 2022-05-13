package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Facetterm
type Facetterm struct { 
	// Term
	Term *string `json:"term,omitempty"`


	// Key
	Key *int `json:"key,omitempty"`


	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Count
	Count *int `json:"count,omitempty"`


	// Time - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Time *time.Time `json:"time,omitempty"`

}

func (o *Facetterm) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Facetterm
	
	Time := new(string)
	if o.Time != nil {
		
		*Time = timeutil.Strftime(o.Time, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Time = nil
	}
	
	return json.Marshal(&struct { 
		Term *string `json:"term,omitempty"`
		
		Key *int `json:"key,omitempty"`
		
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Count *int `json:"count,omitempty"`
		
		Time *string `json:"time,omitempty"`
		*Alias
	}{ 
		Term: o.Term,
		
		Key: o.Key,
		
		Id: o.Id,
		
		Name: o.Name,
		
		Count: o.Count,
		
		Time: Time,
		Alias:    (*Alias)(o),
	})
}

func (o *Facetterm) UnmarshalJSON(b []byte) error {
	var FacettermMap map[string]interface{}
	err := json.Unmarshal(b, &FacettermMap)
	if err != nil {
		return err
	}
	
	if Term, ok := FacettermMap["term"].(string); ok {
		o.Term = &Term
	}
    
	if Key, ok := FacettermMap["key"].(float64); ok {
		KeyInt := int(Key)
		o.Key = &KeyInt
	}
	
	if Id, ok := FacettermMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := FacettermMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Count, ok := FacettermMap["count"].(float64); ok {
		CountInt := int(Count)
		o.Count = &CountInt
	}
	
	if timeString, ok := FacettermMap["time"].(string); ok {
		Time, _ := time.Parse("2006-01-02T15:04:05.999999Z", timeString)
		o.Time = &Time
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Facetterm) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
