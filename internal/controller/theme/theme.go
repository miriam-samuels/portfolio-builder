package theme

import (
	"net/http"

	"github.com/miriam-samuels/portfolio-builder/internal/db"
	"github.com/miriam-samuels/portfolio-builder/internal/helper"
	"github.com/miriam-samuels/portfolio-builder/internal/models/theme"
)

func GetThemes(w http.ResponseWriter, r *http.Request) {
	var themes []theme.Theme

	rows, err := db.Portfolio.Query("SELECT * FROM themes")
	if err != nil {
		helper.SendResponse(w, http.StatusInternalServerError, false, "Internal Server Error", nil)
		return
	}

	for rows.Next() {
		var theme theme.Theme
		rows.Scan(&theme.Id, &theme.Name, &theme.Image)
		themes = append(themes, theme)
	}

	helper.SendResponse(w,http.StatusOK,true,"Request Successful",themes)
}
