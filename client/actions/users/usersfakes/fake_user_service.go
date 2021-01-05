// Code generated by counterfeiter. DO NOT EDIT.
package usersfakes

import (
	"sync"

	"github.com/IBM/satcon-client-go/client/actions/users"
	"github.com/IBM/satcon-client-go/client/types"
)

type FakeUserService struct {
	MeStub        func() (*types.User, error)
	meMutex       sync.RWMutex
	meArgsForCall []struct {
	}
	meReturns struct {
		result1 *types.User
		result2 error
	}
	meReturnsOnCall map[int]struct {
		result1 *types.User
		result2 error
	}
	SignInStub        func(string, string) (*types.Token, error)
	signInMutex       sync.RWMutex
	signInArgsForCall []struct {
		arg1 string
		arg2 string
	}
	signInReturns struct {
		result1 *types.Token
		result2 error
	}
	signInReturnsOnCall map[int]struct {
		result1 *types.Token
		result2 error
	}
	SignUpStub        func(string, string, string, string, string) (*types.Token, error)
	signUpMutex       sync.RWMutex
	signUpArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}
	signUpReturns struct {
		result1 *types.Token
		result2 error
	}
	signUpReturnsOnCall map[int]struct {
		result1 *types.Token
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUserService) Me() (*types.User, error) {
	fake.meMutex.Lock()
	ret, specificReturn := fake.meReturnsOnCall[len(fake.meArgsForCall)]
	fake.meArgsForCall = append(fake.meArgsForCall, struct {
	}{})
	stub := fake.MeStub
	fakeReturns := fake.meReturns
	fake.recordInvocation("Me", []interface{}{})
	fake.meMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserService) MeCallCount() int {
	fake.meMutex.RLock()
	defer fake.meMutex.RUnlock()
	return len(fake.meArgsForCall)
}

func (fake *FakeUserService) MeCalls(stub func() (*types.User, error)) {
	fake.meMutex.Lock()
	defer fake.meMutex.Unlock()
	fake.MeStub = stub
}

func (fake *FakeUserService) MeReturns(result1 *types.User, result2 error) {
	fake.meMutex.Lock()
	defer fake.meMutex.Unlock()
	fake.MeStub = nil
	fake.meReturns = struct {
		result1 *types.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserService) MeReturnsOnCall(i int, result1 *types.User, result2 error) {
	fake.meMutex.Lock()
	defer fake.meMutex.Unlock()
	fake.MeStub = nil
	if fake.meReturnsOnCall == nil {
		fake.meReturnsOnCall = make(map[int]struct {
			result1 *types.User
			result2 error
		})
	}
	fake.meReturnsOnCall[i] = struct {
		result1 *types.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserService) SignIn(arg1 string, arg2 string) (*types.Token, error) {
	fake.signInMutex.Lock()
	ret, specificReturn := fake.signInReturnsOnCall[len(fake.signInArgsForCall)]
	fake.signInArgsForCall = append(fake.signInArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.SignInStub
	fakeReturns := fake.signInReturns
	fake.recordInvocation("SignIn", []interface{}{arg1, arg2})
	fake.signInMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserService) SignInCallCount() int {
	fake.signInMutex.RLock()
	defer fake.signInMutex.RUnlock()
	return len(fake.signInArgsForCall)
}

func (fake *FakeUserService) SignInCalls(stub func(string, string) (*types.Token, error)) {
	fake.signInMutex.Lock()
	defer fake.signInMutex.Unlock()
	fake.SignInStub = stub
}

func (fake *FakeUserService) SignInArgsForCall(i int) (string, string) {
	fake.signInMutex.RLock()
	defer fake.signInMutex.RUnlock()
	argsForCall := fake.signInArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserService) SignInReturns(result1 *types.Token, result2 error) {
	fake.signInMutex.Lock()
	defer fake.signInMutex.Unlock()
	fake.SignInStub = nil
	fake.signInReturns = struct {
		result1 *types.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeUserService) SignInReturnsOnCall(i int, result1 *types.Token, result2 error) {
	fake.signInMutex.Lock()
	defer fake.signInMutex.Unlock()
	fake.SignInStub = nil
	if fake.signInReturnsOnCall == nil {
		fake.signInReturnsOnCall = make(map[int]struct {
			result1 *types.Token
			result2 error
		})
	}
	fake.signInReturnsOnCall[i] = struct {
		result1 *types.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeUserService) SignUp(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string) (*types.Token, error) {
	fake.signUpMutex.Lock()
	ret, specificReturn := fake.signUpReturnsOnCall[len(fake.signUpArgsForCall)]
	fake.signUpArgsForCall = append(fake.signUpArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.SignUpStub
	fakeReturns := fake.signUpReturns
	fake.recordInvocation("SignUp", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.signUpMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserService) SignUpCallCount() int {
	fake.signUpMutex.RLock()
	defer fake.signUpMutex.RUnlock()
	return len(fake.signUpArgsForCall)
}

func (fake *FakeUserService) SignUpCalls(stub func(string, string, string, string, string) (*types.Token, error)) {
	fake.signUpMutex.Lock()
	defer fake.signUpMutex.Unlock()
	fake.SignUpStub = stub
}

func (fake *FakeUserService) SignUpArgsForCall(i int) (string, string, string, string, string) {
	fake.signUpMutex.RLock()
	defer fake.signUpMutex.RUnlock()
	argsForCall := fake.signUpArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeUserService) SignUpReturns(result1 *types.Token, result2 error) {
	fake.signUpMutex.Lock()
	defer fake.signUpMutex.Unlock()
	fake.SignUpStub = nil
	fake.signUpReturns = struct {
		result1 *types.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeUserService) SignUpReturnsOnCall(i int, result1 *types.Token, result2 error) {
	fake.signUpMutex.Lock()
	defer fake.signUpMutex.Unlock()
	fake.SignUpStub = nil
	if fake.signUpReturnsOnCall == nil {
		fake.signUpReturnsOnCall = make(map[int]struct {
			result1 *types.Token
			result2 error
		})
	}
	fake.signUpReturnsOnCall[i] = struct {
		result1 *types.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeUserService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.meMutex.RLock()
	defer fake.meMutex.RUnlock()
	fake.signInMutex.RLock()
	defer fake.signInMutex.RUnlock()
	fake.signUpMutex.RLock()
	defer fake.signUpMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUserService) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ users.UserService = new(FakeUserService)