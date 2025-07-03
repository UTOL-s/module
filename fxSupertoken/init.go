package fxsupertoken

import (
	"fmt"

	"github.com/supertokens/supertokens-golang/ingredients/emaildelivery"
	"github.com/supertokens/supertokens-golang/recipe/passwordless"
	"github.com/supertokens/supertokens-golang/recipe/passwordless/plessmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
	"go.uber.org/fx"
)

// InitSuperTokens initializes SuperTokens with the provided configuration
func InitSuperTokens(config *SuperTokensConfig) error {
	// Check if SuperTokens is properly configured
	if config.ConnectionURI == "" || config.ConnectionURI == "your_supertokens_connection_uri" {
		fmt.Printf("Warning: SuperTokens connection URI is not configured. SuperTokens authentication will be disabled.\n")
		config.IsInitialized = false
		return nil // Don't fail, just skip initialization
	}

	if config.ConnectionAPIKey == "" || config.ConnectionAPIKey == "your_supertokens_api_key" {
		fmt.Printf("Warning: SuperTokens API key is not configured. SuperTokens authentication will be disabled.\n")
		config.IsInitialized = false
		return nil // Don't fail, just skip initialization
	}

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

	config.IsInitialized = true
	fmt.Printf("SuperTokens initialized successfully\n")
	return nil
}

// FxInit is the fx invoke function for initializing SuperTokens
var FxInit = fx.Invoke(InitSuperTokens)
