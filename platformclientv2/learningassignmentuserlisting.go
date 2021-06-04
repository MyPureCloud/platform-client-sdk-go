package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Learningassignmentuserlisting - List of users matching the learning module rule
type Learningassignmentuserlisting struct { 
	// Entities
	Entities *[]Learningassignmentuser `json:"entities,omitempty"`


	// PageSize
	PageSize *int `json:"pageSize,omitempty"`


	// PageNumber
	PageNumber *int `json:"pageNumber,omitempty"`


	// Total - The number of users matching search term
	Total *int `json:"total,omitempty"`


	// UnfilteredTotal - The total number of users
	UnfilteredTotal *int `json:"unfilteredTotal,omitempty"`


	// FirstUri
	FirstUri *string `json:"firstUri,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// NextUri
	NextUri *string `json:"nextUri,omitempty"`


	// PreviousUri
	PreviousUri *string `json:"previousUri,omitempty"`


	// LastUri
	LastUri *string `json:"lastUri,omitempty"`


	// PageCount
	PageCount *int `json:"pageCount,omitempty"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentuserlisting) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
