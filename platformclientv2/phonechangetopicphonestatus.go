package platformclientv2
import (
	"encoding/json"
)

// Phonechangetopicphonestatus
type Phonechangetopicphonestatus struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// OperationalStatus
	OperationalStatus *string `json:"operationalStatus,omitempty"`


	// Edge
	Edge *Phonechangetopicedgereference `json:"edge,omitempty"`


	// Provision
	Provision *Phonechangetopicprovisioninfo `json:"provision,omitempty"`


	// LineStatuses
	LineStatuses *[]Phonechangetopiclinestatus `json:"lineStatuses,omitempty"`


	// EventCreationTime
	EventCreationTime *Phonechangetopicoffsetdatetime `json:"eventCreationTime,omitempty"`

}

// String returns a JSON representation of the model
func (o *Phonechangetopicphonestatus) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
