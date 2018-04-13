package core

import (
	"github.com/erayarslan/furkan-go/info"
	"errors"
	"strings"
)

func Kiss() string {
	info.Happy++
	return info.Lip
}

func Love() error {
	info.Happy--
	return errors.New("no way pointer exception")
}

func WhoAmI() string {
	return info.GetName()
}

func Status() string {
	return info.GetRelationship()
}

func Mood() string {
	var tempHappy = info.Happy
	if tempHappy == 0 {
		return info.Eray
	} else if tempHappy > 0 && tempHappy < 4 {
		return info.Lip
	} else {
		return info.Ahmet
	}
}

func Say(text string) (string, error) {
	if strings.ToLower(text) == info.BadAss {
		if info.Happy-1 != -1 {
			info.Happy--
		}

		return "", errors.New(info.GetName() +
			info.Space +
			"dont" +
			info.Space +
			"apply" +
			info.Space +
			info.Eray +
			info.Space +
			"this" +
			info.Space +
			"text!")
	}

	return text + info.Space + info.GetName(), nil
}

func AmICute() int {
	return info.Yes
}

func I1337() string {
	return "Furk4n"
}
