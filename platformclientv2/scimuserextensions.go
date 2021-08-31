package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Scimuserextensions - Genesys Cloud user extensions to SCIM RFC.
type Scimuserextensions struct { 
	// RoutingSkills - The list of routing skills assigned to a user. Maximum 50 skills.
	RoutingSkills *[]Scimuserroutingskill `json:"routingSkills,omitempty"`


	// RoutingLanguages - The list of routing languages assigned to a user. Maximum 50 languages.
	RoutingLanguages *[]Scimuserroutinglanguage `json:"routingLanguages,omitempty"`


	// ExternalIds - The list of external identifiers assigned to user. Always includes an immutable SCIM authority prefixed with \"x-pc:scimv2:v1\".
	ExternalIds *[]Scimgenesysuserexternalid `json:"externalIds,omitempty"`

}

func (o *Scimuserextensions) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Scimuserextensions
	
	return json.Marshal(&struct { 
		RoutingSkills *[]Scimuserroutingskill `json:"routingSkills,omitempty"`
		
		RoutingLanguages *[]Scimuserroutinglanguage `json:"routingLanguages,omitempty"`
		
		ExternalIds *[]Scimgenesysuserexternalid `json:"externalIds,omitempty"`
		*Alias
	}{ 
		RoutingSkills: o.RoutingSkills,
		
		RoutingLanguages: o.RoutingLanguages,
		
		ExternalIds: o.ExternalIds,
		Alias:    (*Alias)(o),
	})
}

func (o *Scimuserextensions) UnmarshalJSON(b []byte) error {
	var ScimuserextensionsMap map[string]interface{}
	err := json.Unmarshal(b, &ScimuserextensionsMap)
	if err != nil {
		return err
	}
	
	if RoutingSkills, ok := ScimuserextensionsMap["routingSkills"].([]interface{}); ok {
		RoutingSkillsString, _ := json.Marshal(RoutingSkills)
		json.Unmarshal(RoutingSkillsString, &o.RoutingSkills)
	}
	
	if RoutingLanguages, ok := ScimuserextensionsMap["routingLanguages"].([]interface{}); ok {
		RoutingLanguagesString, _ := json.Marshal(RoutingLanguages)
		json.Unmarshal(RoutingLanguagesString, &o.RoutingLanguages)
	}
	
	if ExternalIds, ok := ScimuserextensionsMap["externalIds"].([]interface{}); ok {
		ExternalIdsString, _ := json.Marshal(ExternalIds)
		json.Unmarshal(ExternalIdsString, &o.ExternalIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Scimuserextensions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
