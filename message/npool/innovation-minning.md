# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [npool/innovation-minning.proto](#npool/innovation-minning.proto)
    - [Capital](#innovation.minning.v1.Capital)
    - [CreateCapitalRequest](#innovation.minning.v1.CreateCapitalRequest)
    - [CreateCapitalResponse](#innovation.minning.v1.CreateCapitalResponse)
    - [CreateLaunchTimeRequest](#innovation.minning.v1.CreateLaunchTimeRequest)
    - [CreateLaunchTimeResponse](#innovation.minning.v1.CreateLaunchTimeResponse)
    - [CreateProjectRequest](#innovation.minning.v1.CreateProjectRequest)
    - [CreateProjectResponse](#innovation.minning.v1.CreateProjectResponse)
    - [CreateTeamRequest](#innovation.minning.v1.CreateTeamRequest)
    - [CreateTeamResponse](#innovation.minning.v1.CreateTeamResponse)
    - [CreateTechniqueAnalysisRequest](#innovation.minning.v1.CreateTechniqueAnalysisRequest)
    - [CreateTechniqueAnalysisResponse](#innovation.minning.v1.CreateTechniqueAnalysisResponse)
    - [CreateTrendAnalysisRequest](#innovation.minning.v1.CreateTrendAnalysisRequest)
    - [CreateTrendAnalysisResponse](#innovation.minning.v1.CreateTrendAnalysisResponse)
    - [CreateUserRequest](#innovation.minning.v1.CreateUserRequest)
    - [CreateUserResponse](#innovation.minning.v1.CreateUserResponse)
    - [LaunchTime](#innovation.minning.v1.LaunchTime)
    - [Project](#innovation.minning.v1.Project)
    - [Team](#innovation.minning.v1.Team)
    - [TechniqueAnalysis](#innovation.minning.v1.TechniqueAnalysis)
    - [TrendAnalysis](#innovation.minning.v1.TrendAnalysis)
    - [UpdateCapitalRequest](#innovation.minning.v1.UpdateCapitalRequest)
    - [UpdateCapitalResponse](#innovation.minning.v1.UpdateCapitalResponse)
    - [UpdateLaunchTimeRequest](#innovation.minning.v1.UpdateLaunchTimeRequest)
    - [UpdateLaunchTimeResponse](#innovation.minning.v1.UpdateLaunchTimeResponse)
    - [UpdateProjectRequest](#innovation.minning.v1.UpdateProjectRequest)
    - [UpdateProjectResponse](#innovation.minning.v1.UpdateProjectResponse)
    - [UpdateTeamRequest](#innovation.minning.v1.UpdateTeamRequest)
    - [UpdateTeamResponse](#innovation.minning.v1.UpdateTeamResponse)
    - [UpdateTechniqueAnalysisRequest](#innovation.minning.v1.UpdateTechniqueAnalysisRequest)
    - [UpdateTechniqueAnalysisResponse](#innovation.minning.v1.UpdateTechniqueAnalysisResponse)
    - [UpdateTrendAnalysisRequest](#innovation.minning.v1.UpdateTrendAnalysisRequest)
    - [UpdateTrendAnalysisResponse](#innovation.minning.v1.UpdateTrendAnalysisResponse)
    - [UpdateUserRequest](#innovation.minning.v1.UpdateUserRequest)
    - [UpdateUserResponse](#innovation.minning.v1.UpdateUserResponse)
    - [User](#innovation.minning.v1.User)
    - [VersionResponse](#innovation.minning.v1.VersionResponse)
  
    - [InnovationMinning](#innovation.minning.v1.InnovationMinning)
  
- [Scalar Value Types](#scalar-value-types)



<a name="npool/innovation-minning.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## npool/innovation-minning.proto



<a name="innovation.minning.v1.Capital"></a>

### Capital



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Name | [string](#string) |  |  |
| Introduction | [string](#string) |  |  |
| Logo | [string](#string) |  |  |






<a name="innovation.minning.v1.CreateCapitalRequest"></a>

### CreateCapitalRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Capital](#innovation.minning.v1.Capital) |  |  |






<a name="innovation.minning.v1.CreateCapitalResponse"></a>

### CreateCapitalResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Capital](#innovation.minning.v1.Capital) |  |  |






<a name="innovation.minning.v1.CreateLaunchTimeRequest"></a>

### CreateLaunchTimeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [LaunchTime](#innovation.minning.v1.LaunchTime) |  |  |






<a name="innovation.minning.v1.CreateLaunchTimeResponse"></a>

### CreateLaunchTimeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [LaunchTime](#innovation.minning.v1.LaunchTime) |  |  |






<a name="innovation.minning.v1.CreateProjectRequest"></a>

### CreateProjectRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Project](#innovation.minning.v1.Project) |  |  |






<a name="innovation.minning.v1.CreateProjectResponse"></a>

### CreateProjectResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Project](#innovation.minning.v1.Project) |  |  |






<a name="innovation.minning.v1.CreateTeamRequest"></a>

### CreateTeamRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Team](#innovation.minning.v1.Team) |  |  |






<a name="innovation.minning.v1.CreateTeamResponse"></a>

### CreateTeamResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Team](#innovation.minning.v1.Team) |  |  |






<a name="innovation.minning.v1.CreateTechniqueAnalysisRequest"></a>

### CreateTechniqueAnalysisRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [TechniqueAnalysis](#innovation.minning.v1.TechniqueAnalysis) |  |  |






<a name="innovation.minning.v1.CreateTechniqueAnalysisResponse"></a>

### CreateTechniqueAnalysisResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [TechniqueAnalysis](#innovation.minning.v1.TechniqueAnalysis) |  |  |






<a name="innovation.minning.v1.CreateTrendAnalysisRequest"></a>

### CreateTrendAnalysisRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [TrendAnalysis](#innovation.minning.v1.TrendAnalysis) |  |  |






<a name="innovation.minning.v1.CreateTrendAnalysisResponse"></a>

### CreateTrendAnalysisResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [TrendAnalysis](#innovation.minning.v1.TrendAnalysis) |  |  |






<a name="innovation.minning.v1.CreateUserRequest"></a>

### CreateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [User](#innovation.minning.v1.User) |  |  |






<a name="innovation.minning.v1.CreateUserResponse"></a>

### CreateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [User](#innovation.minning.v1.User) |  |  |






<a name="innovation.minning.v1.LaunchTime"></a>

### LaunchTime



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| ProjectID | [string](#string) |  |  |
| NetworkName | [string](#string) |  |  |
| NetworkVersion | [string](#string) |  |  |
| Introduction | [string](#string) |  |  |
| LaunchTime | [uint32](#uint32) |  |  |
| Incentive | [bool](#bool) |  |  |
| IncentiveTotal | [uint32](#uint32) |  |  |






<a name="innovation.minning.v1.Project"></a>

### Project



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| Name | [string](#string) |  |  |
| GoodID | [string](#string) |  |  |
| TeamID | [string](#string) |  |  |
| CapitalIDs | [string](#string) | repeated |  |
| Introduction | [string](#string) |  |  |
| Logo | [string](#string) |  |  |






<a name="innovation.minning.v1.Team"></a>

### Team



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| TeamName | [string](#string) |  |  |
| Logo | [string](#string) |  |  |
| LeaderID | [string](#string) |  |  |
| MemberIDs | [string](#string) | repeated |  |






<a name="innovation.minning.v1.TechniqueAnalysis"></a>

### TechniqueAnalysis



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| AuthorID | [string](#string) |  |  |
| Title | [string](#string) |  |  |
| Content | [string](#string) |  |  |
| ProjectID | [string](#string) |  |  |
| CreateAt | [uint32](#uint32) |  |  |






<a name="innovation.minning.v1.TrendAnalysis"></a>

### TrendAnalysis



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| ProjectID | [string](#string) |  |  |
| AuthorID | [string](#string) |  |  |
| Title | [string](#string) |  |  |
| Content | [string](#string) |  |  |
| CreateAt | [string](#string) |  |  |






<a name="innovation.minning.v1.UpdateCapitalRequest"></a>

### UpdateCapitalRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Capital](#innovation.minning.v1.Capital) |  |  |






<a name="innovation.minning.v1.UpdateCapitalResponse"></a>

### UpdateCapitalResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Capital](#innovation.minning.v1.Capital) |  |  |






<a name="innovation.minning.v1.UpdateLaunchTimeRequest"></a>

### UpdateLaunchTimeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [LaunchTime](#innovation.minning.v1.LaunchTime) |  |  |






<a name="innovation.minning.v1.UpdateLaunchTimeResponse"></a>

### UpdateLaunchTimeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [LaunchTime](#innovation.minning.v1.LaunchTime) |  |  |






<a name="innovation.minning.v1.UpdateProjectRequest"></a>

### UpdateProjectRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Project](#innovation.minning.v1.Project) |  |  |






<a name="innovation.minning.v1.UpdateProjectResponse"></a>

### UpdateProjectResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Project](#innovation.minning.v1.Project) |  |  |






<a name="innovation.minning.v1.UpdateTeamRequest"></a>

### UpdateTeamRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Team](#innovation.minning.v1.Team) |  |  |






<a name="innovation.minning.v1.UpdateTeamResponse"></a>

### UpdateTeamResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [Team](#innovation.minning.v1.Team) |  |  |






<a name="innovation.minning.v1.UpdateTechniqueAnalysisRequest"></a>

### UpdateTechniqueAnalysisRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [TechniqueAnalysis](#innovation.minning.v1.TechniqueAnalysis) |  |  |






<a name="innovation.minning.v1.UpdateTechniqueAnalysisResponse"></a>

### UpdateTechniqueAnalysisResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [TechniqueAnalysis](#innovation.minning.v1.TechniqueAnalysis) |  |  |






<a name="innovation.minning.v1.UpdateTrendAnalysisRequest"></a>

### UpdateTrendAnalysisRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [TrendAnalysis](#innovation.minning.v1.TrendAnalysis) |  |  |






<a name="innovation.minning.v1.UpdateTrendAnalysisResponse"></a>

### UpdateTrendAnalysisResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [TrendAnalysis](#innovation.minning.v1.TrendAnalysis) |  |  |






<a name="innovation.minning.v1.UpdateUserRequest"></a>

### UpdateUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [User](#innovation.minning.v1.User) |  |  |






<a name="innovation.minning.v1.UpdateUserResponse"></a>

### UpdateUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [User](#innovation.minning.v1.User) |  |  |






<a name="innovation.minning.v1.User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ID | [string](#string) |  |  |
| FirstName | [string](#string) |  |  |
| LastName | [string](#string) |  |  |
| Introduction | [string](#string) |  |  |






<a name="innovation.minning.v1.VersionResponse"></a>

### VersionResponse
request body and response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Info | [string](#string) |  |  |





 

 

 


<a name="innovation.minning.v1.InnovationMinning"></a>

### InnovationMinning
Service Name

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Version | [.google.protobuf.Empty](#google.protobuf.Empty) | [VersionResponse](#innovation.minning.v1.VersionResponse) | Method Version |
| CreateUser | [CreateUserRequest](#innovation.minning.v1.CreateUserRequest) | [CreateUserResponse](#innovation.minning.v1.CreateUserResponse) |  |
| UpdateUser | [UpdateUserRequest](#innovation.minning.v1.UpdateUserRequest) | [UpdateUserResponse](#innovation.minning.v1.UpdateUserResponse) |  |
| CreateTeam | [CreateTeamRequest](#innovation.minning.v1.CreateTeamRequest) | [CreateTeamResponse](#innovation.minning.v1.CreateTeamResponse) |  |
| UpdateTeam | [UpdateTeamRequest](#innovation.minning.v1.UpdateTeamRequest) | [UpdateTeamResponse](#innovation.minning.v1.UpdateTeamResponse) |  |
| CreateTechniqueAnalysis | [CreateTechniqueAnalysisRequest](#innovation.minning.v1.CreateTechniqueAnalysisRequest) | [CreateTechniqueAnalysisResponse](#innovation.minning.v1.CreateTechniqueAnalysisResponse) |  |
| UpdateTechniqueAnalysis | [UpdateTechniqueAnalysisRequest](#innovation.minning.v1.UpdateTechniqueAnalysisRequest) | [UpdateTechniqueAnalysisResponse](#innovation.minning.v1.UpdateTechniqueAnalysisResponse) |  |
| CreateTrendAnalysis | [CreateTrendAnalysisRequest](#innovation.minning.v1.CreateTrendAnalysisRequest) | [CreateTrendAnalysisResponse](#innovation.minning.v1.CreateTrendAnalysisResponse) |  |
| UpdateTrendAnalysis | [UpdateTrendAnalysisRequest](#innovation.minning.v1.UpdateTrendAnalysisRequest) | [UpdateTrendAnalysisResponse](#innovation.minning.v1.UpdateTrendAnalysisResponse) |  |
| CreateCapital | [CreateCapitalRequest](#innovation.minning.v1.CreateCapitalRequest) | [CreateCapitalResponse](#innovation.minning.v1.CreateCapitalResponse) |  |
| UpdateCapital | [UpdateCapitalRequest](#innovation.minning.v1.UpdateCapitalRequest) | [UpdateCapitalResponse](#innovation.minning.v1.UpdateCapitalResponse) |  |
| CreateProject | [CreateProjectRequest](#innovation.minning.v1.CreateProjectRequest) | [CreateProjectResponse](#innovation.minning.v1.CreateProjectResponse) |  |
| UpdateProject | [UpdateProjectRequest](#innovation.minning.v1.UpdateProjectRequest) | [UpdateProjectResponse](#innovation.minning.v1.UpdateProjectResponse) |  |
| CreateLaunchTime | [CreateLaunchTimeRequest](#innovation.minning.v1.CreateLaunchTimeRequest) | [CreateLaunchTimeResponse](#innovation.minning.v1.CreateLaunchTimeResponse) |  |
| UpdateLaunchTime | [UpdateLaunchTimeRequest](#innovation.minning.v1.UpdateLaunchTimeRequest) | [UpdateLaunchTimeResponse](#innovation.minning.v1.UpdateLaunchTimeResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

