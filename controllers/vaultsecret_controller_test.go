//go:build integration
// +build integration

package controllers

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	redhatcopv1alpha1 "github.com/redhat-cop/vault-config-operator/api/v1alpha1"

	"k8s.io/apimachinery/pkg/types"
)

//TODO: Example: https://github.com/kubernetes-sigs/kubebuilder/blob/master/docs/book/src/cronjob-tutorial/testdata/project/controllers/cronjob_controller_test.go
// Define utility constants for object names and testing timeouts/durations and intervals.

var _ = Describe("VaultSecret controller", func() {
	timeout := time.Second * 10
	interval := time.Millisecond * 250
	Context("When creating VaultSecret", func() {
		It("Should be Successful when created", func() {
			By("By creating a new VaultSecret")
			ctx := context.Background()

			instance, err := decoder.GetVaultSecretInstance("../test/vaultsecret/vaultsecret-randomsecret.yaml")
			Expect(err).To(BeNil())
			instance.Namespace = vaultTestNamespaceName

			lookupKey := types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}
			createdVaultSecret := &redhatcopv1alpha1.VaultSecret{}

			// We'll need to retry getting this newly created VaultSecret, given that creation may not immediately happen.
			Eventually(func() bool {
				err := k8sClient.Get(ctx, lookupKey, createdVaultSecret)
				if err != nil {
					return false
				}
				if createdVaultSecret.Status.LastVaultSecretUpdate == nil {
					return false
				}
				return true
			}, timeout, interval).Should(BeTrue())

		})
	})
})

var files = []string{
	"",
	"../test/kv-engine-admin-policy.yaml",
	"../test/secret-writer-policy.yaml",
	"../test/kv-engine-admin-role.yaml",
	"../test/secret-writer-role.yaml",
	"../test/kv-secret-engine.yaml",
	"../test/random-secret.yaml",
}
