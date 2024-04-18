package externaladapter

import (
	"context"
	"fmt"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/infrastructure/external/externalconstant"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/infrastructure/external/externalerror"
	"github.com/hupe1980/go-huggingface"
)

type HuggingFaceClassificationAdapter struct {
	inferenceClient *huggingface.InferenceClient
}

func (h HuggingFaceClassificationAdapter) IsEthical(prompt string) (bool, error) {
	res, err := h.inferenceClient.ZeroShotClassification(context.Background(), &huggingface.ZeroShotClassificationRequest{
		Inputs: []string{prompt},
		Parameters: huggingface.ZeroShotClassificationParameters{
			CandidateLabels: []string{"positive", "negative", "neutral"},
		},
	})

	if err != nil {
		return false, &externalerror.ClassificationError{
			Message: fmt.Sprintf(externalconstant.ClassificationErrorMessage, err),
		}
	}

	labelIndices := make(map[string]uint, len(res[0].Labels))
	for i, label := range res[0].Labels {
		labelIndices[label] = uint(i)
	}

	score, ok := labelIndices["negative"]

	if ok && res[0].Scores[score] > externalconstant.EthicalThreshold {
		return false, nil
	}

	return true, nil
}
