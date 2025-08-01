/*
 * Copyright (c) 2018 VMware, Inc.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and
 * associated documentation files (the "Software"), to deal in the Software without restriction, including
 * without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is furnished to do
 * so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial
 * portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT
 * NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
 * IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
 * WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
 * SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */
package test

import (
	"os"
	"sync"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	chk "github.com/vmware/vmware-go-kcl/clientlibrary/checkpoint"
	cfg "github.com/vmware/vmware-go-kcl/clientlibrary/config"
	par "github.com/vmware/vmware-go-kcl/clientlibrary/partition"
	wk "github.com/vmware/vmware-go-kcl/clientlibrary/worker"
)

func TestWorkerInjectCheckpointer(t *testing.T) {
	kclConfig := cfg.NewKinesisClientLibConfig(appName, streamName, regionName, workerID).
		WithInitialPositionInStream(cfg.LATEST).
		WithMaxRecords(10).
		WithMaxLeasesForWorker(1).
		WithShardSyncIntervalMillis(5000).
		WithFailoverTimeMillis(300000)

	// Configure LocalStack endpoints if environment variables are set
	kclConfig = configureLocalStackEndpoints(kclConfig)

	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	assert.Equal(t, regionName, kclConfig.RegionName)
	assert.Equal(t, streamName, kclConfig.StreamName)

	// configure cloudwatch as metrics system
	kclConfig.WithMonitoringService(getMetricsConfig(kclConfig, metricsSystem))

	// Put some data into stream.
	kc := NewKinesisClient(t, regionName, kclConfig.KinesisEndpoint, kclConfig.KinesisCredentials)
	// publishSomeData(t, kc)
	stop := continuouslyPublishSomeData(t, kc)
	defer stop()

	// custom checkpointer or a mock checkpointer.
	checkpointer := chk.NewDynamoCheckpoint(kclConfig)

	// Inject a custom checkpointer into the worker.
	worker := wk.NewWorker(recordProcessorFactory(t), kclConfig).
		WithCheckpointer(checkpointer)

	err := worker.Start()
	assert.Nil(t, err)

	// wait a few seconds before shutdown processing
	time.Sleep(30 * time.Second)
	worker.Shutdown()

	// verify the checkpointer after graceful shutdown
	status := &par.ShardStatus{
		ID:  shardID,
		Mux: &sync.RWMutex{},
	}
	checkpointer.FetchCheckpoint(status)

	// checkpointer should be the same
	assert.NotEmpty(t, status.Checkpoint)

	// Only the lease owner has been wiped out
	assert.Equal(t, "", status.GetLeaseOwner())

}

func TestWorkerInjectKinesis(t *testing.T) {
	kclConfig := cfg.NewKinesisClientLibConfig(appName, streamName, regionName, workerID).
		WithInitialPositionInStream(cfg.LATEST).
		WithMaxRecords(10).
		WithMaxLeasesForWorker(1).
		WithShardSyncIntervalMillis(5000).
		WithFailoverTimeMillis(300000)

	// Configure LocalStack endpoints if environment variables are set
	kclConfig = configureLocalStackEndpoints(kclConfig)

	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	assert.Equal(t, regionName, kclConfig.RegionName)
	assert.Equal(t, streamName, kclConfig.StreamName)

	// configure cloudwatch as metrics system
	kclConfig.WithMonitoringService(getMetricsConfig(kclConfig, metricsSystem))

	// create custom Kinesis with LocalStack configuration
	kc := NewKinesisClient(t, regionName, kclConfig.KinesisEndpoint, kclConfig.KinesisCredentials)

	// Put some data into stream.
	// publishSomeData(t, kc)
	stop := continuouslyPublishSomeData(t, kc)
	defer stop()

	// Inject a custom checkpointer into the worker.
	worker := wk.NewWorker(recordProcessorFactory(t), kclConfig).
		WithKinesis(kc)

	err := worker.Start()
	assert.Nil(t, err)

	// wait a few seconds before shutdown processing
	time.Sleep(30 * time.Second)
	worker.Shutdown()
}

func TestWorkerInjectKinesisAndCheckpointer(t *testing.T) {
	kclConfig := cfg.NewKinesisClientLibConfig(appName, streamName, regionName, workerID).
		WithInitialPositionInStream(cfg.LATEST).
		WithMaxRecords(10).
		WithMaxLeasesForWorker(1).
		WithShardSyncIntervalMillis(5000).
		WithFailoverTimeMillis(300000)

	// Configure LocalStack endpoints if environment variables are set
	kclConfig = configureLocalStackEndpoints(kclConfig)

	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	assert.Equal(t, regionName, kclConfig.RegionName)
	assert.Equal(t, streamName, kclConfig.StreamName)

	// configure cloudwatch as metrics system
	kclConfig.WithMonitoringService(getMetricsConfig(kclConfig, metricsSystem))

	// create custom Kinesis with LocalStack configuration
	kc := NewKinesisClient(t, regionName, kclConfig.KinesisEndpoint, kclConfig.KinesisCredentials)

	// Put some data into stream.
	// publishSomeData(t, kc)
	stop := continuouslyPublishSomeData(t, kc)
	defer stop()

	// custom checkpointer or a mock checkpointer.
	checkpointer := chk.NewDynamoCheckpoint(kclConfig)

	// Inject both custom checkpointer and kinesis into the worker.
	worker := wk.NewWorker(recordProcessorFactory(t), kclConfig).
		WithKinesis(kc).
		WithCheckpointer(checkpointer)

	err := worker.Start()
	assert.Nil(t, err)

	// wait a few seconds before shutdown processing
	time.Sleep(30 * time.Second)
	worker.Shutdown()
}
