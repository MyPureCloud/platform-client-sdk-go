package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalcontactsunresolvedcontactchangedtopictwitterid
type Externalcontactsunresolvedcontactchangedtopictwitterid struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ScreenName
	ScreenName *string `json:"screenName,omitempty"`


	// Verified
	Verified *bool `json:"verified,omitempty"`


	// ProfileUrl
	ProfileUrl *string `json:"profileUrl,omitempty"`

}

func (o *Externalcontactsunresolvedcontactchangedtopictwitterid) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalcontactsunresolvedcontactchangedtopictwitterid
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ScreenName *string `json:"screenName,omitempty"`
		
		Verified *bool `json:"verified,omitempty"`
		
		ProfileUrl *string `json:"profileUrl,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ScreenName: o.ScreenName,
		
		Verified: o.Verified,
		
		ProfileUrl: o.ProfileUrl,
		Alias:    (*Alias)(o),
	})
}

func (o *Externalcontactsunresolvedcontactchangedtopictwitterid) UnmarshalJSON(b []byte) error {
	var ExternalcontactsunresolvedcontactchangedtopictwitteridMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalcontactsunresolvedcontactchangedtopictwitteridMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ExternalcontactsunresolvedcontactchangedtopictwitteridMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ExternalcontactsunresolvedcontactchangedtopictwitteridMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ScreenName, ok := ExternalcontactsunresolvedcontactchangedtopictwitteridMap["screenName"].(string); ok {
		o.ScreenName = &ScreenName
	}
    
	if Verified, ok := ExternalcontactsunresolvedcontactchangedtopictwitteridMap["verified"].(bool); ok {
		o.Verified = &Verified
	}
    
	if ProfileUrl, ok := ExternalcontactsunresolvedcontactchangedtopictwitteridMap["profileUrl"].(string); ok {
		o.ProfileUrl = &ProfileUrl
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalcontactsunresolvedcontactchangedtopictwitterid) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
