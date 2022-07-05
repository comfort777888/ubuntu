package handlers

import (
	ascii "ascii-art-web/internal/ascii"
)

// CheckUserInput function returns status codes of errors if it cannot process input data.
func CheckUserInput(t string, font string) int {
	// check if press enter for a new line
	for _, v := range t {
		if v == 10 || v == 13 { // checking if it's new line or '\r' (carriage ret)
			continue
		}
		if v < 32 || v > 126 { // if text is cyrrilic
			return 400
		}
	}
	if t == "" { // if we receive empty text form
		return 400
	}
	if font != "standard.txt" && font != "shadow.txt" && font != "thinkertoy.txt" {
		return 400
	}
	if !ascii.ConvertFont(font) {
		return 500
	}
	if ascii.HashSum(font) { // Checks font's file hashsum
		return 500
	}
	return 200
}
