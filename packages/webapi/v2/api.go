package v2

import (
	"github.com/pangpanglabs/echoswagger/v2"

	"github.com/iotaledger/hive.go/core/configuration"
	"github.com/iotaledger/wasp/packages/webapi/v2/controllers/chain"

	loggerpkg "github.com/iotaledger/hive.go/core/logger"
	"github.com/iotaledger/wasp/packages/chains"
	metricspkg "github.com/iotaledger/wasp/packages/metrics"
	"github.com/iotaledger/wasp/packages/peering"
	"github.com/iotaledger/wasp/packages/registry"
	walpkg "github.com/iotaledger/wasp/packages/wal"
	"github.com/iotaledger/wasp/packages/webapi/v2/controllers"
	"github.com/iotaledger/wasp/packages/webapi/v2/interfaces"
	"github.com/iotaledger/wasp/packages/webapi/v2/services"
)

func loadControllers(server echoswagger.ApiRoot, mocker *Mocker, registryProvider registry.Provider, controllersToLoad []interfaces.APIController) {
	/*claimValidator := func(claims *authentication.WaspClaims) bool {
		// The API will be accessible if the token has an 'API' claim
		return claims.HasPermission(permissions.API)
	}*/

	for _, controller := range controllersToLoad {
		controller.RegisterExampleData(mocker)

		publicGroup := server.Group(controller.Name(), "v2")

		controller.RegisterPublic(publicGroup, mocker)

		adminGroup := server.Group(controller.Name(), "v2").
			SetSecurity("Authorization")

		/*authentication.AddAuthentication(adminGroup.EchoGroup(), registryProvider, authentication.AuthConfiguration{
			Scheme: authentication.AuthJWT,
			JWTConfig: authentication.JWTAuthConfiguration{
				Duration: 24 * time.Hour,
			},
		}, claimValidator)
		*/
		controller.RegisterAdmin(adminGroup, mocker)
	}
}

func Init(logger *loggerpkg.Logger,
	server echoswagger.ApiRoot,
	config *configuration.Configuration,
	chainsProvider chains.Provider,
	metrics *metricspkg.Metrics,
	networkProvider peering.NetworkProvider,
	registryProvider registry.Provider,
	wal *walpkg.WAL,
) {
	mocker := NewMocker()
	mocker.LoadMockFiles()

	// Add dependency injection here
	vmService := services.NewVMService(logger, chainsProvider)
	chainService := services.NewChainService(logger, chainsProvider, metrics, registryProvider, vmService, wal)
	nodeService := services.NewNodeService(logger, networkProvider, registryProvider)
	registryService := services.NewRegistryService(logger, chainsProvider, registryProvider)
	offLedgerService := services.NewOffLedgerService(logger, chainService, nodeService)

	controllersToLoad := []interfaces.APIController{
		chain.NewChainController(logger, chainService, nodeService, offLedgerService, registryService, vmService),
		controllers.NewInfoController(logger, config),
	}

	loadControllers(server, mocker, registryProvider, controllersToLoad)
}