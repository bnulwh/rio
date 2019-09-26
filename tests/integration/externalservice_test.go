package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rancher/rio/tests/testutil"

	"github.com/sclevine/spec"
)

func externalServiceTests(t *testing.T, when spec.G, it spec.S) {

	var externalService testutil.TestExternalService

	it.After(func() {
		externalService.Remove()
	})

	when("a config is created with an ip address", func() {
		it("should have its ip address field populated", func() {
			externalService.Create(t, "1.2.3.4")
			assert.Equal(t, "1.2.3.4", externalService.GetFirstIPAddress())
		})
	})
	when("a config is created with a FQDN", func() {
		it("should have its FQDN field populated", func() {
			externalService.Create(t, "test.example.com")
			assert.Equal(t, "test.example.com", externalService.GetFQDN())
		})
	})
}
