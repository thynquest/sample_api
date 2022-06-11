package service

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"go.mongodb.org/mongo-driver/mongo/options"
	"spin.sample.trial/storage"
)

type mockDriver struct {
	mock *mtest.T
}

func newMockDriver(t *mtest.T) *mockDriver {
	return &mockDriver{
		mock: t,
	}
}

func (m *mockDriver) Connect() (context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	return ctx, cancel, m.mock.Client.Connect(ctx)
}

func (m *mockDriver) Close(ctx context.Context, cancel context.CancelFunc) {}

func (m *mockDriver) Insert(ctx context.Context, data interface{}) (interface{}, error) {
	result, err := m.mock.Coll.InsertOne(ctx, data)
	return result, err
}

func (m *mockDriver) Retrieve(ctx context.Context) (*mongo.Cursor, error) {
	options := options.Find()
	cursor, errFind := m.mock.Coll.Find(ctx, bson.D{{Key: "Amount", Value: 12445}}, options)
	return cursor, errFind
}

func TestInsertInvoice(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("TestInsertInvoice", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse(bson.D{{Key: "i", Value: 1}}...))
		driver := newMockDriver(mt)
		st := storage.NewDriverStorage(driver)
		svc := NewInvoiceSvc(st)
		res, err := svc.Create("testcomp", 12345, "2022-04-25", "2022-04-29")
		assert.Equal(mt, true, res)
		assert.Nil(mt, err)
	})
}

func TestRetrieveInvoice(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("TestRetrieveInvoice", func(mt *mtest.T) {
		first := mtest.CreateCursorResponse(
			1,
			"test.mycol",
			mtest.FirstBatch,
			bson.D{
				{Key: "Amount", Value: 12345},
				{Key: "Company", Value: "mycompany"},
				{Key: "IssueDate", Value: "2022-04-25"},
				{Key: "DueDate", Value: "2022-04-29"},
			})
		second := mtest.CreateCursorResponse(
			2,
			"test.mycol",
			mtest.NextBatch,
			bson.D{
				{Key: "Amount", Value: 12445},
				{Key: "Company", Value: "myothercompany"},
				{Key: "IssueDate", Value: "2022-04-25"},
				{Key: "DueDate", Value: "2022-04-29"},
			})
		mt.AddMockResponses(first, second)
		driver := newMockDriver(mt)
		st := storage.NewDriverStorage(driver)
		svc := NewInvoiceSvc(st)
		datas, err := svc.Retrieve()
		assert.Equal(mt, 2, len(datas))
		assert.Nil(mt, err)
	})
}
