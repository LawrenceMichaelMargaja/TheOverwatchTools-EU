package dependencies

import (
	"fmt"
	"github.com/dembygenesis/local.tools/internal/config"
	"github.com/dembygenesis/local.tools/internal/logic_handlers/authlogic"
	"github.com/dembygenesis/local.tools/internal/logic_handlers/capturepageslogic"
	"github.com/dembygenesis/local.tools/internal/persistence/database_helpers/mysql/mysqlconn"
	"github.com/dembygenesis/local.tools/internal/persistence/persistors/mysqlstore"
	"github.com/sarulabs/dingo/v4"
	"github.com/sirupsen/logrus"
)

const (
	logicAuth            = "logic_auth"
	logicCapturePages    = "logic_capture_pages"
	logicCapturePageSets = "logic_capture_pages_sets"
)

func GetLogicHandlers() []dingo.Def {
	return []dingo.Def{
		{
			Name: logicCapturePageSets,
			Build: func(
				cfg *config.App,
				logger *logrus.Entry,
				txProvider *mysqlconn.Provider,
				store *mysqlstore.Repository,
			) (*capturepageslogic.Service, error) {
				logic, err := capturepageslogic.New(&capturepageslogic.Config{
					TxProvider: txProvider,
					Logger:     logger,
					Persistor:  store,
				})
				if err != nil {
					return nil, fmt.Errorf("logiccapturepagesset: %v", err)
				}
				return logic, nil
			},
		},
		{
			Name: logicCapturePages,
			Build: func(
				cfg *config.App,
				logger *logrus.Entry,
				txProvider *mysqlconn.Provider,
				store *mysqlstore.Repository,
			) (*capturepageslogic.Service, error) {
				logic, err := capturepageslogic.New(&capturepageslogic.Config{
					TxProvider: txProvider,
					Logger:     logger,
					Persistor:  store,
				})
				if err != nil {
					return nil, fmt.Errorf("logiccapturepages: %v", err)
				}
				return logic, nil
			},
		},
		{
			Name: logicAuth,
			Build: func(
				cfg *config.App,
				logger *logrus.Entry,
				txProvider *mysqlconn.Provider,
			) (*authlogic.Impl, error) {
				logic, err := authlogic.New(&authlogic.Config{
					TxProvider: txProvider,
					Logger:     logger,
				})
				if err != nil {
					return nil, fmt.Errorf("logicauth: %v", err)
				}
				return logic, nil
			},
		},
	}
}
