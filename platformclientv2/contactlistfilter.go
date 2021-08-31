package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactlistfilter
type Contactlistfilter struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the list.
	Name *string `json:"name,omitempty"`


	// DateCreated - Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version - Required for updates, must match the version number of the most recent update
	Version *int `json:"version,omitempty"`


	// ContactList - The contact list the filter is based on.
	ContactList *Domainentityref `json:"contactList,omitempty"`


	// Clauses - Groups of conditions to filter the contacts by.
	Clauses *[]Contactlistfilterclause `json:"clauses,omitempty"`


	// FilterType - How to join clauses together.
	FilterType *string `json:"filterType,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Contactlistfilter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactlistfilter
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		ContactList *Domainentityref `json:"contactList,omitempty"`
		
		Clauses *[]Contactlistfilterclause `json:"clauses,omitempty"`
		
		FilterType *string `json:"filterType,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: o.Version,
		
		ContactList: o.ContactList,
		
		Clauses: o.Clauses,
		
		FilterType: o.FilterType,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Contactlistfilter) UnmarshalJSON(b []byte) error {
	var ContactlistfilterMap map[string]interface{}
	err := json.Unmarshal(b, &ContactlistfilterMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ContactlistfilterMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := ContactlistfilterMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if dateCreatedString, ok := ContactlistfilterMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := ContactlistfilterMap["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Version, ok := ContactlistfilterMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if ContactList, ok := ContactlistfilterMap["contactList"].(map[string]interface{}); ok {
		ContactListString, _ := json.Marshal(ContactList)
		json.Unmarshal(ContactListString, &o.ContactList)
	}
	
	if Clauses, ok := ContactlistfilterMap["clauses"].([]interface{}); ok {
		ClausesString, _ := json.Marshal(Clauses)
		json.Unmarshal(ClausesString, &o.Clauses)
	}
	
	if FilterType, ok := ContactlistfilterMap["filterType"].(string); ok {
		o.FilterType = &FilterType
	}
	
	if SelfUri, ok := ContactlistfilterMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contactlistfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
