package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Learningmodulerule
type Learningmodulerule struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// IsActive - If true, rule is active
	IsActive *bool `json:"isActive,omitempty"`


	// Parts - The parts of a learning module rule
	Parts *[]Learningmoduleruleparts `json:"parts,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Learningmodulerule) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Learningmodulerule

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		IsActive *bool `json:"isActive,omitempty"`
		
		Parts *[]Learningmoduleruleparts `json:"parts,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		IsActive: u.IsActive,
		
		Parts: u.Parts,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Learningmodulerule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
