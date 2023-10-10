{{- define "chart.envFrom" -}}
envFrom:
- configMapRef:
    name: {{ template "chart.name" . }}-settings
{{- end -}}