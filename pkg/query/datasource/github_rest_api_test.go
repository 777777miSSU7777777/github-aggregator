package datasource

import (
	"context"
	"testing"

	"github.com/777777miSSU7777777/github-aggregator/pkg/entity"

	"github.com/777777miSSU7777777/github-aggregator/pkg/crypto/randutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TestDataSource struct {
	mock.Mock
}

func (ds TestDataSource) GetUser(ctx context.Context, token string) ([]byte, error) {
	args := ds.Called(ctx, token)

	return args.Get(0).([]byte), args.Error(1)
}

func (ds TestDataSource) GetScopes(ctx context.Context, token string) ([]string, error) {
	args := ds.Called(ctx, token)

	return args.Get(0).([]string), args.Error(1)
}

func (ds TestDataSource) GetOrgs(ctx context.Context, token string) ([]byte, error) {
	args := ds.Called(ctx, token)

	return args.Get(0).([]byte), args.Error(1)
}

func (ds TestDataSource) GetOrgsRepos(ctx context.Context, token string, orgs []entity.Organization) ([][]byte, error) {
	args := ds.Called(ctx, token, orgs)

	return args.Get(0).([][]byte), args.Error(1)
}

func (ds TestDataSource) GetReposPullRequests(ctx context.Context, token string, repos []entity.Repository) ([][]byte, error) {
	args := ds.Called(ctx, token, repos)

	return args.Get(0).([][]byte), args.Error(1)
}

func TestGetUser__SameBytes__Equals(t *testing.T) {
	dataSrc := new(TestDataSource)

	randomBytes, _ := randutil.GenerateRandomBytes(16)

	dataSrc.On("GetUser", context.Background(), "123").Return(randomBytes, nil)

	testBytes, _ := dataSrc.GetUser(context.Background(), "123")

	assert.Equal(t, randomBytes, testBytes)
}

func TestGetUser__DifferentBytes__NotEquals(t *testing.T) {
	dataSrc := new(TestDataSource)

	randomBytes, _ := randutil.GenerateRandomBytes(16)

	dataSrc.On("GetUser", context.Background(), "123").Return(randomBytes, nil)

	testBytes, _ := dataSrc.GetUser(context.Background(), "123")

	randomBytes, _ = randutil.GenerateRandomBytes(16)

	assert.NotEqual(t, randomBytes, testBytes)
}

func TestGetScopes__SameStrings__Equals(t *testing.T) {
	dataSrc := new(TestDataSource)

	testStringArray := []string{"1", "2", "3"}

	dataSrc.On("GetScopes", context.Background(), "123").Return(testStringArray, nil)

	stringArray, _ := dataSrc.GetScopes(context.Background(), "123")

	assert.Equal(t, testStringArray, stringArray)
}

func TestGetScopes__DifferentStrings__NotEquals(t *testing.T) {
	dataSrc := new(TestDataSource)

	testStringArray := []string{"1", "2", "3"}

	dataSrc.On("GetScopes", context.Background(), "123").Return(testStringArray, nil)

	stringArray, _ := dataSrc.GetScopes(context.Background(), "123")

	testStringArray = []string{"3", "2", "1"}

	assert.NotEqual(t, testStringArray, stringArray)
}

func TestGetOrgs__SameBytes__Equals(t *testing.T) {
	dataSrc := new(TestDataSource)

	randomBytes, _ := randutil.GenerateRandomBytes(16)

	dataSrc.On("GetOrgs", context.Background(), "123").Return(randomBytes, nil)

	testBytes, _ := dataSrc.GetOrgs(context.Background(), "123")

	assert.Equal(t, randomBytes, testBytes)
}

func TestGetOrgs__DifferentBytes__NotEquals(t *testing.T) {
	dataSrc := new(TestDataSource)

	randomBytes, _ := randutil.GenerateRandomBytes(16)

	dataSrc.On("GetOrgs", context.Background(), "123").Return(randomBytes, nil)

	testBytes, _ := dataSrc.GetOrgs(context.Background(), "123")

	randomBytes, _ = randutil.GenerateRandomBytes(16)

	assert.NotEqual(t, randomBytes, testBytes)
}

func TestGetOrgsRepos__SameBytes__Equals(t *testing.T) {
	dataSrc := new(TestDataSource)

	randomBytes, _ := randutil.GenerateRandomBytes(16)

	dataSrc.On("GetOrgsRepos", context.Background(), "123", []entity.Organization{}).Return([][]byte{randomBytes}, nil)

	testBytes, _ := dataSrc.GetOrgsRepos(context.Background(), "123", []entity.Organization{})

	assert.Equal(t, [][]byte{randomBytes}, testBytes)
}

func TestGetOrgsRepos__DifferentBytes__NotEquals(t *testing.T) {
	dataSrc := new(TestDataSource)

	randomBytes, _ := randutil.GenerateRandomBytes(16)

	dataSrc.On("GetOrgsRepos", context.Background(), "123", []entity.Organization{}).Return([][]byte{randomBytes}, nil)

	testBytes, _ := dataSrc.GetOrgsRepos(context.Background(), "123", []entity.Organization{})

	randomBytes, _ = randutil.GenerateRandomBytes(16)

	assert.NotEqual(t, [][]byte{randomBytes}, testBytes)
}

func TestGetReposPullRequests__SameBytes__Equals(t *testing.T) {
	dataSrc := new(TestDataSource)

	randomBytes, _ := randutil.GenerateRandomBytes(16)

	dataSrc.On("GetReposPullRequests", context.Background(), "123", []entity.Repository{}).Return([][]byte{randomBytes}, nil)

	testBytes, _ := dataSrc.GetReposPullRequests(context.Background(), "123", []entity.Repository{})

	assert.Equal(t, [][]byte{randomBytes}, testBytes)
}

func TestGetReposPullRequests__DifferentBytes__NotEquals(t *testing.T) {
	dataSrc := new(TestDataSource)

	randomBytes, _ := randutil.GenerateRandomBytes(16)

	dataSrc.On("GetReposPullRequests", context.Background(), "123", []entity.Repository{}).Return([][]byte{randomBytes}, nil)

	testBytes, _ := dataSrc.GetReposPullRequests(context.Background(), "123", []entity.Repository{})

	randomBytes, _ = randutil.GenerateRandomBytes(16)

	assert.NotEqual(t, [][]byte{randomBytes}, testBytes)
}
