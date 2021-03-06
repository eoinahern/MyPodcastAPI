package repository

import (
	"log"
	"my_podcast_api/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type EpisodeDB struct {
	*gorm.DB
}

func getEpisode(email string, name string, ref string) *models.Episode {
	return nil
}

func (DB *EpisodeDB) GetAllEpisodes(podcastid int) []models.Episode {

	var episodes []models.Episode
	DB.Where("pod_id = ?", podcastid).Find(&episodes)
	return episodes
}

func (DB *EpisodeDB) AddEpisode(episode models.Episode) error {

	db := DB.Save(&episode)

	if db.Error != nil {
		log.Println(db.Error)
	}

	return db.Error

}

func (DB *EpisodeDB) GetLastEpisode() models.Episode {

	var episode models.Episode
	DB.Last(&episode)
	return episode
}

func UpdateEpisode(episode *models.Episode) {

}
