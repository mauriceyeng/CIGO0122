package main

var isKnightSleep bool
var isPrisonerSleep bool
var isArcherSleep bool
var hasDog bool
var isAllSleep bool

func CanFastAttack() bool {
	return isKnightSleep
}

func CanSpy() bool {
	return !(isKnightSleep && isPrisonerSleep && isArcherSleep)
}

func CanSignalPrisoner() bool {
	return (!isPrisonerSleep) && isArcherSleep

}

func CanFreePrisoner() bool {
	return ((hasDog && isArcherSleep) || (!isPrisonerSleep && isArcherSleep && isKnightSleep))
}

func main() {
	isArcherSleep, isPrisonerSleep, isKnightSleep, hasDog = true, true, true, true
}
