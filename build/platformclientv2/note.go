package platformclientv2
import (
	"time"
	"encoding/json"
	"strconv"
	"strings"
)

// Note
type Note struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// EntityId - The id of the contact or organization to which this note refers. This only needs to be set for input when using the Bulk APIs.
	EntityId *string `json:"entityId,omitempty"`


	// EntityType - This is only need to be set when using Bulk API. Using any other value than contact or organization will result in null being used.
	EntityType *string `json:"entityType,omitempty"`


	// NoteText
	NoteText *string `json:"noteText,omitempty"`


	// ModifyDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifyDate *time.Time `json:"modifyDate,omitempty"`


	// CreateDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreateDate *time.Time `json:"createDate,omitempty"`


	// CreatedBy - When creating or updating a note, only User.id is required. User object is fully populated when expanding a note.
	CreatedBy *User `json:"createdBy,omitempty"`


	// ExternalDataSources - Links to the sources of data (e.g. one source might be a CRM) that contributed data to this record.  Read-only, and only populated when requested via expand param.
	ExternalDataSources *[]Externaldatasource `json:"externalDataSources,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Note) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
