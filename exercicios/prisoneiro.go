package exercicios

// CanFastAttack só pode ser executada quando o cavaleiro está dormindo, pois ele está vulnerável enquanto coloca a armadura.
func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
}

// CanSpy pode ser executada se pelo menos um dos personagens estiver acordado.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	return knightIsAwake || archerIsAwake || prisonerIsAwake
}

// CanSignalPrisoner pode ser executada se o arqueiro estiver acordado e o prisioneiro estiver acordado.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	return !archerIsAwake && prisonerIsAwake
}

// CanFreePrisoner pode ser executada se o prisioneiro estiver acordado e os outros 2 personagens estiverem dormindo
// ou se o cachorro de Annalyn estiver presente e o arqueiro estiver dormindo.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, annalynsDogIsPresent bool) bool {
	if annalynsDogIsPresent {
		return !archerIsAwake
	}
	return prisonerIsAwake && !knightIsAwake && !archerIsAwake
}