package crypto

import (
	"github.com/lucas-clemente/quic-go/qerr"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cert Manager", func() {
	var cm *CertManager

	BeforeEach(func() {
		cm = &CertManager{}
	})

	It("errors when given invalid data", func() {
		err := cm.SetData([]byte("foobar"))
		Expect(err).To(MatchError(qerr.ProofInvalid))
	})

	It("decompresses a certificate chain", func() {
		cert1 := []byte{0xde, 0xca, 0xfb, 0xad}
		cert2 := []byte{0xde, 0xad, 0xbe, 0xef, 0x13, 0x37}
		chain := [][]byte{cert1, cert2}
		compressed, err := compressChain(chain, nil, nil)
		Expect(err).ToNot(HaveOccurred())
		err = cm.SetData(compressed)
		Expect(err).ToNot(HaveOccurred())
		Expect(cm.chain).To(Equal(chain))
	})

	Context("getting the leaf cert", func() {
		It("gets it", func() {
			cert1 := []byte{0xc1}
			cert2 := []byte{0xc2}
			cm.chain = [][]byte{cert1, cert2}
			leafCert := cm.GetLeafCert()
			Expect(leafCert).To(Equal(cert1))
		})

		It("returns nil if the chain hasn't been set yet", func() {
			leafCert := cm.GetLeafCert()
			Expect(leafCert).To(BeNil())
		})
	})
})