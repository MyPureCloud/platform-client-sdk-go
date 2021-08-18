package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Reportmetadata
type Reportmetadata struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Title
	Title *string `json:"title,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Keywords
	Keywords *[]string `json:"keywords,omitempty"`


	// AvailableLocales
	AvailableLocales *[]string `json:"availableLocales,omitempty"`


	// Parameters
	Parameters *[]Parameter `json:"parameters,omitempty"`


	// ExampleUrl
	ExampleUrl *string `json:"exampleUrl,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Reportmetadata) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Reportmetadata

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Keywords *[]string `json:"keywords,omitempty"`
		
		AvailableLocales *[]string `json:"availableLocales,omitempty"`
		
		Parameters *[]Parameter `json:"parameters,omitempty"`
		
		ExampleUrl *string `json:"exampleUrl,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Title: u.Title,
		
		Description: u.Description,
		
		Keywords: u.Keywords,
		
		AvailableLocales: u.AvailableLocales,
		
		Parameters: u.Parameters,
		
		ExampleUrl: u.ExampleUrl,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Reportmetadata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
