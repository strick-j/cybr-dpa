# cybr-dpa  <!-- omit in toc -->
Contains functions for CyberArk DPA API.

## Table of Contents <!-- omit in toc -->
[Usage](#usage)
    - [Service](#service)
    - [Connectors](#connectors)
    - [Discovery](#discovery)
    - [Policies](#policies)
    - [Public Keys](#publickeys)
    - [Settings](#settings)
[Security](#security)


## Usage
```go
import (
    "github.com/strick-j/cybr-dpa"
    "github.com/strick-j/cybr-dpa/types"
)
```

All functions are documented with example usage in their respective go files. General flow for usage will be:
1. Obtain Oauth2 Bearer Token
2. Establish Service with Oauth2 Bearer Token
3. Utilize Service to interact Connectors, Discovery, Policy, Public Keys, and Settings functions

### Service

| Function | Input | Output |
|:--- |:--- |:--- |
| `NewService` | Identity URL (String), Identity API Endpoint (String), Verbose (Bool), Authentication Token [oauth2.token](https://pkg.go.dev/golang.org/x/oauth2#Token) | Service struct containing http.Client |

**Notes:**
1. Your Identity Security Platform Shared Services URL should be in the format TenantID.id.cyberark.cloud
2. The API Endpoint for Dynamic Privilege Access should be "api"

### Connectors
| Function | Input | Output |
|:--- |:--- |:--- |
| `GenerateScript` | Struct containing ConnectorOS and ConntectorType | GenerateScriptResponse Struct, Error Response Struct, or Error |

### Discovery
| Function | Input | Output |
|:--- |:--- |:--- |
| `ListTargetSets` | Ordered map of key value pairs for query | ListTargetSetResponse Struct, Error Response Struct, or Error |
| `AddTargetSets` | Struct containing required information | AddTargetSetResponse Struct, Error Response Struct, or Error |
| `DeleteTargetSets` | Slice containing strings | DeleteTargetSetResponse Struct, Error Response Struct, or Error |

**Notes:**
1. Example ordered map for List Target Sets: 
```go
query := map[string]string{"name":"example.com"}
```
### Policies
| Function | Input | Output |
|:--- |:--- |:--- |
| `ListPolicies` | nil | List Policies Struct, Error Response Struct, or Error |
| `GetPolicy` | String containing policy id | Policy Struct, Error Response Struct, or Error |
| `AddPolicy` | Struct containing new policy | AddPolicy Struct, Error Response Struct, or Error |
| `UpdatePolicy` | Struct containing policy settings, string containing policy id | Policy Struct, Error Response Struct, or Error |
| `DeletePolicy` | String containig policy id | Error Response Struct, or Error |

### Public Keys
| Function | Input | Output |
|:--- |:--- |:--- |
| `GetPublicKey` | Ordered map of key value pairs for query | PublicKey Struct, Error Response Struct, or Error |
| `GetPublicKeyScript` | Ordered map of key value pairs for query | PublicKeyScript Struct, Error Response Struct, or Error |

**Notes:**
1. Example ordered map for both functions:
```go
query := map[string]string{"workspaceId":"12347578363","workspaceType":"AWS"}
```

### Settings
| Function | Input | Output |
|:--- |:--- |:--- |
| `ListSettings` | nil | Settings Struct, Error Response Struct, or Error |
| `ListSettingsFeature` | String containing desired Setting | Feature Setting Struct, Error Response Struct, or Error |
| `UpdateSettingsSets` | Struct containing Settings to Update | DeleteTargetSetResponse Struct, Error Response Struct, or Error |

**Notes:**
1. Valid feature names for ListSettingsFeature are: 'MFA_CACHING', 'STANDING_ACCESS', 'SSH_COMMAND_AUDIT', 'RDP_FILE_TRANSFER', 'CERTIFICATE_VALIDATION'

## Secrurity
If there is a security concern or bug discovered, please responsibly disclose all information to joe (dot) strickland (at) cyberark (dot) com.