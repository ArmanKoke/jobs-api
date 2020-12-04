package utils

import (
	"io/ioutil"
	"jobs-api/utils/logger"
	"net/http"

	"go.uber.org/zap"
)

// Swagger struct
type Swagger struct {
	Log logger.LogEnv
}

// GetSwagger returns json
func (s *Swagger) GetSwagger(w http.ResponseWriter, r *http.Request) {
	log := s.Log.ZapLogger()
	defer log.Sync()

	swaggerSpec, err := ioutil.ReadFile("./swagger.json")
	if err != nil {
		log.Error("Can not load Swagger specification:", zap.Error(err))
	}

	if swaggerSpec == nil || len(swaggerSpec) == 0 {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("Swagger specification is not ready yet. Check the service log for warnings."))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if r.Method != "HEAD" {
		_, _ = w.Write(swaggerSpec)
	}
}
