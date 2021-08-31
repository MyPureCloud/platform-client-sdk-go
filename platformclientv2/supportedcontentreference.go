package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Supportedcontentreference - Reference to supported content profile associated with the integration
type Supportedcontentreference struct { 
	// Id - The SupportedContent unique identifier associated with this integration
	Id *string `json:"id,omitempty"`


	// Name - The SupportedContent profile name
	Name *string `json:"name,omitempty"`


	// SelfUri - The SupportedContent profile URI
	SelfUri *string `json:"selfUri,omitempty"`


	// MediaTypes - Media types definition for the supported content
	MediaTypes *Mediatypes `json:"mediaTypes,omitempty"`

}

func (o *Supportedcontentreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Supportedcontentreference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		MediaTypes *Mediatypes `json:"mediaTypes,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SelfUri: o.SelfUri,
		
		MediaTypes: o.MediaTypes,
		Alias:    (*Alias)(o),
	})
}

func (o *Supportedcontentreference) UnmarshalJSON(b []byte) error {
	var SupportedcontentreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &SupportedcontentreferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := SupportedcontentreferenceMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := SupportedcontentreferenceMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if SelfUri, ok := SupportedcontentreferenceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	
	if MediaTypes, ok := SupportedcontentreferenceMap["mediaTypes"].(map[string]interface{}); ok {
		MediaTypesString, _ := json.Marshal(MediaTypes)
		json.Unmarshal(MediaTypesString, &o.MediaTypes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Supportedcontentreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
