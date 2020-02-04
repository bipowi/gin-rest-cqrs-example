package application

import (
	"mime/multipart"
)

// Update update account by accountID
func (service *Service) Update(
	accountID string,
	email string,
	provider string,
	socialID string,
	password string,
	image *multipart.FileHeader,
	gender string,
	intereste string,
) {
	oldData := service.ReadAccountByID(accountID)
	if oldData == nil {
		return
	}
	hashedPassword, hashedSocialID := getHashedPasswordAndSocialID(password, socialID)
	imageKey := ""
	if image != nil {
		imageKey = service.infrastructure.AWS.S3.Upload(image)
	}
	service.infrastructure.Repository.Save(
		accountID,
		email,
		provider,
		hashedSocialID,
		hashedPassword,
		imageKey,
		gender,
		intereste,
	)
	service.infrastructure.Email.Send([]string{email}, "Account is updated.")
}