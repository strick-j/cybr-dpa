package types

type Settings struct {
	MfaCaching            MfaCaching            `json:"mfaCaching,omitempty"`
	SSHCommandAudit       SSHCommandAudit       `json:"sshCommandAudit,omitempty"`
	StandingAccess        StandingAccess        `json:"standingAccess,omitempty"`
	RdpFileTransfer       RdpFileTransfer       `json:"rdpFileTransfer,omitempty"`
	CertificateValidation CertificateValidation `json:"certificateValidation,omitempty"`
}

type MfaCaching struct {
	IsMfaCachingEnabled  bool `json:"isMfaCachingEnabled,omitempty"`
	KeyExpirationTimeSec int  `json:"keyExpirationTimeSec,omitempty"`
}

type SSHCommandAudit struct {
	IsCommandParsingForAuditEnabled bool   `json:"isCommandParsingForAuditEnabled,omitempty"`
	ShellPromptForAudit             string `json:"shellPromptForAudit,omitempty"`
}

type StandingAccess struct {
	StandingAccessAvailable bool `json:"standingAccessAvailable,omitempty"`
	SessionMaxDuration      int  `json:"sessionMaxDuration,omitempty"`
	SessionIdleTime         int  `json:"sessionIdleTime,omitempty"`
}

type RdpFileTransfer struct {
	Enabled bool `json:"enabled,omitempty"`
}

type CertificateValidation struct {
	Enabled bool `json:"enabled,omitempty"`
}

type FeatureSetting struct {
	FeatureName string      `json:"featureName,omitempty"`
	FeatureConf FeatureConf `json:"featureConf,omitempty"`
}

type FeatureConf struct {
	IsMfaCachingEnabled             bool   `json:"isMfaCachingEnabled,omitempty"`
	KeyExpirationTimeSec            int    `json:"keyExpirationTimeSec,omitempty"`
	IsCommandParsingForAuditEnabled bool   `json:"isCommandParsingForAuditEnabled,omitempty"`
	ShellPromptForAudit             string `json:"shellPromptForAudit,omitempty"`
}
