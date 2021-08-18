package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Domainedgesoftwareversiondto
type Domainedgesoftwareversiondto struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// EdgeVersion
	EdgeVersion *string `json:"edgeVersion,omitempty"`


	// PublishDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	PublishDate *time.Time `json:"publishDate,omitempty"`


	// EdgeUri
	EdgeUri *string `json:"edgeUri,omitempty"`


	// LatestRelease
	LatestRelease *bool `json:"latestRelease,omitempty"`


	// Current
	Current *bool `json:"current,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Domainedgesoftwareversiondto) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Domainedgesoftwareversiondto

	
	PublishDate := new(string)
	if u.PublishDate != nil {
		
		*PublishDate = timeutil.Strftime(u.PublishDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		PublishDate = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		EdgeVersion *string `json:"edgeVersion,omitempty"`
		
		PublishDate *string `json:"publishDate,omitempty"`
		
		EdgeUri *string `json:"edgeUri,omitempty"`
		
		LatestRelease *bool `json:"latestRelease,omitempty"`
		
		Current *bool `json:"current,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		EdgeVersion: u.EdgeVersion,
		
		PublishDate: PublishDate,
		
		EdgeUri: u.EdgeUri,
		
		LatestRelease: u.LatestRelease,
		
		Current: u.Current,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Domainedgesoftwareversiondto) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
