package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectflownotificationflownotification
type Architectflownotificationflownotification struct { 
	// Id - The flow ID
	Id *string `json:"id,omitempty"`


	// Name - The flow name
	Name *string `json:"name,omitempty"`


	// Description - The flow description
	Description *string `json:"description,omitempty"`


	// Deleted - The flow deleted state
	Deleted *bool `json:"deleted,omitempty"`


	// CheckedInVersion
	CheckedInVersion *Architectflownotificationflowversion `json:"checkedInVersion,omitempty"`


	// SavedVersion - A bare-bones flow version object
	SavedVersion *Architectflownotificationflowversion `json:"savedVersion,omitempty"`


	// PublishedVersion - A bare-bones flow version object
	PublishedVersion *Architectflownotificationflowversion `json:"publishedVersion,omitempty"`


	// CurrentOperation
	CurrentOperation *Architectflownotificationarchitectoperation `json:"currentOperation,omitempty"`

}

func (o *Architectflownotificationflownotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectflownotificationflownotification
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		Deleted *bool `json:"deleted,omitempty"`
		
		CheckedInVersion *Architectflownotificationflowversion `json:"checkedInVersion,omitempty"`
		
		SavedVersion *Architectflownotificationflowversion `json:"savedVersion,omitempty"`
		
		PublishedVersion *Architectflownotificationflowversion `json:"publishedVersion,omitempty"`
		
		CurrentOperation *Architectflownotificationarchitectoperation `json:"currentOperation,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		Deleted: o.Deleted,
		
		CheckedInVersion: o.CheckedInVersion,
		
		SavedVersion: o.SavedVersion,
		
		PublishedVersion: o.PublishedVersion,
		
		CurrentOperation: o.CurrentOperation,
		Alias:    (*Alias)(o),
	})
}

func (o *Architectflownotificationflownotification) UnmarshalJSON(b []byte) error {
	var ArchitectflownotificationflownotificationMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectflownotificationflownotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ArchitectflownotificationflownotificationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ArchitectflownotificationflownotificationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := ArchitectflownotificationflownotificationMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if Deleted, ok := ArchitectflownotificationflownotificationMap["deleted"].(bool); ok {
		o.Deleted = &Deleted
	}
    
	if CheckedInVersion, ok := ArchitectflownotificationflownotificationMap["checkedInVersion"].(map[string]interface{}); ok {
		CheckedInVersionString, _ := json.Marshal(CheckedInVersion)
		json.Unmarshal(CheckedInVersionString, &o.CheckedInVersion)
	}
	
	if SavedVersion, ok := ArchitectflownotificationflownotificationMap["savedVersion"].(map[string]interface{}); ok {
		SavedVersionString, _ := json.Marshal(SavedVersion)
		json.Unmarshal(SavedVersionString, &o.SavedVersion)
	}
	
	if PublishedVersion, ok := ArchitectflownotificationflownotificationMap["publishedVersion"].(map[string]interface{}); ok {
		PublishedVersionString, _ := json.Marshal(PublishedVersion)
		json.Unmarshal(PublishedVersionString, &o.PublishedVersion)
	}
	
	if CurrentOperation, ok := ArchitectflownotificationflownotificationMap["currentOperation"].(map[string]interface{}); ok {
		CurrentOperationString, _ := json.Marshal(CurrentOperation)
		json.Unmarshal(CurrentOperationString, &o.CurrentOperation)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Architectflownotificationflownotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
