package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Greetingmediainfo
type Greetingmediainfo struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// MediaFileUri
	MediaFileUri *string `json:"mediaFileUri,omitempty"`


	// MediaImageUri
	MediaImageUri *string `json:"mediaImageUri,omitempty"`

}

func (o *Greetingmediainfo) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Greetingmediainfo
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		MediaFileUri *string `json:"mediaFileUri,omitempty"`
		
		MediaImageUri *string `json:"mediaImageUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		MediaFileUri: o.MediaFileUri,
		
		MediaImageUri: o.MediaImageUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Greetingmediainfo) UnmarshalJSON(b []byte) error {
	var GreetingmediainfoMap map[string]interface{}
	err := json.Unmarshal(b, &GreetingmediainfoMap)
	if err != nil {
		return err
	}
	
	if Id, ok := GreetingmediainfoMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if MediaFileUri, ok := GreetingmediainfoMap["mediaFileUri"].(string); ok {
		o.MediaFileUri = &MediaFileUri
	}
	
	if MediaImageUri, ok := GreetingmediainfoMap["mediaImageUri"].(string); ok {
		o.MediaImageUri = &MediaImageUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Greetingmediainfo) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
