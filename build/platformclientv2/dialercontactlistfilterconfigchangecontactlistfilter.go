package platformclientv2
import (
	"time"
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
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialercontactlistfilterconfigchangecontactlistfilter) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
