package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Chatmessageuser
type Chatmessageuser struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DisplayName
	DisplayName *string `json:"displayName,omitempty"`


	// Username
	Username *string `json:"username,omitempty"`


	// Images
	Images *[]Userimage `json:"images,omitempty"`

}

func (o *Chatmessageuser) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Chatmessageuser
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		
		Username *string `json:"username,omitempty"`
		
		Images *[]Userimage `json:"images,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DisplayName: o.DisplayName,
		
		Username: o.Username,
		
		Images: o.Images,
		Alias:    (*Alias)(o),
	})
}

func (o *Chatmessageuser) UnmarshalJSON(b []byte) error {
	var ChatmessageuserMap map[string]interface{}
	err := json.Unmarshal(b, &ChatmessageuserMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ChatmessageuserMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := ChatmessageuserMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if DisplayName, ok := ChatmessageuserMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
	
	if Username, ok := ChatmessageuserMap["username"].(string); ok {
		o.Username = &Username
	}
	
	if Images, ok := ChatmessageuserMap["images"].([]interface{}); ok {
		ImagesString, _ := json.Marshal(Images)
		json.Unmarshal(ImagesString, &o.Images)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Chatmessageuser) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
