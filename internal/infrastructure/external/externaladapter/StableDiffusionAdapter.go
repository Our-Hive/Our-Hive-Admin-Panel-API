package externaladapter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/configuration"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/infrastructure/external/externalconstant"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/infrastructure/external/externalerror"
	"io"
	"net/http"
)

type StableDiffusionAdapter struct {
	token string
}

func NewStableDiffusionAdapter() *StableDiffusionAdapter {
	return &StableDiffusionAdapter{token: configuration.HUGGING_FACE_TOKEN}
}

func (s StableDiffusionAdapter) GenerateImage(prompt string) (image []byte, err error) {
	body := map[string]string{
		"inputs": prompt,
	}

	jsonData, err := json.Marshal(body)

	if err != nil {
		return nil, &externalerror.MarshallingError{
			Message: fmt.Sprintf(externalconstant.MarshallingErrorMessage, err),
		}
	}

	req, err := http.NewRequest("POST", externalconstant.StableDiffusionApiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, &externalerror.RequestCreationError{
			Message: fmt.Sprintf(externalconstant.RequestCreationErrorMessage, err),
		}
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, &externalerror.RequestSendingError{
			Message: fmt.Sprintf(externalconstant.RequestSendingErrorMessage, err),
		}
	}

	image, err = io.ReadAll(res.Body)

	err = res.Body.Close()

	if err != nil {
		return nil, &externalerror.ResponseClosingError{
			Message: fmt.Sprintf(externalconstant.ResponseClosingErrorMessage, err),
		}
	}

	return image, nil
}
