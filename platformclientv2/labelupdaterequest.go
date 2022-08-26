package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Labelupdaterequest
type Labelupdaterequest struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the label.
	Name *string `json:"name,omitempty"`


	// Color - The color for the label.
	Color *string `json:"color,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Labelupdaterequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Labelupdaterequest
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Color *string `json:"color,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Color: o.Color,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Labelupdaterequest) UnmarshalJSON(b []byte) error {
	var LabelupdaterequestMap map[string]interface{}
	err := json.Unmarshal(b, &LabelupdaterequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := LabelupdaterequestMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := LabelupdaterequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Color, ok := LabelupdaterequestMap["color"].(string); ok {
		o.Color = &Color
	}
    
	if SelfUri, ok := LabelupdaterequestMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Labelupdaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
