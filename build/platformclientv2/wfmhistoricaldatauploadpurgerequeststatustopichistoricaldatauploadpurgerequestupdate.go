package platformclientv2
import (
	"encoding/json"
)

// Wfmhistoricaldatauploadpurgerequeststatustopichistoricaldatauploadpurgerequestupdate
type Wfmhistoricaldatauploadpurgerequeststatustopichistoricaldatauploadpurgerequestupdate struct { 
	// Status
	Status *string `json:"status,omitempty"`

}

// String returns a JSON representation of the model
func (o *Wfmhistoricaldatauploadpurgerequeststatustopichistoricaldatauploadpurgerequestupdate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
