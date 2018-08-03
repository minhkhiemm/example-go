// github.com/matryer/moq

package book

import (
	"context"
	"sync"

	"github.com/minhkhiemm/example-go/domain"
)

var (
	lockServiceMockCreate  sync.RWMutex
	lockServiceMockDelete  sync.RWMutex
	lockServiceMockFind    sync.RWMutex
	lockServiceMockFindAll sync.RWMutex
	lockServiceMockUpdate  sync.RWMutex
)

// ServiceMock is a mock implementation of Service.
//
//     func TestSomethingThatUsesService(t *testing.T) {
//
//         // make and configure a mocked Service
//         mockedService := &ServiceMock{
//             CreateFunc: func(ctx context.Context, p *domain.Book) error {
// 	               panic("TODO: mock out the Create method")
//             },
//             DeleteFunc: func(ctx context.Context, p *domain.Book) error {
// 	               panic("TODO: mock out the Delete method")
//             },
//             FindFunc: func(ctx context.Context, p *domain.Book) (*domain.Book, error) {
// 	               panic("TODO: mock out the Find method")
//             },
//             FindAllFunc: func(ctx context.Context) ([]domain.Book, error) {
// 	               panic("TODO: mock out the FindAll method")
//             },
//             UpdateFunc: func(ctx context.Context, p *domain.Book) (*domain.Book, error) {
// 	               panic("TODO: mock out the Update method")
//             },
//         }
//
//         // TODO: use mockedService in code that requires Service
//         //       and then make assertions.
//
//     }
type ServiceMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, p *domain.Book) error

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(ctx context.Context, p *domain.Book) error

	// FindFunc mocks the Find method.
	FindFunc func(ctx context.Context, p *domain.Book) (*domain.Book, error)

	// FindAllFunc mocks the FindAll method.
	FindAllFunc func(ctx context.Context) ([]domain.Book, error)

	// UpdateFunc mocks the Update method.
	UpdateFunc func(ctx context.Context, p *domain.Book) (*domain.Book, error)

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// P is the p argument value.
			P *domain.Book
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// P is the p argument value.
			P *domain.Book
		}
		// Find holds details about calls to the Find method.
		Find []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// P is the p argument value.
			P *domain.Book
		}
		// FindAll holds details about calls to the FindAll method.
		FindAll []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// P is the p argument value.
			P *domain.Book
		}
	}
}

// Create calls CreateFunc.
func (mock *ServiceMock) Create(ctx context.Context, p *domain.Book) error {
	if mock.CreateFunc == nil {
		panic("ServiceMock.CreateFunc: method is nil but Service.Create was just called")
	}
	callInfo := struct {
		Ctx context.Context
		P   *domain.Book
	}{
		Ctx: ctx,
		P:   p,
	}
	lockServiceMockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	lockServiceMockCreate.Unlock()
	return mock.CreateFunc(ctx, p)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//     len(mockedService.CreateCalls())
func (mock *ServiceMock) CreateCalls() []struct {
	Ctx context.Context
	P   *domain.Book
} {
	var calls []struct {
		Ctx context.Context
		P   *domain.Book
	}
	lockServiceMockCreate.RLock()
	calls = mock.calls.Create
	lockServiceMockCreate.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *ServiceMock) Delete(ctx context.Context, p *domain.Book) error {
	if mock.DeleteFunc == nil {
		panic("ServiceMock.DeleteFunc: method is nil but Service.Delete was just called")
	}
	callInfo := struct {
		Ctx context.Context
		P   *domain.Book
	}{
		Ctx: ctx,
		P:   p,
	}
	lockServiceMockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	lockServiceMockDelete.Unlock()
	return mock.DeleteFunc(ctx, p)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//     len(mockedService.DeleteCalls())
func (mock *ServiceMock) DeleteCalls() []struct {
	Ctx context.Context
	P   *domain.Book
} {
	var calls []struct {
		Ctx context.Context
		P   *domain.Book
	}
	lockServiceMockDelete.RLock()
	calls = mock.calls.Delete
	lockServiceMockDelete.RUnlock()
	return calls
}

// Find calls FindFunc.
func (mock *ServiceMock) Find(ctx context.Context, p *domain.Book) (*domain.Book, error) {
	if mock.FindFunc == nil {
		panic("ServiceMock.FindFunc: method is nil but Service.Find was just called")
	}
	callInfo := struct {
		Ctx context.Context
		P   *domain.Book
	}{
		Ctx: ctx,
		P:   p,
	}
	lockServiceMockFind.Lock()
	mock.calls.Find = append(mock.calls.Find, callInfo)
	lockServiceMockFind.Unlock()
	return mock.FindFunc(ctx, p)
}

// FindCalls gets all the calls that were made to Find.
// Check the length with:
//     len(mockedService.FindCalls())
func (mock *ServiceMock) FindCalls() []struct {
	Ctx context.Context
	P   *domain.Book
} {
	var calls []struct {
		Ctx context.Context
		P   *domain.Book
	}
	lockServiceMockFind.RLock()
	calls = mock.calls.Find
	lockServiceMockFind.RUnlock()
	return calls
}

// FindAll calls FindAllFunc.
func (mock *ServiceMock) FindAll(ctx context.Context) ([]domain.Book, error) {
	if mock.FindAllFunc == nil {
		panic("ServiceMock.FindAllFunc: method is nil but Service.FindAll was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	lockServiceMockFindAll.Lock()
	mock.calls.FindAll = append(mock.calls.FindAll, callInfo)
	lockServiceMockFindAll.Unlock()
	return mock.FindAllFunc(ctx)
}

// FindAllCalls gets all the calls that were made to FindAll.
// Check the length with:
//     len(mockedService.FindAllCalls())
func (mock *ServiceMock) FindAllCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	lockServiceMockFindAll.RLock()
	calls = mock.calls.FindAll
	lockServiceMockFindAll.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *ServiceMock) Update(ctx context.Context, p *domain.Book) (*domain.Book, error) {
	if mock.UpdateFunc == nil {
		panic("ServiceMock.UpdateFunc: method is nil but Service.Update was just called")
	}
	callInfo := struct {
		Ctx context.Context
		P   *domain.Book
	}{
		Ctx: ctx,
		P:   p,
	}
	lockServiceMockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	lockServiceMockUpdate.Unlock()
	return mock.UpdateFunc(ctx, p)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//     len(mockedService.UpdateCalls())
func (mock *ServiceMock) UpdateCalls() []struct {
	Ctx context.Context
	P   *domain.Book
} {
	var calls []struct {
		Ctx context.Context
		P   *domain.Book
	}
	lockServiceMockUpdate.RLock()
	calls = mock.calls.Update
	lockServiceMockUpdate.RUnlock()
	return calls
}
