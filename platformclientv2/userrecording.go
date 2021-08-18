package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userrecording
type Userrecording struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// ContentUri
	ContentUri *string `json:"contentUri,omitempty"`


	// Workspace
	Workspace *Domainentityref `json:"workspace,omitempty"`


	// CreatedBy
	CreatedBy *Domainentityref `json:"createdBy,omitempty"`


	// Conversation
	Conversation *Conversation `json:"conversation,omitempty"`


	// ContentLength
	ContentLength *int `json:"contentLength,omitempty"`


	// DurationMilliseconds
	DurationMilliseconds *int `json:"durationMilliseconds,omitempty"`


	// Thumbnails
	Thumbnails *[]Documentthumbnail `json:"thumbnails,omitempty"`


	// Read
	Read *bool `json:"read,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Userrecording) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userrecording

	
	DateCreated := new(string)
	if u.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(u.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if u.DateModified != nil {
		
		*DateModified = timeutil.Strftime(u.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		ContentUri *string `json:"contentUri,omitempty"`
		
		Workspace *Domainentityref `json:"workspace,omitempty"`
		
		CreatedBy *Domainentityref `json:"createdBy,omitempty"`
		
		Conversation *Conversation `json:"conversation,omitempty"`
		
		ContentLength *int `json:"contentLength,omitempty"`
		
		DurationMilliseconds *int `json:"durationMilliseconds,omitempty"`
		
		Thumbnails *[]Documentthumbnail `json:"thumbnails,omitempty"`
		
		Read *bool `json:"read,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		ContentUri: u.ContentUri,
		
		Workspace: u.Workspace,
		
		CreatedBy: u.CreatedBy,
		
		Conversation: u.Conversation,
		
		ContentLength: u.ContentLength,
		
		DurationMilliseconds: u.DurationMilliseconds,
		
		Thumbnails: u.Thumbnails,
		
		Read: u.Read,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userrecording) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
