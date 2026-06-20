package Pair_test

import (
	Pair "github.com/go-composites/pair/src"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Pair", func() {

	ginkgo.It("constructor exposes both slots", func() {
		p := Pair.New("key", 42)
		gomega.Expect(p.First()).To(gomega.Equal("key"))
		gomega.Expect(p.Second()).To(gomega.Equal(42))
	})

	ginkgo.It("reports that a real Pair is not null", func() {
		gomega.Expect(Pair.New("a", "b").IsNull()).To(gomega.BeFalse())
	})

	ginkgo.It("renders its Go string form", func() {
		gomega.Expect(Pair.New("k", 7).ToGoString()).To(gomega.Equal("(k, 7)"))
	})

	ginkgo.It("converts to a two-element Array", func() {
		array := Pair.New("k", 7).ToArray()
		gomega.Expect(array.Len()).To(gomega.Equal(2))
		gomega.Expect(array.Fetch(0).Payload()).To(gomega.Equal("k"))
		gomega.Expect(array.Fetch(1).Payload()).To(gomega.Equal(7))
	})

	ginkgo.Context("Equal", func() {
		ginkgo.It("is true for two Pairs with deeply-equal slots", func() {
			gomega.Expect(Pair.New("k", 7).Equal(Pair.New("k", 7))).To(gomega.BeTrue())
		})
		ginkgo.It("is false when the first slot differs", func() {
			gomega.Expect(Pair.New("k", 7).Equal(Pair.New("x", 7))).To(gomega.BeFalse())
		})
		ginkgo.It("is false when the second slot differs", func() {
			gomega.Expect(Pair.New("k", 7).Equal(Pair.New("k", 9))).To(gomega.BeFalse())
		})
		ginkgo.It("is false against the null Pair", func() {
			gomega.Expect(Pair.New("k", 7).Equal(Pair.Null())).To(gomega.BeFalse())
		})
	})

	ginkgo.Context("the Null-Object Pair", func() {
		ginkgo.It("reports that it is null", func() {
			gomega.Expect(Pair.Null().IsNull()).To(gomega.BeTrue())
		})
		ginkgo.It("has nil slots", func() {
			gomega.Expect(Pair.Null().First()).To(gomega.BeNil())
			gomega.Expect(Pair.Null().Second()).To(gomega.BeNil())
		})
		ginkgo.It("equals nothing", func() {
			gomega.Expect(Pair.Null().Equal(Pair.New("k", 7))).To(gomega.BeFalse())
			gomega.Expect(Pair.Null().Equal(Pair.Null())).To(gomega.BeFalse())
		})
		ginkgo.It("converts to the empty Array", func() {
			gomega.Expect(Pair.Null().ToArray().IsEmpty()).To(gomega.BeTrue())
		})
		ginkgo.It("renders as the null literal", func() {
			gomega.Expect(Pair.Null().ToGoString()).To(gomega.Equal("(null)"))
		})
	})
})
