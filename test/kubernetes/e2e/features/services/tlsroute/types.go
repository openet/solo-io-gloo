package tlsroute

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	testmatchers "github.com/solo-io/gloo/test/gomega/matchers"

	"github.com/solo-io/skv2/codegen/util"

	"github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// Constants used by TestConfigureTLSRouteBackingDestinationsWithSingleService
	singleSvcNsName          = "single-tls-route"
	singleSvcGatewayName     = "single-tls-gateway"
	singleSvcListenerName443 = "listener-443"
	singleSvcName            = "single-svc"
	singleSvcTLSRouteName    = "single-tls-route"

	// Constants used by TestConfigureTLSRouteBackingDestinationsWithMultiServices
	multiSvcNsName           = "multi-tls-route"
	multiSvcGatewayName      = "multi-tls-gateway"
	multiSvcListenerName6443 = "listener-6443"
	multiSvcListenerName8443 = "listener-8443"
	multiSvc1Name            = "multi-svc-1"
	multiSvc2Name            = "multi-svc-2"
	multiSvcTLSRouteName1    = "tls-route-1"
	multiSvcTLSRouteName2    = "tls-route-2"

	// Constants for CrossNamespaceTLSRouteWithReferenceGrant
	crossNsTestName           = "CrossNamespaceTLSRouteWithReferenceGrant"
	crossNsClientName         = "cross-namespace-allowed-client-ns"
	crossNsBackendNsName      = "cross-namespace-allowed-backend-ns"
	crossNsGatewayName        = "gateway"
	crossNsListenerName       = "listener-8443"
	crossNsBackendSvcName     = "backend-svc"
	crossNsTLSRouteName       = "tls-route"
	crossNsReferenceGrantName = "reference-grant"

	// Constants for CrossNamespaceTLSRouteWithoutReferenceGrant
	crossNsNoRefGrantTestName       = "CrossNamespaceTLSRouteWithoutReferenceGrant"
	crossNsNoRefGrantClientNsName   = "client-ns-no-refgrant"
	crossNsNoRefGrantBackendNsName  = "backend-ns-no-refgrant"
	crossNsNoRefGrantGatewayName    = "gateway"
	crossNsNoRefGrantListenerName   = "listener-8443"
	crossNsNoRefGrantBackendSvcName = "backend-svc"
	crossNsNoRefGrantTLSRouteName   = "tls-route"
)

var (
	// Variables used by TestConfigureTCPRouteBackingDestinationsWithSingleService
	multiSvcNsManifest               = filepath.Join(util.MustGetThisDir(), "testdata", "multi-ns.yaml")
	multiSvcGatewayAndClientManifest = filepath.Join(util.MustGetThisDir(), "testdata", "multi-listener-gateway-and-client.yaml")
	multiSvcBackendManifest          = filepath.Join(util.MustGetThisDir(), "testdata", "multi-backend-service.yaml")
	multiSvcTlsRouteManifest         = filepath.Join(util.MustGetThisDir(), "testdata", "multi-tlsroute.yaml")

	// Variables used by TestConfigureTCPRouteBackingDestinationsWithMultiServices
	singleSvcNsManifest               = filepath.Join(util.MustGetThisDir(), "testdata", "single-ns.yaml")
	singleSvcGatewayAndClientManifest = filepath.Join(util.MustGetThisDir(), "testdata", "single-listener-gateway-and-client.yaml")
	singleSvcBackendManifest          = filepath.Join(util.MustGetThisDir(), "testdata", "single-backend-service.yaml")
	singleSvcTLSRouteManifest         = filepath.Join(util.MustGetThisDir(), "testdata", "single-tlsroute.yaml")
	singleSecretManifest              = filepath.Join(util.MustGetThisDir(), "testdata", "tls-secret.yaml")

	// Manifests for CrossNamespaceTLSRouteWithReferenceGrant
	crossNsClientNsManifest   = filepath.Join(util.MustGetThisDir(), "testdata", "cross-ns-client-ns.yaml")
	crossNsBackendNsManifest  = filepath.Join(util.MustGetThisDir(), "testdata", "cross-ns-backend-ns.yaml")
	crossNsGatewayManifest    = filepath.Join(util.MustGetThisDir(), "testdata", "cross-ns-gateway-and-client.yaml")
	crossNsBackendSvcManifest = filepath.Join(util.MustGetThisDir(), "testdata", "cross-ns-backend-service.yaml")
	crossNsTLSRouteManifest   = filepath.Join(util.MustGetThisDir(), "testdata", "cross-ns-tlsroute.yaml")
	crossNsRefGrantManifest   = filepath.Join(util.MustGetThisDir(), "testdata", "cross-ns-referencegrant.yaml")

	// Manifests for CrossNamespaceTCPRouteWithoutReferenceGrant
	crossNsNoRefGrantClientNsManifest   = filepath.Join(util.MustGetThisDir(), "testdata", "cross-ns-no-refgrant-client-ns.yaml")
	crossNsNoRefGrantBackendNsManifest  = filepath.Join(util.MustGetThisDir(), "testdata", "cross-ns-no-refgrant-backend-ns.yaml")
	crossNsNoRefGrantGatewayManifest    = filepath.Join(util.MustGetThisDir(), "testdata", "cross-ns-no-refgrant-gateway-and-client.yaml")
	crossNsNoRefGrantBackendSvcManifest = filepath.Join(util.MustGetThisDir(), "testdata", "cross-ns-no-refgrant-backend-service.yaml")
	crossNsNoRefGrantTLSRouteManifest   = filepath.Join(util.MustGetThisDir(), "testdata", "cross-ns-no-refgrant-tlsroute.yaml")

	// Assertion test timers
	ctxTimeout = 5 * time.Minute
	timeout    = 60 * time.Second

	// Proxy resources to be translated
	singleSvcNS = &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: singleSvcNsName,
		},
	}

	singleGlooProxy = metav1.ObjectMeta{
		Name:      "gloo-proxy-single-tls-gateway",
		Namespace: singleSvcNsName,
	}
	singleSvcProxyDeployment = &appsv1.Deployment{ObjectMeta: singleGlooProxy}
	singleSvcProxyService    = &corev1.Service{ObjectMeta: singleGlooProxy}

	multiSvcNS = &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: multiSvcNsName,
		},
	}

	multiGlooProxy = metav1.ObjectMeta{
		Name:      "gloo-proxy-multi-tls-gateway",
		Namespace: multiSvcNsName,
	}
	multiProxyDeployment = &appsv1.Deployment{ObjectMeta: multiGlooProxy}
	multiProxyService    = &corev1.Service{ObjectMeta: multiGlooProxy}

	// Expected curl responses from tests
	expectedSingleSvcResp = &testmatchers.HttpResponse{
		StatusCode: http.StatusOK,
		Body: gomega.SatisfyAll(
			gomega.MatchRegexp(fmt.Sprintf(`"namespace"\s*:\s*"%s"`, singleSvcNsName)),
			gomega.MatchRegexp(`"service"\s*:\s*"single-svc"`),
		),
	}

	crossNsGlooProxy = metav1.ObjectMeta{
		Name:      "gloo-proxy-gateway",
		Namespace: crossNsClientName,
	}
	crossNsProxyDeployment = &appsv1.Deployment{ObjectMeta: crossNsGlooProxy}
	crossNsProxyService    = &corev1.Service{ObjectMeta: crossNsGlooProxy}

	crossNsNoRefGrantGlooProxy = metav1.ObjectMeta{
		Name:      "gloo-proxy-gateway",
		Namespace: crossNsNoRefGrantClientNsName,
	}
	crossNsNoRefGrantProxyDeployment = &appsv1.Deployment{ObjectMeta: crossNsNoRefGrantGlooProxy}
	crossNsNoRefGrantProxyService    = &corev1.Service{ObjectMeta: crossNsNoRefGrantGlooProxy}

	expectedMultiSvc1Resp = &testmatchers.HttpResponse{
		StatusCode: http.StatusOK,
		Body: gomega.SatisfyAll(
			gomega.MatchRegexp(fmt.Sprintf(`"namespace"\s*:\s*"%s"`, multiSvcNsName)),
			gomega.MatchRegexp(fmt.Sprintf(`"service"\s*:\s*"%s"`, multiSvc1Name)),
		),
	}

	expectedMultiSvc2Resp = &testmatchers.HttpResponse{
		StatusCode: http.StatusOK,
		Body: gomega.SatisfyAll(
			gomega.MatchRegexp(fmt.Sprintf(`"namespace"\s*:\s*"%s"`, multiSvcNsName)),
			gomega.MatchRegexp(fmt.Sprintf(`"service"\s*:\s*"%s"`, multiSvc2Name)),
		),
	}

	expectedCrossNsResp = &testmatchers.HttpResponse{
		StatusCode: http.StatusOK,
		Body: gomega.SatisfyAll(
			gomega.MatchRegexp(fmt.Sprintf(`"namespace"\s*:\s*"%s"`, crossNsBackendNsName)),
			gomega.MatchRegexp(fmt.Sprintf(`"service"\s*:\s*"%s"`, crossNsBackendSvcName)),
		),
	}
)
