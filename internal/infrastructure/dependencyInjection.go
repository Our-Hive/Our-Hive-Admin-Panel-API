package infrastructure

import (
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"context"
	firebase "firebase.google.com/go"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/application/handler"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/configuration"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/domain/api/usecase"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/infrastructure/external/externaladapter"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/infrastructure/input/rest/controller"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/infrastructure/output/fbstorage/bucket"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/infrastructure/output/fbstorage/storageadapter"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/infrastructure/output/firestore/firestoreadapter"
	"github.com/Our-Hive/Our-Hive-Admin-Panel-API/internal/infrastructure/output/firestore/repository"
	"google.golang.org/api/option"
)

var ctx = context.Background()
var bucketHandle *storage.BucketHandle
var firestoreClient *firestore.Client

func InitializeFirebase() {
	sa := option.WithCredentialsFile(configuration.FB_ADMIN_SDK_PATH)
	app, err := firebase.NewApp(ctx, nil, sa)

	if err != nil {
		panic(err)
	}

	firestoreClient, err = app.Firestore(ctx)

	if err != nil {
		panic(err)
	}

	storageClient, err := app.Storage(ctx)

	bucketHandle, err = storageClient.Bucket(configuration.FB_STORAGE_BUCKET)

	if err != nil {
		panic(err)
	}

}

func InitializeImageUseCase() *usecase.ImageUseCase {
	imageRepository := repository.NewImageFireStoreRepository(firestoreClient, ctx)
	imagePersistenceAdapter := firestoreadapter.NewImagePersistenceAdapter(imageRepository)

	return usecase.NewImageUseCase(imagePersistenceAdapter)
}

func InitializeUploadUseCase() *usecase.UploadUseCase {
	imageUseCase := InitializeImageUseCase()
	imageDataUseCase := InitializeImageDataUseCase()

	return usecase.NewUploadUseCase(imageUseCase, imageDataUseCase)
}

func InitializeImageDataUseCase() *usecase.ImageDataUseCase {
	imageBucket := bucket.NewImageBucket(bucketHandle, ctx)
	imageStorageAdapter := storageadapter.NewImageStorageAdapter(imageBucket)

	return usecase.NewImageDataUseCase(imageStorageAdapter)
}

func InitializeGenerationController() *controller.GenerationController {
	imageUseCase := InitializeImageUseCase()

	imageDataUseCase := InitializeImageDataUseCase()

	huggingFaceClassificationAdapter := externaladapter.NewHuggingFaceClassificationAdapter()
	stableDiffusionAdapter := externaladapter.NewStableDiffusionAdapter()

	generationUseCase := usecase.NewGenerationUseCase(imageUseCase, imageDataUseCase, huggingFaceClassificationAdapter, stableDiffusionAdapter)
	generarationHandler := handler.NewGenerationHandler(generationUseCase)

	return controller.NewGenerationController(generarationHandler)
}

func InitializeImageController() *controller.ImageController {
	imageUseCase := InitializeImageUseCase()
	imageHandler := handler.NewImageHandler(imageUseCase)
	uploadUseCase := InitializeUploadUseCase()
	uploadHandler := handler.NewUploadHandler(uploadUseCase)

	return controller.NewImageController(imageHandler, uploadHandler)
}

func InitializeContactLineController() *controller.ContactLineController {
	contactLineRepository := repository.NewContactLineRepository(firestoreClient, ctx)
	contactLinePersistenceAdapter := firestoreadapter.NewContactLinePersistenceAdapter(contactLineRepository)
	contactLineUseCase := usecase.NewContactLineUseCase(contactLinePersistenceAdapter)
	contactLineHandler := handler.NewContactLineHandler(contactLineUseCase)

	return controller.NewContactLineController(contactLineHandler)
}

func InitializeRecommendedContentController() *controller.RecommendedContentController {
	recommendedContentRepository := repository.NewRecommendedContentRepository(firestoreClient, ctx)
	recommendedContentPersistenceAdapter := firestoreadapter.NewDigitalContentPersistenceAdapter(recommendedContentRepository)
	recommendedContentUseCase := usecase.NewDigitalContentUseCase(recommendedContentPersistenceAdapter)
	recommendedContentHandler := handler.NewRecommendedContentHandler(recommendedContentUseCase)

	return controller.NewRecommendedContentController(recommendedContentHandler)
}
