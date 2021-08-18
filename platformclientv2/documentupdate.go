package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentupdate
type Documentupdate struct { 
	// ChangeNumber
	ChangeNumber *int `json:"changeNumber,omitempty"`


	// Name - The name of the document
	Name *string `json:"name,omitempty"`


	// Read
	Read *bool `json:"read,omitempty"`


	// AddTags
	AddTags *[]string `json:"addTags,omitempty"`


	// RemoveTags
	RemoveTags *[]string `json:"removeTags,omitempty"`


	// AddTagIds
	AddTagIds *[]string `json:"addTagIds,omitempty"`


	// RemoveTagIds
	RemoveTagIds *[]string `json:"removeTagIds,omitempty"`


	// UpdateAttributes
	UpdateAttributes *[]Documentattribute `json:"updateAttributes,omitempty"`


	// RemoveAttributes
	RemoveAttributes *[]string `json:"removeAttributes,omitempty"`

}

func (u *Documentupdate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Documentupdate

	

	return json.Marshal(&struct { 
		ChangeNumber *int `json:"changeNumber,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Read *bool `json:"read,omitempty"`
		
		AddTags *[]string `json:"addTags,omitempty"`
		
		RemoveTags *[]string `json:"removeTags,omitempty"`
		
		AddTagIds *[]string `json:"addTagIds,omitempty"`
		
		RemoveTagIds *[]string `json:"removeTagIds,omitempty"`
		
		UpdateAttributes *[]Documentattribute `json:"updateAttributes,omitempty"`
		
		RemoveAttributes *[]string `json:"removeAttributes,omitempty"`
		*Alias
	}{ 
		ChangeNumber: u.ChangeNumber,
		
		Name: u.Name,
		
		Read: u.Read,
		
		AddTags: u.AddTags,
		
		RemoveTags: u.RemoveTags,
		
		AddTagIds: u.AddTagIds,
		
		RemoveTagIds: u.RemoveTagIds,
		
		UpdateAttributes: u.UpdateAttributes,
		
		RemoveAttributes: u.RemoveAttributes,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Documentupdate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
