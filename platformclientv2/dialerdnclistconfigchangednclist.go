package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerdnclistconfigchangednclist
type Dialerdnclistconfigchangednclist struct { 
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


	// ImportStatus
	ImportStatus *Dialerdnclistconfigchangeimportstatus `json:"importStatus,omitempty"`


	// Size
	Size *int `json:"size,omitempty"`


	// DncSourceType
	DncSourceType *string `json:"dncSourceType,omitempty"`


	// LoginId
	LoginId *string `json:"loginId,omitempty"`


	// DncCodes
	DncCodes *[]string `json:"dncCodes,omitempty"`


	// LicenseId
	LicenseId *string `json:"licenseId,omitempty"`


	// ContactMethod
	ContactMethod *string `json:"contactMethod,omitempty"`


	// Division
	Division *Dialerdnclistconfigchangeurireference `json:"division,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dialerdnclistconfigchangednclist) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
