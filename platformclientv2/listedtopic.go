package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Listedtopic
type Listedtopic struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// Published
	Published *bool `json:"published,omitempty"`


	// Strictness
	Strictness *string `json:"strictness,omitempty"`


	// ProgramsCount
	ProgramsCount *int `json:"programsCount,omitempty"`


	// Tags
	Tags *[]string `json:"tags,omitempty"`


	// Dialect
	Dialect *string `json:"dialect,omitempty"`


	// Participants
	Participants *string `json:"participants,omitempty"`


	// PhrasesCount
	PhrasesCount *int `json:"phrasesCount,omitempty"`


	// ModifiedBy
	ModifiedBy *Addressableentityref `json:"modifiedBy,omitempty"`


	// DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Listedtopic) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
