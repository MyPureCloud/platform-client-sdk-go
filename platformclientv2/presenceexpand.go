package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Presenceexpand
type Presenceexpand struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Presences - An array of user presences
	Presences *[]Userpresence `json:"presences,omitempty"`


	// OutOfOffices - An array of out of office statuses
	OutOfOffices *[]Outofoffice `json:"outOfOffices,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Presenceexpand) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Presenceexpand
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Presences *[]Userpresence `json:"presences,omitempty"`
		
		OutOfOffices *[]Outofoffice `json:"outOfOffices,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Presences: o.Presences,
		
		OutOfOffices: o.OutOfOffices,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Presenceexpand) UnmarshalJSON(b []byte) error {
	var PresenceexpandMap map[string]interface{}
	err := json.Unmarshal(b, &PresenceexpandMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PresenceexpandMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := PresenceexpandMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Presences, ok := PresenceexpandMap["presences"].([]interface{}); ok {
		PresencesString, _ := json.Marshal(Presences)
		json.Unmarshal(PresencesString, &o.Presences)
	}
	
	if OutOfOffices, ok := PresenceexpandMap["outOfOffices"].([]interface{}); ok {
		OutOfOfficesString, _ := json.Marshal(OutOfOffices)
		json.Unmarshal(OutOfOfficesString, &o.OutOfOffices)
	}
	
	if SelfUri, ok := PresenceexpandMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Presenceexpand) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
