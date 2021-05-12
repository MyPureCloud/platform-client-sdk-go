package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Localencryptionconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
