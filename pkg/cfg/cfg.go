// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cfg

var (
	_features = newEmptyFeatures()

	// Use alias of Features.Use func
	Use = _features.Use

	// UseDeafult alias of Features.UseDefault func
	UseDefault = _features.UseDefault

	// As alias of Features.Cfg func
	As = _features.Cfg

	// If alias of Features.CfgIf func
	If = _features.CfgIf

	// In alias of Features.CfgIn func
	In = _features.CfgIn

	// Be alias of Feaures.CfgBe func
	Be = _features.CfgBe

	// Not alias of Features.CfgNot func
	Not = _features.CfgNot
)

// Initialize initialize features in cfg pkg
func Initialize(suites map[string][]string, kv map[string]string) {
	_features = NewFeatures(suites, kv)
	{
		// must re-assign variable below
		Use = _features.Use
		UseDefault = _features.UseDefault
		As = _features.Cfg
		If = _features.CfgIf
		In = _features.CfgIn
		Be = _features.CfgBe
		Not = _features.CfgNot
	}
}
