package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Contactlistdivisionview) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactlistdivisionview

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Division *Division `json:"division,omitempty"`
		
		ColumnNames *[]string `json:"columnNames,omitempty"`
		
		PhoneColumns *[]Contactphonenumbercolumn `json:"phoneColumns,omitempty"`
		
		ImportStatus *Importstatus `json:"importStatus,omitempty"`
		
		Size *int `json:"size,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Division: u.Division,
		
		ColumnNames: u.ColumnNames,
		
		PhoneColumns: u.PhoneColumns,
		
		ImportStatus: u.ImportStatus,
		
		Size: u.Size,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Contactlistdivisionview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
