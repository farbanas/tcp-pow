package mappers_test

import (
	"encoding/json"
	"net"
	"tcp-pow/internal/mappers"
	"tcp-pow/internal/models"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMappers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mapper Suite")
}

var _ = Describe("Given a connection mapper", func() {
	Describe("When map to internal data is called", func() {
		var connectionMapper *mappers.ConnectionMapper
		var listener net.Listener
		var conn net.Conn

		BeforeEach(func() {
			connectionMapper = mappers.NewConnectionMapper()

			var err error
			listener, err = net.Listen("tcp", ":8080")
			Expect(err).To(BeNil())
		})

		AfterEach(func() {
			err := listener.Close()
			Expect(err).To(BeNil())
		})

		Context("And reading from connection returns an error", func() {
			var clientConn net.Conn
			BeforeEach(func() {
				var err error
				clientConn, err = net.Dial("tcp", "localhost:8080")
				Expect(err).To(BeNil())

				conn, err = listener.Accept()
				Expect(err).To(BeNil())

				clientConn.Close()
			})

			It("Then mapper should return that connection is closed", func() {
				data, done, err := connectionMapper.MapToInternalData(conn)
				Expect(data).To(BeEquivalentTo(models.TCPData{}))
				Expect(done).To(BeTrue())
				Expect(err).To(BeNil())
			})
		})

		Context("And marshaling to internal data returns an error", func() {
			var clientConn net.Conn
			BeforeEach(func() {
				var err error
				clientConn, err = net.Dial("tcp", "localhost:8080")
				Expect(err).To(BeNil())

				conn, err = listener.Accept()
				Expect(err).To(BeNil())

				clientConn.Write([]byte("this is a test\n"))
			})

			AfterEach(func() {
				clientConn.Close()
			})

			It("Then mapper should return an error", func() {
				data, done, err := connectionMapper.MapToInternalData(conn)
				Expect(data).To(BeEquivalentTo(models.TCPData{}))
				Expect(done).To(BeFalse())
				Expect(err).ToNot(BeNil())
			})
		})

		Context("And everything is fine", func() {
			var clientConn net.Conn
			BeforeEach(func() {
				var err error
				clientConn, err = net.Dial("tcp", "localhost:8080")
				Expect(err).To(BeNil())

				conn, err = listener.Accept()
				Expect(err).To(BeNil())

				data := models.TCPData{
					PackageType: models.SendChallenge,
				}

				dataBytes, err := json.Marshal(data)
				Expect(err).To(BeNil())

				dataBytes = append(dataBytes, '\n')
				clientConn.Write([]byte(dataBytes))
			})

			AfterEach(func() {
				clientConn.Close()
			})

			It("Then mapper should return marshaled data", func() {
				expectedData := models.TCPData{
					PackageType: models.SendChallenge,
				}

				data, done, err := connectionMapper.MapToInternalData(conn)
				Expect(data).To(BeEquivalentTo(expectedData))
				Expect(done).To(BeFalse())
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("When map from internal data is called", func() {
		var connectionMapper *mappers.ConnectionMapper
		var data models.TCPData
		BeforeEach(func() {
			connectionMapper = mappers.NewConnectionMapper()
			data = models.TCPData{
				PackageType: models.RequestChallenge,
			}
		})

		Context("And data can be marshaled to json", func() {

			It("Then mapper return marshaled data", func() {
				expectedResult, err := json.Marshal(data)
				Expect(err).To(BeNil())

				result, err := connectionMapper.MapFromInternalData(data)
				Expect(result).To(BeEquivalentTo(expectedResult))
				Expect(err).To(BeNil())
			})
		})
	})
})
