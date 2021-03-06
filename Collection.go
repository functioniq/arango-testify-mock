// Code generated by mockery v1.0.0. DO NOT EDIT.

package arangomock

import context "context"
import driver "github.com/arangodb/go-driver"
import mock "github.com/stretchr/testify/mock"

// Collection is an autogenerated mock type for the Collection type
type Collection struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx
func (_m *Collection) Count(ctx context.Context) (int64, error) {
	ret := _m.Called(ctx)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context) int64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDocument provides a mock function with given fields: ctx, document
func (_m *Collection) CreateDocument(ctx context.Context, document interface{}) (driver.DocumentMeta, error) {
	ret := _m.Called(ctx, document)

	var r0 driver.DocumentMeta
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) driver.DocumentMeta); ok {
		r0 = rf(ctx, document)
	} else {
		r0 = ret.Get(0).(driver.DocumentMeta)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, document)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDocuments provides a mock function with given fields: ctx, documents
func (_m *Collection) CreateDocuments(ctx context.Context, documents interface{}) (driver.DocumentMetaSlice, driver.ErrorSlice, error) {
	ret := _m.Called(ctx, documents)

	var r0 driver.DocumentMetaSlice
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) driver.DocumentMetaSlice); ok {
		r0 = rf(ctx, documents)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.DocumentMetaSlice)
		}
	}

	var r1 driver.ErrorSlice
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) driver.ErrorSlice); ok {
		r1 = rf(ctx, documents)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(driver.ErrorSlice)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, interface{}) error); ok {
		r2 = rf(ctx, documents)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Database provides a mock function with given fields:
func (_m *Collection) Database() driver.Database {
	ret := _m.Called()

	var r0 driver.Database
	if rf, ok := ret.Get(0).(func() driver.Database); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Database)
		}
	}

	return r0
}

// DocumentExists provides a mock function with given fields: ctx, key
func (_m *Collection) DocumentExists(ctx context.Context, key string) (bool, error) {
	ret := _m.Called(ctx, key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnsureFullTextIndex provides a mock function with given fields: ctx, fields, options
func (_m *Collection) EnsureFullTextIndex(ctx context.Context, fields []string, options *driver.EnsureFullTextIndexOptions) (driver.Index, bool, error) {
	ret := _m.Called(ctx, fields, options)

	var r0 driver.Index
	if rf, ok := ret.Get(0).(func(context.Context, []string, *driver.EnsureFullTextIndexOptions) driver.Index); ok {
		r0 = rf(ctx, fields, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Index)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(context.Context, []string, *driver.EnsureFullTextIndexOptions) bool); ok {
		r1 = rf(ctx, fields, options)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, []string, *driver.EnsureFullTextIndexOptions) error); ok {
		r2 = rf(ctx, fields, options)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EnsureGeoIndex provides a mock function with given fields: ctx, fields, options
func (_m *Collection) EnsureGeoIndex(ctx context.Context, fields []string, options *driver.EnsureGeoIndexOptions) (driver.Index, bool, error) {
	ret := _m.Called(ctx, fields, options)

	var r0 driver.Index
	if rf, ok := ret.Get(0).(func(context.Context, []string, *driver.EnsureGeoIndexOptions) driver.Index); ok {
		r0 = rf(ctx, fields, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Index)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(context.Context, []string, *driver.EnsureGeoIndexOptions) bool); ok {
		r1 = rf(ctx, fields, options)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, []string, *driver.EnsureGeoIndexOptions) error); ok {
		r2 = rf(ctx, fields, options)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EnsureHashIndex provides a mock function with given fields: ctx, fields, options
func (_m *Collection) EnsureHashIndex(ctx context.Context, fields []string, options *driver.EnsureHashIndexOptions) (driver.Index, bool, error) {
	ret := _m.Called(ctx, fields, options)

	var r0 driver.Index
	if rf, ok := ret.Get(0).(func(context.Context, []string, *driver.EnsureHashIndexOptions) driver.Index); ok {
		r0 = rf(ctx, fields, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Index)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(context.Context, []string, *driver.EnsureHashIndexOptions) bool); ok {
		r1 = rf(ctx, fields, options)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, []string, *driver.EnsureHashIndexOptions) error); ok {
		r2 = rf(ctx, fields, options)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EnsurePersistentIndex provides a mock function with given fields: ctx, fields, options
func (_m *Collection) EnsurePersistentIndex(ctx context.Context, fields []string, options *driver.EnsurePersistentIndexOptions) (driver.Index, bool, error) {
	ret := _m.Called(ctx, fields, options)

	var r0 driver.Index
	if rf, ok := ret.Get(0).(func(context.Context, []string, *driver.EnsurePersistentIndexOptions) driver.Index); ok {
		r0 = rf(ctx, fields, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Index)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(context.Context, []string, *driver.EnsurePersistentIndexOptions) bool); ok {
		r1 = rf(ctx, fields, options)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, []string, *driver.EnsurePersistentIndexOptions) error); ok {
		r2 = rf(ctx, fields, options)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// EnsureSkipListIndex provides a mock function with given fields: ctx, fields, options
func (_m *Collection) EnsureSkipListIndex(ctx context.Context, fields []string, options *driver.EnsureSkipListIndexOptions) (driver.Index, bool, error) {
	ret := _m.Called(ctx, fields, options)

	var r0 driver.Index
	if rf, ok := ret.Get(0).(func(context.Context, []string, *driver.EnsureSkipListIndexOptions) driver.Index); ok {
		r0 = rf(ctx, fields, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Index)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(context.Context, []string, *driver.EnsureSkipListIndexOptions) bool); ok {
		r1 = rf(ctx, fields, options)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, []string, *driver.EnsureSkipListIndexOptions) error); ok {
		r2 = rf(ctx, fields, options)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ImportDocuments provides a mock function with given fields: ctx, documents, options
func (_m *Collection) ImportDocuments(ctx context.Context, documents interface{}, options *driver.ImportDocumentOptions) (driver.ImportDocumentStatistics, error) {
	ret := _m.Called(ctx, documents, options)

	var r0 driver.ImportDocumentStatistics
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, *driver.ImportDocumentOptions) driver.ImportDocumentStatistics); ok {
		r0 = rf(ctx, documents, options)
	} else {
		r0 = ret.Get(0).(driver.ImportDocumentStatistics)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, *driver.ImportDocumentOptions) error); ok {
		r1 = rf(ctx, documents, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Index provides a mock function with given fields: ctx, name
func (_m *Collection) Index(ctx context.Context, name string) (driver.Index, error) {
	ret := _m.Called(ctx, name)

	var r0 driver.Index
	if rf, ok := ret.Get(0).(func(context.Context, string) driver.Index); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.Index)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IndexExists provides a mock function with given fields: ctx, name
func (_m *Collection) IndexExists(ctx context.Context, name string) (bool, error) {
	ret := _m.Called(ctx, name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Indexes provides a mock function with given fields: ctx
func (_m *Collection) Indexes(ctx context.Context) ([]driver.Index, error) {
	ret := _m.Called(ctx)

	var r0 []driver.Index
	if rf, ok := ret.Get(0).(func(context.Context) []driver.Index); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]driver.Index)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Load provides a mock function with given fields: ctx
func (_m *Collection) Load(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *Collection) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Properties provides a mock function with given fields: ctx
func (_m *Collection) Properties(ctx context.Context) (driver.CollectionProperties, error) {
	ret := _m.Called(ctx)

	var r0 driver.CollectionProperties
	if rf, ok := ret.Get(0).(func(context.Context) driver.CollectionProperties); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(driver.CollectionProperties)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadDocument provides a mock function with given fields: ctx, key, result
func (_m *Collection) ReadDocument(ctx context.Context, key string, result interface{}) (driver.DocumentMeta, error) {
	ret := _m.Called(ctx, key, result)

	var r0 driver.DocumentMeta
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) driver.DocumentMeta); ok {
		r0 = rf(ctx, key, result)
	} else {
		r0 = ret.Get(0).(driver.DocumentMeta)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, interface{}) error); ok {
		r1 = rf(ctx, key, result)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadDocuments provides a mock function with given fields: ctx, keys, results
func (_m *Collection) ReadDocuments(ctx context.Context, keys []string, results interface{}) (driver.DocumentMetaSlice, driver.ErrorSlice, error) {
	ret := _m.Called(ctx, keys, results)

	var r0 driver.DocumentMetaSlice
	if rf, ok := ret.Get(0).(func(context.Context, []string, interface{}) driver.DocumentMetaSlice); ok {
		r0 = rf(ctx, keys, results)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.DocumentMetaSlice)
		}
	}

	var r1 driver.ErrorSlice
	if rf, ok := ret.Get(1).(func(context.Context, []string, interface{}) driver.ErrorSlice); ok {
		r1 = rf(ctx, keys, results)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(driver.ErrorSlice)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, []string, interface{}) error); ok {
		r2 = rf(ctx, keys, results)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Remove provides a mock function with given fields: ctx
func (_m *Collection) Remove(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveDocument provides a mock function with given fields: ctx, key
func (_m *Collection) RemoveDocument(ctx context.Context, key string) (driver.DocumentMeta, error) {
	ret := _m.Called(ctx, key)

	var r0 driver.DocumentMeta
	if rf, ok := ret.Get(0).(func(context.Context, string) driver.DocumentMeta); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(driver.DocumentMeta)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveDocuments provides a mock function with given fields: ctx, keys
func (_m *Collection) RemoveDocuments(ctx context.Context, keys []string) (driver.DocumentMetaSlice, driver.ErrorSlice, error) {
	ret := _m.Called(ctx, keys)

	var r0 driver.DocumentMetaSlice
	if rf, ok := ret.Get(0).(func(context.Context, []string) driver.DocumentMetaSlice); ok {
		r0 = rf(ctx, keys)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.DocumentMetaSlice)
		}
	}

	var r1 driver.ErrorSlice
	if rf, ok := ret.Get(1).(func(context.Context, []string) driver.ErrorSlice); ok {
		r1 = rf(ctx, keys)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(driver.ErrorSlice)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, []string) error); ok {
		r2 = rf(ctx, keys)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ReplaceDocument provides a mock function with given fields: ctx, key, document
func (_m *Collection) ReplaceDocument(ctx context.Context, key string, document interface{}) (driver.DocumentMeta, error) {
	ret := _m.Called(ctx, key, document)

	var r0 driver.DocumentMeta
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) driver.DocumentMeta); ok {
		r0 = rf(ctx, key, document)
	} else {
		r0 = ret.Get(0).(driver.DocumentMeta)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, interface{}) error); ok {
		r1 = rf(ctx, key, document)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplaceDocuments provides a mock function with given fields: ctx, keys, documents
func (_m *Collection) ReplaceDocuments(ctx context.Context, keys []string, documents interface{}) (driver.DocumentMetaSlice, driver.ErrorSlice, error) {
	ret := _m.Called(ctx, keys, documents)

	var r0 driver.DocumentMetaSlice
	if rf, ok := ret.Get(0).(func(context.Context, []string, interface{}) driver.DocumentMetaSlice); ok {
		r0 = rf(ctx, keys, documents)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.DocumentMetaSlice)
		}
	}

	var r1 driver.ErrorSlice
	if rf, ok := ret.Get(1).(func(context.Context, []string, interface{}) driver.ErrorSlice); ok {
		r1 = rf(ctx, keys, documents)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(driver.ErrorSlice)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, []string, interface{}) error); ok {
		r2 = rf(ctx, keys, documents)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Revision provides a mock function with given fields: ctx
func (_m *Collection) Revision(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetProperties provides a mock function with given fields: ctx, options
func (_m *Collection) SetProperties(ctx context.Context, options driver.SetCollectionPropertiesOptions) error {
	ret := _m.Called(ctx, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, driver.SetCollectionPropertiesOptions) error); ok {
		r0 = rf(ctx, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Statistics provides a mock function with given fields: ctx
func (_m *Collection) Statistics(ctx context.Context) (driver.CollectionStatistics, error) {
	ret := _m.Called(ctx)

	var r0 driver.CollectionStatistics
	if rf, ok := ret.Get(0).(func(context.Context) driver.CollectionStatistics); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(driver.CollectionStatistics)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Status provides a mock function with given fields: ctx
func (_m *Collection) Status(ctx context.Context) (driver.CollectionStatus, error) {
	ret := _m.Called(ctx)

	var r0 driver.CollectionStatus
	if rf, ok := ret.Get(0).(func(context.Context) driver.CollectionStatus); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(driver.CollectionStatus)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Truncate provides a mock function with given fields: ctx
func (_m *Collection) Truncate(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unload provides a mock function with given fields: ctx
func (_m *Collection) Unload(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateDocument provides a mock function with given fields: ctx, key, update
func (_m *Collection) UpdateDocument(ctx context.Context, key string, update interface{}) (driver.DocumentMeta, error) {
	ret := _m.Called(ctx, key, update)

	var r0 driver.DocumentMeta
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) driver.DocumentMeta); ok {
		r0 = rf(ctx, key, update)
	} else {
		r0 = ret.Get(0).(driver.DocumentMeta)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, interface{}) error); ok {
		r1 = rf(ctx, key, update)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDocuments provides a mock function with given fields: ctx, keys, updates
func (_m *Collection) UpdateDocuments(ctx context.Context, keys []string, updates interface{}) (driver.DocumentMetaSlice, driver.ErrorSlice, error) {
	ret := _m.Called(ctx, keys, updates)

	var r0 driver.DocumentMetaSlice
	if rf, ok := ret.Get(0).(func(context.Context, []string, interface{}) driver.DocumentMetaSlice); ok {
		r0 = rf(ctx, keys, updates)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(driver.DocumentMetaSlice)
		}
	}

	var r1 driver.ErrorSlice
	if rf, ok := ret.Get(1).(func(context.Context, []string, interface{}) driver.ErrorSlice); ok {
		r1 = rf(ctx, keys, updates)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(driver.ErrorSlice)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, []string, interface{}) error); ok {
		r2 = rf(ctx, keys, updates)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
