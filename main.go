package main

import (
	"emailn/internal/domain/campaign"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/rs/xid"
)

func main() {

	contacts := []campaign.Contact{
		{Email: "test@example.com"},
	}
	content := "This is a test campaign content."
	name := "Test Campaign"
	createdOn := time.Now()
	id := xid.New().String()

	campaign := campaign.Campaign{
		Contacts:  contacts,
		Content:   content,
		Name:      name,
		CreatedOn: createdOn,
		ID:        id,
	}
	validate := validator.New()
	err := validate.Struct(campaign)
	if err == nil {
		println("Validation passed")
	} else {
		validationErrors := err.(validator.ValidationErrors)
		for _, v := range validationErrors {
			switch v.Tag() {
			case "required":
				println(v.StructField() + " is required")
			case "min":
				println(v.StructField() + " does not meet the minimum length")
			case "max":
				println(v.StructField() + " exceeds the maximum length")
			case "email":
				println(v.StructField() + " is not a valid email")
			default:
				println(v.StructField() + " is invalid: " + v.Tag())
			}
		}
	}

}
