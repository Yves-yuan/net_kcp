/**
 * Auto generated, do not edit it
 */
package data

import (
	. "server/core/data/container"
)

type LoadedHook func()

var GameDataManager struct {
	PassBallContainer
	GlobalConstantContainer
	OpenBallPositionContainer
	FastDribbleContainer
	CatchBallContainer
	RecvBallAnimCfgContainer
	OffBallDecelContainer
	OffballfaststopContainer
	OffballStartContainer
	OffballAccContainer
	OffballRunContainer
	OffBallNaturalStopContainer
	NormalDribbleContainer
	NormalDribblePosContainer
	SlidingContainer
	PlayerAnimIDNameContainer
	PlayerAnimConfigContainer
	SceneContainer
	RuleContainer
	CornerBallPositionContainer
	InterceptContainer
	StartdribbleContainer
}

var hooks []LoadedHook

func AfterLoad(f LoadedHook) {
	hooks = append(hooks, f)
}

func LoadAll() {
	GameDataManager.PassBallContainer.LoadDataFromBin()
	GameDataManager.GlobalConstantContainer.LoadDataFromBin()
	GameDataManager.OpenBallPositionContainer.LoadDataFromBin()
	GameDataManager.FastDribbleContainer.LoadDataFromBin()
	GameDataManager.CatchBallContainer.LoadDataFromBin()
	GameDataManager.RecvBallAnimCfgContainer.LoadDataFromBin()
	GameDataManager.OffBallDecelContainer.LoadDataFromBin()
	GameDataManager.OffballfaststopContainer.LoadDataFromBin()
	GameDataManager.OffballStartContainer.LoadDataFromBin()
	GameDataManager.OffballAccContainer.LoadDataFromBin()
	GameDataManager.OffballRunContainer.LoadDataFromBin()
	GameDataManager.OffBallNaturalStopContainer.LoadDataFromBin()
	GameDataManager.NormalDribbleContainer.LoadDataFromBin()
	GameDataManager.NormalDribblePosContainer.LoadDataFromBin()
	GameDataManager.SlidingContainer.LoadDataFromBin()
	GameDataManager.PlayerAnimIDNameContainer.LoadDataFromBin()
	GameDataManager.PlayerAnimConfigContainer.LoadDataFromBin()
	GameDataManager.SceneContainer.LoadDataFromBin()
	GameDataManager.RuleContainer.LoadDataFromBin()
	GameDataManager.CornerBallPositionContainer.LoadDataFromBin()
	GameDataManager.InterceptContainer.LoadDataFromBin()
	GameDataManager.StartdribbleContainer.LoadDataFromBin()

	// Run hooks
	for _, f := range hooks {
		f()
	}
}
