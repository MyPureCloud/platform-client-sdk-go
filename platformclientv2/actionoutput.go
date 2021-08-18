package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Actionoutput - Output definition of Action.
type Actionoutput struct { 
	// SuccessSchema - JSON schema that defines the transformed, successful result that will be sent back to the caller. If the 'flatten' query parameter is omitted or false, this field will be returned. Either successSchema or successSchemaFlattened will be returned, not both.
	SuccessSchema *Jsonschemadocument `json:"successSchema,omitempty"`


	// SuccessSchemaUri - URI to retrieve success schema
	SuccessSchemaUri *string `json:"successSchemaUri,omitempty"`


	// ErrorSchema - JSON schema that defines the body of response when request is not successful. If the 'flatten' query parameter is omitted or false, this field will be returned. Either errorSchema or errorSchemaFlattened will be returned, not both.
	ErrorSchema *Jsonschemadocument `json:"errorSchema,omitempty"`


	// ErrorSchemaUri - URI to retrieve error schema
	ErrorSchemaUri *string `json:"errorSchemaUri,omitempty"`


	// SuccessSchemaFlattened - JSON schema that defines the transformed, successful result that will be sent back to the caller. The schema is transformed based on Architect's flattened format. If the 'flatten' query parameter is supplied as true, this field will be returned. Either successSchema or successSchemaFlattened will be returned, not both.
	SuccessSchemaFlattened *Jsonschemadocument `json:"successSchemaFlattened,omitempty"`


	// ErrorSchemaFlattened - JSON schema that defines the body of response when request is not successful. The schema is transformed based on Architect's flattened format. If the 'flatten' query parameter is supplied as true, this field will be returned. Either errorSchema or errorSchemaFlattened will be returned, not both.
	ErrorSchemaFlattened *interface{} `json:"errorSchemaFlattened,omitempty"`

}

func (u *Actionoutput) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Actionoutput

	

	return json.Marshal(&struct { 
		SuccessSchema *Jsonschemadocument `json:"successSchema,omitempty"`
		
		SuccessSchemaUri *string `json:"successSchemaUri,omitempty"`
		
		ErrorSchema *Jsonschemadocument `json:"errorSchema,omitempty"`
		
		ErrorSchemaUri *string `json:"errorSchemaUri,omitempty"`
		
		SuccessSchemaFlattened *Jsonschemadocument `json:"successSchemaFlattened,omitempty"`
		
		ErrorSchemaFlattened *interface{} `json:"errorSchemaFlattened,omitempty"`
		*Alias
	}{ 
		SuccessSchema: u.SuccessSchema,
		
		SuccessSchemaUri: u.SuccessSchemaUri,
		
		ErrorSchema: u.ErrorSchema,
		
		ErrorSchemaUri: u.ErrorSchemaUri,
		
		SuccessSchemaFlattened: u.SuccessSchemaFlattened,
		
		ErrorSchemaFlattened: u.ErrorSchemaFlattened,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Actionoutput) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
