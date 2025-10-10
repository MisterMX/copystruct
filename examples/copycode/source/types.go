package source

import "github.com/mistermx/copystruct/examples/copycode/thirdparty"

// User represents a system user.
type User struct {
	// ID is the unique identifier
	ID int `json:"id"`
	// Name of the user
	Name    string  `json:"name"`
	Profile Profile `json:",inline"` // embedded profile

	Import thirdparty.SomeType `json:"import"`
}

// Profile contains extra info.
type Profile struct {
	Bio string // short bio
}

// OtherType is copied separately
type OtherType struct {
	Hello string `json:"hello,omitempty"`
}

// IgnoredType is not copied
type IgnoredType struct {
	Foo     string
	Profile Profile `json:",inline"` // embedded profile
}
