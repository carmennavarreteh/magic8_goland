package magic8

import (
	"math/rand"
)

var list_answer = []string{
	//affirmative answer
	"It is certain",
	"It is decidedly so",
	"Without a doubt",
	"Yes definitely",
	"You may rely on it",
	"As I see it, yes",
	"Most likely",
	"Outlook good",
	"Yes",
	"Signs point to yes",

	// non-committal answers
	"Reply hazy, try again",
	"Ask again later",
	"Better not tell you now",
	"Cannot predict now",
	"Concentrate and ask again",

	// negative answers
	"Don't count on it",
	"My reply is no",
	"My sources say no",
	"Outlook not so good",
	"Very doubtful",
}

type magic8_ struct {
	answer []string
	seed   int64
}

func (m *magic8_) Answer() string {

	return m.answer[rand.Intn(len(m.answer))]
}

func New(seed int64) *magic8_ {

	rand.Seed(seed)
	return &magic8_{
		answer: list_answer,
		seed:   seed,
	}
}
