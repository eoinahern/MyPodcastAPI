package repository

import "my_podcast_api/models"

type PodcastDB struct {
	*DB
}

func (DB *PodcastDB) getAll() {

}

func (DB *PodcastDB) getPodcast(id string) *models.Podcast {
	return nil
}
