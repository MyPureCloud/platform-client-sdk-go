package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Language
type Language struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The language name.
	Name *string `json:"name,omitempty"`


	// DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// State
	State *string `json:"state,omitempty"`


	// Version
	Version *string `json:"version,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Language) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Language
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Version *string `json:"version,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateModified: DateModified,
		
		State: o.State,
		
		Version: o.Version,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Language) UnmarshalJSON(b []byte) error {
	var LanguageMap map[string]interface{}
	err := json.Unmarshal(b, &LanguageMap)
	if err != nil {
		return err
	}
	
	if Id, ok := LanguageMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := LanguageMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateModifiedString, ok := LanguageMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if State, ok := LanguageMap["state"].(string); ok {
		o.State = &State
	}
    
	if Version, ok := LanguageMap["version"].(string); ok {
		o.Version = &Version
	}
    
	if SelfUri, ok := LanguageMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Language) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
