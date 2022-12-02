package day

import (
	"hyoa/aoc2022/internal/utils"
	"strings"
)

type Day2 struct {
	data   []string
	shapes map[ShapeKind]Shape
}

var ShapeValue = map[ShapeKind]int{Rock: 1, Paper: 2, Scissors: 3}
var ShapeSymbol = map[string]ShapeKind{"A": Rock, "B": Paper, "C": Scissors, "X": Rock, "Y": Paper, "Z": Scissors}
var OutcomeSymbol = map[string]Outcome{"X": Lost, "Y": Draw, "Z": Win}

type ShapeKind string
type Outcome int

type Shape struct {
	kind ShapeKind
	win  ShapeKind
	lost ShapeKind
}

const (
	Rock     ShapeKind = "rock"
	Paper    ShapeKind = "paper"
	Scissors ShapeKind = "scissors"
	Win      Outcome   = 6
	Lost     Outcome   = 0
	Draw     Outcome   = 3
)

func (d *Day2) Init(path string) error {
	d.data = utils.ReadTextFileLinesAsString(path)

	shapesRef := make(map[ShapeKind]Shape)
	rock := Shape{kind: Rock, win: Scissors, lost: Paper}
	paper := Shape{kind: Paper, win: Rock, lost: Scissors}
	scissors := Shape{kind: Scissors, win: Paper, lost: Rock}

	shapesRef[Rock] = rock
	shapesRef[Paper] = paper
	shapesRef[Scissors] = scissors

	d.shapes = shapesRef

	return nil
}

func (d *Day2) Step1() (int, error) {
	tt := 0
	for _, roundShapes := range d.data {
		shapes := strings.Split(roundShapes, " ")
		opponentShape := ShapeSymbol[shapes[0]]
		meShape := ShapeSymbol[shapes[1]]

		outcome := d.outcomeRound(meShape, opponentShape)
		tt += int(outcome) + ShapeValue[meShape]
	}

	return tt, nil
}

func (d *Day2) Step2() (int, error) {
	tt := 0
	for _, roundShapes := range d.data {
		shapes := strings.Split(roundShapes, " ")
		opponentShape := ShapeSymbol[shapes[0]]
		outcome := OutcomeSymbol[shapes[1]]

		switch outcome {
		case Win:
			tt += ShapeValue[d.shapes[opponentShape].lost] + int(Win)
		case Lost:
			tt += ShapeValue[d.shapes[opponentShape].win] + int(Lost)
		default:
			tt += ShapeValue[opponentShape] + int(Draw)
		}
	}

	return tt, nil
}

func (d *Day2) outcomeRound(me, opponent ShapeKind) Outcome {
	if me == opponent {
		return Draw
	}

	if me == d.shapes[opponent].lost {
		return Win
	}

	return Lost
}
