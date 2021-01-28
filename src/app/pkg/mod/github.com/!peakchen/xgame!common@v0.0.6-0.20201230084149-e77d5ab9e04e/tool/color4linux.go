package tool

//add by stefan for linux system.

import (
	"fmt"
)

/*
Foreground	Background 	Color
30			40			black
31			41			red
32			42			green
33			43			yellow
34			44			blue
35			45			Fuchsia
36			46			cyan
37			47			white
*/

const (
	LinuxForeground_Black   = 0x1E
	LinuxForeground_Red     = 0x1F
	LinuxForeground_GREEN   = 0x20
	LinuxForeground_YELLOW  = 0x21
	LinuxForeground_BLUE    = 0x22
	LinuxForeground_FUCHSIA = 0x23
	LinuxForeground_CYAN    = 0x24
	LinuxForeground_WHITE   = 0x25
)

const (
	LinuxBackground_Black   = 0x28
	LinuxBackground_Red     = 0x29
	LinuxBackground_GREEN   = 0x2a
	LinuxBackground_YELLOW  = 0x2b
	LinuxBackground_BLUE    = 0x2c
	LinuxBackground_FUCHSIA = 0x2d
	LinuxBackground_CYAN    = 0x2e
	LinuxBackground_WHITE   = 0x2f
)

/*
coding   meaning
	0 	Terminal default settings 	(终端默认设置)
	1 	Highlight	  				(高亮显示)
	4 	Use underline 				(使用下划线)
	5 	flashes     				(闪烁)
	7 	Highlighted 				(反白显示)
	8 	Invisible 					(不可见)
*/

/*
	@param 1: background color
	@param 2: foreground color
	@param 3: content
*/
func LinuxColorPrint(bg int, fg int, str string) {
	//SGR params x1b = 033
	strFmt := fmt.Sprintf("\x1b[%dm\x1b[%dm%s\x1b[0m", bg, fg, str)
	fmt.Println(strFmt)
}

/*
	default background color: black
	@param 1: foreground color
	@param 2: content
*/
func LinuxDefaultBGPrint(fg int, str string) {
	strFmt := fmt.Sprintf("\x1b[40m\x1b[%dm%s\x1b[0m", fg, str)
	fmt.Println(strFmt)
}

/*
	default background color: black
	default foreground color: white
	@param 1: content
*/
func LinuxDefaultColorPrint(str string) {
	strFmt := fmt.Sprintf("\x1b[m%s\x1b[0m", str)
	fmt.Println(strFmt)
}
