package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialercontactlistfilterconfigchangecontactlistfilter
type Dialercontactlistfilterconfigchangecontactlistfilter struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version
	Version *int `json:"version,omitempty"`


	// ContactList
	ContactList *Dialercontactlistfilterconfigchangeurireference `json:"contactList,omitempty"`


	// ContactListColumns
	ContactListColumns *[]string `json:"contactListColumns,omitempty"`


	// Clauses
	Clauses *[]Dialercontactlistfilterconfigchangefilterclause `json:"clauses,omitempty"`


	// FilterType
	FilterType *string `json:"filterType,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (u *Dialercontactlistfilterconfigchangecontactlistfilter) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialercontactlistfilterconfigchangecontactlistfilter

	
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
		
		Version *int `json:"version,omitempty"`
		
		ContactList *Dialercontactlistfilterconfigchangeurireference `json:"contactList,omitempty"`
		
		ContactListColumns *[]string `json:"contactListColumns,omitempty"`
		
		Clauses *[]Dialercontactlistfilterconfigchangefilterclause `json:"clauses,omitempty"`
		
		FilterType *string `json:"filterType,omitempty"`
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Version: u.Version,
		
		ContactList: u.ContactList,
		
		ContactListColumns: u.ContactListColumns,
		
		Clauses: u.Clauses,
		
		FilterType: u.FilterType,
		
		AdditionalProperties: u.AdditionalProperties,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Dialercontactlistfilterconfigchangecontactlistfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
