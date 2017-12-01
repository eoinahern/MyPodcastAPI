package repository

import (
	"my_podcast_api/models"

	"github.com/jinzhu/gorm"
)

type PodcastDB struct {
	*gorm.DB
}

func (DB *PodcastDB) getAll() {

}

func (DB *PodcastDB) getPodcast(id string) *models.Podcast {
	return nil
}
