package platformclientv2
import (
	"time"
	"encoding/json"
)

// Dnclistcreate
type Dnclistcreate struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the DncList.
	Name *string `json:"name,omitempty"`


	// DateCreated - Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`


	// DateModified - Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateModified *time.Time `json:"dateModified,omitempty"`


	// Version - Required for updates, must match the version number of the most recent update
	Version *int32 `json:"version,omitempty"`


	// ImportStatus - The status of the import process
	ImportStatus *Importstatus `json:"importStatus,omitempty"`


	// Size - The total number of phone numbers in the DncList.
	Size *int64 `json:"size,omitempty"`


	// DncSourceType - The type of the DncList.
	DncSourceType *string `json:"dncSourceType,omitempty"`


	// LoginId - A dnc.com loginId. Required if the dncSourceType is dnc.com.
	LoginId *string `json:"loginId,omitempty"`


	// DncCodes - The list of dnc.com codes to be treated as DNC. Required if the dncSourceType is dnc.com.
	DncCodes *[]string `json:"dncCodes,omitempty"`


	// LicenseId - A gryphon license number. Required if the dncSourceType is gryphon.
	LicenseId *string `json:"licenseId,omitempty"`


	// Division - The division this DncList belongs to.
	Division *Domainentityref `json:"division,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Dnclistcreate) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
