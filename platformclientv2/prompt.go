package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Prompt
type Prompt struct { 
	// Id - The prompt identifier
	Id *string `json:"id,omitempty"`


	// Name - The prompt name.
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Resources - List of resources associated with this prompt
	Resources *[]Promptasset `json:"resources,omitempty"`


	// CurrentOperation - Current prompt operation status
	CurrentOperation *Operation `json:"currentOperation,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Prompt) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Prompt

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Resources *[]Promptasset `json:"resources,omitempty"`
		
		CurrentOperation *Operation `json:"currentOperation,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Description: u.Description,
		
		Resources: u.Resources,
		
		CurrentOperation: u.CurrentOperation,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Prompt) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
