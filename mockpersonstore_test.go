package main

// AUTOGENERATED BY MOQ
// github.com/matryer/moq

// PersonStoreMock is a mock implementation of PersonStore.
//
//     func TestSomethingThatUsesPersonStore(t *testing.T) {
//
//         // make and configure a mocked PersonStore
//         mockedPersonStore := &PersonStoreMock{ 
//             CreateFunc: func(ctx context.Context, person *Person, confirm bool) error {
// 	               panic("TODO: mock out the Create function")
//             },
//             GetFunc: func(ctx context.Context, id string) (*Person, error) {
// 	               panic("TODO: mock out the Get function")
//             },
//         }
//
//         // TODO: use mockedPersonStore in code that requires PersonStore
//     
//     }
type PersonStoreMock struct {
	// CreateFunc mocks the Create function.
	CreateFunc func(ctx context.Context, person *Person, confirm bool) error
	// GetFunc mocks the Get function.
	GetFunc func(ctx context.Context, id string) (*Person, error)
}

// Create calls CreateFunc.
func (mock *PersonStoreMock) Create(ctx context.Context, person *Person, confirm bool) error {
	if mock.Create == nil {
		panic("moq: CreateFunc is nil but was just called")
	}
	return mock.CreateFunc(ctx, person, confirm)
}

// Get calls GetFunc.
func (mock *PersonStoreMock) Get(ctx context.Context, id string) (*Person, error) {
	if mock.Get == nil {
		panic("moq: GetFunc is nil but was just called")
	}
	return mock.GetFunc(ctx, id)
}