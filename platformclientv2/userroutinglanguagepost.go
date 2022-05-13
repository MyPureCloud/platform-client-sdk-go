package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userroutinglanguagepost - Represents an organization language assigned to a user. When assigning to a user specify the organization langauge id as the id.
type Userroutinglanguagepost struct { 
	// Id - The id of the existing routing language to add to the user
	Id *string `json:"id,omitempty"`


	// Proficiency - Proficiency is a rating from 0.0 to 5.0 on how competent an agent is for a particular language. It is used when a queue is set to \"Best available language\" mode to allow acd interactions to target agents with higher proficiency ratings.
	Proficiency *float64 `json:"proficiency,omitempty"`


	// LanguageUri - URI to the organization language used by this user language.
	LanguageUri *string `json:"languageUri,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Userroutinglanguagepost) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userroutinglanguagepost
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Proficiency *float64 `json:"proficiency,omitempty"`
		
		LanguageUri *string `json:"languageUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Proficiency: o.Proficiency,
		
		LanguageUri: o.LanguageUri,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Userroutinglanguagepost) UnmarshalJSON(b []byte) error {
	var UserroutinglanguagepostMap map[string]interface{}
	err := json.Unmarshal(b, &UserroutinglanguagepostMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UserroutinglanguagepostMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Proficiency, ok := UserroutinglanguagepostMap["proficiency"].(float64); ok {
		o.Proficiency = &Proficiency
	}
    
	if LanguageUri, ok := UserroutinglanguagepostMap["languageUri"].(string); ok {
		o.LanguageUri = &LanguageUri
	}
    
	if SelfUri, ok := UserroutinglanguagepostMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Userroutinglanguagepost) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
