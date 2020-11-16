package wireguard

var (
	ClientCfgTpl = `#{{ .Client.Identifier }}
[Interface]
Address = {{ .Client.IPsStr }}
PrivateKey = {{ .Client.PrivateKey }}
{{if .Server.DNSStr -}}
DNS = {{ .Server.DNSStr }}
{{- end}}
{{- if ne .Server.Mtu 0 -}}
MTU = {{.Server.Mtu}}
{{- end}}

[Peer]
PublicKey = {{ .Server.PublicKey }}
{{- if .Client.PresharedKey -}}
PresharedKey = {{ .Client.PresharedKey }}
{{- end -}}
AllowedIPs = {{ .Client.AllowedIPsStr }}
Endpoint = {{ .Server.Endpoint }}
{{if and (ne .Server.PersistentKeepalive 0) (not .Client.IgnorePersistentKeepalive) -}}
PersistentKeepalive = {{.Server.PersistentKeepalive}}
{{- end}}
`
	DeviceCfgTpl = `# AUTOGENERATED FILE - DO NOT EDIT
# Updated: {{ .Server.UpdatedAt }} / Created: {{ .Server.CreatedAt }}
[Interface]
{{- range .Server.IPs }}
Address = {{ . }}
{{- end}}
ListenPort = {{ .Server.ListenPort }}
PrivateKey = {{ .Server.PrivateKey }}
{{- if ne .Server.Mtu 0 -}}
MTU = {{.Server.Mtu}}
{{- end -}}
PreUp = {{ .Server.PreUp }}
PostUp = {{ .Server.PostUp }}
PreDown = {{ .Server.PreDown }}
PostDown = {{ .Server.PostDown }}

{{range .Clients}}
{{if not .DeactivatedAt -}}
# {{.Identifier}} / {{.Email}} / Updated: {{.UpdatedAt}} / Created: {{.CreatedAt}}
[Peer]
PublicKey = {{ .PublicKey }}
{{- if .PresharedKey -}}
PresharedKey = {{ .PresharedKey }}
{{- end -}}
AllowedIPs = {{ StringsJoin .IPs ", " }}
{{- end}}
{{end}}`
)
