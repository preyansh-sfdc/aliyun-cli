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
package openapi

import (
	"github.com/aliyun/aliyun-cli/cli"
	"github.com/aliyun/aliyun-cli/i18n"
)

func AddFlags(fs *cli.FlagSet) {
	fs.Add(NewSecureFlag())
	fs.Add(NewForceFlag())
	fs.Add(NewEndpointFlag())
	fs.Add(NewVersionFlag())
	fs.Add(NewHeaderFlag())
	fs.Add(NewBodyFlag())
	fs.Add(NewBodyFileFlag())
	fs.Add(PagerFlag)
	fs.Add(NewAcceptFlag())
	fs.Add(NewOutputFlag())
	fs.Add(WaiterFlag)
	fs.Add(NewDryRunFlag())
	fs.Add(NewQuietFlag())
	fs.Add(NewRoaFlag())
}

const (
	SecureFlagName   = "secure"
	ForceFlagName    = "force"
	EndpointFlagName = "endpoint"
	VersionFlagName  = "version"
	HeaderFlagName   = "header"
	BodyFlagName     = "body"
	BodyFileFlagName = "body-file"
	AcceptFlagName   = "accept"
	RoaFlagName      = "roa"
	DryRunFlagName   = "dryrun"
	QuietFlagName    = "quiet"
	OutputFlagName   = "output"
)

func OutputFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(OutputFlagName)
}

func SecureFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(SecureFlagName)
}

func ForceFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(ForceFlagName)
}

func EndpointFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(EndpointFlagName)
}

func VersionFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(VersionFlagName)
}

func HeaderFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(HeaderFlagName)
}

func BodyFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(BodyFlagName)
}

func BodyFileFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(BodyFileFlagName)
}

func AcceptFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(AcceptFlagName)
}

func RoaFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(RoaFlagName)
}

func DryRunFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(DryRunFlagName)
}

func QuietFlag(fs *cli.FlagSet) *cli.Flag {
	return fs.Get(QuietFlagName)
}

// TODO next version
//VerboseFlag = &cli.Flag{Category: "caller",
//	Name: "verbose",
//	Shorthand: 'v',
//	AssignedMode: cli.AssignedNone,
//	Short: i18n.T(
//		"add `--verbose` to enable verbose mode",
//		"?????? `--verbose` ??????????????????",
//	),
//}

func NewSecureFlag() *cli.Flag {
	return &cli.Flag{Category: "caller",
		Name: SecureFlagName, AssignedMode: cli.AssignedNone,
		Short: i18n.T(
			"use `--secure` to force https",
			"?????? `--secure` ??????????????????https????????????")}
}

func NewForceFlag() *cli.Flag {
	return &cli.Flag{Category: "caller",
		Name: ForceFlagName, AssignedMode: cli.AssignedNone,
		Short: i18n.T(
			"use `--force` to skip api and parameters check",
			"?????? `--force` ???????????????API???????????????????????????")}
}

func NewEndpointFlag() *cli.Flag {
	return &cli.Flag{Category: "caller",
		Name: EndpointFlagName, AssignedMode: cli.AssignedOnce,
		Short: i18n.T(
			"use `--endpoint <endpoint>` to assign endpoint",
			"?????? `--endpoint <endpoint>` ????????????????????????")}
}

func NewVersionFlag() *cli.Flag {
	return &cli.Flag{Category: "caller",
		Name: VersionFlagName, AssignedMode: cli.AssignedOnce,
		Short: i18n.T(
			"use `--version <YYYY-MM-DD>` to assign product api version",
			"?????? `--version <YYYY-MM-DD>` ??????????????????API??????")}
}

func NewHeaderFlag() *cli.Flag {
	return &cli.Flag{Category: "caller",
		Name: HeaderFlagName, AssignedMode: cli.AssignedRepeatable,
		Short: i18n.T(
			"use `--header X-foo=bar` to add custom HTTP header, repeatable",
			"?????? `--header X-foo=bar` ??????????????????HTTP???, ???????????????")}
}

func NewBodyFlag() *cli.Flag {
	return &cli.Flag{Category: "caller",
		Name: BodyFlagName, AssignedMode: cli.AssignedOnce,
		Short: i18n.T(
			"use `--body $(cat foo.json)` to assign http body in RESTful call",
			"?????? `--body $(cat foo.json)` ????????????RESTful????????????HTTP??????")}
}

func NewBodyFileFlag() *cli.Flag {
	return &cli.Flag{Category: "caller",
		Name: BodyFileFlagName, AssignedMode: cli.AssignedOnce, Hidden: true,
		Short: i18n.T(
			"assign http body in Restful call with local file",
			"?????? `--body-file foo.json` ?????????????????????")}
}

func NewAcceptFlag() *cli.Flag {
	return &cli.Flag{Category: "caller",
		Name: AcceptFlagName, AssignedMode: cli.AssignedOnce, Hidden: true,
		Short: i18n.T(
			"add `--accept {json|xml}` to add Accept header",
			"?????? `--accept {json|xml}` ?????????Accept???")}
}

func NewRoaFlag() *cli.Flag {
	return &cli.Flag{Category: "caller",
		Name: RoaFlagName, AssignedMode: cli.AssignedOnce, Hidden: true,
		Short: i18n.T(
			"use `--roa {GET|PUT|POST|DELETE}` to assign restful call.[DEPRECATED]",
			"?????? `--roa {GET|PUT|POST|DELETE}` ??????restful????????????[?????????]",
		),
	}
}

func NewDryRunFlag() *cli.Flag {
	return &cli.Flag{Category: "caller",
		Name:         DryRunFlagName,
		AssignedMode: cli.AssignedNone,
		Short: i18n.T(
			"add `--dryrun` to validate and print request without running.",
			"?????? `--dryrun` ?????????????????????????????????????????????????????????",
		),
		ExcludeWith: []string{PagerFlag.Name, WaiterFlag.Name},
	}
}

func NewQuietFlag() *cli.Flag {
	return &cli.Flag{Category: "caller",
		Name:         QuietFlagName,
		Shorthand:    'q',
		AssignedMode: cli.AssignedNone,
		Short: i18n.T(
			"add `--quiet` to hide normal output",
			"?????? `--quiet` ??????????????????",
		),
		ExcludeWith: []string{DryRunFlagName},
	}
}
