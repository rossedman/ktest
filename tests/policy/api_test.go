package core_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var _ = Describe("client", func() {
	Context("twilio-system namespace exists", func() {
		It("returns a 200 OK response", func() {
			// creates the in-cluster config
			config, _ := rest.InClusterConfig()
			// creates the clientset
			clientset, _ := kubernetes.NewForConfig(config)
			_, err := clientset.CoreV1().Namespaces().Get(context.TODO(), "twilio-system-policy", metav1.GetOptions{})
			Expect(err).To(BeNil())
		})
	})
})
