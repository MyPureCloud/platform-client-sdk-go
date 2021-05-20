package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialersequencescheduleconfigchangesequenceschedule
type Dialersequencescheduleconfigchangesequenceschedule struct { 
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


	// Intervals
	Intervals *[]Dialersequencescheduleconfigchangescheduleinterval `json:"intervals,omitempty"`


	// TimeZone
	TimeZone *string `json:"timeZone,omitempty"`


	// Sequence
	Sequence *Dialersequencescheduleconfigchangeurireference `json:"sequence,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialersequencescheduleconfigchangesequenceschedule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
