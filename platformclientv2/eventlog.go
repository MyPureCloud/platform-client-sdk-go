package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Eventlog
type Eventlog struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ErrorEntity
	ErrorEntity *Domainentityref `json:"errorEntity,omitempty"`


	// RelatedEntity
	RelatedEntity *Domainentityref `json:"relatedEntity,omitempty"`


	// Timestamp - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Timestamp *time.Time `json:"timestamp,omitempty"`


	// Level
	Level *string `json:"level,omitempty"`


	// Category
	Category *string `json:"category,omitempty"`


	// CorrelationId
	CorrelationId *string `json:"correlationId,omitempty"`


	// EventMessage
	EventMessage *Eventmessage `json:"eventMessage,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Eventlog) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Eventlog

	
	Timestamp := new(string)
	if u.Timestamp != nil {
		
		*Timestamp = timeutil.Strftime(u.Timestamp, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Timestamp = nil
	}
	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ErrorEntity *Domainentityref `json:"errorEntity,omitempty"`
		
		RelatedEntity *Domainentityref `json:"relatedEntity,omitempty"`
		
		Timestamp *string `json:"timestamp,omitempty"`
		
		Level *string `json:"level,omitempty"`
		
		Category *string `json:"category,omitempty"`
		
		CorrelationId *string `json:"correlationId,omitempty"`
		
		EventMessage *Eventmessage `json:"eventMessage,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		ErrorEntity: u.ErrorEntity,
		
		RelatedEntity: u.RelatedEntity,
		
		Timestamp: Timestamp,
		
		Level: u.Level,
		
		Category: u.Category,
		
		CorrelationId: u.CorrelationId,
		
		EventMessage: u.EventMessage,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Eventlog) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
