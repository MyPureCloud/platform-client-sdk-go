package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Script
type Script struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// VersionId
	VersionId *string `json:"versionId,omitempty"`


	// CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`


	// ModifiedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`


	// PublishedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	PublishedDate *time.Time `json:"publishedDate,omitempty"`


	// VersionDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	VersionDate *time.Time `json:"versionDate,omitempty"`


	// StartPageId
	StartPageId *string `json:"startPageId,omitempty"`


	// StartPageName
	StartPageName *string `json:"startPageName,omitempty"`


	// Features
	Features *interface{} `json:"features,omitempty"`


	// Variables
	Variables *interface{} `json:"variables,omitempty"`


	// CustomActions
	CustomActions *interface{} `json:"customActions,omitempty"`


	// Pages
	Pages *[]Page `json:"pages,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Script) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Script

	
	CreatedDate := new(string)
	if u.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(u.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	ModifiedDate := new(string)
	if u.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(u.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	PublishedDate := new(string)
	if u.PublishedDate != nil {
		
		*PublishedDate = timeutil.Strftime(u.PublishedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		PublishedDate = nil
	}
	
	VersionDate := new(string)
	if u.VersionDate != nil {
		
		*VersionDate = timeutil.Strftime(u.VersionDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		VersionDate = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VersionId *string `json:"versionId,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		PublishedDate *string `json:"publishedDate,omitempty"`
		
		VersionDate *string `json:"versionDate,omitempty"`
		
		StartPageId *string `json:"startPageId,omitempty"`
		
		StartPageName *string `json:"startPageName,omitempty"`
		
		Features *interface{} `json:"features,omitempty"`
		
		Variables *interface{} `json:"variables,omitempty"`
		
		CustomActions *interface{} `json:"customActions,omitempty"`
		
		Pages *[]Page `json:"pages,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		VersionId: u.VersionId,
		
		CreatedDate: CreatedDate,
		
		ModifiedDate: ModifiedDate,
		
		PublishedDate: PublishedDate,
		
		VersionDate: VersionDate,
		
		StartPageId: u.StartPageId,
		
		StartPageName: u.StartPageName,
		
		Features: u.Features,
		
		Variables: u.Variables,
		
		CustomActions: u.CustomActions,
		
		Pages: u.Pages,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Script) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
