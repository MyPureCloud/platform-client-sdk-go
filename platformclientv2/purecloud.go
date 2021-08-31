package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Purecloud
type Purecloud struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Disabled
	Disabled *bool `json:"disabled,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Purecloud) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Purecloud
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Disabled *bool `json:"disabled,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Disabled: o.Disabled,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Purecloud) UnmarshalJSON(b []byte) error {
	var PurecloudMap map[string]interface{}
	err := json.Unmarshal(b, &PurecloudMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PurecloudMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := PurecloudMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Disabled, ok := PurecloudMap["disabled"].(bool); ok {
		o.Disabled = &Disabled
	}
	
	if SelfUri, ok := PurecloudMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Purecloud) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
