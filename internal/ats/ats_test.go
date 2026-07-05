package ats_test

import (
	"github.com/divya/resume-analyser/internal/ats"
	"github.com/divya/resume-analyser/internal/matcher"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ATS Service", func() {

	var service *ats.Service

	BeforeEach(func() {
		service = ats.NewService()
	})

	Context("when all required skills match", func() {

		It("should return grade A+", func() {

			match := &matcher.MatchResult{
				SkillScore: 95,
			}

			result := service.Calculate(match)

			Expect(result.Score).To(Equal(95))
			Expect(result.Grade).To(Equal("A+"))
		})

	})
	Context("Score is above 95", func() {
		It("shld mark the resume as perfect resume", func() {
			match := &matcher.MatchResult{
				SkillScore: 100,
			}
			result := service.Calculate(match)
			Expect(result.Badge).To(Equal("Perfect Match"))
		})
	})
	Context("Score 82 ", func() {
		It("should return grade A", func() {
			match := &matcher.MatchResult{
				SkillScore: 82,
			}
			result := service.Calculate(match)
			Expect(result.Grade).To(Equal("A"))
		})
	})

	Context("Score 68", func() {
		It("should return C", func() {
			match := &matcher.MatchResult{
				SkillScore: 68,
			}
			result := service.Calculate(match)
			Expect(result.Grade).To(Equal("C"))
		})
	})

	Context("Score 40", func() {
		It("Should return grade F", func() {
			match := &matcher.MatchResult{
				SkillScore: 40,
			}
			result := service.Calculate(match)
			Expect(result.Grade).To(Equal("F"))
		})
	})

})
