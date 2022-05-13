package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Activitycodereference
type Activitycodereference struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// SecondaryPresences - The secondary presences of this activity code.
	SecondaryPresences *[]Secondarypresence `json:"secondaryPresences,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Activitycodereference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Activitycodereference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SecondaryPresences *[]Secondarypresence `json:"secondaryPresences,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SecondaryPresences: o.SecondaryPresences,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Activitycodereference) UnmarshalJSON(b []byte) error {
	var ActivitycodereferenceMap map[string]interface{}
	err := json.Unmarshal(b, &ActivitycodereferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ActivitycodereferenceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ActivitycodereferenceMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SecondaryPresences, ok := ActivitycodereferenceMap["secondaryPresences"].([]interface{}); ok {
		SecondaryPresencesString, _ := json.Marshal(SecondaryPresences)
		json.Unmarshal(SecondaryPresencesString, &o.SecondaryPresences)
	}
	
	if SelfUri, ok := ActivitycodereferenceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Activitycodereference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
