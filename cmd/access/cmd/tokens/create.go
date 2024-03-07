package tokens

import (
	"context"
	"strconv"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	indentv1 "go.indent.com/indent-go/api/indent/v1"
	"go.indent.com/indent-go/pkg/cliutil"
	"go.indent.com/indent-go/pkg/common"
)

const (
	// defaultServiceAccount is the default service account to use when creating tokens.
	defaultServiceAccount = "access-cli-default-service-account"

	// defaultServiceAccountDisplayName is the default service account display name to use when creating tokens.
	defaultServiceAccountDisplayName = "Default Service Account"

	// defaultExpiryDays is the default number of days until a token expires.
	defaultExpiryDays = 25
)

// NewCreateOptions returns a CreateOptions with the defaults set.
func NewCreateOptions() *CreateOptions {
	return &CreateOptions{
		CreateTokenRequest: new(indentv1.CreateTokenRequest),
	}
}

// CreateOptions defines how a token is created.
type CreateOptions struct {
	*indentv1.CreateTokenRequest

	// CreateAccessToken indicates that an access token should be created.
	CreateAccessToken bool
}

// NewCmdCreate returns a command that creates tokens.
func NewCmdCreate(f cliutil.Factory) *cobra.Command {
	opts := NewCreateOptions()
	cmd := &cobra.Command{
		Use:   "create [service_account]",
		Short: "Creates new refresh tokens",
		Long:  `Creates new refresh tokens for a service account`,
		Run: func(cmd *cobra.Command, args []string) {
			logger := f.Logger()
			client := f.API(cmd.Context()).Accounts()

			opts.ServiceAccountId = serviceAccountID(cmd.Context(), logger, f, client, args)
			logger = logger.With(zap.Uint64("serviceAccountID", opts.ServiceAccountId))

			logger.Debug("Creating token")
			opts.SpaceName = f.Config().Space
			token, err := client.CreateToken(cmd.Context(), opts.CreateTokenRequest)
			if err != nil {
				logger.Fatal("Failed to create token", zap.Error(err))
			}

			logger = logger.With(zap.String("refreshToken", token.GetRefreshToken()))
			logger.Info("Refresh token created")

			if opts.CreateAccessToken {
				logger.Debug("Creating access token")
				tokenClient := f.API(cmd.Context()).Tokens()
				accessToken, err := tokenClient.ExchangeToken(cmd.Context(), &indentv1.ExchangeTokenRequest{
					RefreshToken: token.GetRefreshToken(),
				})
				if err != nil {
					logger.Fatal("Failed to create access token", zap.Error(err))
				}

				logger = logger.With(zap.String("accessToken", accessToken.GetAccessToken()))
				logger.Info("Access token created")
			}
		},
	}

	flags := cmd.Flags()
	flags.Uint64Var(&opts.ExpiryDays, "expiry-days", defaultExpiryDays, "Number of days until token expires")
	flags.BoolVar(&opts.CreateAccessToken, "access-token", false, "Create an access token also")
	return cmd
}

func serviceAccountID(ctx context.Context, logger *zap.Logger, f cliutil.Factory, client indentv1.AccountAPIClient,
	args []string) (svcAccountID uint64) {
	var err error
	if len(args) != 0 {
		if svcAccountID, err = strconv.ParseUint(args[0], common.Base10, common.BitSize64); err != nil {
			logger.Fatal("Failed to parse service account ID", zap.Error(err))
		}
	}

	// if service account is specified, use it
	if svcAccountID != 0 {
		return svcAccountID
	}

	// use default service account if none specified
	logger.Debug("Looking up default service account")
	space := f.Config().Space
	var resp *indentv1.ListServiceAccountsResponse
	if resp, err = client.ListServicesAccounts(ctx, &indentv1.ListServiceAccountsRequest{
		SpaceName: space,
	}); err != nil {
		logger.Fatal("Failed to list service accounts", zap.Error(err))
	}

	for _, serviceAccount := range resp.GetAccounts() {
		meta := serviceAccount.GetMeta()
		if meta.GetSpace() == space && meta.GetName() == defaultServiceAccount {
			return serviceAccount.GetServiceAccountId()
		}
	}

	logger.Debug("Creating default service account")
	var svcAcct *indentv1.ServiceAccount
	if svcAcct, err = client.CreateServiceAccount(ctx, &indentv1.CreateServiceAccountRequest{
		SpaceName:   space,
		Name:        defaultServiceAccount,
		DisplayName: defaultServiceAccountDisplayName,
	}); err != nil {
		logger.Fatal("Failed to create default service account", zap.Error(err))
	}
	svcAccountID = svcAcct.GetServiceAccountId()
	logger.Debug("Default service account created", zap.Uint64("svcAccountID", svcAccountID))
	return svcAccountID
}
