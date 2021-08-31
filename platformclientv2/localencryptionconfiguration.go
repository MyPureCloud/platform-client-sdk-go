package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Localencryptionconfiguration
type Localencryptionconfiguration struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Url - The url for decryption. This must specify the path to where Purecloud can requests decryption
	Url *string `json:"url,omitempty"`


	// ApiId - The api id for Hawk Authentication.
	ApiId *string `json:"apiId,omitempty"`


	// ApiKey - The api shared symmetric key used for hawk authentication
	ApiKey *string `json:"apiKey,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Localencryptionconfiguration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Localencryptionconfiguration
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Url *string `json:"url,omitempty"`
		
		ApiId *string `json:"apiId,omitempty"`
		
		ApiKey *string `json:"apiKey,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Url: o.Url,
		
		ApiId: o.ApiId,
		
		ApiKey: o.ApiKey,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Localencryptionconfiguration) UnmarshalJSON(b []byte) error {
	var LocalencryptionconfigurationMap map[string]interface{}
	err := json.Unmarshal(b, &LocalencryptionconfigurationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := LocalencryptionconfigurationMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := LocalencryptionconfigurationMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Url, ok := LocalencryptionconfigurationMap["url"].(string); ok {
		o.Url = &Url
	}
	
	if ApiId, ok := LocalencryptionconfigurationMap["apiId"].(string); ok {
		o.ApiId = &ApiId
	}
	
	if ApiKey, ok := LocalencryptionconfigurationMap["apiKey"].(string); ok {
		o.ApiKey = &ApiKey
	}
	
	if SelfUri, ok := LocalencryptionconfigurationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Localencryptionconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
