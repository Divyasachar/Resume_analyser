package upload

import (
	"errors"
	"mime/multipart"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Upload Service", func() {

	var storage *MockStorage

	var service *Service

	BeforeEach(func() {

		storage = &MockStorage{}

		service = NewService(storage)
	})
	Context("when storage saves successfully", func() {

		It("returns the upload response", func() {

			storage.Path = "uploads/resume.pdf"

			file := &multipart.FileHeader{
				Filename: "resume.pdf",
				Size:     1000,
			}

			response, err := service.Save(file)

			Expect(err).To(BeNil())

			Expect(storage.SaveCalled).To(BeTrue())
			Expect(storage.LastFile.Filename).To(Equal("resume.pdf"))
			Expect(response.Path).To(Equal("uploads/resume.pdf"))
		})

	})
	Context("when storage returns an error", func() {

		It("returns an error", func() {

			storage.Err = errors.New("disk full")

			file := &multipart.FileHeader{
				Filename: "resume.pdf",
			}

			_, err := service.Save(file)

			Expect(err).NotTo(BeNil())
		})

	})
	Context("When I save twice", func() {
		It("should return save count as 2", func() {
			file := &multipart.FileHeader{
				Filename: "resume.pdf",
			}
			storage.Save(file)
			_, err := storage.Save(file)
			Expect(err).To(BeNil())
			Expect(storage.SaveCount).To(Equal(2))
		})
	})
})

