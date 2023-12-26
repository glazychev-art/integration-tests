// Code generated by gotestmd DO NOT EDIT.
package features

import (
	"github.com/stretchr/testify/suite"

	"github.com/networkservicemesh/integration-tests/extensions/base"
	"github.com/networkservicemesh/integration-tests/suites/basic"
)

type Suite struct {
	base.Suite
	basicSuite basic.Suite
}

func (s *Suite) SetupSuite() {
	parents := []interface{}{&s.Suite, &s.basicSuite}
	for _, p := range parents {
		if v, ok := p.(suite.TestingSuite); ok {
			v.SetT(s.T())
		}
		if v, ok := p.(suite.SetupAllSuite); ok {
			v.SetupSuite()
		}
	}
}
func (s *Suite) TestScaled_registry() {
	r := s.Runner("../deployments-k8s/examples/features/scaled-registry")
	s.T().Cleanup(func() {
		r.Run(`kubectl get pods -A -o wide`)
		r.Run(`kubectl describe pod -n nsm-system`)
		r.Run(`kubectl delete ns ns-scaled-registry`)
		r.Run(`kubectl scale --replicas=1 deployments/registry-k8s -n nsm-system`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/features/scaled-registry?ref=790b99a8932a75bfd40ee0b4e58ce4c8779d2e83`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-scaled-registry`)
	r.Run(`NSE=$(kubectl get pod -n ns-scaled-registry --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}' -l app=nse-kernel)`)
	r.Run(`kubectl get nses -A | grep $NSE`)
	r.Run(`kubectl scale --replicas=0 deployments/registry-k8s -n nsm-system`)
	r.Run(`kubectl wait --for=delete --timeout=1m pod -l app=registry -n nsm-system`)
	r.Run(`kubectl get nses -A | grep $NSE`)
	r.Run(`kubectl scale --replicas=2 deployments/registry-k8s -n nsm-system`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=registry -n nsm-system`)
	r.Run(`kubectl scale --replicas=0 deployments/nse-kernel -n ns-scaled-registry`)
	r.Run(`kubectl get nses -A | grep $NSE` + "\n" + `if [[ "$?" == "1" ]]; then echo OK; else echo "nse entry still exists"; false; fi`)
}
