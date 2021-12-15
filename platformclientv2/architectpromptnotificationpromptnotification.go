package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectpromptnotificationpromptnotification
type Architectpromptnotificationpromptnotification struct { 
	// Id - The prompt ID
	Id *string `json:"id,omitempty"`


	// Name - The prompt name
	Name *string `json:"name,omitempty"`


	// Description - The prompt description
	Description *string `json:"description,omitempty"`


	// CurrentOperation
	CurrentOperation *Architectpromptnotificationarchitectoperation `json:"currentOperation,omitempty"`

}

func (o *Architectpromptnotificationpromptnotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectpromptnotificationpromptnotification
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		CurrentOperation *Architectpromptnotificationarchitectoperation `json:"currentOperation,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		CurrentOperation: o.CurrentOperation,
		Alias:    (*Alias)(o),
	})
}

func (o *Architectpromptnotificationpromptnotification) UnmarshalJSON(b []byte) error {
	var ArchitectpromptnotificationpromptnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectpromptnotificationpromptnotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ArchitectpromptnotificationpromptnotificationMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := ArchitectpromptnotificationpromptnotificationMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Description, ok := ArchitectpromptnotificationpromptnotificationMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if CurrentOperation, ok := ArchitectpromptnotificationpromptnotificationMap["currentOperation"].(map[string]interface{}); ok {
		CurrentOperationString, _ := json.Marshal(CurrentOperation)
		json.Unmarshal(CurrentOperationString, &o.CurrentOperation)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Architectpromptnotificationpromptnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
