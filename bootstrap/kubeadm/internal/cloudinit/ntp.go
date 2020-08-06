/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cloudinit

const (
	ntpTemplate = `{{ define "ntp" -}}
{{- if . }}
ntp:
  {{ if .Enabled -}}
  enabled: true
  {{ end -}}
  {{ if .NTPClient -}}
  ntp_client: {{ .NTPClient }}
  {{ end -}}
  servers:{{ range .Servers }}
    - {{ . }}
  {{- end }}
  pools:{{ range .Pools }}
    - {{ . }}
  {{- end -}}
  {{- with .NTPConfig }}
  config:
    {{- if .ConfPath }}
    confpath: {{ .ConfPath }}
    {{- end -}}
    {{- if .CheckEXE }}
    check_exe: {{ .CheckEXE }}
    {{- end }}
    packages:{{ range .Packages }}
      - {{ . }}
    {{- end }}
    {{- if .ServiceName }}
    service_name: {{ .ServiceName }}
    {{- end }}
    {{- if .Template }}
    template: |
{{ Indent 8 .Template }}
    {{- end -}}
  {{- end -}}
{{- end -}}
{{- end -}}
`
)
