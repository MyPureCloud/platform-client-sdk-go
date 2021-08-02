package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Nludomain
type Nludomain struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the NLU domain.
	Name *string `json:"name,omitempty"`


	// Language - The language culture of the NLU domain, e.g. `en-us`, `de-de`.
	Language *string `json:"language,omitempty"`


	// DraftVersion - The draft version of that NLU domain.
	DraftVersion *Nludomainversion `json:"draftVersion,omitempty"`


	// LastPublishedVersion - The last published version of that NLU domain.
	LastPublishedVersion *Nludomainversion `json:"lastPublishedVersion,omitempty"`


	// DateCreated - The date when the NLU domain was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - The date when the NLU domain was updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// EngineVersion - The version of the NLU engine to use.
	EngineVersion *string `json:"engineVersion,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Nludomain) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
