package deployer

import (
	"encoding/json"
	"fmt"
	"maps"
	"regexp"
	"slices"
	"sort"
	"strings"

	"github.com/rotisserie/eris"
	"github.com/solo-io/gloo/projects/gateway2/api/v1alpha1"
	"github.com/solo-io/gloo/projects/gateway2/ports"
	"github.com/solo-io/gloo/projects/gateway2/translator/types"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/utils/ptr"
)

// This file contains helper functions that generate helm values in the format needed
// by the deployer.

var ComponentLogLevelEmptyError = func(key string, value string) error {
	return eris.Errorf("an empty key or value was provided in componentLogLevels: key=%s, value=%s", key, value)
}

// Extract the listener ports from a Gateway and corresponding listener sets. These will be used to populate:
// 1. the ports exposed on the envoy container
// 2. the ports exposed on the proxy service
func getPortsValues(cgw *types.ConsolidatedGateway, gwp *v1alpha1.GatewayParameters) []helmPort {
	gwPorts := map[uint16]*helmPort{}
	for i, cl := range cgw.GetConsolidatedListeners() {
		listener := cl.Listener
		portName := string(listener.Name)
		if cl.ListenerSet != nil {
			// This ensures a unique name per port even if a GW and LS have listeners with the same name
			portName = fmt.Sprintf("%d-%s", i, listener.Name)
		}
		appendPortValue(gwPorts, uint16(listener.Port), portName, gwp)
	}
	finalPorts := make([]helmPort, len(gwPorts))
	for i, port := range slices.Sorted(maps.Keys(gwPorts)) {
		finalPorts[i] = *gwPorts[port]
	}
	return finalPorts
}

func sanitizePortName(name string) string {
	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z0-9-]+`)
	str := nonAlphanumericRegex.ReplaceAllString(name, "-")
	doubleHyphen := regexp.MustCompile(`-{2,}`)
	str = doubleHyphen.ReplaceAllString(str, "-")

	// This is a kubernetes spec requirement.
	maxPortNameLength := 15
	if len(str) > maxPortNameLength {
		str = str[:maxPortNameLength]
	}
	return str
}

func appendPortValue(gwPorts map[uint16]*helmPort, port uint16, name string, gwp *v1alpha1.GatewayParameters) {
	// only process this port if we haven't already processed a listener with the same port
	if _, ok := gwPorts[port]; ok {
		return
	}

	targetPort := ports.TranslatePort(port)
	portName := sanitizePortName(name)
	protocol := "TCP"

	// Search for static NodePort set from the GatewayParameters spec
	// If not found the default value of `nil` will not render anything.
	var nodePort *uint16 = nil
	if gwp.Spec.GetKube().GetService().GetType() != nil && *(gwp.Spec.GetKube().GetService().GetType()) == corev1.ServiceTypeNodePort {
		if idx := slices.IndexFunc(gwp.Spec.GetKube().GetService().GetPorts(), func(p *v1alpha1.Port) bool {
			return p.GetPort() == uint16(port)
		}); idx != -1 {
			nodePort = ptr.To(uint16(*gwp.Spec.GetKube().GetService().GetPorts()[idx].GetNodePort()))
		}
	}

	gwPorts[port] = &helmPort{
		Port:       &port,
		TargetPort: &targetPort,
		Name:       &portName,
		Protocol:   &protocol,
		NodePort:   nodePort,
	}
}

// TODO: Removing until autoscaling is re-added.
// See: https://github.com/solo-io/solo-projects/issues/5948
// Convert autoscaling values from GatewayParameters into helm values to be used by the deployer.
// func getAutoscalingValues(autoscaling *v1.Autoscaling) *helmAutoscaling {
// 	hpaConfig := autoscaling.HorizontalPodAutoscaler
// 	if hpaConfig == nil {
// 		return nil
// 	}

// 	trueVal := true
// 	autoscalingVals := &helmAutoscaling{
// 		Enabled: &trueVal,
// 	}
// 	autoscalingVals.MinReplicas = hpaConfig.MinReplicas
// 	autoscalingVals.MaxReplicas = hpaConfig.MaxReplicas
// 	autoscalingVals.TargetCPUUtilizationPercentage = hpaConfig.TargetCpuUtilizationPercentage
// 	autoscalingVals.TargetMemoryUtilizationPercentage = hpaConfig.TargetMemoryUtilizationPercentage

// 	return autoscalingVals
// }

// Convert service values from GatewayParameters into helm values to be used by the deployer.
func getServiceValues(svcConfig *v1alpha1.Service) *helmService {
	// convert the service type enum to its string representation;
	// if type is not set, it will default to 0 ("ClusterIP")
	var svcType *string
	if svcConfig.GetType() != nil {
		svcType = ptr.To(string(*svcConfig.GetType()))
	}
	var svcExternalTrafficPolicy *string
	if svcConfig.GetExternalTrafficPolicy() != nil {
		svcExternalTrafficPolicy = ptr.To(string(*svcConfig.GetExternalTrafficPolicy()))
	}
	return &helmService{
		Type:                  svcType,
		ClusterIP:             svcConfig.GetClusterIP(),
		ExtraAnnotations:      svcConfig.GetExtraAnnotations(),
		ExtraLabels:           svcConfig.GetExtraLabels(),
		ExternalTrafficPolicy: svcExternalTrafficPolicy,
	}
}

// Convert service account values from GatewayParameters into helm values to be used by the deployer.
func getServiceAccountValues(svcAccountConfig *v1alpha1.ServiceAccount) *helmServiceAccount {
	return &helmServiceAccount{
		ExtraAnnotations: svcAccountConfig.GetExtraAnnotations(),
		ExtraLabels:      svcAccountConfig.GetExtraLabels(),
	}
}

// Convert sds values from GatewayParameters into helm values to be used by the deployer.
func getSdsContainerValues(sdsContainerConfig *v1alpha1.SdsContainer) *helmSdsContainer {
	if sdsContainerConfig == nil {
		return nil
	}

	vals := &helmSdsContainer{
		Image:           getImageValues(sdsContainerConfig.GetImage()),
		Resources:       sdsContainerConfig.GetResources(),
		SecurityContext: sdsContainerConfig.GetSecurityContext(),
		SdsBootstrap:    &sdsBootstrap{},
	}

	if bootstrap := sdsContainerConfig.GetBootstrap(); bootstrap != nil {
		vals.SdsBootstrap = &sdsBootstrap{
			LogLevel: bootstrap.GetLogLevel(),
		}
	}

	return vals
}

func getIstioContainerValues(config *v1alpha1.IstioContainer) *helmIstioContainer {
	if config == nil {
		return nil
	}

	return &helmIstioContainer{
		Image:                 getImageValues(config.GetImage()),
		LogLevel:              config.GetLogLevel(),
		Resources:             config.GetResources(),
		SecurityContext:       config.GetSecurityContext(),
		IstioDiscoveryAddress: config.GetIstioDiscoveryAddress(),
		IstioMetaMeshId:       config.GetIstioMetaMeshId(),
		IstioMetaClusterId:    config.GetIstioMetaClusterId(),
	}
}

// Convert istio values from GatewayParameters into helm values to be used by the deployer.
func getIstioValues(istioIntegrationEnabled bool, istioConfig *v1alpha1.IstioIntegration) *helmIstio {
	// if istioConfig is nil, istio sds is disabled and values can be ignored
	if istioConfig == nil {
		return &helmIstio{
			Enabled: ptr.To(istioIntegrationEnabled),
		}
	}

	return &helmIstio{
		Enabled: ptr.To(istioIntegrationEnabled),
	}
}

// Converts AwsInfo (which come from Settings values) into aws helm values
func getAwsValues(awsInfo *AwsInfo) *helmAws {
	if awsInfo != nil {
		return &helmAws{
			EnableServiceAccountCredentials: &awsInfo.EnableServiceAccountCredentials,
			StsClusterName:                  &awsInfo.StsClusterName,
			StsUri:                          &awsInfo.StsUri,
		}
	}
	return nil
}

// Get the image values for the envoy container in the proxy deployment.
func getImageValues(image *v1alpha1.Image) *helmImage {
	if image == nil {
		return &helmImage{}
	}

	helmImage := &helmImage{
		Registry:   image.GetRegistry(),
		Repository: image.GetRepository(),
		Tag:        image.GetTag(),
		Digest:     image.GetDigest(),
	}
	if image.GetPullPolicy() != nil {
		helmImage.PullPolicy = ptr.To(string(*image.GetPullPolicy()))
	}

	return helmImage
}

// Get the stats values for the envoy listener in the configmap for bootstrap.
func getStatsValues(statsConfig *v1alpha1.StatsConfig) *helmStatsConfig {
	if statsConfig == nil {
		return nil
	}
	return &helmStatsConfig{
		Enabled:            statsConfig.GetEnabled(),
		RoutePrefixRewrite: statsConfig.GetRoutePrefixRewrite(),
		EnableStatsRoute:   statsConfig.GetEnableStatsRoute(),
		StatsPrefixRewrite: statsConfig.GetStatsRoutePrefixRewrite(),
	}
}

// ComponentLogLevelsToString converts the key-value pairs in the map into a string of the
// format: key1:value1,key2:value2,key3:value3, where the keys are sorted alphabetically.
// If an empty map is passed in, then an empty string is returned.
// Map keys and values may not be empty.
// No other validation is currently done on the keys/values.
func ComponentLogLevelsToString(vals map[string]string) (string, error) {
	if len(vals) == 0 {
		return "", nil
	}

	parts := make([]string, 0, len(vals))
	for k, v := range vals {
		if k == "" || v == "" {
			return "", ComponentLogLevelEmptyError(k, v)
		}
		parts = append(parts, fmt.Sprintf("%s:%s", k, v))
	}
	sort.Strings(parts)
	return strings.Join(parts, ","), nil
}

func getAIExtensionValues(config *v1alpha1.AiExtension) (*helmAIExtension, error) {
	if config == nil {
		return nil, nil
	}

	// If we don't do this check, a byte array containing the characters "null" will be rendered
	// This will not be marshallable by the component so instead we render nothing.
	var statsByt []byte
	if config.GetStats() != nil {
		var err error
		statsByt, err = json.Marshal(config.GetStats())
		if err != nil {
			return nil, err
		}
	}
	var tracingByt []byte
	if config.GetTracing() != nil {
		var err error
		tracingByt, err = json.Marshal(config.GetTracing())
		if err != nil {
			return nil, err
		}
	}

	return &helmAIExtension{
		Enabled:         *config.GetEnabled(),
		Image:           getImageValues(config.GetImage()),
		SecurityContext: config.GetSecurityContext(),
		Resources:       config.GetResources(),
		Env:             config.GetEnv(),
		Ports:           config.GetPorts(),
		Stats:           statsByt,
		Tracing:         tracingByt,
	}, nil
}
