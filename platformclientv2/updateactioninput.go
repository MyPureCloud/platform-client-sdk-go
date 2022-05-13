package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Updateactioninput
type Updateactioninput struct { 
	// Category - Category of action, Can be up to 256 characters long
	Category *string `json:"category,omitempty"`


	// Name - Name of action, Can be up to 256 characters long
	Name *string `json:"name,omitempty"`


	// Config - Configuration to support request and response processing
	Config *Actionconfig `json:"config,omitempty"`


	// Version - Version of this action
	Version *int `json:"version,omitempty"`

}

func (o *Updateactioninput) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Updateactioninput
	
	return json.Marshal(&struct { 
		Category *string `json:"category,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Config *Actionconfig `json:"config,omitempty"`
		
		Version *int `json:"version,omitempty"`
		*Alias
	}{ 
		Category: o.Category,
		
		Name: o.Name,
		
		Config: o.Config,
		
		Version: o.Version,
		Alias:    (*Alias)(o),
	})
}

func (o *Updateactioninput) UnmarshalJSON(b []byte) error {
	var UpdateactioninputMap map[string]interface{}
	err := json.Unmarshal(b, &UpdateactioninputMap)
	if err != nil {
		return err
	}
	
	if Category, ok := UpdateactioninputMap["category"].(string); ok {
		o.Category = &Category
	}
    
	if Name, ok := UpdateactioninputMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Config, ok := UpdateactioninputMap["config"].(map[string]interface{}); ok {
		ConfigString, _ := json.Marshal(Config)
		json.Unmarshal(ConfigString, &o.Config)
	}
	
	if Version, ok := UpdateactioninputMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Updateactioninput) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
