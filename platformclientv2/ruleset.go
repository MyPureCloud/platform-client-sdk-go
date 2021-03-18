package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Ruleset
type Ruleset struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the RuleSet.
	Name *string `json:"name,omitempty"`


	// DateCreated - Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version - Required for updates, must match the version number of the most recent update
	Version *int `json:"version,omitempty"`


	// ContactList - A ContactList to provide user-interface suggestions for contact columns on relevant conditions and actions.
	ContactList *Domainentityref `json:"contactList,omitempty"`


	// Queue - A Queue to provide user-interface suggestions for wrap-up codes on relevant conditions and actions.
	Queue *Domainentityref `json:"queue,omitempty"`


	// Rules - The list of rules.
	Rules *[]Dialerrule `json:"rules,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Ruleset) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
