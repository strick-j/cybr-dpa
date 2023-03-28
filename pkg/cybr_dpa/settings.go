package cybr_dpa

import (
	"context"
	"fmt"

	"github.com/strick-j/cybr-dpa/pkg/cybr_dpa/types"
)

var (
	Settings types.Settings
)

// GetSettingsByFeature: API for DPA service settings and configuration, including MFA caching. To run this API, you must have the DpaAdmin role.
//
// Example Usage:
//
//	getSettingsByFeature, err := s.GetSettingsByFeature(context.Background, "MFA_CACHING")
func (s *Service) GetSettingsByFeature(ctx context.Context, featureName string) (*types.Settings, error) {
	if err := s.client.Get(ctx, fmt.Sprintf("/%s/%s", "settings", featureName), &Settings); err != nil {
		return nil, fmt.Errorf("failed to get Connector Install Script: %w", err)
	}

	return &Settings, nil
}

// PutSettingsByFeature: Overrides all settings for the specified feature. Unspecified settings are restored to their default values.
//
//	settingsUpdate := types.FeatureConf {
//		IsMfaCachingEnabled: true,
//		KeyExpirationTimeSec: 900,
//	}
//	putSettingsByFeature, err := s.PutSettingsByFeature(context.Background, "MFA_CACHING", settingsUpdate)
func (s *Service) PutSettingsByFeature(ctx context.Context, featureName string, featureConf types.FeatureConf) (*types.Settings, error) {
	if err := s.client.Put(ctx, fmt.Sprintf("/%s/%s", "settings", featureName), featureConf, &Settings); err != nil {
		return nil, fmt.Errorf("failed to get Connector Install Script: %w", err)
	}

	return &Settings, nil
}

// PatchSettingsByFeature: Overrides all settings for the specified feature. Unspecified settings are restored to their default values.
//
//	settingsUpdate := types.FeatureConf {
//		IsMfaCachingEnabled: true,
//		KeyExpirationTimeSec: 900,
//	}
//	patchSettingsByFeature, err := s.PatchSettingsByFeature(context.Background, "MFA_CACHING", settingsUpdate)
func (s *Service) PatchSettingsByFeature(ctx context.Context, featureName string, featureConf types.FeatureConf) (*types.Settings, error) {
	if err := s.client.Patch(ctx, fmt.Sprintf("/%s/%s", "settings", featureName), featureConf, &Settings); err != nil {
		return nil, fmt.Errorf("failed to get Connector Install Script: %w", err)
	}

	return &Settings, nil
}
