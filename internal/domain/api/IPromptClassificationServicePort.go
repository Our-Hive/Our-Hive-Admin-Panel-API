package api

type IPromptClassificationServicePort interface {
	IsEthical(prompt string) (bool, error)
}
