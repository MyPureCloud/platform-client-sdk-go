package platformclientv2
import (
	"encoding/json"
)

// Architectflownotificationhomeorganization
type Architectflownotificationhomeorganization struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ThirdPartyOrgName
	ThirdPartyOrgName *string `json:"thirdPartyOrgName,omitempty"`

}

// String returns a JSON representation of the model
func (o *Architectflownotificationhomeorganization) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
