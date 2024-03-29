package board

var (
	translaterX = map[int]string{
		0: "1",
		1: "2",
		2: "3",
	}
	translaterY = map[int]string{
		0: "A",
		1: "B",
		2: "C",
	}
)

func TranslateStepToLetter(step stepCoordinates) StepLetterCoordinates {
	translatedX := translaterX[step.X]
	translatedY := translaterY[step.Y]

	return StepLetterCoordinates{X: translatedX, Y: translatedY}
}
