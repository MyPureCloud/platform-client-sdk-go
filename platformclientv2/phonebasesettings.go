package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Phonebasesettings
type Phonebasesettings struct { 
	// Id - The globally unique identifier for this phone base settings
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Phonebasesettings) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Phonebasesettings
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Phonebasesettings) UnmarshalJSON(b []byte) error {
	var PhonebasesettingsMap map[string]interface{}
	err := json.Unmarshal(b, &PhonebasesettingsMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PhonebasesettingsMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := PhonebasesettingsMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SelfUri, ok := PhonebasesettingsMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Phonebasesettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
