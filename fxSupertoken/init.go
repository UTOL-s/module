package fxsupertoken

import (
	"github.com/supertokens/supertokens-golang/ingredients/emaildelivery"
	"github.com/supertokens/supertokens-golang/recipe/passwordless"
	"github.com/supertokens/supertokens-golang/recipe/passwordless/plessmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
	"go.uber.org/fx"
)

// InitSuperTokens initializes SuperTokens with the provided configuration
func InitSuperTokens(config *SuperTokensConfig) error {
	supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: config.ConnectionURI,
			APIKey:        config.ConnectionAPIKey,
		},
		AppInfo: supertokens.AppInfo{
			AppName:         config.AppName,
			APIDomain:       config.APIDomain,
			WebsiteDomain:   config.WebsiteDomain,
			APIBasePath:     &config.APIBasePath,
			WebsiteBasePath: &config.WebBasePath,
		},
		RecipeList: []supertokens.Recipe{
			passwordless.Init(
				plessmodels.TypeInput{
					FlowType: "USER_INPUT_CODE_AND_MAGIC_LINK",
					ContactMethodEmailOrPhone: plessmodels.ContactMethodEmailOrPhoneConfig{
						Enabled: true,
					},

					//Override: &plessmodels.OverrideStruct{
					//	Functions: OverRideSignIn,
					//	APIs:      nil,
					//},

					EmailDelivery: &emaildelivery.TypeInput{

						Service: passwordless.MakeSMTPService(emaildelivery.SMTPServiceConfig{

							Settings: emaildelivery.SMTPSettings{
								Host: config.EmailHost,
								From: emaildelivery.SMTPFrom{
									Name:  "OTP",
									Email: config.Email,
								},
								Port: 465,
								//Username: &smtpUsername, // this is optional. In case not given, from.email will be used
								Username: &config.Email,
								Password: config.EmailPassword,
								Secure:   false,
							},
						}),
					},
				},
			),
		},
	})

	return nil
}

// FxInit is the fx invoke function for initializing SuperTokens
var FxInit = fx.Invoke(InitSuperTokens)
