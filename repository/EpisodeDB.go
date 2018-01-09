package repository

import (
	"podcast_api/models"

	"github.com/jinzhu/gorm"
)

type EpisodeDB struct {
	*gorm.DB
}

func getEpisode(email string, name string, ref string) *models.Episode {
	return nil
}

func getAllEpisodes(email string, podcastname string) []models.Episode {
	return []models.Episode{}
}

func AddEpisode(email string, podcastName string) {

}

func UpdateEpisode(episode *models.Episode) {

}
