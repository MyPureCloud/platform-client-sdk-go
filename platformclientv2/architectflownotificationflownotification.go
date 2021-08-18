package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectflownotificationflownotification
type Architectflownotificationflownotification struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Deleted
	Deleted *bool `json:"deleted,omitempty"`


	// CheckedInVersion
	CheckedInVersion *Architectflownotificationflowversion `json:"checkedInVersion,omitempty"`


	// SavedVersion
	SavedVersion *Architectflownotificationflowversion `json:"savedVersion,omitempty"`


	// PublishedVersion
	PublishedVersion *Architectflownotificationflowversion `json:"publishedVersion,omitempty"`


	// CurrentOperation
	CurrentOperation *Architectflownotificationarchitectoperation `json:"currentOperation,omitempty"`

}

func (u *Architectflownotificationflownotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectflownotificationflownotification

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Deleted *bool `json:"deleted,omitempty"`
		
		CheckedInVersion *Architectflownotificationflowversion `json:"checkedInVersion,omitempty"`
		
		SavedVersion *Architectflownotificationflowversion `json:"savedVersion,omitempty"`
		
		PublishedVersion *Architectflownotificationflowversion `json:"publishedVersion,omitempty"`
		
		CurrentOperation *Architectflownotificationarchitectoperation `json:"currentOperation,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Description: u.Description,
		
		Deleted: u.Deleted,
		
		CheckedInVersion: u.CheckedInVersion,
		
		SavedVersion: u.SavedVersion,
		
		PublishedVersion: u.PublishedVersion,
		
		CurrentOperation: u.CurrentOperation,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Architectflownotificationflownotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
