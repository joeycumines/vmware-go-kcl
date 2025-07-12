package cloudwatch

import (
	creds "github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMonitoringService(t *testing.T) {
	const region = `us-west-2`
	svc := NewMonitoringService(region, creds.AnonymousCredentials)
	assert.Equal(t, DEFAULT_CLOUDWATCH_METRICS_BUFFER_DURATION, svc.bufferDuration)
	assert.Equal(t, `us-west-2`, svc.region)
	if svc.credentials != creds.AnonymousCredentials {
		t.Errorf("Expected credentials to be anonymous, got %v", svc.credentials)
	}
}
