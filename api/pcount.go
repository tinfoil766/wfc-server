package api

import (
	"net/http"
	"time"
	"wwfc/database"

	"github.com/linkdata/deadlock"
)

type PCountRequest struct{}

type PCountResponse struct {
	Count   int
	Success bool
	Error   string
}

var PCountRoute = MakeRouteSpec[PCountRequest, PCountResponse](
	false,
	"/api/pcount",
	HandlePCount,
	http.MethodGet,
)

var (
	mu        = deadlock.Mutex{}
	lastQuery = time.Now()
	lastCount = -1
)

func HandlePCount(_ any, _ bool, r *http.Request) (any, int, error) {
	res := PCountResponse{}

	mu.Lock()
	defer mu.Unlock()

	// If it is less than 10 seconds since the last query, do not recalculate
	if lastCount != -1 && time.Since(lastQuery) < time.Duration(10*1e9) {
		res.Count = lastCount
		return res, http.StatusOK, nil
	}

	count, err := database.CountTotalUsers(pool, ctx)
	lastQuery = time.Now()

	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	res.Count = count
	lastCount = count
	return res, http.StatusOK, nil
}
