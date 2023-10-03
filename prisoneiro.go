package exercicios

// CanQuickAttack pode ser executado apenas quando o cavaleiro está dormindo, pois ele está vulnerável enquanto coloca a armadura.
func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
}

// CanSpy pode ser executado se pelo menos um dos personagens estiver acordado.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	return knightIsAwake || archerIsAwake || prisonerIsAwake
}

// CanSignalPrisoner pode ser executado se o arqueiro estiver acordado e o prisioneiro estiver acordado.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	return archerIsAwake && prisonerIsAwake
}

// CanReleasePrisoner pode ser executado se o prisioneiro estiver acordado e os outros 2 personagens estiverem dormindo
// ou se o cachorro de Annalyn estiver com ela e o arqueiro estiver dormindo.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, annalynsDogIsPresent bool) bool {
	if annalynsDogIsPresent {
		return !archerIsAwake
	}
	return prisonerIsAwake && !knightIsAwake && !archerIsAwake
}
