package devoxxfr2018

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
)

func TestConference(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("junit.xml")
	//junitReporter := reporters.NewJUnitReporter(fmt.Sprintf("junit_%d.xml", config.GinkgoConfig.ParallelNode))
	RunSpecsWithDefaultAndCustomReporters(t, "Conference CFP Test Suite", []Reporter{junitReporter})
}

var _ = Describe("Conference", func() {
	Context("initially", func() {
		conf := Conference{}
		It("has 0 talk", func() {
			Expect(conf.TalkCount()).Should(BeZero())
		})
		It("has 0 speaker", func() {
			Expect(conf.SpeakerCount()).Should(BeZero())
		})
	})

	Context("when a speaker is added", func() {

		speaker := NewSpeaker("Jean-Christophe", "Sirot", "jc@devoxx.fr", "Docker")
		var conf Conference

		BeforeEach(func() {
			conf = Conference{}
			conf.AddSpeaker(speaker)
		})

		It("has 1 speaker", func() {
			Expect(conf.SpeakerCount()).Should(Equal(1))
		})

		It("should return the speaker by ID", func() {
			actual, err := conf.GetSpeaker(speaker.ID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(*actual).Should(Equal(speaker))
		})

		It("should fail when getting an unknown speaker", func() {
			actual, err := conf.GetSpeaker("foobar")
			Expect(err).Should(HaveOccurred())
			Expect(actual).Should(BeNil())
		})

		It("should return 0 speaker when removed", func() {
			conf.RemoveSpeaker(speaker.ID)
			_, err := conf.GetSpeaker("foobar")
			Expect(err).Should(HaveOccurred())
			Expect(conf.SpeakerCount()).Should(BeZero())
		})
	})

	Context("when some talks are added", func() {

		var conf Conference
		var id string

		BeforeEach(func() {
			conf = Conference{}
			By("Adding 2 speakers")
			speaker1 := NewSpeaker("Jean-Christophe", "Sirot", "jc@devoxx.fr", "Docker")
			speaker2 := NewSpeaker("Charles", "Sabourdin", "charles@devoxx.fr", "Indep")
			conf.AddSpeaker(speaker1)
			conf.AddSpeaker(speaker2)
			By("Adding 2 talks")
			talk1 := NewTalk("Tester le code Go avec Ginkgo et Gomega", "Des trucs en Go", speaker1)
			talk2 := NewTalk("Java dans Docker : Bonnes pratiques", "Des trucs en Java avec Docker", speaker1, speaker2)
			conf.AddTalk(talk1)
			conf.AddTalk(talk2)
			id = speaker1.ID
		})

		It("has 2 talks", func() {
			Expect(conf.TalkCount()).Should(Equal(2))
		})

		It("cannot remove a speaker", func() {
			err := conf.RemoveSpeaker(id)
			Expect(err).Should(HaveOccurred())
		})
	})

	Context("when a speaker is asynchronously added", func() {

		conf := Conference{}
		speaker := NewSpeaker("Jean-Christophe", "Sirot", "jc@devoxx.fr", "Docker")

		It("has 1 speaker after some time", func() {
			conf.AddSpeakerAsync(speaker)
			Eventually(func() int {
				return conf.SpeakerCount()
			}, 10*time.Second, 1*time.Second).Should(Equal(1))
		})
	})
})
