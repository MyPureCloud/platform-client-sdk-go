package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Contactlistdivisionview
type Contactlistdivisionview struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Division - The division to which this entity belongs.
	Division *Division `json:"division,omitempty"`


	// ColumnNames - The names of the contact data columns.
	ColumnNames *[]string `json:"columnNames,omitempty"`


	// PhoneColumns - Indicates which columns are phone numbers.
	PhoneColumns *[]Contactphonenumbercolumn `json:"phoneColumns,omitempty"`


	// ImportStatus - The status of the import process.
	ImportStatus *Importstatus `json:"importStatus,omitempty"`


	// Size - The number of contacts in the ContactList.
	Size *int `json:"size,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

// String returns a JSON representation of the model
func (o *Contactlistdivisionview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
