package label_test

import (
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/synnaxlabs/synnax/pkg/distribution/ontology"
	"github.com/synnaxlabs/synnax/pkg/distribution/ontology/group"
	"github.com/synnaxlabs/x/color"
	"github.com/synnaxlabs/x/gorp"
	"github.com/synnaxlabs/x/kv/memkv"
	"github.com/synnaxlabs/x/query"
	. "github.com/synnaxlabs/x/testutil"

	"github.com/synnaxlabs/synnax/pkg/label"
)

var _ = Describe("Label", Ordered, func() {
	var (
		db  *gorp.DB
		svc *label.Service
		w   label.Writer
		tx  gorp.Tx
	)
	BeforeAll(func() {
		db = gorp.Wrap(memkv.New())
		otg := MustSucceed(ontology.Open(ctx, ontology.Config{DB: db}))
		g := MustSucceed(group.OpenService(group.Config{DB: db, Ontology: otg}))
		svc = MustSucceed(label.OpenService(ctx, label.Config{
			DB:       db,
			Ontology: otg,
			Group:    g,
		}))
	})
	AfterAll(func() {
		Expect(db.Close()).To(Succeed())
	})
	BeforeEach(func() {
		tx = db.OpenTx()
		w = svc.NewWriter(tx)
	})
	AfterEach(func() {
		Expect(tx.Close()).To(Succeed())
	})
	Describe("Create", func() {
		It("Should create a new label", func() {
			l := &label.Label{
				Name:  "Label",
				Color: color.Color("#000000"),
			}
			Expect(w.Create(ctx, l)).To(Succeed())
			Expect(l.Key).ToNot(Equal(uuid.Nil))
		})
	})
	Describe("Delete", func() {
		It("Should delete a label", func() {
			l := &label.Label{
				Name:  "Label",
				Color: color.Color("#000000"),
			}
			Expect(w.Create(ctx, l)).To(Succeed())
			Expect(w.Delete(ctx, l.Key)).To(Succeed())
			Expect(svc.NewRetrieve().WhereKeys(l.Key).Exec(ctx, nil)).To(MatchError(query.NotFound))
		})
	})
	Describe("Label", func() {
		It("Should get a label", func() {
			l := &label.Label{
				Name:  "Label",
				Color: color.Color("#000000"),
			}
			Expect(w.Create(ctx, l)).To(Succeed())
			labeled := &label.Label{
				Name:  "Labeled",
				Color: color.Color("#000000"),
			}
			Expect(w.Create(ctx, labeled)).To(Succeed())
			Expect(w.Label(ctx, label.OntologyID(labeled.Key), []uuid.UUID{l.Key})).To(Succeed())
			labels := MustSucceed(svc.RetrieveFor(ctx, label.OntologyID(labeled.Key), tx))
			Expect(labels).To(HaveLen(1))
			Expect(labels[0].Key).To(Equal(l.Key))
		})
	})
	Describe("RemoveLabel", func() {
		It("Should remove a label", func() {
			l := &label.Label{
				Name:  "Label",
				Color: color.Color("#000000"),
			}
			Expect(w.Create(ctx, l)).To(Succeed())
			labeled := &label.Label{
				Name:  "Labeled",
				Color: color.Color("#000000"),
			}
			Expect(w.Create(ctx, labeled)).To(Succeed())
			Expect(w.Label(ctx, label.OntologyID(labeled.Key), []uuid.UUID{l.Key})).To(Succeed())
			labels := MustSucceed(svc.RetrieveFor(ctx, label.OntologyID(labeled.Key), tx))
			Expect(labels).To(HaveLen(1))
			Expect(labels[0].Key).To(Equal(l.Key))
			Expect(w.RemoveLabel(ctx, label.OntologyID(labeled.Key), []uuid.UUID{l.Key})).To(Succeed())
			labels = MustSucceed(svc.RetrieveFor(ctx, label.OntologyID(labeled.Key), tx))
			Expect(labels).To(HaveLen(0))
		})
	})
})