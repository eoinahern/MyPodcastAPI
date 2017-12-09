package repository

import (
	"my_podcast_api/models"

	"github.com/jinzhu/gorm"
)

type PodcastDB struct {
	*gorm.DB
}

func (DB *PodcastDB) GetAll() []models.Podcast {

	var podcasts []models.Podcast
	DB.Find(&podcasts)

	return podcasts
}

func (DB *PodcastDB) GetPodcast(userName string, podcastName string) *models.Podcast {
	var podcast models.Podcast
	DB.Where("user_email = ? AND name = ?", userName, podcastName).First(&podcast)
	return &podcast
}

func (DB *PodcastDB) CheckPodcastUserName(userName string, podcastName string) bool {

	var podcast models.Podcast
	DB.Where("user_email = ? AND name = ? ", userName, podcastName).First(&podcast)

	if len(podcast.Name) == 0 {
		return true
	} else {
		return false
	}

}

func (DB *PodcastDB) UpdatePodcast(podcast models.Podcast) bool {
	return true
}
