package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createresponseassetresponse
type Createresponseassetresponse struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Url - Pre-signed URL to PUT the file to
	Url *string `json:"url,omitempty"`


	// Headers - Required headers when uploading a file through PUT request to the URL
	Headers *map[string]string `json:"headers,omitempty"`

}

func (o *Createresponseassetresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createresponseassetresponse
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		Headers *map[string]string `json:"headers,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Url: o.Url,
		
		Headers: o.Headers,
		Alias:    (*Alias)(o),
	})
}

func (o *Createresponseassetresponse) UnmarshalJSON(b []byte) error {
	var CreateresponseassetresponseMap map[string]interface{}
	err := json.Unmarshal(b, &CreateresponseassetresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CreateresponseassetresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Url, ok := CreateresponseassetresponseMap["url"].(string); ok {
		o.Url = &Url
	}
    
	if Headers, ok := CreateresponseassetresponseMap["headers"].(map[string]interface{}); ok {
		HeadersString, _ := json.Marshal(Headers)
		json.Unmarshal(HeadersString, &o.Headers)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createresponseassetresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
