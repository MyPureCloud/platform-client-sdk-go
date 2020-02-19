package platformclientv2
import (
	"encoding/json"
)

// Architectflowoutcomenotificationhomeorganization
type Architectflowoutcomenotificationhomeorganization struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ThirdPartyOrgName
	ThirdPartyOrgName *string `json:"thirdPartyOrgName,omitempty"`

}

// String returns a JSON representation of the model
func (o *Architectflowoutcomenotificationhomeorganization) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
