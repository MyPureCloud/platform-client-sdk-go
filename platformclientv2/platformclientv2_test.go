package platformclientv2

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

type testConfig struct {
	environment         string
	clientID            string
	clientSecret        string
	logLevel            LoggingLevel
	userEmail           string
	usersAPI            *UsersApi
	userID              string
	userName            string
	userDepartment      string
	userProfileSkill    string
	busyPresenceID      string
	availablePresenceID string
}

type testSerializationStruct struct {
	IntProp          int       `json:"int,omitempty"`
	IntPropArr       []int     `json:"intArr,omitempty"`
	IntPropPtr       *int      `json:"intPtr,omitempty"`
	IntPropArrPtr    *[]int    `json:"intArrPtr,omitempty"`
	StringProp       string    `json:"string,omitempty"`
	StringPropArr    []string  `json:"stringArr,omitempty"`
	StringPropPtr    *string   `json:"stringPtr,omitempty"`
	StringPropArrPtr *[]string `json:"stringArrPtr,omitempty"`
	BoolProp         bool      `json:"bool,omitempty"`
	BoolPropArr      []bool    `json:"boolArr,omitempty"`
	BoolPropPtr      *bool     `json:"boolPtr,omitempty"`
	BoolPropArrPtr   *[]bool   `json:"boolArrPtr,omitempty"`
}

var config testConfig

func TestEnvVars(t *testing.T) {
	// Get
	config = testConfig{
		environment:         "https://api." + os.Getenv("PURECLOUD_ENVIRONMENT"),
		clientID:            os.Getenv("PURECLOUD_CLIENT_ID"),
		clientSecret:        os.Getenv("PURECLOUD_CLIENT_SECRET"),
		userName:            "GO SDK Tester",
		userDepartment:      "Ministry of Testing",
		userProfileSkill:    "Testmaster",
		busyPresenceID:      "31fe3bac-dea6-44b7-bed7-47f91660a1a0",
		availablePresenceID: "6a3af858-942f-489d-9700-5f9bcdcdae9b",
		logLevel:            LNone,
	}
	config.userEmail = fmt.Sprintf("%v@%v", uuid.New().String(), config.environment[12:])

	// Check
	if config.environment == "" {
		t.Error("Not set: PURECLOUD_ENVIRONMENT")
	}
	if config.clientID == "" {
		t.Error("Not set: PURECLOUD_CLIENT_ID")
	}
	if config.clientSecret == "" {
		t.Error("Not set: PURECLOUD_CLIENT_SECRET")
	}
	if config.userEmail == "@"+config.environment[12:] {
		t.Error("Invalid user email")
	}

	// Setup
	GetDefaultConfiguration().BasePath = config.environment
	GetDefaultConfiguration().LoggingConfiguration.LogLevel = config.logLevel
	config.usersAPI = NewUsersApi()

	// Log
	t.Logf("Enviornment: %v", config.environment)
	t.Logf("clientID: %v", config.clientID)
	t.Logf("userEmail: %v", config.userEmail)
}

func TestDefaultValueSerialization(t *testing.T) {
	expected := `{"intPtr":0,"intArrPtr":[],"stringPtr":"","stringArrPtr":[],"boolPtr":false,"boolArrPtr":[]}`
	intPropArrPtr := make([]int, 0)
	stringPropArrPtr := make([]string, 0)
	boolPropArrPtr := make([]bool, 0)
	intProp := 0
	v := testSerializationStruct{
		IntProp:          intProp,
		IntPropArr:       make([]int, 0),
		IntPropPtr:       &intProp,
		IntPropArrPtr:    &intPropArrPtr,
		StringProp:       "",
		StringPropArr:    make([]string, 0),
		StringPropPtr:    String(""),
		StringPropArrPtr: &stringPropArrPtr,
		BoolProp:         false,
		BoolPropArr:      make([]bool, 0),
		BoolPropPtr:      Bool(false),
		BoolPropArrPtr:   &boolPropArrPtr,
	}
	j, _ := json.Marshal(v)
	s := string(j)
	if s != expected {
		t.Log("testSerializationStruct did not serialize correctly")
		t.Logf("Expected: %v", expected)
		t.Logf("Actual:   %v", s)
		t.FailNow()
	}
}

func TestValueSerialization(t *testing.T) {
	expected := `{"int":10,"intArr":[0,0],"intPtr":10,"intArrPtr":[0,0],"string":"asdf","stringArr":["",""],"stringPtr":"asdf","stringArrPtr":["",""],"bool":true,"boolArr":[false,false],"boolPtr":true,"boolArrPtr":[false,false]}`
	intPropArrPtr := make([]int, 2)
	stringPropArrPtr := make([]string, 2)
	boolPropArrPtr := make([]bool, 2)
	intProp := 10
	v := testSerializationStruct{
		IntProp:          10,
		IntPropArr:       make([]int, 2),
		IntPropPtr:       &intProp,
		IntPropArrPtr:    &intPropArrPtr,
		StringProp:       "asdf",
		StringPropArr:    make([]string, 2),
		StringPropPtr:    String("asdf"),
		StringPropArrPtr: &stringPropArrPtr,
		BoolProp:         true,
		BoolPropArr:      make([]bool, 2),
		BoolPropPtr:      Bool(true),
		BoolPropArrPtr:   &boolPropArrPtr,
	}
	j, _ := json.Marshal(v)
	s := string(j)
	if s != expected {
		t.Log("testSerializationStruct did not serialize correctly")
		t.Logf("Expected: %v", expected)
		t.Logf("Actual:   %v", s)
		t.FailNow()
	}
}

func TestCustomSerialization1(t *testing.T) {
	patchBody := Patchactionmap{}

	// Use SetField to trigger custom serialization
	patchBody.SetField("Id", String("111-111"))

	// These manually set fields won't appear in the output when SetField is used
	patchBody.Version = Int(1)
	patchBody.DisplayName = String("manually set display name")

	// Validate serialization
	j, err := json.Marshal(patchBody)
	if err != nil {
		t.Error(err)
	}
	s := string(j)
	expected := `{"id":"111-111"}`
	if s != expected {
		t.Log("Patchactionmap did not serialize correctly")
		t.Logf("Expected: %v", expected)
		t.Logf("Actual:   %v", s)
		t.FailNow()
	}
}

func TestCustomSerialization2(t *testing.T) {
	patchBody := Patchactionmap{}

	// These manually set fields won't appear in the output when SetField is used
	patchBody.Id = String("999-999")
	patchBody.Version = Int(1)
	patchBody.DisplayName = String("manually set display name")
	patchBody.IsActive = Bool(true)

	// Use SetField to trigger custom serialization
	patchBody.SetField("Id", String("222-222"))
	patchBody.SetField("Version", Int(2))
	patchBody.SetField("DisplayName", nil)

	// Validate serialization
	j, err := json.Marshal(patchBody)
	if err != nil {
		t.Error(err)
	}
	s := string(j)
	expected := `{"displayName":null,"id":"222-222","version":2}`
	if s != expected {
		t.Log("Patchactionmap did not serialize correctly")
		t.Logf("Expected: %v", expected)
		t.Logf("Actual:   %v", s)
		t.FailNow()
	}
}

func TestCustomSerialization3(t *testing.T) {
	patchBody := Patchactionmap{}

	// These manually set fields won't appear in the output when SetField is used
	patchBody.Version = Int(1)
	patchBody.DisplayName = String("manually set display name")

	// Use SetField to trigger custom serialization
	patchBody.SetField("Id", String("333-333"))
	patchBody.SetField("Version", Int(3))
	patchBody.SetField("DisplayName", nil)
	parsedDate, _ := time.Parse(time.RFC3339, "2023-02-01T08:52:04.571291Z")
	patchBody.SetField("EndDate", &parsedDate)

	// Use default serialization on nested struct
	patchAction := Patchaction{}
	patchAction.ActionTargetId = String("12345")
	patchAction.IsPacingEnabled = Bool(false)
	patchAction.MediaType = nil
	patchBody.SetField("Action", &patchAction)

	// Validate serialization
	j, err := json.Marshal(patchBody)
	if err != nil {
		t.Error(err)
	}
	s := string(j)
	expected := `{"action":{"actionTargetId":"12345","isPacingEnabled":false},"displayName":null,"endDate":"2023-02-01T08:52:04.571291Z","id":"333-333","version":3}`
	if s != expected {
		t.Log("Patchactionmap did not serialize correctly")
		t.Logf("Expected: %v", expected)
		t.Logf("Actual:   %v", s)
		t.FailNow()
	}
}

func TestCustomSerialization4(t *testing.T) {
	patchBody := Patchactionmap{}

	// These manually set fields won't appear in the output when SetField is used
	patchBody.Version = Int(1)
	patchBody.DisplayName = String("manually set display name")

	// Use SetField to trigger custom serialization
	patchBody.SetField("Id", String("444-444"))
	patchBody.SetField("Version", Int(4))
	patchBody.SetField("DisplayName", nil)
	patchBody.SetField("EndDate", nil)

	// Use custom serialization on nested struct
	patchAction := Patchaction{}
	patchAction.SetField("ActionTargetId", String("12345"))
	patchAction.IsPacingEnabled = Bool(false)
	patchAction.SetField("MediaType", nil)
	patchBody.SetField("Action", &patchAction)

	// Validate serialization
	j, err := json.Marshal(patchBody)
	if err != nil {
		t.Error(err)
	}
	s := string(j)
	expected := `{"action":{"actionTargetId":"12345","mediaType":null},"displayName":null,"endDate":null,"id":"444-444","version":4}`
	if s != expected {
		t.Log("Patchactionmap did not serialize correctly")
		t.Logf("Expected: %v", expected)
		t.Logf("Actual:   %v", s)
		t.FailNow()
	}
}

func TestAuthentication(t *testing.T) {
	err := GetDefaultConfiguration().AuthorizeClientCredentials(config.clientID, config.clientSecret)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateUser(t *testing.T) {
	// Create user
	password := uuid.New().String() + "!@#$1234asdfASDF"
	newUser := Createuser{Name: &config.userName, Email: &config.userEmail, Password: &password}

	user, response, err := config.usersAPI.PostUsers(newUser)
	if err != nil {
		t.Error(err)
	} else if response != nil && response.Error != nil {
		t.Error(response.Error)
	} else {
		// Validate response
		if *user.Name != config.userName {
			t.Error("Data mismatch: user.Name")
		}
		if *user.Email != config.userEmail {
			t.Error("Data mismatch: user.Email")
		}

		// Store user ID
		config.userID = *user.Id
		t.Logf("New user's ID: %v", *user.Id)
	}
}

func TestUpdateUser(t *testing.T) {
	// Update user
	version := 1
	updateUser := Updateuser{Department: &config.userDepartment, Version: &version}

	user, response, err := config.usersAPI.PatchUser(config.userID, updateUser)
	if err != nil {
		t.Error(err)
	} else if response != nil && response.Error != nil {
		t.Error(response.Error)
	} else {
		// Validate response
		if *user.Name != config.userName {
			t.Error("Data mismatch: user.Name")
		}
		if *user.Email != config.userEmail {
			t.Error("Data mismatch: user.Email")
		}
		if *user.Department != config.userDepartment {
			t.Error("Data mismatch: user.Department")
		}
	}
}

func TestSetProfileSkills(t *testing.T) {
	// Update user
	skills, response, err := config.usersAPI.PutUserProfileskills(config.userID, []string{config.userProfileSkill})
	if err != nil {
		t.Error(err)
	} else if response != nil && response.Error != nil {
		t.Error(response.Error)
	} else {
		// Validate response
		if len(skills) != 1 {
			t.Errorf("Skills array contained the wrong number of elements. Expected 1: %v", skills)
		} else if skills[0] != config.userProfileSkill {
			t.Errorf("Skill did not match. Expected %v, actual: %v", config.userProfileSkill, skills[0])
		}
	}
}

func TestGetUser(t *testing.T) {
	// Get user
	user, response, err := config.usersAPI.GetUser(config.userID, []string{"profileSkills"}, "", "")
	if err != nil {
		t.Error(err)
	} else if response != nil && response.Error != nil {
		t.Error(response.Error)
	} else {
		// Validate response
		if *user.Name != config.userName {
			t.Error("Data mismatch: user.Name")
		}
		if *user.Email != config.userEmail {
			t.Error("Data mismatch: user.Email")
		}
		if *user.Department != config.userDepartment {
			t.Error("Data mismatch: user.Department")
		}
		// Commented out until the issue with APIs to send the latest Version of the User is fixed.
		// if user.ProfileSkills == nil || len(*user.ProfileSkills) != 1 || (*user.ProfileSkills)[0] != config.userProfileSkill {
		// 	t.Error("Data mismatch: user.ProfileSkills")
		// }
	}
}

func TestDeleteUser(t *testing.T) {
	// Delete user
	_, response, err := config.usersAPI.DeleteUser(config.userID)
	if err != nil {
		t.Error(err)
	} else if response != nil && response.Error != nil {
		t.Error(response.Error)
	}
}

func Example_authorizeDefaultConfiguration() {
	// Use the default config instance and retrieve settings from env vars
	config := GetDefaultConfiguration()
	config.BasePath = "https://api." + os.Getenv("PURECLOUD_ENVIRONMENT") // e.g. PURECLOUD_ENVIRONMENT=mypurecloud.com
	clientID := os.Getenv("PURECLOUD_CLIENT_ID")
	clientSecret := os.Getenv("PURECLOUD_CLIENT_SECRET")

	// Authorize the configuration
	err := config.AuthorizeClientCredentials(clientID, clientSecret)
	if err != nil {
		panic(err)
	}

	// Create an API instance using the default config
	usersAPI := NewUsersApi()
	fmt.Printf("Users API type: %v", reflect.TypeOf(usersAPI).String())
	// Output: Users API type: UsersAPI

	// Make requests using usersAPI
}

func Example_authorizeNewConfiguration() {
	// Create a new config instance and retrieve settings from env vars
	config := NewConfiguration()
	config.BasePath = "https://api." + os.Getenv("PURECLOUD_ENVIRONMENT") // e.g. PURECLOUD_ENVIRONMENT=mypurecloud.com
	clientID := os.Getenv("PURECLOUD_CLIENT_ID")
	clientSecret := os.Getenv("PURECLOUD_CLIENT_SECRET")

	// Authorize the configuration
	err := config.AuthorizeClientCredentials(clientID, clientSecret)
	if err != nil {
		panic(err)
	}

	// Create an API instance using the config instance
	usersAPI := NewUsersApiWithConfig(config)
	fmt.Printf("Users API type: %v", reflect.TypeOf(usersAPI).String())
	// Output: Users API type: UsersAPI

	// Make requests using usersAPI
}
