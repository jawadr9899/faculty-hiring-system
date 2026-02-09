package types

import (
	"uhs/internal/models"
	"uhs/internal/services"
)

type UserOps services.DatabaseOperations[models.User]
type AnalyticsOps services.DatabaseOperations[models.Analytics]
type PositionOps services.DatabaseOperations[models.Position]
type CvOps services.DatabaseOperations[models.Cv]
type ApplicationOps services.DatabaseOperations[models.Application]