package config

const (
	SINGLE = 1*iota + 1
	MULTI
)

var AsciiArt = `
     _______..__   __.      ___       __  ___  _______
    /       ||  \ |  |     /   \     |  |/  / |   ____|
   |   (----` + "`" + `|   \|  |    /  ^  \    |  '  /  |  |__
    \   \    |  . ` + "`" + `  |   /  /_\  \   |    <   |   __|
.----)   |   |  |\   |  /  _____  \  |  .  \  |  |____
|_______/    |__| \__| /__/     \__\ |__|\__\ |_______|
`

var MultiArrowUp = []rune{119, 1094}

var MultiArrowDown = []rune{115, 1099}

var MultiArrowLeft = []rune{97, 1092}

var MultiArrowRight = []rune{100, 1074}

const (
	KeyArrowUp = 0 + iota
	KeyArrowDown
	KeyArrowLeft
	KeyArrowRight
)
