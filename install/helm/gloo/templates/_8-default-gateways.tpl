{{- define "httpGatewayTemplate" }}
{{- $contextValues := index . 0 }}
{{- $gatewayType := index . 1 }}
{{- $allGatewaySettings := index . 2 }}
{{- $tracingProvider := index . 3 }}
{{- $gateway := dict }}

{{- if hasKey $allGatewaySettings $gatewayType }}
  {{- $gatewaySettings := index $allGatewaySettings $gatewayType }}
  {{- $_ := set $gateway "httpGateway" $gatewaySettings }}
{{- else if $tracingProvider }}
  {{- $_ := set $gateway "httpGateway" (dict "options" (dict "httpConnectionManagerSettings" (dict "tracing" $tracingProvider))) }}
{{- else }}
  {{- $_ := set $gateway "httpGateway" (dict) }}
{{- end }}

{{- if and ($contextValues.global).nfType $contextValues.global.nfInstanceId }}
  {{- $_ := merge $gateway.httpGateway (dict "options" (dict "httpConnectionManagerSettings" (dict "serverName" (printf "%s-%s" $contextValues.global.nfType $contextValues.global.nfInstanceId)))) }}
{{- end }}

{{- if ($contextValues.httpConnectionManager).idleTimeout }}
  {{- $_ := merge $gateway.httpGateway (dict "options" (dict "httpConnectionManagerSettings" (dict "idleTimeout" $contextValues.httpConnectionManager.idleTimeout))) }}
{{- end }}

{{- if $contextValues.gatewayProxyExtensions }}
  {{- $_ := merge $gateway.httpGateway (dict "options" (dict "extensions" (dict "configs" $contextValues.gatewayProxyExtensions))) }}
{{- end }}

{{- toYaml $gateway | indent 2 }}
{{- end }}

{{- define "defaultGateway.gateway" -}}
{{- $name := (index . 1) }}
{{- $spec := (index . 2) }}
{{- with (first .) }}
{{- $gatewaySettings := $spec.gatewaySettings }}
{{- if $gatewaySettings.enabled }}
apiVersion: gateway.solo.io/v1
kind: Gateway
metadata:
  name: {{ $name | kebabcase }}
  namespace: {{ $spec.namespace | default .Release.Namespace }}
  labels:
{{ include "gloo.labels" . | indent 4}}
spec:
  {{- if $gatewaySettings.ipv4Only }}
  bindAddress: "0.0.0.0"
  {{- else }}
  bindAddress: "::"
  {{- end }}
  bindPort: {{ $spec.podTemplate.httpPort }}
{{- if $gatewaySettings.httpHybridGateway }}
{{ toYaml $gatewaySettings.httpHybridGateway | indent 2}}
{{- end }}
{{- "\n" }}
{{- template "httpGatewayTemplate" (list .Values "customHttpGateway" $gatewaySettings ($spec.tracing).provider) -}}
{{- if or ($gatewaySettings.options) ($gatewaySettings.accessLoggingService) }}
  options:
{{- end }}
  {{- if $gatewaySettings.options }}
  {{ toYaml $gatewaySettings.options | nindent 4 }}
  {{- end }}
  {{- if $gatewaySettings.accessLoggingService }}
    accessLoggingService:
  {{- toYaml $gatewaySettings.accessLoggingService | nindent 6 }}
  {{- end }}
  useProxyProto: {{ $gatewaySettings.useProxyProto }}
  ssl: false
  proxyNames:
  - {{ $name | kebabcase }}
{{- end }}{{/* $gatewaySettings.enabled */}}
{{- end }}{{/* with */}}
{{- end }}{{/* define "defaultGateway.gateway" */}}


{{- define "defaultGateway.sslGateway" -}}
{{- $name := (index . 1) }}
{{- $spec := (index . 2) }}
{{- with (first .) }}
{{- $gatewaySettings := $spec.gatewaySettings }}
{{- if $gatewaySettings.enabled }}
apiVersion: gateway.solo.io/v1
kind: Gateway
metadata:
  name: {{ $name | kebabcase }}-ssl
  namespace: {{ $spec.namespace | default .Release.Namespace }}
  labels:
{{ include "gloo.labels" . | indent 4}}
spec:
  {{- if $gatewaySettings.ipv4Only }}
  bindAddress: "0.0.0.0"
  {{- else }}
  bindAddress: "::"
  {{- end }}
  bindPort: {{ $spec.podTemplate.httpsPort }}
{{- if $gatewaySettings.httpsHybridGateway }}
{{ toYaml $gatewaySettings.httpsHybridGateway | indent 2}}
{{- end }}
{{- "\n" }}
{{- template "httpGatewayTemplate" (list .Values "customHttpsGateway" $gatewaySettings ($spec.tracing).provider) -}}
{{- if or ($gatewaySettings.options) ($gatewaySettings.accessLoggingService) }}
  options:
{{- end }}
  {{- if $gatewaySettings.options }}
  {{ toYaml $gatewaySettings.options | nindent 4 }}
  {{- end }}
  {{- if $gatewaySettings.accessLoggingService }}
    accessLoggingService:
  {{- toYaml $gatewaySettings.accessLoggingService | nindent 6 }}
  {{- end }}
  useProxyProto: {{ $gatewaySettings.useProxyProto }}
  ssl: true
  proxyNames:
  - {{ $name | kebabcase }}
{{- end }}{{/* $gatewaySettings.enabled */}}
{{- end }}{{/* with */}}
{{- end }}{{/* define "defaultGatway.sslGateway" */}}

{{- define "defaultGateway.failoverGateway" -}}
{{- $name := (index . 1) }}
{{- $spec := (index . 2) }}
{{- with (first .) }}
{{- $gatewaySettings := $spec.gatewaySettings }}
{{- if $gatewaySettings.enabled }}
apiVersion: gateway.solo.io/v1
kind: Gateway
metadata:
  name: {{ $name | kebabcase }}-failover
  namespace: {{ $spec.namespace | default .Release.Namespace }}
  labels:
{{ include "gloo.labels" . | indent 4}}
spec:
{{- if $gatewaySettings.ipv4Only }}
  bindAddress: "0.0.0.0"
{{- else }}
  bindAddress: "::"
{{- end }}
  bindPort: {{ $spec.failover.port }}
  tcpGateway:
    tcpHosts:
    - name: failover
      sslConfig:
        secretRef:
          name: {{ $spec.failover.secretName }}
          namespace: {{ .Release.Namespace }}
      destination:
        forwardSniClusterName: {}
  proxyNames:
  - {{ $name | kebabcase }}
{{- end }}{{/* $gatewaySettings.enabled */}}
{{- end }}{{/* with */}}
{{- end }}{{/* define "defaultGateway.failoverGateway" */}}

{{- define "gloo.customResources.defaultGateways" -}}
{{- $gatewayProxy := .Values.gatewayProxies.gatewayProxy }}
{{- range $name, $gatewaySpec := .Values.gatewayProxies }}
{{- $spec := include "gloo.util.mergeOverwriteWithOmit" (list $gatewaySpec $gatewayProxy) | fromJson }}
{{- $gatewaySettings := $spec.gatewaySettings }}
{{- if and $spec.gatewaySettings (not $gatewaySpec.disabled) }}
{{- $ctx := (list $ $name $spec) }}
{{- if not $gatewaySettings.disableGeneratedGateways }}
{{- if not $gatewaySettings.disableHttpGateway }}
{{- $defaultGatewayOverride := $spec.gatewaySettings.httpGatewayKubeOverride }}
---
{{ include "gloo.util.merge" (list $ctx $defaultGatewayOverride "defaultGateway.gateway") -}}
{{- end }}{{/* if not $gatewaySettings.disableHttpGateway */}}
{{- if not $gatewaySettings.disableHttpsGateway }}
{{- $sslGatewayOverride := $spec.gatewaySettings.httpsGatewayKubeOverride }}
---
{{ include "gloo.util.merge" (list $ctx $sslGatewayOverride "defaultGateway.sslGateway") -}}
{{- end }}{{/* if not $gatewaySettings.disableHttpsGateway  */}}
{{- end }}{{/* if not $gatewaySettings.disableGeneratedGateways */}}
{{- if $spec.failover }}
{{- if $spec.failover.enabled }}
{{- $failoverGatewayOverride := $spec.failover.kubeResourceOverride }}
---
{{ include "gloo.util.merge" (list $ctx $failoverGatewayOverride "defaultGateway.failoverGateway") -}}
{{- end }}{{/* if $spec.failover.enabled */}}
{{- end }}{{/* if $spec.failover */}}
{{- end }}{{/* if $spec.gatewaySettings and (not $spec.disabled) */}}
{{- end }}{{/* range gateways */}}
{{- end }}{{/* define "gloo.customResources.defaultGateways" */}}
