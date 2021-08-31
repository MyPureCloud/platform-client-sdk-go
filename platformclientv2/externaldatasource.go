package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externaldatasource - Describes a link to a record in an external system that contributed data to a Relate record
type Externaldatasource struct { 
	// Platform - The platform that was the source of the data.  Example: a CRM like SALESFORCE.
	Platform *string `json:"platform,omitempty"`


	// Url - An URL that links to the source record that contributed data to the associated entity.
	Url *string `json:"url,omitempty"`

}

func (o *Externaldatasource) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externaldatasource
	
	return json.Marshal(&struct { 
		Platform *string `json:"platform,omitempty"`
		
		Url *string `json:"url,omitempty"`
		*Alias
	}{ 
		Platform: o.Platform,
		
		Url: o.Url,
		Alias:    (*Alias)(o),
	})
}

func (o *Externaldatasource) UnmarshalJSON(b []byte) error {
	var ExternaldatasourceMap map[string]interface{}
	err := json.Unmarshal(b, &ExternaldatasourceMap)
	if err != nil {
		return err
	}
	
	if Platform, ok := ExternaldatasourceMap["platform"].(string); ok {
		o.Platform = &Platform
	}
	
	if Url, ok := ExternaldatasourceMap["url"].(string); ok {
		o.Url = &Url
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Externaldatasource) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
