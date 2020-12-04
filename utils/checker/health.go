package health

import (
	"encoding/json"
	"jobs-api/utils/env"
	"log"
	"net/http"
	"time"
)

var version = "1.0" //todo env
var description = "jobs-api"

//Health отображает информацию для мониторинга
type Health struct {
	debug string
}

//Create Создание инстанса
func Create() (health *Health) {
	health = new(Health)
	health.debug = env.AppDebugMode()
	return
}

type data struct {
	Version       string `json:"version"`
	Description   string `json:"description"`
	IsDebug       string `json:"debug"`
	LocalDateTime string `json:"localDateTime"`
	UTCDateTime   string `json:"utcDateTime"`
}

// Response data
// swagger:response checkHealth
type Response struct {
	Status bool
	Data   data
}

// Get returns service health
// swagger:route GET /health checkHealth
//
// Check health of the service
//     Responses:
//       200: checkHealth
func (svc Health) Get(w http.ResponseWriter, r *http.Request) {
	var resp Response
	hd := data{version,
		description,
		svc.debug,
		time.Now().Local().Format("02.01.2006 15:04:05"),
		time.Now().UTC().Format("02.01.2006 15:04:05")}
	resp.Status = true
	resp.Data = hd
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Printf("Error in health.Get method: %s", err)
	}
}
