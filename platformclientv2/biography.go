package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Biography
type Biography struct { 
	// Biography - Personal detailed description
	Biography *string `json:"biography,omitempty"`


	// Interests
	Interests *[]string `json:"interests,omitempty"`


	// Hobbies
	Hobbies *[]string `json:"hobbies,omitempty"`


	// Spouse
	Spouse *string `json:"spouse,omitempty"`


	// Education - User education details
	Education *[]Education `json:"education,omitempty"`

}

func (o *Biography) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Biography
	
	return json.Marshal(&struct { 
		Biography *string `json:"biography,omitempty"`
		
		Interests *[]string `json:"interests,omitempty"`
		
		Hobbies *[]string `json:"hobbies,omitempty"`
		
		Spouse *string `json:"spouse,omitempty"`
		
		Education *[]Education `json:"education,omitempty"`
		*Alias
	}{ 
		Biography: o.Biography,
		
		Interests: o.Interests,
		
		Hobbies: o.Hobbies,
		
		Spouse: o.Spouse,
		
		Education: o.Education,
		Alias:    (*Alias)(o),
	})
}

func (o *Biography) UnmarshalJSON(b []byte) error {
	var BiographyMap map[string]interface{}
	err := json.Unmarshal(b, &BiographyMap)
	if err != nil {
		return err
	}
	
	if Biography, ok := BiographyMap["biography"].(string); ok {
		o.Biography = &Biography
	}
	
	if Interests, ok := BiographyMap["interests"].([]interface{}); ok {
		InterestsString, _ := json.Marshal(Interests)
		json.Unmarshal(InterestsString, &o.Interests)
	}
	
	if Hobbies, ok := BiographyMap["hobbies"].([]interface{}); ok {
		HobbiesString, _ := json.Marshal(Hobbies)
		json.Unmarshal(HobbiesString, &o.Hobbies)
	}
	
	if Spouse, ok := BiographyMap["spouse"].(string); ok {
		o.Spouse = &Spouse
	}
	
	if Education, ok := BiographyMap["education"].([]interface{}); ok {
		EducationString, _ := json.Marshal(Education)
		json.Unmarshal(EducationString, &o.Education)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Biography) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
