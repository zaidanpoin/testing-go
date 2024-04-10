package kalkulator

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Kalkulator", func() {
	Describe("ketika dua bilangan dijumlahkan", func() {
		It("harus menghasilkan jumlah yang benar", func() {

			Expect(HitungNilai(1, 2)).To(Equal(3))
		})
	})

})
