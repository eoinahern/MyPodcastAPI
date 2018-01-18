package repository

import (
	"podcast_api/models"

	_ "github.com/go-sql-driver/mysql"
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
