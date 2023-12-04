package exercicios

import "fmt"

func main() {
	// Tamanho da caixa de sapato e do sapato em metros
	tamanhoCaixa := 5
	tamanhoSapato := 3

	// Peso inicial da caixa de sapato e do sapato em quilos
	pesoInicialCaixa := tamanhoCaixa
	pesoInicialSapato := tamanhoSapato

	// Número de quilômetros percorridos pelo entregador
	kmPercorridos := 4

	// Cálculo do peso adicional por quilômetro percorrido
	pesoAdicional := kmPercorridos

	// Calcula o peso total da caixa de sapato e do sapato após a entrega
	pesoTotalCaixa := pesoInicialCaixa + pesoAdicional
	pesoTotalSapato := pesoInicialSapato + pesoAdicional

	// Calcula o peso total da caixa de sapato e do sapato juntos
	pesoTotal := pesoTotalCaixa + pesoTotalSapato

	fmt.Printf("Peso total da caixa de sapato após a entrega: %d kg\n", pesoTotalCaixa)
	fmt.Printf("Peso total do sapato após a entrega: %d kg\n", pesoTotalSapato)
	fmt.Printf("Peso total da caixa de sapato e do sapato juntos após a entrega: %d kg\n", pesoTotal)
}
