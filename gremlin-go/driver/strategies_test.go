/*
Licensed to the Apache Software Foundation (ASF) Under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package gremlingo

import (
	"crypto/tls"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getModernGraph(t *testing.T, url string, auth *AuthInfo, tls *tls.Config) *GraphTraversalSource {
	remote, err := NewDriverRemoteConnection(url,
		func(settings *DriverRemoteConnectionSettings) {
			settings.TlsConfig = tls
			settings.AuthInfo = auth
			settings.TraversalSource = testServerModernGraphAlias
		})
	assert.Nil(t, err)
	assert.NotNil(t, remote)
	g := Traversal_().WithRemote(remote)

	return g
}

func TestStrategy(t *testing.T) {
	testNoAuthUrl := getEnvOrDefaultString("GREMLIN_SERVER_URL", noAuthUrl)

	t.Run("Test read with ConnectiveStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		count, err := g.WithStrategies(ConnectiveStrategy()).V().Count().ToList()
		assert.Nil(t, err)
		assert.NotNil(t, count)
		assert.Equal(t, 1, len(count))
		val, err := count[0].GetInt32()
		assert.Nil(t, err)
		assert.Equal(t, int32(6), val)
	})

	t.Run("Test read with OptionsStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		count, err := g.WithStrategies(OptionsStrategy(map[string]interface{}{"a": "b"})).V().Count().ToList()
		assert.Nil(t, err)
		assert.NotNil(t, count)
		assert.Equal(t, 1, len(count))
		val, err := count[0].GetInt32()
		assert.Nil(t, err)
		assert.Equal(t, int32(6), val)
	})

	t.Run("Test read with PartitionStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		count, err := g.WithStrategies(PartitionStrategy("partition", "write", []string{"read"}, true)).V().Count().ToList()
		assert.Nil(t, err)
		assert.NotNil(t, count)
		assert.Equal(t, 1, len(count))
		val, err := count[0].GetInt32()
		assert.Nil(t, err)
		assert.Equal(t, int32(0), val)
	})

	t.Run("Test read with SeedStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		count, err := g.WithStrategies(SeedStrategy(1)).V().Count().ToList()
		assert.Nil(t, err)
		assert.NotNil(t, count)
		assert.Equal(t, 1, len(count))
		val, err := count[0].GetInt32()
		assert.Nil(t, err)
		assert.Equal(t, int32(6), val)
	})

	t.Run("Test read with SubgraphStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		count, err := g.WithStrategies(SubgraphStrategy(T__.HasLabel(testLabel), nil, nil)).V().Count().ToList()
		assert.Nil(t, err)
		assert.NotNil(t, count)
		assert.Equal(t, 1, len(count))
		val, err := count[0].GetInt32()
		assert.Nil(t, err)
		assert.Equal(t, int32(0), val)
	})

	t.Run("Test bytecode generation for MatchAlgorithmStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		bytecode := g.WithStrategies(MatchAlgorithmStrategy("greedy")).bytecode
		assert.Equal(t, 1, len(bytecode.sourceInstructions))
		assert.Equal(t, 1, len(bytecode.sourceInstructions[0].arguments))
		assert.Equal(t, "withStrategies", bytecode.sourceInstructions[0].operator)
		assert.Equal(t, "org.apache.tinkerpop.gremlin.process.traversal.strategy.finalization.MatchAlgorithmStrategy",
			bytecode.sourceInstructions[0].arguments[0].(*traversalStrategy).name)
		assert.Equal(t, map[string]interface{}{"matchAlgorithm": "greedy"},
			bytecode.sourceInstructions[0].arguments[0].(*traversalStrategy).configuration)
	})

	t.Run("Test read with AdjacentToIncidentStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		count, err := g.WithStrategies(AdjacentToIncidentStrategy()).V().Count().ToList()
		assert.Nil(t, err)
		assert.NotNil(t, count)
		assert.Equal(t, 1, len(count))
		val, err := count[0].GetInt32()
		assert.Nil(t, err)
		assert.Equal(t, int32(6), val)
	})

	t.Run("Test read with ByModulatorOptimizationStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		count, err := g.WithStrategies(ByModulatorOptimizationStrategy()).V().Count().ToList()
		assert.Nil(t, err)
		assert.NotNil(t, count)
		assert.Equal(t, 1, len(count))
		val, err := count[0].GetInt32()
		assert.Nil(t, err)
		assert.Equal(t, int32(6), val)
	})

	t.Run("Test read with CountStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		count, err := g.WithStrategies(CountStrategy()).V().Count().ToList()
		assert.Nil(t, err)
		assert.NotNil(t, count)
		assert.Equal(t, 1, len(count))
		val, err := count[0].GetInt32()
		assert.Nil(t, err)
		assert.Equal(t, int32(6), val)
	})

	t.Run("Test read with EarlyLimitStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		count, err := g.WithStrategies(EarlyLimitStrategy()).V().Count().ToList()
		assert.Nil(t, err)
		assert.NotNil(t, count)
		assert.Equal(t, 1, len(count))
		val, err := count[0].GetInt32()
		assert.Nil(t, err)
		assert.Equal(t, int32(6), val)
	})

	t.Run("Test read with FilterRankingStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		count, err := g.WithStrategies(FilterRankingStrategy()).V().Count().ToList()
		assert.Nil(t, err)
		assert.NotNil(t, count)
		assert.Equal(t, 1, len(count))
		val, err := count[0].GetInt32()
		assert.Nil(t, err)
		assert.Equal(t, int32(6), val)
	})

	t.Run("Test read with IdentityRemovalStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		count, err := g.WithStrategies(IdentityRemovalStrategy()).V().Count().ToList()
		assert.Nil(t, err)
		assert.NotNil(t, count)
		assert.Equal(t, 1, len(count))
		val, err := count[0].GetInt32()
		assert.Nil(t, err)
		assert.Equal(t, int32(6), val)
	})

	t.Run("Test read with IncidentToAdjacentStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		count, err := g.WithStrategies(IncidentToAdjacentStrategy()).V().Count().ToList()
		assert.Nil(t, err)
		assert.NotNil(t, count)
		assert.Equal(t, 1, len(count))
		val, err := count[0].GetInt32()
		assert.Nil(t, err)
		assert.Equal(t, int32(6), val)
	})

	t.Run("Test read with InlineFilterStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		count, err := g.WithStrategies(InlineFilterStrategy()).V().Count().ToList()
		assert.Nil(t, err)
		assert.NotNil(t, count)
		assert.Equal(t, 1, len(count))
		val, err := count[0].GetInt32()
		assert.Nil(t, err)
		assert.Equal(t, int32(6), val)
	})

	t.Run("Test read with LazyBarrierStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		count, err := g.WithStrategies(LazyBarrierStrategy()).V().Count().ToList()
		assert.Nil(t, err)
		assert.NotNil(t, count)
		assert.Equal(t, 1, len(count))
		val, err := count[0].GetInt32()
		assert.Nil(t, err)
		assert.Equal(t, int32(6), val)
	})

	t.Run("Test read with MatchPredicateStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		count, err := g.WithStrategies(MatchPredicateStrategy()).V().Count().ToList()
		assert.Nil(t, err)
		assert.NotNil(t, count)
		assert.Equal(t, 1, len(count))
		val, err := count[0].GetInt32()
		assert.Nil(t, err)
		assert.Equal(t, int32(6), val)
	})

	t.Run("Test read with OrderLimitStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		count, err := g.WithStrategies(OrderLimitStrategy()).V().Count().ToList()
		assert.Nil(t, err)
		assert.NotNil(t, count)
		assert.Equal(t, 1, len(count))
		val, err := count[0].GetInt32()
		assert.Nil(t, err)
		assert.Equal(t, int32(6), val)
	})

	t.Run("Test read with PathProcessorStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		count, err := g.WithStrategies(PathProcessorStrategy()).V().Count().ToList()
		assert.Nil(t, err)
		assert.NotNil(t, count)
		assert.Equal(t, 1, len(count))
		val, err := count[0].GetInt32()
		assert.Nil(t, err)
		assert.Equal(t, int32(6), val)
	})

	t.Run("Test read with PathRetractionStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		count, err := g.WithStrategies(PathRetractionStrategy()).V().Count().ToList()
		assert.Nil(t, err)
		assert.NotNil(t, count)
		assert.Equal(t, 1, len(count))
		val, err := count[0].GetInt32()
		assert.Nil(t, err)
		assert.Equal(t, int32(6), val)
	})

	t.Run("Test read with ProductiveByStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		count, err := g.WithStrategies(ProductiveByStrategy([]string{"a", "b"})).V().Count().ToList()
		assert.Nil(t, err)
		assert.NotNil(t, count)
		assert.Equal(t, 1, len(count))
		val, err := count[0].GetInt32()
		assert.Nil(t, err)
		assert.Equal(t, int32(6), val)
	})

	t.Run("Test read with RepeatUnrollStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		count, err := g.WithStrategies(RepeatUnrollStrategy()).V().Count().ToList()
		assert.Nil(t, err)
		assert.NotNil(t, count)
		assert.Equal(t, 1, len(count))
		val, err := count[0].GetInt32()
		assert.Nil(t, err)
		assert.Equal(t, int32(6), val)
	})

	t.Run("Test read with EdgeLabelVerificationStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		count, err := g.WithStrategies(EdgeLabelVerificationStrategy(true, true)).V().Count().ToList()
		assert.Nil(t, err)
		assert.NotNil(t, count)
		assert.Equal(t, 1, len(count))
		val, err := count[0].GetInt32()
		assert.Nil(t, err)
		assert.Equal(t, int32(6), val)
	})

	t.Run("Test read with LambdaRestrictionStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		count, err := g.WithStrategies(LambdaRestrictionStrategy()).V().Count().ToList()
		assert.Nil(t, err)
		assert.NotNil(t, count)
		assert.Equal(t, 1, len(count))
		val, err := count[0].GetInt32()
		assert.Nil(t, err)
		assert.Equal(t, int32(6), val)
	})

	t.Run("Test read with TestReadOnlyStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		count, err := g.WithStrategies(ReadOnlyStrategy()).V().Count().ToList()
		assert.Nil(t, err)
		assert.NotNil(t, count)
		assert.Equal(t, 1, len(count))
		val, err := count[0].GetInt32()
		assert.Nil(t, err)
		assert.Equal(t, int32(6), val)
	})

	t.Run("Test write with TestReadOnlyStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		promise := g.WithStrategies(ReadOnlyStrategy()).AddV("person").Property("name", "foo").Iterate()
		assert.NotNil(t, <-promise)
	})

	t.Run("Test read with ReservedKeysVerificationStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		strategy := ReservedKeysVerificationStrategy(true, true, []string{"xyz"})
		count, err := g.WithStrategies(strategy).V().Count().ToList()
		assert.Nil(t, err)
		assert.NotNil(t, count)
		assert.Equal(t, 1, len(count))
		val, err := count[0].GetInt32()
		assert.Nil(t, err)
		assert.Equal(t, int32(6), val)
	})

	t.Run("Test read with RepeatUnrollStrategy", func(t *testing.T) {
		g := getModernGraph(t, testNoAuthUrl, &AuthInfo{}, &tls.Config{})
		defer g.remoteConnection.Close()

		count, err := g.WithStrategies(RepeatUnrollStrategy()).V().Count().ToList()
		assert.Nil(t, err)
		assert.NotNil(t, count)
		assert.Equal(t, 1, len(count))
		val, err := count[0].GetInt32()
		assert.Nil(t, err)
		assert.Equal(t, int32(6), val)
	})
}
