package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userroutinglanguage - Represents an organization language assigned to a user. When assigning to a user specify the organization language id as the id.
type Userroutinglanguage struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Proficiency - A rating from 0.0 to 5.0 that indicates how fluent an agent is in a particular language. ACD interactions are routed to agents with higher proficiency ratings.
	Proficiency *float64 `json:"proficiency,omitempty"`


	// State - Activate or deactivate this routing language.
	State *string `json:"state,omitempty"`


	// LanguageUri - URI to the organization language used by this user language.
	LanguageUri *string `json:"languageUri,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Userroutinglanguage) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userroutinglanguage
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Proficiency *float64 `json:"proficiency,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		LanguageUri *string `json:"languageUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Proficiency: o.Proficiency,
		
		State: o.State,
		
		LanguageUri: o.LanguageUri,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Userroutinglanguage) UnmarshalJSON(b []byte) error {
	var UserroutinglanguageMap map[string]interface{}
	err := json.Unmarshal(b, &UserroutinglanguageMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UserroutinglanguageMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := UserroutinglanguageMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Proficiency, ok := UserroutinglanguageMap["proficiency"].(float64); ok {
		o.Proficiency = &Proficiency
	}
	
	if State, ok := UserroutinglanguageMap["state"].(string); ok {
		o.State = &State
	}
	
	if LanguageUri, ok := UserroutinglanguageMap["languageUri"].(string); ok {
		o.LanguageUri = &LanguageUri
	}
	
	if SelfUri, ok := UserroutinglanguageMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userroutinglanguage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
