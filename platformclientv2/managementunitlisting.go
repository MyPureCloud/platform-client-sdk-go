package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Managementunitlisting
type Managementunitlisting struct { 
	// Entities
	Entities *[]Managementunit `json:"entities,omitempty"`


	// PageSize - Deprecated, paging is not supported
	PageSize *int `json:"pageSize,omitempty"`


	// PageNumber - Deprecated, paging is not supported
	PageNumber *int `json:"pageNumber,omitempty"`


	// Total - Deprecated, paging is not supported
	Total *int `json:"total,omitempty"`


	// FirstUri - Deprecated, paging is not supported
	FirstUri *string `json:"firstUri,omitempty"`


	// NextUri - Deprecated, paging is not supported
	NextUri *string `json:"nextUri,omitempty"`


	// PageCount - Deprecated, paging is not supported
	PageCount *int `json:"pageCount,omitempty"`


	// PreviousUri - Deprecated, paging is not supported
	PreviousUri *string `json:"previousUri,omitempty"`


	// LastUri - Deprecated, paging is not supported
	LastUri *string `json:"lastUri,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Managementunitlisting) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Managementunitlisting

	

	return json.Marshal(&struct { 
		Entities *[]Managementunit `json:"entities,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		Total *int `json:"total,omitempty"`
		
		FirstUri *string `json:"firstUri,omitempty"`
		
		NextUri *string `json:"nextUri,omitempty"`
		
		PageCount *int `json:"pageCount,omitempty"`
		
		PreviousUri *string `json:"previousUri,omitempty"`
		
		LastUri *string `json:"lastUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Entities: u.Entities,
		
		PageSize: u.PageSize,
		
		PageNumber: u.PageNumber,
		
		Total: u.Total,
		
		FirstUri: u.FirstUri,
		
		NextUri: u.NextUri,
		
		PageCount: u.PageCount,
		
		PreviousUri: u.PreviousUri,
		
		LastUri: u.LastUri,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Managementunitlisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
