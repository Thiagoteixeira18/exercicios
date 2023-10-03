package exercicios

// OvenTime é a constante que representa o tempo necessário para assar a lasanha em minutos.
const OvenTime = 40

// RemainingOvenTime retorna os minutos restantes com base nos minutos "reais" já no forno.
func RemainingOvenTime(actualMinutesInOven int) int {
    return OvenTime - actualMinutesInOven
}

// PreparationTime calcula o tempo necessário para preparar a lasanha com base na quantidade de camadas.
func PreparationTime(numberOfLayers int) int {
    // Supondo que leva 2 minutos para preparar cada camada.
    return 2 * numberOfLayers
}

// ElapsedTime calcula o tempo total decorrido cozinhando a lasanha,
// incluindo o tempo de preparo e o tempo em que a lasanha está assando no forno.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
    preparationTime := PreparationTime(numberOfLayers)
    
    // Apenas retornamos a soma do tempo de preparo com o tempo real no forno.
    return preparationTime + actualMinutesInOven
}
