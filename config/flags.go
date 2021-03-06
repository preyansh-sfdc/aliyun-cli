// Copyright (c) 2009-present, Alibaba Cloud All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package config

import (
	"github.com/aliyun/aliyun-cli/cli"
	"github.com/aliyun/aliyun-cli/i18n"
)

const (
	ProfileFlagName         = "profile"
	ModeFlagName            = "mode"
	AccessKeyIdFlagName     = "access-key-id"
	AccessKeySecretFlagName = "access-key-secret"
	StsTokenFlagName        = "sts-token"
	StsRegionFlagName       = "sts-region"
	RamRoleNameFlagName     = "ram-role-name"
	RamRoleArnFlagName      = "ram-role-arn"
	RoleSessionNameFlagName = "role-session-name"
	SourceProfileFlagName   = "source-profile"
	PrivateKeyFlagName      = "private-key"
	KeyPairNameFlagName     = "key-pair-name"
	RegionFlagName          = "region"
	LanguageFlagName        = "language"
	ReadTimeoutFlagName     = "read-timeout"
	ConnectTimeoutFlagName  = "connect-timeout"
	RetryCountFlagName      = "retry-count"
	SkipSecureVerifyName    = "skip-secure-verify"
	ConfigurePathFlagName   = "config-path"
	ExpiredSecondsFlagName  = "expired-seconds"
	ProcessCommandFlagName  = "process-command"
)

func AddFlags(fs *cli.FlagSet) {
	fs.Add(NewModeFlag())
	fs.Add(NewProfileFlag())
	fs.Add(NewLanguageFlag())
	fs.Add(NewRegionFlag())
	fs.Add(NewConfigurePathFlag())
	fs.Add(NewAccessKeyIdFlag())
	fs.Add(NewAccessKeySecretFlag())
	fs.Add(NewStsTokenFlag())
	fs.Add(NewStsRegionFlag())
	fs.Add(NewRamRoleNameFlag())
	fs.Add(NewRamRoleArnFlag())
	fs.Add(NewRoleSessionNameFlag())
	fs.Add(NewPrivateKeyFlag())
	fs.Add(NewKeyPairNameFlag())
	fs.Add(NewReadTimeoutFlag())
	fs.Add(NewConnectTimeoutFlag())
	fs.Add(NewRetryCountFlag())
	fs.Add(NewSkipSecureVerify())
	fs.Add(NewExpiredSecondsFlag())
	fs.Add(NewProcessCommandFlag())
}

func ConnectTimeoutFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(ConnectTimeoutFlagName)
}

func ProfileFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(ProfileFlagName)
}

func ModeFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(ModeFlagName)
}

func AccessKeyIdFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(AccessKeyIdFlagName)
}

func AccessKeySecretFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(AccessKeySecretFlagName)
}

func StsTokenFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(StsTokenFlagName)
}

func StsRegionFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(StsRegionFlagName)
}

func RamRoleNameFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(RamRoleNameFlagName)
}

func RamRoleArnFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(RamRoleArnFlagName)
}

func SourceProfileFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(SourceProfileFlagName)
}

func RoleSessionNameFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(RoleSessionNameFlagName)
}

func PrivateKeyFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(PrivateKeyFlagName)
}

func KeyPairNameFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(KeyPairNameFlagName)
}

func RegionFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(RegionFlagName)
}

func LanguageFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(LanguageFlagName)
}

func ReadTimeoutFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(ReadTimeoutFlagName)
}

func RetryCountFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(RetryCountFlagName)
}

func SkipSecureVerify(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(SkipSecureVerifyName)
}
func ConfigurePathFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(ConfigurePathFlagName)
}
func ExpiredSecondsFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(ExpiredSecondsFlagName)
}
func ProcessCommandFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(ProcessCommandFlagName)
}

//var OutputFlag = &cli.Flag{Category: "config",
//	Name: "output", AssignedMode: cli.AssignedOnce, Hidden: true,
//	Usage: i18n.T(
//		"* assign output format, only support json",
//		"* ??????????????????, ???????????????Json")}

//varfs.Add(cli.Flag{Name: "site", AssignedMode: cli.AssignedOnce,
//Usage: i18n.T("assign site, support china/international/japan", "")})

func NewProfileFlag() *cli.Flag {
	return &cli.Flag{
		Category:     "config",
		Name:         ProfileFlagName,
		Shorthand:    'p',
		DefaultValue: "default", Persistent: true,
		Short: i18n.T(
			"use `--profile <profileName>` to select profile",
			"?????? `--profile <profileName>` ????????????????????????")}
}

func NewModeFlag() *cli.Flag {
	return &cli.Flag{
		Category: "config",
		Name:     ModeFlagName, DefaultValue: "AK", Persistent: true,
		Short: i18n.T(
			"use `--mode {AK|StsToken|RamRoleArn|EcsRamRole|RsaKeyPair|RamRoleArnWithRoleName}` to assign authenticate mode",
			"?????? `--mode {AK|StsToken|RamRoleArn|EcsRamRole|RsaKeyPair|RamRoleArnWithRoleName}` ??????????????????")}
}

func NewAccessKeyIdFlag() *cli.Flag {
	return &cli.Flag{
		Category:     "config",
		Name:         AccessKeyIdFlagName,
		AssignedMode: cli.AssignedOnce,
		Short: i18n.T(
			"use `--access-key-id <AccessKeyId>` to assign AccessKeyId, required in AK/StsToken/RamRoleArn mode",
			"?????? `--access-key-id <AccessKeyId>` ??????AccessKeyId")}
}

func NewAccessKeySecretFlag() *cli.Flag {
	return &cli.Flag{
		Category:     "config",
		Name:         AccessKeySecretFlagName,
		AssignedMode: cli.AssignedOnce,
		Short: i18n.T(
			"use `--access-key-secret <AccessKeySecret>` to assign AccessKeySecret",
			"?????? `--access-key-secret <AccessKeySecret>` ??????AccessKeySecret")}
}

func NewStsTokenFlag() *cli.Flag {
	return &cli.Flag{
		Category:     "config",
		Name:         StsTokenFlagName,
		AssignedMode: cli.AssignedOnce,
		Short: i18n.T(
			"use `--sts-token <StsToken>` to assign StsToken",
			"?????? `--sts-token <StsToken>` ??????StsToken")}
}

func NewStsRegionFlag() *cli.Flag {
	return &cli.Flag{
		Category:     "config",
		Name:         StsRegionFlagName,
		AssignedMode: cli.AssignedOnce,
		Short: i18n.T(
			"use `--sts-region <StsRegion>` to assign StsRegion",
			"?????? `--sts-region <StsRegion>` ??????StsRegion")}
}

func NewRamRoleNameFlag() *cli.Flag {
	return &cli.Flag{
		Category:     "config",
		Name:         RamRoleNameFlagName,
		AssignedMode: cli.AssignedOnce,
		Persistent:   true,
		Short: i18n.T(
			"use `--ram-role-name <RamRoleName>` to assign RamRoleName",
			"?????? `--ram-role-name <RamRoleName>` ??????RamRoleName")}
}

func NewRamRoleArnFlag() *cli.Flag {
	return &cli.Flag{
		Category:     "config",
		Name:         RamRoleArnFlagName,
		AssignedMode: cli.AssignedOnce,
		Short: i18n.T(
			"use `--ram-role-arn <RamRoleArn>` to assign RamRoleArn",
			"?????? `--ram-role-arn <RamRoleArn>` ??????RamRoleArn")}
}

func NewRoleSessionNameFlag() *cli.Flag {
	return &cli.Flag{Category: "config",
		Name:         RoleSessionNameFlagName,
		AssignedMode: cli.AssignedOnce,
		Short: i18n.T(
			"use `--role-session-name <RoleSessionName>` to assign RoleSessionName",
			"?????? `--role-session-name <RoleSessionName>` ??????RoleSessionName")}
}
func NewExpiredSecondsFlag() *cli.Flag {
	return &cli.Flag{
		Category:     "config",
		Name:         ExpiredSecondsFlagName,
		AssignedMode: cli.AssignedOnce,
		Short: i18n.T(
			"use `--expired-seconds <seconds>` to specify expiration time",
			"?????? `--expired-seconds <seconds>` ????????????????????????")}
}
func NewPrivateKeyFlag() *cli.Flag {
	return &cli.Flag{Category: "config",
		Name:         PrivateKeyFlagName,
		AssignedMode: cli.AssignedOnce,
		Short: i18n.T(
			"use `--private-key <PrivateKey>` to assign RSA PrivateKey",
			"?????? `--private-key <PrivateKey>` ??????RSA??????")}
}

func NewKeyPairNameFlag() *cli.Flag {
	return &cli.Flag{Category: "config",
		Name:         KeyPairNameFlagName,
		AssignedMode: cli.AssignedOnce,
		Short: i18n.T(
			"use `--key-pair-name <KeyPairName>` to assign KeyPairName",
			"?????? `--key-pair-name <KeyPairName>` ??????KeyPairName")}
}

func NewProcessCommandFlag() *cli.Flag {
	return &cli.Flag{
		Category:     "config",
		Name:         ProcessCommandFlagName,
		AssignedMode: cli.AssignedOnce,
		Short: i18n.T(
			"use `--process-command <ProcessCommand>` to specify external program execution command",
			"?????? `--process-command <ProcessCommand>` ??????????????????????????????",
		),
	}
}

func NewRegionFlag() *cli.Flag {
	return &cli.Flag{Category: "config",
		Name:         RegionFlagName,
		AssignedMode: cli.AssignedOnce,
		Short: i18n.T(
			"use `--region <regionId>` to assign region",
			"?????? `--region <regionId>` ?????????????????????")}
}

func NewLanguageFlag() *cli.Flag {
	return &cli.Flag{
		Category:     "config",
		Name:         LanguageFlagName,
		AssignedMode: cli.AssignedOnce,
		Short: i18n.T(
			"use `--language [en|zh]` to assign language",
			"?????? `--language [en|zh]` ???????????????")}
}
func NewConfigurePathFlag() *cli.Flag {
	return &cli.Flag{
		Category:     "config",
		Name:         ConfigurePathFlagName,
		AssignedMode: cli.AssignedOnce,
		Persistent:   true,
		Short: i18n.T(
			"use `--config-path` to specify the configuration file path",
			"?????? `--config-path` ????????????????????????",
		),
	}
}
func NewReadTimeoutFlag() *cli.Flag {
	return &cli.Flag{
		Category:     "caller",
		Name:         ReadTimeoutFlagName,
		AssignedMode: cli.AssignedOnce,
		Aliases:      []string{"retry-timeout"},
		Short: i18n.T(
			"use `--read-timeout <seconds>` to set I/O timeout(seconds)",
			"?????? `--read-timeout <seconds>` ??????I/O????????????(???)"),
	}
}

func NewConnectTimeoutFlag() *cli.Flag {
	return &cli.Flag{
		Category:     "caller",
		Name:         ConnectTimeoutFlagName,
		AssignedMode: cli.AssignedOnce,
		Short: i18n.T(
			"use `--connect-timeout <seconds>` to set connect timeout(seconds)",
			"?????? `--connect-timeout <seconds>` ??????????????????????????????(???)"),
	}
}

func NewRetryCountFlag() *cli.Flag {
	return &cli.Flag{
		Category:     "caller",
		Name:         RetryCountFlagName,
		AssignedMode: cli.AssignedOnce,
		Short: i18n.T(
			"use `--retry-count <count>` to set retry count",
			"?????? `--retry-count <count>` ??????????????????"),
	}
}

func NewSkipSecureVerify() *cli.Flag {
	return &cli.Flag{
		Category:     "caller",
		Name:         SkipSecureVerifyName,
		AssignedMode: cli.AssignedNone,
		Persistent:   true,
		Short: i18n.T(
			"use `--skip-secure-verify` to skip https certification validate [Not recommended]",
			"?????? `--skip-secure-verify` ??????https??????????????? [???????????????]",
		),
	}
}
