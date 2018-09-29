package lcd

import (
	"bufio"
	"strconv"
	"strings"
)

func Display(scanner *bufio.Scanner, animate bool) {
	lcd := NewLCD(50, 6)
	presenter := initializePresenter(animate)

	for scanner.Scan() {
		if line := scanner.Text(); len(strings.TrimSpace(line)) > 0 {
			performInstruction(line, lcd)
			presenter.Present(lcd)
		}
	}

	presenter.Finalize()
}

func initializePresenter(animate bool) Presenter {
	if animate {
		return NewAnimatedPresenter()
	} else {
		return NewStaticPresenter()
	}
}

func performInstruction(line string, lcd *LCD) {
	line = strings.Replace(line, "x=", "", -1)
	line = strings.Replace(line, "y=", "", -1)
	line = strings.Replace(line, "x", " ", -1)
	fields := strings.Fields(line)

	if strings.HasPrefix(line, "rect") {
		columns := parseInt(fields[1])
		rows := parseInt(fields[2])
		lcd.RectangleOn(columns, rows)
	} else if strings.HasPrefix(line, "rotate row") {
		row := parseInt(fields[2])
		rotation := parseInt(fields[4])
		lcd.RotateRow(row, rotation)
	} else if strings.HasPrefix(line, "rotate column") {
		column := parseInt(fields[2])
		rotation := parseInt(fields[4])
		lcd.RotateColumn(column, rotation)
	}
}

func parseInt(value string) int {
	parsed, _ := strconv.Atoi(value)
	return parsed
}
