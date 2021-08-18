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

func (u *Facetterm) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Facetterm

	
	Time := new(string)
	if u.Time != nil {
		
		*Time = timeutil.Strftime(u.Time, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Term: u.Term,
		
		Key: u.Key,
		
		Id: u.Id,
		
		Name: u.Name,
		
		Count: u.Count,
		
		Time: Time,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Facetterm) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
