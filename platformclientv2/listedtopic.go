package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
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

func (u *Listedtopic) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Listedtopic

	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Published *bool `json:"published,omitempty"`
		
		Strictness *string `json:"strictness,omitempty"`
		
		ProgramsCount *int `json:"programsCount,omitempty"`
		
		Tags *[]string `json:"tags,omitempty"`
		
		Dialect *string `json:"dialect,omitempty"`
		
		Participants *string `json:"participants,omitempty"`
		
		PhrasesCount *int `json:"phrasesCount,omitempty"`
		
		ModifiedBy *Addressableentityref `json:"modifiedBy,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Description: u.Description,
		
		Published: u.Published,
		
		Strictness: u.Strictness,
		
		ProgramsCount: u.ProgramsCount,
		
		Tags: u.Tags,
		
		Dialect: u.Dialect,
		
		Participants: u.Participants,
		
		PhrasesCount: u.PhrasesCount,
		
		ModifiedBy: u.ModifiedBy,
		
		DateModified: DateModified,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Listedtopic) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
