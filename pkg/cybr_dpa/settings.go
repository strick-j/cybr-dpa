package cybr_dpa

import (
	"context"
	"fmt"
)

type Settings struct {
	FeatureName string      `json:"feature_name,omitempty"`
	FeatureConf FeatureConf `json:"feature_conf,omitempty"`
	Code        string      `json:"code,omitempty"`
	Message     string      `json:"message,omitempty"`
	Description string      `json:"description,omitempty"`
}

type FeatureConf struct {
	IsMfaCachingEnabled  bool `json:"is_mfa_caching_enabled,omitempty"`
	KeyExpirationTimeSec int  `json:"key_expiration_time_sec,omitempty"`
}

var (
	settings Settings
)

// GetSettingsByFeature: API for DPA service settings and configuration, including MFA caching. To run this API, you must have the DpaAdmin role.
//
// Example Usage:
//
//	getSettingsByFeature, err := s.GetSettingsByFeature(context.Background, "MFA_CACHING")
func (s *Service) GetSettingsByFeature(ctx context.Context, featureName string) (*Settings, error) {
	if err := s.client.Get(ctx, fmt.Sprintf("/%s/%s", "settings", featureName), &settings); err != nil {
		return nil, fmt.Errorf("failed to get Connector Install Script: %w", err)
	}

	return &settings, nil
}

// PutSettingsByFeature: Overrides all settings for the specified feature. Unspecified settings are restored to their default values.
//
//	settingsUpdate := types.FeatureConf {
//	  IsMfaCachingEnabled: true,
//	  KeyExpirationTimeSec: 900,
//	}
//
//	putSettingsByFeature, err := s.PutSettingsByFeature(context.Background, "MFA_CACHING", settingsUpdate)
func (s *Service) PutSettingsByFeature(ctx context.Context, featureName string, featureConf FeatureConf) (*Settings, error) {
	if err := s.client.Put(ctx, fmt.Sprintf("/%s/%s", "settings", featureName), featureConf, &settings); err != nil {
		return nil, fmt.Errorf("failed to get Connector Install Script: %w", err)
	}

	return &settings, nil
}

// PatchSettingsByFeature: Overrides all settings for the specified feature. Unspecified settings are restored to their default values.
//
//	settingsUpdate := types.FeatureConf {
//	  IsMfaCachingEnabled: true,
//	  KeyExpirationTimeSec: 900,
//	}
//
//	patchSettingsByFeature, err := s.PatchSettingsByFeature(context.Background, "MFA_CACHING", settingsUpdate)
func (s *Service) PatchSettingsByFeature(ctx context.Context, featureName string, featureConf FeatureConf) (*Settings, error) {
	if err := s.client.Patch(ctx, fmt.Sprintf("/%s/%s", "settings", featureName), featureConf, &settings); err != nil {
		return nil, fmt.Errorf("failed to get Connector Install Script: %w", err)
	}

	return &settings, nil
}
