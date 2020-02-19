package platformclientv2
import (
	"time"
	"encoding/json"
)

// Trustuserdetails
type Trustuserdetails struct { 
	// DateCreated - Date Trust User was added. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// CreatedBy - User that added trusted user.
	CreatedBy *Orguser `json:"createdBy,omitempty"`

}

// String returns a JSON representation of the model
func (o *Trustuserdetails) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
