package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Nludomain) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nludomain

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Language *string `json:"language,omitempty"`
		
		DraftVersion *Nludomainversion `json:"draftVersion,omitempty"`
		
		LastPublishedVersion *Nludomainversion `json:"lastPublishedVersion,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		EngineVersion *string `json:"engineVersion,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Language: u.Language,
		
		DraftVersion: u.DraftVersion,
		
		LastPublishedVersion: u.LastPublishedVersion,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		EngineVersion: u.EngineVersion,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Nludomain) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
