package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i route256/checkout/internal/domain/checkout.Repository -o ./mocks/repository_minimock.go -n RepositoryMock

import (
	"context"
	"route256/checkout/internal/models"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// RepositoryMock implements checkout.Repository
type RepositoryMock struct {
	t minimock.Tester

	funcAddToCart          func(ctx context.Context, user int64, sku uint32, count uint32) (err error)
	inspectFuncAddToCart   func(ctx context.Context, user int64, sku uint32, count uint32)
	afterAddToCartCounter  uint64
	beforeAddToCartCounter uint64
	AddToCartMock          mRepositoryMockAddToCart

	funcClearCart          func(ctx context.Context, user int64) (err error)
	inspectFuncClearCart   func(ctx context.Context, user int64)
	afterClearCartCounter  uint64
	beforeClearCartCounter uint64
	ClearCartMock          mRepositoryMockClearCart

	funcDeleteFromCart          func(ctx context.Context, user int64, sku uint32, count uint32) (err error)
	inspectFuncDeleteFromCart   func(ctx context.Context, user int64, sku uint32, count uint32)
	afterDeleteFromCartCounter  uint64
	beforeDeleteFromCartCounter uint64
	DeleteFromCartMock          mRepositoryMockDeleteFromCart

	funcListCart          func(ctx context.Context, user int64) (ia1 []models.Item, err error)
	inspectFuncListCart   func(ctx context.Context, user int64)
	afterListCartCounter  uint64
	beforeListCartCounter uint64
	ListCartMock          mRepositoryMockListCart
}

// NewRepositoryMock returns a mock for checkout.Repository
func NewRepositoryMock(t minimock.Tester) *RepositoryMock {
	m := &RepositoryMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AddToCartMock = mRepositoryMockAddToCart{mock: m}
	m.AddToCartMock.callArgs = []*RepositoryMockAddToCartParams{}

	m.ClearCartMock = mRepositoryMockClearCart{mock: m}
	m.ClearCartMock.callArgs = []*RepositoryMockClearCartParams{}

	m.DeleteFromCartMock = mRepositoryMockDeleteFromCart{mock: m}
	m.DeleteFromCartMock.callArgs = []*RepositoryMockDeleteFromCartParams{}

	m.ListCartMock = mRepositoryMockListCart{mock: m}
	m.ListCartMock.callArgs = []*RepositoryMockListCartParams{}

	return m
}

type mRepositoryMockAddToCart struct {
	mock               *RepositoryMock
	defaultExpectation *RepositoryMockAddToCartExpectation
	expectations       []*RepositoryMockAddToCartExpectation

	callArgs []*RepositoryMockAddToCartParams
	mutex    sync.RWMutex
}

// RepositoryMockAddToCartExpectation specifies expectation struct of the Repository.AddToCart
type RepositoryMockAddToCartExpectation struct {
	mock    *RepositoryMock
	params  *RepositoryMockAddToCartParams
	results *RepositoryMockAddToCartResults
	Counter uint64
}

// RepositoryMockAddToCartParams contains parameters of the Repository.AddToCart
type RepositoryMockAddToCartParams struct {
	ctx   context.Context
	user  int64
	sku   uint32
	count uint32
}

// RepositoryMockAddToCartResults contains results of the Repository.AddToCart
type RepositoryMockAddToCartResults struct {
	err error
}

// Expect sets up expected params for Repository.AddToCart
func (mmAddToCart *mRepositoryMockAddToCart) Expect(ctx context.Context, user int64, sku uint32, count uint32) *mRepositoryMockAddToCart {
	if mmAddToCart.mock.funcAddToCart != nil {
		mmAddToCart.mock.t.Fatalf("RepositoryMock.AddToCart mock is already set by Set")
	}

	if mmAddToCart.defaultExpectation == nil {
		mmAddToCart.defaultExpectation = &RepositoryMockAddToCartExpectation{}
	}

	mmAddToCart.defaultExpectation.params = &RepositoryMockAddToCartParams{ctx, user, sku, count}
	for _, e := range mmAddToCart.expectations {
		if minimock.Equal(e.params, mmAddToCart.defaultExpectation.params) {
			mmAddToCart.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAddToCart.defaultExpectation.params)
		}
	}

	return mmAddToCart
}

// Inspect accepts an inspector function that has same arguments as the Repository.AddToCart
func (mmAddToCart *mRepositoryMockAddToCart) Inspect(f func(ctx context.Context, user int64, sku uint32, count uint32)) *mRepositoryMockAddToCart {
	if mmAddToCart.mock.inspectFuncAddToCart != nil {
		mmAddToCart.mock.t.Fatalf("Inspect function is already set for RepositoryMock.AddToCart")
	}

	mmAddToCart.mock.inspectFuncAddToCart = f

	return mmAddToCart
}

// Return sets up results that will be returned by Repository.AddToCart
func (mmAddToCart *mRepositoryMockAddToCart) Return(err error) *RepositoryMock {
	if mmAddToCart.mock.funcAddToCart != nil {
		mmAddToCart.mock.t.Fatalf("RepositoryMock.AddToCart mock is already set by Set")
	}

	if mmAddToCart.defaultExpectation == nil {
		mmAddToCart.defaultExpectation = &RepositoryMockAddToCartExpectation{mock: mmAddToCart.mock}
	}
	mmAddToCart.defaultExpectation.results = &RepositoryMockAddToCartResults{err}
	return mmAddToCart.mock
}

// Set uses given function f to mock the Repository.AddToCart method
func (mmAddToCart *mRepositoryMockAddToCart) Set(f func(ctx context.Context, user int64, sku uint32, count uint32) (err error)) *RepositoryMock {
	if mmAddToCart.defaultExpectation != nil {
		mmAddToCart.mock.t.Fatalf("Default expectation is already set for the Repository.AddToCart method")
	}

	if len(mmAddToCart.expectations) > 0 {
		mmAddToCart.mock.t.Fatalf("Some expectations are already set for the Repository.AddToCart method")
	}

	mmAddToCart.mock.funcAddToCart = f
	return mmAddToCart.mock
}

// When sets expectation for the Repository.AddToCart which will trigger the result defined by the following
// Then helper
func (mmAddToCart *mRepositoryMockAddToCart) When(ctx context.Context, user int64, sku uint32, count uint32) *RepositoryMockAddToCartExpectation {
	if mmAddToCart.mock.funcAddToCart != nil {
		mmAddToCart.mock.t.Fatalf("RepositoryMock.AddToCart mock is already set by Set")
	}

	expectation := &RepositoryMockAddToCartExpectation{
		mock:   mmAddToCart.mock,
		params: &RepositoryMockAddToCartParams{ctx, user, sku, count},
	}
	mmAddToCart.expectations = append(mmAddToCart.expectations, expectation)
	return expectation
}

// Then sets up Repository.AddToCart return parameters for the expectation previously defined by the When method
func (e *RepositoryMockAddToCartExpectation) Then(err error) *RepositoryMock {
	e.results = &RepositoryMockAddToCartResults{err}
	return e.mock
}

// AddToCart implements checkout.Repository
func (mmAddToCart *RepositoryMock) AddToCart(ctx context.Context, user int64, sku uint32, count uint32) (err error) {
	mm_atomic.AddUint64(&mmAddToCart.beforeAddToCartCounter, 1)
	defer mm_atomic.AddUint64(&mmAddToCart.afterAddToCartCounter, 1)

	if mmAddToCart.inspectFuncAddToCart != nil {
		mmAddToCart.inspectFuncAddToCart(ctx, user, sku, count)
	}

	mm_params := &RepositoryMockAddToCartParams{ctx, user, sku, count}

	// Record call args
	mmAddToCart.AddToCartMock.mutex.Lock()
	mmAddToCart.AddToCartMock.callArgs = append(mmAddToCart.AddToCartMock.callArgs, mm_params)
	mmAddToCart.AddToCartMock.mutex.Unlock()

	for _, e := range mmAddToCart.AddToCartMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmAddToCart.AddToCartMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAddToCart.AddToCartMock.defaultExpectation.Counter, 1)
		mm_want := mmAddToCart.AddToCartMock.defaultExpectation.params
		mm_got := RepositoryMockAddToCartParams{ctx, user, sku, count}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmAddToCart.t.Errorf("RepositoryMock.AddToCart got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmAddToCart.AddToCartMock.defaultExpectation.results
		if mm_results == nil {
			mmAddToCart.t.Fatal("No results are set for the RepositoryMock.AddToCart")
		}
		return (*mm_results).err
	}
	if mmAddToCart.funcAddToCart != nil {
		return mmAddToCart.funcAddToCart(ctx, user, sku, count)
	}
	mmAddToCart.t.Fatalf("Unexpected call to RepositoryMock.AddToCart. %v %v %v %v", ctx, user, sku, count)
	return
}

// AddToCartAfterCounter returns a count of finished RepositoryMock.AddToCart invocations
func (mmAddToCart *RepositoryMock) AddToCartAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddToCart.afterAddToCartCounter)
}

// AddToCartBeforeCounter returns a count of RepositoryMock.AddToCart invocations
func (mmAddToCart *RepositoryMock) AddToCartBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddToCart.beforeAddToCartCounter)
}

// Calls returns a list of arguments used in each call to RepositoryMock.AddToCart.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAddToCart *mRepositoryMockAddToCart) Calls() []*RepositoryMockAddToCartParams {
	mmAddToCart.mutex.RLock()

	argCopy := make([]*RepositoryMockAddToCartParams, len(mmAddToCart.callArgs))
	copy(argCopy, mmAddToCart.callArgs)

	mmAddToCart.mutex.RUnlock()

	return argCopy
}

// MinimockAddToCartDone returns true if the count of the AddToCart invocations corresponds
// the number of defined expectations
func (m *RepositoryMock) MinimockAddToCartDone() bool {
	for _, e := range m.AddToCartMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddToCartMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddToCartCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAddToCart != nil && mm_atomic.LoadUint64(&m.afterAddToCartCounter) < 1 {
		return false
	}
	return true
}

// MinimockAddToCartInspect logs each unmet expectation
func (m *RepositoryMock) MinimockAddToCartInspect() {
	for _, e := range m.AddToCartMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepositoryMock.AddToCart with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddToCartMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddToCartCounter) < 1 {
		if m.AddToCartMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepositoryMock.AddToCart")
		} else {
			m.t.Errorf("Expected call to RepositoryMock.AddToCart with params: %#v", *m.AddToCartMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAddToCart != nil && mm_atomic.LoadUint64(&m.afterAddToCartCounter) < 1 {
		m.t.Error("Expected call to RepositoryMock.AddToCart")
	}
}

type mRepositoryMockClearCart struct {
	mock               *RepositoryMock
	defaultExpectation *RepositoryMockClearCartExpectation
	expectations       []*RepositoryMockClearCartExpectation

	callArgs []*RepositoryMockClearCartParams
	mutex    sync.RWMutex
}

// RepositoryMockClearCartExpectation specifies expectation struct of the Repository.ClearCart
type RepositoryMockClearCartExpectation struct {
	mock    *RepositoryMock
	params  *RepositoryMockClearCartParams
	results *RepositoryMockClearCartResults
	Counter uint64
}

// RepositoryMockClearCartParams contains parameters of the Repository.ClearCart
type RepositoryMockClearCartParams struct {
	ctx  context.Context
	user int64
}

// RepositoryMockClearCartResults contains results of the Repository.ClearCart
type RepositoryMockClearCartResults struct {
	err error
}

// Expect sets up expected params for Repository.ClearCart
func (mmClearCart *mRepositoryMockClearCart) Expect(ctx context.Context, user int64) *mRepositoryMockClearCart {
	if mmClearCart.mock.funcClearCart != nil {
		mmClearCart.mock.t.Fatalf("RepositoryMock.ClearCart mock is already set by Set")
	}

	if mmClearCart.defaultExpectation == nil {
		mmClearCart.defaultExpectation = &RepositoryMockClearCartExpectation{}
	}

	mmClearCart.defaultExpectation.params = &RepositoryMockClearCartParams{ctx, user}
	for _, e := range mmClearCart.expectations {
		if minimock.Equal(e.params, mmClearCart.defaultExpectation.params) {
			mmClearCart.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmClearCart.defaultExpectation.params)
		}
	}

	return mmClearCart
}

// Inspect accepts an inspector function that has same arguments as the Repository.ClearCart
func (mmClearCart *mRepositoryMockClearCart) Inspect(f func(ctx context.Context, user int64)) *mRepositoryMockClearCart {
	if mmClearCart.mock.inspectFuncClearCart != nil {
		mmClearCart.mock.t.Fatalf("Inspect function is already set for RepositoryMock.ClearCart")
	}

	mmClearCart.mock.inspectFuncClearCart = f

	return mmClearCart
}

// Return sets up results that will be returned by Repository.ClearCart
func (mmClearCart *mRepositoryMockClearCart) Return(err error) *RepositoryMock {
	if mmClearCart.mock.funcClearCart != nil {
		mmClearCart.mock.t.Fatalf("RepositoryMock.ClearCart mock is already set by Set")
	}

	if mmClearCart.defaultExpectation == nil {
		mmClearCart.defaultExpectation = &RepositoryMockClearCartExpectation{mock: mmClearCart.mock}
	}
	mmClearCart.defaultExpectation.results = &RepositoryMockClearCartResults{err}
	return mmClearCart.mock
}

// Set uses given function f to mock the Repository.ClearCart method
func (mmClearCart *mRepositoryMockClearCart) Set(f func(ctx context.Context, user int64) (err error)) *RepositoryMock {
	if mmClearCart.defaultExpectation != nil {
		mmClearCart.mock.t.Fatalf("Default expectation is already set for the Repository.ClearCart method")
	}

	if len(mmClearCart.expectations) > 0 {
		mmClearCart.mock.t.Fatalf("Some expectations are already set for the Repository.ClearCart method")
	}

	mmClearCart.mock.funcClearCart = f
	return mmClearCart.mock
}

// When sets expectation for the Repository.ClearCart which will trigger the result defined by the following
// Then helper
func (mmClearCart *mRepositoryMockClearCart) When(ctx context.Context, user int64) *RepositoryMockClearCartExpectation {
	if mmClearCart.mock.funcClearCart != nil {
		mmClearCart.mock.t.Fatalf("RepositoryMock.ClearCart mock is already set by Set")
	}

	expectation := &RepositoryMockClearCartExpectation{
		mock:   mmClearCart.mock,
		params: &RepositoryMockClearCartParams{ctx, user},
	}
	mmClearCart.expectations = append(mmClearCart.expectations, expectation)
	return expectation
}

// Then sets up Repository.ClearCart return parameters for the expectation previously defined by the When method
func (e *RepositoryMockClearCartExpectation) Then(err error) *RepositoryMock {
	e.results = &RepositoryMockClearCartResults{err}
	return e.mock
}

// ClearCart implements checkout.Repository
func (mmClearCart *RepositoryMock) ClearCart(ctx context.Context, user int64) (err error) {
	mm_atomic.AddUint64(&mmClearCart.beforeClearCartCounter, 1)
	defer mm_atomic.AddUint64(&mmClearCart.afterClearCartCounter, 1)

	if mmClearCart.inspectFuncClearCart != nil {
		mmClearCart.inspectFuncClearCart(ctx, user)
	}

	mm_params := &RepositoryMockClearCartParams{ctx, user}

	// Record call args
	mmClearCart.ClearCartMock.mutex.Lock()
	mmClearCart.ClearCartMock.callArgs = append(mmClearCart.ClearCartMock.callArgs, mm_params)
	mmClearCart.ClearCartMock.mutex.Unlock()

	for _, e := range mmClearCart.ClearCartMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmClearCart.ClearCartMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmClearCart.ClearCartMock.defaultExpectation.Counter, 1)
		mm_want := mmClearCart.ClearCartMock.defaultExpectation.params
		mm_got := RepositoryMockClearCartParams{ctx, user}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmClearCart.t.Errorf("RepositoryMock.ClearCart got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmClearCart.ClearCartMock.defaultExpectation.results
		if mm_results == nil {
			mmClearCart.t.Fatal("No results are set for the RepositoryMock.ClearCart")
		}
		return (*mm_results).err
	}
	if mmClearCart.funcClearCart != nil {
		return mmClearCart.funcClearCart(ctx, user)
	}
	mmClearCart.t.Fatalf("Unexpected call to RepositoryMock.ClearCart. %v %v", ctx, user)
	return
}

// ClearCartAfterCounter returns a count of finished RepositoryMock.ClearCart invocations
func (mmClearCart *RepositoryMock) ClearCartAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmClearCart.afterClearCartCounter)
}

// ClearCartBeforeCounter returns a count of RepositoryMock.ClearCart invocations
func (mmClearCart *RepositoryMock) ClearCartBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmClearCart.beforeClearCartCounter)
}

// Calls returns a list of arguments used in each call to RepositoryMock.ClearCart.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmClearCart *mRepositoryMockClearCart) Calls() []*RepositoryMockClearCartParams {
	mmClearCart.mutex.RLock()

	argCopy := make([]*RepositoryMockClearCartParams, len(mmClearCart.callArgs))
	copy(argCopy, mmClearCart.callArgs)

	mmClearCart.mutex.RUnlock()

	return argCopy
}

// MinimockClearCartDone returns true if the count of the ClearCart invocations corresponds
// the number of defined expectations
func (m *RepositoryMock) MinimockClearCartDone() bool {
	for _, e := range m.ClearCartMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ClearCartMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterClearCartCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcClearCart != nil && mm_atomic.LoadUint64(&m.afterClearCartCounter) < 1 {
		return false
	}
	return true
}

// MinimockClearCartInspect logs each unmet expectation
func (m *RepositoryMock) MinimockClearCartInspect() {
	for _, e := range m.ClearCartMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepositoryMock.ClearCart with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ClearCartMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterClearCartCounter) < 1 {
		if m.ClearCartMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepositoryMock.ClearCart")
		} else {
			m.t.Errorf("Expected call to RepositoryMock.ClearCart with params: %#v", *m.ClearCartMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcClearCart != nil && mm_atomic.LoadUint64(&m.afterClearCartCounter) < 1 {
		m.t.Error("Expected call to RepositoryMock.ClearCart")
	}
}

type mRepositoryMockDeleteFromCart struct {
	mock               *RepositoryMock
	defaultExpectation *RepositoryMockDeleteFromCartExpectation
	expectations       []*RepositoryMockDeleteFromCartExpectation

	callArgs []*RepositoryMockDeleteFromCartParams
	mutex    sync.RWMutex
}

// RepositoryMockDeleteFromCartExpectation specifies expectation struct of the Repository.DeleteFromCart
type RepositoryMockDeleteFromCartExpectation struct {
	mock    *RepositoryMock
	params  *RepositoryMockDeleteFromCartParams
	results *RepositoryMockDeleteFromCartResults
	Counter uint64
}

// RepositoryMockDeleteFromCartParams contains parameters of the Repository.DeleteFromCart
type RepositoryMockDeleteFromCartParams struct {
	ctx   context.Context
	user  int64
	sku   uint32
	count uint32
}

// RepositoryMockDeleteFromCartResults contains results of the Repository.DeleteFromCart
type RepositoryMockDeleteFromCartResults struct {
	err error
}

// Expect sets up expected params for Repository.DeleteFromCart
func (mmDeleteFromCart *mRepositoryMockDeleteFromCart) Expect(ctx context.Context, user int64, sku uint32, count uint32) *mRepositoryMockDeleteFromCart {
	if mmDeleteFromCart.mock.funcDeleteFromCart != nil {
		mmDeleteFromCart.mock.t.Fatalf("RepositoryMock.DeleteFromCart mock is already set by Set")
	}

	if mmDeleteFromCart.defaultExpectation == nil {
		mmDeleteFromCart.defaultExpectation = &RepositoryMockDeleteFromCartExpectation{}
	}

	mmDeleteFromCart.defaultExpectation.params = &RepositoryMockDeleteFromCartParams{ctx, user, sku, count}
	for _, e := range mmDeleteFromCart.expectations {
		if minimock.Equal(e.params, mmDeleteFromCart.defaultExpectation.params) {
			mmDeleteFromCart.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDeleteFromCart.defaultExpectation.params)
		}
	}

	return mmDeleteFromCart
}

// Inspect accepts an inspector function that has same arguments as the Repository.DeleteFromCart
func (mmDeleteFromCart *mRepositoryMockDeleteFromCart) Inspect(f func(ctx context.Context, user int64, sku uint32, count uint32)) *mRepositoryMockDeleteFromCart {
	if mmDeleteFromCart.mock.inspectFuncDeleteFromCart != nil {
		mmDeleteFromCart.mock.t.Fatalf("Inspect function is already set for RepositoryMock.DeleteFromCart")
	}

	mmDeleteFromCart.mock.inspectFuncDeleteFromCart = f

	return mmDeleteFromCart
}

// Return sets up results that will be returned by Repository.DeleteFromCart
func (mmDeleteFromCart *mRepositoryMockDeleteFromCart) Return(err error) *RepositoryMock {
	if mmDeleteFromCart.mock.funcDeleteFromCart != nil {
		mmDeleteFromCart.mock.t.Fatalf("RepositoryMock.DeleteFromCart mock is already set by Set")
	}

	if mmDeleteFromCart.defaultExpectation == nil {
		mmDeleteFromCart.defaultExpectation = &RepositoryMockDeleteFromCartExpectation{mock: mmDeleteFromCart.mock}
	}
	mmDeleteFromCart.defaultExpectation.results = &RepositoryMockDeleteFromCartResults{err}
	return mmDeleteFromCart.mock
}

// Set uses given function f to mock the Repository.DeleteFromCart method
func (mmDeleteFromCart *mRepositoryMockDeleteFromCart) Set(f func(ctx context.Context, user int64, sku uint32, count uint32) (err error)) *RepositoryMock {
	if mmDeleteFromCart.defaultExpectation != nil {
		mmDeleteFromCart.mock.t.Fatalf("Default expectation is already set for the Repository.DeleteFromCart method")
	}

	if len(mmDeleteFromCart.expectations) > 0 {
		mmDeleteFromCart.mock.t.Fatalf("Some expectations are already set for the Repository.DeleteFromCart method")
	}

	mmDeleteFromCart.mock.funcDeleteFromCart = f
	return mmDeleteFromCart.mock
}

// When sets expectation for the Repository.DeleteFromCart which will trigger the result defined by the following
// Then helper
func (mmDeleteFromCart *mRepositoryMockDeleteFromCart) When(ctx context.Context, user int64, sku uint32, count uint32) *RepositoryMockDeleteFromCartExpectation {
	if mmDeleteFromCart.mock.funcDeleteFromCart != nil {
		mmDeleteFromCart.mock.t.Fatalf("RepositoryMock.DeleteFromCart mock is already set by Set")
	}

	expectation := &RepositoryMockDeleteFromCartExpectation{
		mock:   mmDeleteFromCart.mock,
		params: &RepositoryMockDeleteFromCartParams{ctx, user, sku, count},
	}
	mmDeleteFromCart.expectations = append(mmDeleteFromCart.expectations, expectation)
	return expectation
}

// Then sets up Repository.DeleteFromCart return parameters for the expectation previously defined by the When method
func (e *RepositoryMockDeleteFromCartExpectation) Then(err error) *RepositoryMock {
	e.results = &RepositoryMockDeleteFromCartResults{err}
	return e.mock
}

// DeleteFromCart implements checkout.Repository
func (mmDeleteFromCart *RepositoryMock) DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint32) (err error) {
	mm_atomic.AddUint64(&mmDeleteFromCart.beforeDeleteFromCartCounter, 1)
	defer mm_atomic.AddUint64(&mmDeleteFromCart.afterDeleteFromCartCounter, 1)

	if mmDeleteFromCart.inspectFuncDeleteFromCart != nil {
		mmDeleteFromCart.inspectFuncDeleteFromCart(ctx, user, sku, count)
	}

	mm_params := &RepositoryMockDeleteFromCartParams{ctx, user, sku, count}

	// Record call args
	mmDeleteFromCart.DeleteFromCartMock.mutex.Lock()
	mmDeleteFromCart.DeleteFromCartMock.callArgs = append(mmDeleteFromCart.DeleteFromCartMock.callArgs, mm_params)
	mmDeleteFromCart.DeleteFromCartMock.mutex.Unlock()

	for _, e := range mmDeleteFromCart.DeleteFromCartMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDeleteFromCart.DeleteFromCartMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDeleteFromCart.DeleteFromCartMock.defaultExpectation.Counter, 1)
		mm_want := mmDeleteFromCart.DeleteFromCartMock.defaultExpectation.params
		mm_got := RepositoryMockDeleteFromCartParams{ctx, user, sku, count}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDeleteFromCart.t.Errorf("RepositoryMock.DeleteFromCart got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDeleteFromCart.DeleteFromCartMock.defaultExpectation.results
		if mm_results == nil {
			mmDeleteFromCart.t.Fatal("No results are set for the RepositoryMock.DeleteFromCart")
		}
		return (*mm_results).err
	}
	if mmDeleteFromCart.funcDeleteFromCart != nil {
		return mmDeleteFromCart.funcDeleteFromCart(ctx, user, sku, count)
	}
	mmDeleteFromCart.t.Fatalf("Unexpected call to RepositoryMock.DeleteFromCart. %v %v %v %v", ctx, user, sku, count)
	return
}

// DeleteFromCartAfterCounter returns a count of finished RepositoryMock.DeleteFromCart invocations
func (mmDeleteFromCart *RepositoryMock) DeleteFromCartAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeleteFromCart.afterDeleteFromCartCounter)
}

// DeleteFromCartBeforeCounter returns a count of RepositoryMock.DeleteFromCart invocations
func (mmDeleteFromCart *RepositoryMock) DeleteFromCartBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeleteFromCart.beforeDeleteFromCartCounter)
}

// Calls returns a list of arguments used in each call to RepositoryMock.DeleteFromCart.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDeleteFromCart *mRepositoryMockDeleteFromCart) Calls() []*RepositoryMockDeleteFromCartParams {
	mmDeleteFromCart.mutex.RLock()

	argCopy := make([]*RepositoryMockDeleteFromCartParams, len(mmDeleteFromCart.callArgs))
	copy(argCopy, mmDeleteFromCart.callArgs)

	mmDeleteFromCart.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteFromCartDone returns true if the count of the DeleteFromCart invocations corresponds
// the number of defined expectations
func (m *RepositoryMock) MinimockDeleteFromCartDone() bool {
	for _, e := range m.DeleteFromCartMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteFromCartMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteFromCartCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDeleteFromCart != nil && mm_atomic.LoadUint64(&m.afterDeleteFromCartCounter) < 1 {
		return false
	}
	return true
}

// MinimockDeleteFromCartInspect logs each unmet expectation
func (m *RepositoryMock) MinimockDeleteFromCartInspect() {
	for _, e := range m.DeleteFromCartMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepositoryMock.DeleteFromCart with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteFromCartMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteFromCartCounter) < 1 {
		if m.DeleteFromCartMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepositoryMock.DeleteFromCart")
		} else {
			m.t.Errorf("Expected call to RepositoryMock.DeleteFromCart with params: %#v", *m.DeleteFromCartMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDeleteFromCart != nil && mm_atomic.LoadUint64(&m.afterDeleteFromCartCounter) < 1 {
		m.t.Error("Expected call to RepositoryMock.DeleteFromCart")
	}
}

type mRepositoryMockListCart struct {
	mock               *RepositoryMock
	defaultExpectation *RepositoryMockListCartExpectation
	expectations       []*RepositoryMockListCartExpectation

	callArgs []*RepositoryMockListCartParams
	mutex    sync.RWMutex
}

// RepositoryMockListCartExpectation specifies expectation struct of the Repository.ListCart
type RepositoryMockListCartExpectation struct {
	mock    *RepositoryMock
	params  *RepositoryMockListCartParams
	results *RepositoryMockListCartResults
	Counter uint64
}

// RepositoryMockListCartParams contains parameters of the Repository.ListCart
type RepositoryMockListCartParams struct {
	ctx  context.Context
	user int64
}

// RepositoryMockListCartResults contains results of the Repository.ListCart
type RepositoryMockListCartResults struct {
	ia1 []models.Item
	err error
}

// Expect sets up expected params for Repository.ListCart
func (mmListCart *mRepositoryMockListCart) Expect(ctx context.Context, user int64) *mRepositoryMockListCart {
	if mmListCart.mock.funcListCart != nil {
		mmListCart.mock.t.Fatalf("RepositoryMock.ListCart mock is already set by Set")
	}

	if mmListCart.defaultExpectation == nil {
		mmListCart.defaultExpectation = &RepositoryMockListCartExpectation{}
	}

	mmListCart.defaultExpectation.params = &RepositoryMockListCartParams{ctx, user}
	for _, e := range mmListCart.expectations {
		if minimock.Equal(e.params, mmListCart.defaultExpectation.params) {
			mmListCart.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmListCart.defaultExpectation.params)
		}
	}

	return mmListCart
}

// Inspect accepts an inspector function that has same arguments as the Repository.ListCart
func (mmListCart *mRepositoryMockListCart) Inspect(f func(ctx context.Context, user int64)) *mRepositoryMockListCart {
	if mmListCart.mock.inspectFuncListCart != nil {
		mmListCart.mock.t.Fatalf("Inspect function is already set for RepositoryMock.ListCart")
	}

	mmListCart.mock.inspectFuncListCart = f

	return mmListCart
}

// Return sets up results that will be returned by Repository.ListCart
func (mmListCart *mRepositoryMockListCart) Return(ia1 []models.Item, err error) *RepositoryMock {
	if mmListCart.mock.funcListCart != nil {
		mmListCart.mock.t.Fatalf("RepositoryMock.ListCart mock is already set by Set")
	}

	if mmListCart.defaultExpectation == nil {
		mmListCart.defaultExpectation = &RepositoryMockListCartExpectation{mock: mmListCart.mock}
	}
	mmListCart.defaultExpectation.results = &RepositoryMockListCartResults{ia1, err}
	return mmListCart.mock
}

// Set uses given function f to mock the Repository.ListCart method
func (mmListCart *mRepositoryMockListCart) Set(f func(ctx context.Context, user int64) (ia1 []models.Item, err error)) *RepositoryMock {
	if mmListCart.defaultExpectation != nil {
		mmListCart.mock.t.Fatalf("Default expectation is already set for the Repository.ListCart method")
	}

	if len(mmListCart.expectations) > 0 {
		mmListCart.mock.t.Fatalf("Some expectations are already set for the Repository.ListCart method")
	}

	mmListCart.mock.funcListCart = f
	return mmListCart.mock
}

// When sets expectation for the Repository.ListCart which will trigger the result defined by the following
// Then helper
func (mmListCart *mRepositoryMockListCart) When(ctx context.Context, user int64) *RepositoryMockListCartExpectation {
	if mmListCart.mock.funcListCart != nil {
		mmListCart.mock.t.Fatalf("RepositoryMock.ListCart mock is already set by Set")
	}

	expectation := &RepositoryMockListCartExpectation{
		mock:   mmListCart.mock,
		params: &RepositoryMockListCartParams{ctx, user},
	}
	mmListCart.expectations = append(mmListCart.expectations, expectation)
	return expectation
}

// Then sets up Repository.ListCart return parameters for the expectation previously defined by the When method
func (e *RepositoryMockListCartExpectation) Then(ia1 []models.Item, err error) *RepositoryMock {
	e.results = &RepositoryMockListCartResults{ia1, err}
	return e.mock
}

// ListCart implements checkout.Repository
func (mmListCart *RepositoryMock) ListCart(ctx context.Context, user int64) (ia1 []models.Item, err error) {
	mm_atomic.AddUint64(&mmListCart.beforeListCartCounter, 1)
	defer mm_atomic.AddUint64(&mmListCart.afterListCartCounter, 1)

	if mmListCart.inspectFuncListCart != nil {
		mmListCart.inspectFuncListCart(ctx, user)
	}

	mm_params := &RepositoryMockListCartParams{ctx, user}

	// Record call args
	mmListCart.ListCartMock.mutex.Lock()
	mmListCart.ListCartMock.callArgs = append(mmListCart.ListCartMock.callArgs, mm_params)
	mmListCart.ListCartMock.mutex.Unlock()

	for _, e := range mmListCart.ListCartMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ia1, e.results.err
		}
	}

	if mmListCart.ListCartMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmListCart.ListCartMock.defaultExpectation.Counter, 1)
		mm_want := mmListCart.ListCartMock.defaultExpectation.params
		mm_got := RepositoryMockListCartParams{ctx, user}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmListCart.t.Errorf("RepositoryMock.ListCart got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmListCart.ListCartMock.defaultExpectation.results
		if mm_results == nil {
			mmListCart.t.Fatal("No results are set for the RepositoryMock.ListCart")
		}
		return (*mm_results).ia1, (*mm_results).err
	}
	if mmListCart.funcListCart != nil {
		return mmListCart.funcListCart(ctx, user)
	}
	mmListCart.t.Fatalf("Unexpected call to RepositoryMock.ListCart. %v %v", ctx, user)
	return
}

// ListCartAfterCounter returns a count of finished RepositoryMock.ListCart invocations
func (mmListCart *RepositoryMock) ListCartAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmListCart.afterListCartCounter)
}

// ListCartBeforeCounter returns a count of RepositoryMock.ListCart invocations
func (mmListCart *RepositoryMock) ListCartBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmListCart.beforeListCartCounter)
}

// Calls returns a list of arguments used in each call to RepositoryMock.ListCart.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmListCart *mRepositoryMockListCart) Calls() []*RepositoryMockListCartParams {
	mmListCart.mutex.RLock()

	argCopy := make([]*RepositoryMockListCartParams, len(mmListCart.callArgs))
	copy(argCopy, mmListCart.callArgs)

	mmListCart.mutex.RUnlock()

	return argCopy
}

// MinimockListCartDone returns true if the count of the ListCart invocations corresponds
// the number of defined expectations
func (m *RepositoryMock) MinimockListCartDone() bool {
	for _, e := range m.ListCartMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ListCartMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterListCartCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcListCart != nil && mm_atomic.LoadUint64(&m.afterListCartCounter) < 1 {
		return false
	}
	return true
}

// MinimockListCartInspect logs each unmet expectation
func (m *RepositoryMock) MinimockListCartInspect() {
	for _, e := range m.ListCartMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepositoryMock.ListCart with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ListCartMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterListCartCounter) < 1 {
		if m.ListCartMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepositoryMock.ListCart")
		} else {
			m.t.Errorf("Expected call to RepositoryMock.ListCart with params: %#v", *m.ListCartMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcListCart != nil && mm_atomic.LoadUint64(&m.afterListCartCounter) < 1 {
		m.t.Error("Expected call to RepositoryMock.ListCart")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *RepositoryMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockAddToCartInspect()

		m.MinimockClearCartInspect()

		m.MinimockDeleteFromCartInspect()

		m.MinimockListCartInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *RepositoryMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *RepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAddToCartDone() &&
		m.MinimockClearCartDone() &&
		m.MinimockDeleteFromCartDone() &&
		m.MinimockListCartDone()
}