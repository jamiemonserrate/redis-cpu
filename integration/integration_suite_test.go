package integration_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/cf-redis-broker/integration"

	"testing"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var (
	binPath     string
	redisRunner *integration.RedisRunner
)

var _ = BeforeSuite(func() {
	// redisRunner = &integration.RedisRunner{}
	// startRedis()
})

var _ = AfterSuite(func() {
	// gexec.CleanupBuildArtifacts()
	// stopRedis()
})

// func stopRedis() {
// 	redisRunner.Stop()
// }

// func startRedis() {
// 	redisRunner.Start([]string{"--port", "6480"})
// }
