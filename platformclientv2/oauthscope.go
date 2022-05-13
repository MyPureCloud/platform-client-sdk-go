package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Oauthscope
type Oauthscope struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Oauthscope) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Oauthscope
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Description: o.Description,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Oauthscope) UnmarshalJSON(b []byte) error {
	var OauthscopeMap map[string]interface{}
	err := json.Unmarshal(b, &OauthscopeMap)
	if err != nil {
		return err
	}
	
	if Id, ok := OauthscopeMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Description, ok := OauthscopeMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if SelfUri, ok := OauthscopeMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Oauthscope) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
