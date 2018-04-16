package devoxxfr2018

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Conference contains speakers and talks
type Conference struct {
	ID       string
	Speakers map[string]Speaker
	Talks    map[string]Talk
}

// Talk represents any talk given at Devoxx
type Talk struct {
	ID          string
	Title       string
	Description string
	Speakers    []string
}

// Speaker represents any Devoxx speaker
type Speaker struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Company   string
}

func (conf *Conference) init() {
	if conf.Speakers == nil {
		conf.Speakers = map[string]Speaker{}
	}
	if conf.Talks == nil {
		conf.Talks = map[string]Talk{}
	}
}

func NewSpeaker(firstName, lastName, email, company string) Speaker {
	return Speaker{
		ID:        uuid.New().String(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Company:   company,
	}
}

func NewTalk(title, description string, speakers ...Speaker) Talk {
	ids := make([]string, len(speakers))
	for i, v := range speakers {
		ids[i] = v.ID
	}
	return Talk{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
		Speakers:    ids,
	}
}

func (conf *Conference) AddSpeaker(speaker Speaker) {
	conf.init()
	if _, ok := conf.Speakers[speaker.ID]; !ok {
		conf.Speakers[speaker.ID] = speaker
	}
}

func (conf *Conference) AddSpeakerAsync(speaker Speaker) <-chan bool {
	conf.init()
	c := make(chan bool, 0)
	go func() {
		defer close(c)
		time.Sleep(5 * time.Second)
		if _, ok := conf.Speakers[speaker.ID]; !ok {
			conf.Speakers[speaker.ID] = speaker
			c <- true
		}
	}()
	return c
}

func (conf *Conference) RemoveSpeaker(id string) error {
	conf.init()
	for _, talk := range conf.Talks {
		for _, sid := range talk.Speakers {
			if sid == id {
				return fmt.Errorf("Speaker with id %s cannot be removed because it is assigned to a talk", id)
			}
		}
	}
	if _, ok := conf.Speakers[id]; ok {
		delete(conf.Speakers, id)
	}
	return nil
}

func (conf *Conference) GetSpeaker(id string) (*Speaker, error) {
	conf.init()
	if _, ok := conf.Speakers[id]; ok {
		speaker := conf.Speakers[id]
		return &speaker, nil
	}
	return nil, fmt.Errorf("Speaker with id %s not found", id)
}

func (conf *Conference) AddTalk(talk Talk) {
	conf.init()
	if _, ok := conf.Talks[talk.ID]; !ok {
		conf.Talks[talk.ID] = talk
	}
}

func (conf *Conference) RemoveTalk(id string) {
	conf.init()
	if _, ok := conf.Talks[id]; ok {
		delete(conf.Talks, id)
	}
}

func (conf *Conference) getSpeakers(id string) []Speaker {
	conf.init()
	if _, ok := conf.Talks[id]; ok {
		speakers := make([]Speaker, len(conf.Talks[id].Speakers))
		for _, sid := range conf.Talks[id].Speakers {
			speakers = append(speakers, conf.Speakers[sid])
		}
		return speakers
	}
	return nil
}

func (conf *Conference) TalkCount() int {
	conf.init()
	return len(conf.Talks)
}

func (conf *Conference) SpeakerCount() int {
	conf.init()
	return len(conf.Speakers)
}
