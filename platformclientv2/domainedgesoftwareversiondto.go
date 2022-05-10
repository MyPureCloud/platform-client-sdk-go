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


	// Current
	Current *bool `json:"current,omitempty"`


	// LatestRelease
	LatestRelease *bool `json:"latestRelease,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Domainedgesoftwareversiondto) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Domainedgesoftwareversiondto
	
	PublishDate := new(string)
	if o.PublishDate != nil {
		
		*PublishDate = timeutil.Strftime(o.PublishDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		PublishDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		EdgeVersion *string `json:"edgeVersion,omitempty"`
		
		PublishDate *string `json:"publishDate,omitempty"`
		
		EdgeUri *string `json:"edgeUri,omitempty"`
		
		Current *bool `json:"current,omitempty"`
		
		LatestRelease *bool `json:"latestRelease,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		EdgeVersion: o.EdgeVersion,
		
		PublishDate: PublishDate,
		
		EdgeUri: o.EdgeUri,
		
		Current: o.Current,
		
		LatestRelease: o.LatestRelease,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Domainedgesoftwareversiondto) UnmarshalJSON(b []byte) error {
	var DomainedgesoftwareversiondtoMap map[string]interface{}
	err := json.Unmarshal(b, &DomainedgesoftwareversiondtoMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DomainedgesoftwareversiondtoMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := DomainedgesoftwareversiondtoMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if EdgeVersion, ok := DomainedgesoftwareversiondtoMap["edgeVersion"].(string); ok {
		o.EdgeVersion = &EdgeVersion
	}
	
	if publishDateString, ok := DomainedgesoftwareversiondtoMap["publishDate"].(string); ok {
		PublishDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", publishDateString)
		o.PublishDate = &PublishDate
	}
	
	if EdgeUri, ok := DomainedgesoftwareversiondtoMap["edgeUri"].(string); ok {
		o.EdgeUri = &EdgeUri
	}
	
	if Current, ok := DomainedgesoftwareversiondtoMap["current"].(bool); ok {
		o.Current = &Current
	}
	
	if LatestRelease, ok := DomainedgesoftwareversiondtoMap["latestRelease"].(bool); ok {
		o.LatestRelease = &LatestRelease
	}
	
	if SelfUri, ok := DomainedgesoftwareversiondtoMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Domainedgesoftwareversiondto) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
