package thirdparty

// Some type from an external, third party package that won't be included in
// the resulting package but imported instead.
type SomeType struct {
	Foo string
}
