package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Signedurlresponse
type Signedurlresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Url - Url of the downloaded pcap file
	Url *string `json:"url,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Signedurlresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Signedurlresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Url: o.Url,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Signedurlresponse) UnmarshalJSON(b []byte) error {
	var SignedurlresponseMap map[string]interface{}
	err := json.Unmarshal(b, &SignedurlresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SignedurlresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := SignedurlresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Url, ok := SignedurlresponseMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if SelfUri, ok := SignedurlresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Signedurlresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
