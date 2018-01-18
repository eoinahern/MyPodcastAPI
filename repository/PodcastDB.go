package repository

import (
	"fmt"
	"log"
	"my_podcast_api/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type PodcastDB struct {
	*gorm.DB
}

func (DB *PodcastDB) GetAll() []models.SecurePodcast {

	var podcasts []models.SecurePodcast
	rows, err := DB.Raw("SELECT podcast_id, icon, name, episode_num from podcasts").Rows()

	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var pod models.SecurePodcast
		rows.Scan(&pod.PodcastID, &pod.Icon, &pod.Name, &pod.EpisodeNum)
		podcasts = append(podcasts, pod)
	}

	return podcasts
}

func (DB *PodcastDB) GetPodcast(userName string, podcastName string) *models.Podcast {
	var podcast models.Podcast
	DB.Where("user_email = ? AND name = ?", userName, podcastName).First(&podcast)
	return &podcast
}

func (DB *PodcastDB) CheckPodcastCreated(userName string) models.Podcast {

	var podcast models.Podcast
	DB.Where("user_email = ?", userName).First(&podcast)

	return podcast

}

func (DB *PodcastDB) UpdatePodcast(podcast models.Podcast) bool {
	return true
}

func (DB *PodcastDB) CreatePodcast(podcast models.Podcast) error {

	db := DB.Save(&podcast)

	if db.Error != nil {
		fmt.Println(db.Error)
		log.Println(db.Error)
	}

	return db.Error
}
