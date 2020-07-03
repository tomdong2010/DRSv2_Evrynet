package collateral_test

import (
	"github.com/Evrynetlabs/evrynet-node/common"
	"github.com/Evrynetlabs/evrynet-node/core/types"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/velo-protocol/DRSv2_Evrynet/go/cmd/gvel/layers/commands/collateral"
	"github.com/velo-protocol/DRSv2_Evrynet/go/cmd/gvel/layers/mocks"
	"github.com/velo-protocol/DRSv2_Evrynet/go/cmd/gvel/utils/console"
	mockutils "github.com/velo-protocol/DRSv2_Evrynet/go/cmd/gvel/utils/mocks"
	"math/big"
	"testing"
)

type helper struct {
	commandHandler *collateral.CommandHandler
	loggerHook     *test.Hook
	tableLogHook   *test.Hook
	mockLogic      *mocks.MockLogic
	mockPrompt     *mockutils.MockPrompt
	mockConfig     *mockutils.MockConfiguration
	mockController *gomock.Controller
	mockTx         *types.Transaction
	done           func()
}

func initTest(t *testing.T) *helper {
	mockCtrl := gomock.NewController(t)
	mockLogic := mocks.NewMockLogic(mockCtrl)
	mockPrompt := mockutils.NewMockPrompt(mockCtrl)
	mockConfig := mockutils.NewMockConfiguration(mockCtrl)

	logger, hook := test.NewNullLogger()
	console.Logger = logger

	// table logger
	tableLogger, tableLogHook := test.NewNullLogger()
	console.TableLogger = tableLogger

	// to omit what loader print
	console.DefaultLoadWriter = console.Logger.Out

	return &helper{
		commandHandler: collateral.NewCommandHandler(mockLogic, mockPrompt, mockConfig),
		mockLogic:      mockLogic,
		mockPrompt:     mockPrompt,
		mockConfig:     mockConfig,
		mockController: mockCtrl,
		mockTx:         types.NewTransaction(1, common.Address{}, big.NewInt(1), 1, big.NewInt(1), []byte{}),
		loggerHook:     hook,
		tableLogHook:   tableLogHook,
		done: func() {
			mockCtrl.Finish()
			hook.Reset()
		},
	}
}
