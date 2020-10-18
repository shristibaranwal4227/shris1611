package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllMeetings(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	meetings := []model.Project{}
	db.Find(&projects)
	respondJSON(w, http.StatusOK, projects)
}

func CreateMeeting(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	meeting := model.Project{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&project); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&project).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, project)
}

func GetMeeting(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	meeting := getMeetingOr404(db, title, w, r)
	if meeting == nil {
		return
	}
	respondJSON(w, http.StatusOK, project)
}

func UpdateMeeting(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	meeting := getMeetingOr404(db, title, w, r)
	if meeting == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&meeting); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&meeting).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, project)
}

func DeleteMeeting(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	meeting := getMeetingOr404(db, title, w, r)
	if meeting == nil {
		return
	}
	if err := db.Delete(&meeting).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func ArchiveProject(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	meeting := getMeetingOr404(db, title, w, r)
	if meeting == nil {
		return
	}
	meeting.Archive()
	if err := db.Save(&meeting).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, project)
}

func RestoreMeeting(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	title := vars["title"]
	project := getMeetingOr404(db, title, w, r)
	if meeting == nil {
		return
	}
	meeting.Restore()
	if err := db.Save(&meeting).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, project)
}

// getProjectOr404 gets a project instance if exists, or respond the 404 error otherwise
func getMeetingOr404(db *gorm.DB, title string, w http.ResponseWriter, r *http.Request) *model.Project {
	meeting := model.Meeting{}
	if err := db.First(&meeting, model.Meeting{Title: title}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &meeting
}

