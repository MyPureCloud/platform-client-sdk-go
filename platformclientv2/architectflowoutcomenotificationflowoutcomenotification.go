package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Architectflowoutcomenotificationflowoutcomenotification
type Architectflowoutcomenotificationflowoutcomenotification struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Description
	Description *string `json:"description,omitempty"`


	// CurrentOperation
	CurrentOperation *Architectflowoutcomenotificationarchitectoperation `json:"currentOperation,omitempty"`

}

func (o *Architectflowoutcomenotificationflowoutcomenotification) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Architectflowoutcomenotificationflowoutcomenotification
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		CurrentOperation *Architectflowoutcomenotificationarchitectoperation `json:"currentOperation,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		CurrentOperation: o.CurrentOperation,
		Alias:    (*Alias)(o),
	})
}

func (o *Architectflowoutcomenotificationflowoutcomenotification) UnmarshalJSON(b []byte) error {
	var ArchitectflowoutcomenotificationflowoutcomenotificationMap map[string]interface{}
	err := json.Unmarshal(b, &ArchitectflowoutcomenotificationflowoutcomenotificationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ArchitectflowoutcomenotificationflowoutcomenotificationMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := ArchitectflowoutcomenotificationflowoutcomenotificationMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Description, ok := ArchitectflowoutcomenotificationflowoutcomenotificationMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if CurrentOperation, ok := ArchitectflowoutcomenotificationflowoutcomenotificationMap["currentOperation"].(map[string]interface{}); ok {
		CurrentOperationString, _ := json.Marshal(CurrentOperation)
		json.Unmarshal(CurrentOperationString, &o.CurrentOperation)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Architectflowoutcomenotificationflowoutcomenotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
