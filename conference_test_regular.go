package devoxxfr2018

import (
	"fmt"
	"testing"
)

func TestConferenceInitiallyHas0Talks(t *testing.T) {
	conf := Conference{}
	actual := conf.TalkCount()
	if actual != 0 {
		t.Fatal(fmt.Sprintf("Expected 0 talks, got %d.", actual))
	}
}

func TestConferenceInitiallyHas0Speakers(t *testing.T) {
	conf := Conference{}
	actual := conf.SpeakerCount()
	if actual != 0 {
		t.Fatal(fmt.Sprintf("Expected 0 speakers, got %d.", actual))
	}
}

func TestConferenceWhenASpeakerIsAddedHas1Speaker(t *testing.T) {
	conf := Conference{}
	speaker := NewSpeaker("Jean-Christophe", "Sirot", "jc@devoxx.fr", "Docker")
	conf.AddSpeaker(speaker)
	actual := conf.SpeakerCount()
	if actual != 1 {
		t.Fatal(fmt.Sprintf("Expected 1 speakers, got %d.", actual))
	}
}

func TestConferenceWhenASpeakerIsAddedShouldReturnTheSpeakerByID(t *testing.T) {
	conf := Conference{}
	speaker := NewSpeaker("Jean-Christophe", "Sirot", "jc@devoxx.fr", "Docker")
	conf.AddSpeaker(speaker)
	actual, err := conf.GetSpeaker(speaker.ID)
	if err != nil {
		t.Fatal(err.Error())
	}
	if *actual != speaker {
		t.Fatal(fmt.Sprintf("Actual speaker %v do not match expected %v.", *actual, speaker))
	}
}

func TestConferenceFailedWhenGettingAnUnknownSpeaker(t *testing.T) {
	conf := Conference{}
	speaker := NewSpeaker("Jean-Christophe", "Sirot", "jc@devoxx.fr", "Docker")
	conf.AddSpeaker(speaker)
	_, err := conf.GetSpeaker("foobar")
	if err == nil {
		t.Fatal("No error returned but one error was expeted")
	}
}

func TestConferenceWhenASpeakerIsAddedAndRemovedShouldReturn0Speaker(t *testing.T) {
	conf := Conference{}
	speaker := NewSpeaker("Jean-Christophe", "Sirot", "jc@devoxx.fr", "Docker")
	conf.AddSpeaker(speaker)
	conf.RemoveSpeaker(speaker.ID)
	_, err := conf.GetSpeaker(speaker.ID)
	if err == nil {
		t.Fatal("No error returned but one error was expeted")
	}
	actual := conf.SpeakerCount()
	if actual != 0 {
		t.Fatal(fmt.Sprintf("Expected 0 speakers, got %d.", actual))
	}
}
