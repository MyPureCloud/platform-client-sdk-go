package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Topic
type Topic struct { 
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


	// Programs
	Programs *[]Baseprogramentity `json:"programs,omitempty"`


	// Tags
	Tags *[]string `json:"tags,omitempty"`


	// Dialect
	Dialect *string `json:"dialect,omitempty"`


	// Participants
	Participants *string `json:"participants,omitempty"`


	// Phrases
	Phrases *[]Phrase `json:"phrases,omitempty"`


	// ModifiedBy
	ModifiedBy *Addressableentityref `json:"modifiedBy,omitempty"`


	// DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// PublishedBy
	PublishedBy *Addressableentityref `json:"publishedBy,omitempty"`


	// DatePublished - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DatePublished *time.Time `json:"datePublished,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Topic) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Topic

	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DatePublished := new(string)
	if u.DatePublished != nil {
		
		*DatePublished = timeutil.Strftime(u.DatePublished, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DatePublished = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Published *bool `json:"published,omitempty"`
		
		Strictness *string `json:"strictness,omitempty"`
		
		Programs *[]Baseprogramentity `json:"programs,omitempty"`
		
		Tags *[]string `json:"tags,omitempty"`
		
		Dialect *string `json:"dialect,omitempty"`
		
		Participants *string `json:"participants,omitempty"`
		
		Phrases *[]Phrase `json:"phrases,omitempty"`
		
		ModifiedBy *Addressableentityref `json:"modifiedBy,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		PublishedBy *Addressableentityref `json:"publishedBy,omitempty"`
		
		DatePublished *string `json:"datePublished,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Description: u.Description,
		
		Published: u.Published,
		
		Strictness: u.Strictness,
		
		Programs: u.Programs,
		
		Tags: u.Tags,
		
		Dialect: u.Dialect,
		
		Participants: u.Participants,
		
		Phrases: u.Phrases,
		
		ModifiedBy: u.ModifiedBy,
		
		DateModified: DateModified,
		
		PublishedBy: u.PublishedBy,
		
		DatePublished: DatePublished,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Topic) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
