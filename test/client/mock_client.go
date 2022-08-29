package client

import (
	"testing"

	"github.com/fi-ts/cloud-go/api/client"
	accountingmock "github.com/fi-ts/cloud-go/test/mocks/accounting"
	clustermock "github.com/fi-ts/cloud-go/test/mocks/cluster"
	databasemock "github.com/fi-ts/cloud-go/test/mocks/database"
	healthmock "github.com/fi-ts/cloud-go/test/mocks/health"
	ipmock "github.com/fi-ts/cloud-go/test/mocks/ip"
	masterdatamock "github.com/fi-ts/cloud-go/test/mocks/masterdata"
	projectmock "github.com/fi-ts/cloud-go/test/mocks/project"
	s3mock "github.com/fi-ts/cloud-go/test/mocks/s3"
	tenantmock "github.com/fi-ts/cloud-go/test/mocks/tenant"
	versionmock "github.com/fi-ts/cloud-go/test/mocks/version"
	volumemock "github.com/fi-ts/cloud-go/test/mocks/volume"

	"github.com/stretchr/testify/mock"
)

type CloudMockFns struct {
	Accounting func(mock *mock.Mock)
	Cluster    func(mock *mock.Mock)
	Database   func(mock *mock.Mock)
	Health     func(mock *mock.Mock)
	IP         func(mock *mock.Mock)
	Masterdata func(mock *mock.Mock)
	Project    func(mock *mock.Mock)
	S3         func(mock *mock.Mock)
	Tenant     func(mock *mock.Mock)
	Version    func(mock *mock.Mock)
	Volume     func(mock *mock.Mock)
}

type CloudMockClient struct {
	accounting accountingmock.ClientService
	cluster    clustermock.ClientService
	database   databasemock.ClientService
	health     healthmock.ClientService
	ip         ipmock.ClientService
	masterdata masterdatamock.ClientService
	project    projectmock.ClientService
	s3         s3mock.ClientService
	tenant     tenantmock.ClientService
	version    versionmock.ClientService
	volume     volumemock.ClientService
}

func NewCloudMockClient(mockFns *CloudMockFns) (*CloudMockClient, *client.CloudAPI) {
	mockClient := &CloudMockClient{
		accounting: accountingmock.ClientService{},
		cluster:    clustermock.ClientService{},
		database:   databasemock.ClientService{},
		health:     healthmock.ClientService{},
		ip:         ipmock.ClientService{},
		masterdata: masterdatamock.ClientService{},
		project:    projectmock.ClientService{},
		s3:         s3mock.ClientService{},
		tenant:     tenantmock.ClientService{},
		version:    versionmock.ClientService{},
		volume:     volumemock.ClientService{},
	}

	client := &client.CloudAPI{
		Accounting: &mockClient.accounting,
		Cluster:    &mockClient.cluster,
		Database:   &mockClient.database,
		Health:     &mockClient.health,
		IP:         &mockClient.ip,
		Masterdata: &mockClient.masterdata,
		Project:    &mockClient.project,
		S3:         &mockClient.s3,
		Tenant:     &mockClient.tenant,
		Version:    &mockClient.version,
		Volume:     &mockClient.volume,
	}

	if mockFns == nil {
		return mockClient, client
	}

	if mockFns.Accounting != nil {
		mockFns.Accounting(&mockClient.accounting.Mock)
	}
	if mockFns.Cluster != nil {
		mockFns.Cluster(&mockClient.cluster.Mock)
	}
	if mockFns.Database != nil {
		mockFns.Database(&mockClient.database.Mock)
	}
	if mockFns.Health != nil {
		mockFns.Health(&mockClient.health.Mock)
	}
	if mockFns.IP != nil {
		mockFns.IP(&mockClient.ip.Mock)
	}
	if mockFns.Masterdata != nil {
		mockFns.Masterdata(&mockClient.masterdata.Mock)
	}
	if mockFns.Project != nil {
		mockFns.Project(&mockClient.project.Mock)
	}
	if mockFns.S3 != nil {
		mockFns.S3(&mockClient.s3.Mock)
	}
	if mockFns.Tenant != nil {
		mockFns.Tenant(&mockClient.tenant.Mock)
	}
	if mockFns.Version != nil {
		mockFns.Version(&mockClient.version.Mock)
	}
	if mockFns.Volume != nil {
		mockFns.Volume(&mockClient.volume.Mock)
	}

	return mockClient, client
}

func (c *CloudMockClient) AssertExpectations(t *testing.T) {
	_ = c.accounting.AssertExpectations(t)
	_ = c.cluster.AssertExpectations(t)
	_ = c.database.AssertExpectations(t)
	_ = c.health.AssertExpectations(t)
	_ = c.ip.AssertExpectations(t)
	_ = c.masterdata.AssertExpectations(t)
	_ = c.project.AssertExpectations(t)
	_ = c.s3.AssertExpectations(t)
	_ = c.tenant.AssertExpectations(t)
	_ = c.version.AssertExpectations(t)
	_ = c.volume.AssertExpectations(t)
}
