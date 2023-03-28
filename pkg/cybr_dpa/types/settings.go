package types

type Settings struct {
	FeatureName string      `json:"feature_name,omitempty"`
	FeatureConf FeatureConf `json:"feature_conf,omitempty"`
	Code        string      `json:"code,omitempty"`
	Message     string      `json:"message,omitempty"`
	Description string      `json:"description,omitempty"`
}

type FeatureConf struct {
	IsMfaCachingEnabled  bool `json:"isMfaCachingEnabled,omitempty"`
	KeyExpirationTimeSec int  `json:"keyExpirationTimeSec,omitempty"`
}
