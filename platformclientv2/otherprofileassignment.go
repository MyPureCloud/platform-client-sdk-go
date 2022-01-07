package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Otherprofileassignment
type Otherprofileassignment struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// CurrentProfile - The current performance profile that this user belongs to
	CurrentProfile *Domainentityref `json:"currentProfile,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Otherprofileassignment) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Otherprofileassignment
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		CurrentProfile *Domainentityref `json:"currentProfile,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		CurrentProfile: o.CurrentProfile,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Otherprofileassignment) UnmarshalJSON(b []byte) error {
	var OtherprofileassignmentMap map[string]interface{}
	err := json.Unmarshal(b, &OtherprofileassignmentMap)
	if err != nil {
		return err
	}
	
	if Id, ok := OtherprofileassignmentMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if CurrentProfile, ok := OtherprofileassignmentMap["currentProfile"].(map[string]interface{}); ok {
		CurrentProfileString, _ := json.Marshal(CurrentProfile)
		json.Unmarshal(CurrentProfileString, &o.CurrentProfile)
	}
	
	if SelfUri, ok := OtherprofileassignmentMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Otherprofileassignment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
