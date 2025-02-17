package folders

import (
	"encoding/json"
	"media_batch_converter/backend/internal/services"
	"net/http"

	"github.com/sirupsen/logrus"
)

// GetFolders
// @Tags         folders
// @Summary      Get all folders.
// @Description  Retrieve the list of available folders.
// @Success      200  {array}  models.Folder
// @Failure      500  "Something went wrong"
// @Router       /folders [get]
func GetFolders(w http.ResponseWriter, r *http.Request) {
	allFolders, err := services.GetAllFolders()
	if err != nil {
		logrus.Errorf("error fetching folders: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(allFolders)
	_, _ = w.Write(body)
}
