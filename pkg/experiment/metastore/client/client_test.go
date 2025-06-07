package metastoreclient

import (
	"context"
	"sync"
	"testing"

	"github.com/grafana/dskit/flagext"
	"github.com/grafana/dskit/grpcclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	metastorev1 "github.com/grafana/pyroscope/api/gen/proto/go/metastore/v1"
	"github.com/grafana/pyroscope/pkg/test"
	"github.com/grafana/pyroscope/pkg/test/mocks/mockdiscovery"
)

const nServers = 3

func TestUnavailable(t *testing.T) {
	d := mockdiscovery.NewMockDiscovery(t)
	d.On("Subscribe", mock.Anything).Return()
	l := test.NewTestingLogger(t)
	c := New(l, grpcclient.Config{}, d)
	ports, err := test.GetFreePorts(nServers)
	assert.NoError(t, err)

	d.On("Rediscover").Run(func(args mock.Arguments) {
	}).Return()

	c.updateServers(createServers(ports))
	res, err := c.AddBlock(context.Background(), &metastorev1.AddBlockRequest{})
	require.Error(t, err)
	require.Nil(t, res)

}

func TestUnavailable_Rediscover_Wrong_Leader(t *testing.T) {
	t.Run("AddBlock", func(t *testing.T) {
		testRediscoverWrongLeader(t, func(c *Client) {
			res, err := c.AddBlock(context.Background(), &metastorev1.AddBlockRequest{})
			require.NoError(t, err)
			require.NotNil(t, res)
		})
	})
	t.Run("QueryMetadata", func(t *testing.T) {
		testRediscoverWrongLeader(t, func(c *Client) {
			res, err := c.QueryMetadata(context.Background(), &metastorev1.QueryMetadataRequest{})
			require.NoError(t, err)
			require.NotNil(t, res)
		})
	})
	t.Run("PollCompactionJobs", func(t *testing.T) {
		testRediscoverWrongLeader(t, func(c *Client) {
			res, err := c.PollCompactionJobs(context.Background(), &metastorev1.PollCompactionJobsRequest{})
			require.NoError(t, err)
			require.NotNil(t, res)
		})
	})
	t.Run("GetProfileStats", func(t *testing.T) {
		testRediscoverWrongLeader(t, func(c *Client) {
			res, err := c.GetTenant(context.Background(), &metastorev1.GetTenantRequest{})
			require.NoError(t, err)
			require.NotNil(t, res)
		})
	})
}

func testRediscoverWrongLeader(t *testing.T, f func(c *Client)) {
	d := mockdiscovery.NewMockDiscovery(t)
	d.On("Subscribe", mock.Anything).Return()
	l := test.NewTestingLogger(t)
	config := &grpcclient.Config{}
	flagext.DefaultValues(config)
	ports, err := test.GetFreePorts(nServers * 2)
	assert.NoError(t, err)

	p1 := ports[:nServers]
	dServers1 := createServers(p1)

	p2 := ports[nServers:]
	dServers2 := createServers(p2)
	mockServers2, dialOpt := createMockServers(t, l, dServers2)
	defer mockServers2.Close()

	c := New(l, *config, d, dialOpt...)
	m := sync.Mutex{}
	verify := func() {}
	initWrongLeaderCalled := false
	d.On("Rediscover", mock.Anything).Run(func(args mock.Arguments) {
		m.Lock()
		defer m.Unlock()
		if !initWrongLeaderCalled {
			initWrongLeaderCalled = true
			verify = mockServers2.InitWrongLeader()
			// call updateServers twice
			c.updateServers(dServers2)
			c.updateServers(dServers2)
		}
	}).Return()

	c.updateServers(dServers1)
	f(c)
	verify()
}

func TestServerError(t *testing.T) {
	d := mockdiscovery.NewMockDiscovery(t)
	d.On("Subscribe", mock.Anything).Return()
	l := test.NewTestingLogger(t)
	c := New(l, grpcclient.Config{}, d)

	d.On("Rediscover").Run(func(args mock.Arguments) {
	}).Return()

	res, err := c.AddBlock(context.Background(), &metastorev1.AddBlockRequest{})
	require.Error(t, err)
	require.Nil(t, res)
}
