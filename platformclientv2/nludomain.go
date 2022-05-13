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

func (o *Nludomain) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Nludomain
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: o.Id,
		
		Name: o.Name,
		
		Language: o.Language,
		
		DraftVersion: o.DraftVersion,
		
		LastPublishedVersion: o.LastPublishedVersion,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		EngineVersion: o.EngineVersion,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Nludomain) UnmarshalJSON(b []byte) error {
	var NludomainMap map[string]interface{}
	err := json.Unmarshal(b, &NludomainMap)
	if err != nil {
		return err
	}
	
	if Id, ok := NludomainMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := NludomainMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Language, ok := NludomainMap["language"].(string); ok {
		o.Language = &Language
	}
    
	if DraftVersion, ok := NludomainMap["draftVersion"].(map[string]interface{}); ok {
		DraftVersionString, _ := json.Marshal(DraftVersion)
		json.Unmarshal(DraftVersionString, &o.DraftVersion)
	}
	
	if LastPublishedVersion, ok := NludomainMap["lastPublishedVersion"].(map[string]interface{}); ok {
		LastPublishedVersionString, _ := json.Marshal(LastPublishedVersion)
		json.Unmarshal(LastPublishedVersionString, &o.LastPublishedVersion)
	}
	
	if dateCreatedString, ok := NludomainMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := NludomainMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if EngineVersion, ok := NludomainMap["engineVersion"].(string); ok {
		o.EngineVersion = &EngineVersion
	}
    
	if SelfUri, ok := NludomainMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Nludomain) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
