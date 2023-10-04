package rps

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/yessikaplata/rps-game/rps/mocks"
)

type rpsTestSuite struct {
	suite.Suite
	ramdomMock         *mocks.RandomGenerator
	rPSGame            RPSGame
	roundExpected      Round
	round              Round
	computerChoice     int
	computerChoiceName string
}

func (suite *rpsTestSuite) SetupSuite() {
}

func (suite *rpsTestSuite) SetupTest() {
	suite.ramdomMock = new(mocks.RandomGenerator)
}

func (suite *rpsTestSuite) TearDownTest() {
	suite.ramdomMock.AssertExpectations(suite.T())
}

func TestPlayRound(t *testing.T) {
	suite.Run(t, new(rpsTestSuite))
}

// ************************************* TESTS *********************************//
func (suite *rpsTestSuite) TestWhenPlayerChooseRockAndComputerChooseRock_thenDraw() {
	suite.givenARPSGame()
	suite.andComputerChoise(0, "The Computer chosen Rock")
	suite.andRoundExpected(suite.getRoundDraw("Try again!"))
	suite.whenPlayRound(0)
	suite.thenPlaySuccessfully()
}

func (suite *rpsTestSuite) TestWhenPlayerChooseRockAndComputerChoosePaper_thenLose() {
	suite.givenARPSGame()
	suite.andComputerChoise(1, "The Computer chosen Paper")
	suite.andRoundExpected(suite.getRoundLose("You are out of luck!"))
	suite.whenPlayRound(0)
	suite.thenPlaySuccessfully()
}

func (suite *rpsTestSuite) TestWhenPlayerChooseRockAndComputerChooseScissors_thenWin() {
	suite.givenARPSGame()
	suite.andComputerChoise(2, "The Computer chosen Scissors")
	suite.andRoundExpected(suite.getRoundWin("Well done!"))
	suite.whenPlayRound(0)
	suite.thenPlaySuccessfully()
}

func (suite *rpsTestSuite) TestWhenPlayerChoosePaperAndComputerChooseRock_thenWin() {
	suite.givenARPSGame()
	suite.andComputerChoise(0, "The Computer chosen Rock")
	suite.andRoundExpected(suite.getRoundWin("You got it!"))
	suite.whenPlayRound(1)
	suite.thenPlaySuccessfully()
}

func (suite *rpsTestSuite) TestWhenPlayerChoosePaperAndComputerChoosePaper_thenDraw() {
	suite.givenARPSGame()
	suite.andComputerChoise(1, "The Computer chosen Paper")
	suite.andRoundExpected(suite.getRoundDraw("You are out of luck!"))
	suite.whenPlayRound(1)
	suite.thenPlaySuccessfully()
}

func (suite *rpsTestSuite) TestWhenPlayerChoosePaperAndComputerChooseScissors_thenLose() {
	suite.givenARPSGame()
	suite.andComputerChoise(2, "The Computer chosen Scissors")
	suite.andRoundExpected(suite.getRoundLose("Is not your day!"))
	suite.whenPlayRound(1)
	suite.thenPlaySuccessfully()
}

func (suite *rpsTestSuite) TestWhenPlayerChooseScissorsAndComputerChooseScissors_thenDraw() {
	suite.givenARPSGame()
	suite.andComputerChoise(2, "The Computer chosen Scissors")
	suite.andRoundExpected(suite.getRoundDraw("Is not your day!"))
	suite.whenPlayRound(2)
	suite.thenPlaySuccessfully()
}

func (suite *rpsTestSuite) TestWhenPlayerChooseScissorsAndComputerChooseRock_thenLose() {
	suite.givenARPSGame()
	suite.andComputerChoise(0, "The Computer chosen Rock")
	suite.andRoundExpected(suite.getRoundLose("Try again!"))
	suite.whenPlayRound(2)
	suite.thenPlaySuccessfully()
}

func (suite *rpsTestSuite) TestWhenPlayerChooseScissorsAndComputerChoosePaper_thenWin() {
	suite.givenARPSGame()
	suite.andComputerChoise(1, "The Computer chosen Paper")
	suite.andRoundExpected(suite.getRoundWin("Good Job!"))
	suite.whenPlayRound(2)
	suite.thenPlaySuccessfully()
}

// ************************************* GIVEN *********************************//
func (suite *rpsTestSuite) givenARPSGame() {
	suite.rPSGame = NewRPSGame(suite.ramdomMock)
}

func (suite *rpsTestSuite) andRoundExpected(round Round) {
	suite.roundExpected = round
}

func (suite *rpsTestSuite) andComputerChoise(option int, optionName string) {
	suite.computerChoice = option
	suite.computerChoiceName = optionName
	suite.ramdomMock.
		On("Intn", 3).
		Return(option)

}

// ************************************* WHEN *********************************//
func (suite *rpsTestSuite) whenPlayRound(option int) {
	suite.round = suite.rPSGame.PlayRound(option)
}

// ************************************* THEN *********************************//
func (suite *rpsTestSuite) thenPlaySuccessfully() {
	suite.Equal(suite.round, suite.roundExpected)
}

// ************************************* AUX *********************************//

func (suite *rpsTestSuite) getRoundDraw(message string) Round {
	return Round{
		Message:            message,
		ComputerChoice:     suite.computerChoice,
		ComputerChoiceName: suite.computerChoiceName,
		RoundResult:        "It is a draw",
		ComputerScore:      "0",
		PlayerScore:        "0",
	}
}

func (suite *rpsTestSuite) getRoundLose(message string) Round {
	return Round{
		Message:            message,
		ComputerChoice:     suite.computerChoice,
		ComputerChoiceName: suite.computerChoiceName,
		RoundResult:        "You have lost",
		ComputerScore:      "1",
		PlayerScore:        "0",
	}
}

func (suite *rpsTestSuite) getRoundWin(message string) Round {
	return Round{
		Message:            message,
		ComputerChoice:     suite.computerChoice,
		ComputerChoiceName: suite.computerChoiceName,
		RoundResult:        "You have won",
		ComputerScore:      "0",
		PlayerScore:        "1",
	}
}
