// Copyright 2023 Synnax Labs, Inc.
//
// Use of this software is governed by the Business Source License included in the file
// licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with the Business Source
// License, use of this software will be governed by the Apache License, Version 2.0,
// included in the file licenses/APL.txt.

package domain_test

import (
	"encoding/binary"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/samber/lo"
	"github.com/synnaxlabs/cesium/internal/core"
	"github.com/synnaxlabs/cesium/internal/domain"
	"github.com/synnaxlabs/x/config"
	xfs "github.com/synnaxlabs/x/io/fs"
	"github.com/synnaxlabs/x/telem"
	. "github.com/synnaxlabs/x/testutil"
	"github.com/synnaxlabs/x/validate"
	"os"
	"sync"
	"time"
)

func extractPointer(f xfs.File) (p struct {
	telem.TimeRange
	fileKey uint16
	offset  uint32
	length  uint32
}) {
	var b = make([]byte, 26)
	_, err := f.ReadAt(b, 0)
	Expect(err).ToNot(HaveOccurred())
	p.TimeRange.Start = telem.TimeStamp(binary.LittleEndian.Uint64(b[0:8]))
	p.TimeRange.End = telem.TimeStamp(binary.LittleEndian.Uint64(b[8:16]))
	p.fileKey = binary.LittleEndian.Uint16(b[16:18])
	p.offset = binary.LittleEndian.Uint32(b[18:22])
	p.length = binary.LittleEndian.Uint32(b[22:26])

	return
}

var _ = Describe("WriterBehavior", func() {
	for fsName, makeFS := range fileSystems {
		fs := makeFS()
		Context("FS: "+fsName, func() {
			var db *domain.DB
			BeforeEach(func() {
				db = MustSucceed(domain.Open(domain.Config{FS: MustSucceed(fs.Sub(rootPath))}))
			})
			AfterEach(func() {
				Expect(db.Close()).To(Succeed())
				Expect(fs.Remove(rootPath)).To(Succeed())
			})
			Describe("Start Validation", func() {
				Context("No domain overlap", func() {
					It("Should successfully open the writer", func() {
						w := MustSucceed(db.NewWriter(ctx, domain.WriterConfig{
							Start: 10 * telem.SecondTS,
						}))
						Expect(w.Close(ctx)).To(Succeed())
					})
				})
				Context("TimeRange overlap", func() {
					It("Should fail to open the writer", func() {
						w := MustSucceed(db.NewWriter(
							ctx,
							domain.WriterConfig{
								Start: 10 * telem.SecondTS,
							}))
						Expect(w.Write([]byte{1, 2, 3, 4, 5, 6})).To(Equal(6))
						Expect(w.Commit(ctx, 15*telem.SecondTS)).To(Succeed())
						Expect(w.Close(ctx)).To(Succeed())
						_, err := db.NewWriter(ctx, domain.WriterConfig{
							Start: 10 * telem.SecondTS,
						})
						Expect(err).To(HaveOccurredAs(domain.ErrDomainOverlap))
					})
				})
			})
			Describe("End Validation", func() {
				Context("No domain overlap", func() {
					It("Should successfully commit", func() {
						w := MustSucceed(db.NewWriter(ctx, domain.WriterConfig{
							Start: 10 * telem.SecondTS,
						}))
						MustSucceed(w.Write([]byte{1, 2, 3, 4, 5, 6}))
						Expect(w.Commit(ctx, 20*telem.SecondTS)).To(Succeed())
						Expect(w.Close(ctx)).To(Succeed())
					})
				})
				Context("TimeRange overlap", func() {
					It("Should fail to commit", func() {
						w := MustSucceed(db.NewWriter(ctx, domain.WriterConfig{
							Start: 10 * telem.SecondTS,
						}))
						MustSucceed(w.Write([]byte{1, 2, 3, 4, 5, 6}))
						Expect(w.Commit(ctx, 20*telem.SecondTS)).To(Succeed())
						Expect(w.Close(ctx)).To(Succeed())
						w = MustSucceed(db.NewWriter(ctx, domain.WriterConfig{
							Start: 4 * telem.SecondTS,
						}))
						MustSucceed(w.Write([]byte{1, 2, 3, 4, 5, 6}))
						Expect(w.Commit(ctx, 15*telem.SecondTS)).To(HaveOccurredAs(domain.ErrDomainOverlap))
						Expect(w.Close(ctx)).To(Succeed())
					})
					It("Should fail to commit an update to a writer", func() {
						w := MustSucceed(db.NewWriter(ctx, domain.WriterConfig{
							Start: 10 * telem.SecondTS,
						}))
						MustSucceed(w.Write([]byte{1, 2, 3, 4, 5, 6}))
						Expect(w.Commit(ctx, 20*telem.SecondTS)).To(Succeed())
						Expect(w.Close(ctx)).To(Succeed())
						w = MustSucceed(db.NewWriter(ctx, domain.WriterConfig{
							Start: 4 * telem.SecondTS,
						}))
						MustSucceed(w.Write([]byte{1, 2, 3, 4}))
						Expect(w.Commit(ctx, 8*telem.SecondTS)).To(Succeed())
						Expect(w.Commit(ctx, 15*telem.SecondTS)).To(HaveOccurredAs(domain.ErrDomainOverlap))
						Expect(w.Close(ctx)).To(Succeed())
					})
				})
				Context("Writing past preset end", func() {
					It("Should fail to commit", func() {
						w := MustSucceed(db.NewWriter(ctx, domain.WriterConfig{
							Start: 10 * telem.SecondTS,
							End:   20 * telem.SecondTS,
						}))
						MustSucceed(w.Write([]byte{1, 2, 3, 4, 5, 6}))
						Expect(w.Commit(ctx, 30*telem.SecondTS)).To(MatchError(ContainSubstring("commit timestamp cannot be greater than preset end")))
						Expect(w.Close(ctx)).To(Succeed())
					})
				})
				Context("Commit before start", func() {
					It("Should fail to commit", func() {
						w := MustSucceed(db.NewWriter(ctx, domain.WriterConfig{
							Start: 10 * telem.SecondTS,
						}))
						MustSucceed(w.Write([]byte{1, 2, 3, 4, 5, 6}))
						Expect(w.Commit(ctx, 5*telem.SecondTS)).To(HaveOccurredAs(validate.Error))
						Expect(w.Close(ctx)).To(Succeed())
					})
				})
				Describe("End of one domain is the start of another", func() {
					It("Should successfully commit", func() {
						w := MustSucceed(db.NewWriter(ctx, domain.WriterConfig{
							Start: 10 * telem.SecondTS,
						}))
						MustSucceed(w.Write([]byte{1, 2, 3, 4, 5, 6}))
						Expect(w.Commit(ctx, 20*telem.SecondTS)).To(Succeed())
						Expect(w.Close(ctx)).To(Succeed())
						w = MustSucceed(db.NewWriter(ctx, domain.WriterConfig{
							Start: 20 * telem.SecondTS,
						}))
						MustSucceed(w.Write([]byte{1, 2, 3, 4, 5, 6}))
						Expect(w.Commit(ctx, 30*telem.SecondTS)).To(Succeed())
						Expect(w.Close(ctx)).To(Succeed())
					})
				})
				Context("Multi Commit", func() {
					It("Should correctly commit a writer multiple times", func() {
						w := MustSucceed(db.NewWriter(ctx, domain.WriterConfig{
							Start: 10 * telem.SecondTS,
						}))
						MustSucceed(w.Write([]byte{1, 2, 3, 4, 5, 6}))
						Expect(w.Commit(ctx, 20*telem.SecondTS)).To(Succeed())
						MustSucceed(w.Write([]byte{1, 2, 3, 4, 5, 6}))
						Expect(w.Commit(ctx, 30*telem.SecondTS)).To(Succeed())
						Expect(w.Close(ctx)).To(Succeed())
					})
					Context("Commit before previous commit", func() {
						It("Should fail to commit", func() {
							w := MustSucceed(db.NewWriter(ctx, domain.WriterConfig{
								Start: 10 * telem.SecondTS,
							}))
							MustSucceed(w.Write([]byte{1, 2, 3, 4, 5, 6}))
							Expect(w.Commit(ctx, 15*telem.SecondTS)).To(Succeed())
							Expect(w.Commit(ctx, 14*telem.SecondTS)).To(HaveOccurredAs(validate.Error))
							Expect(w.Close(ctx)).To(Succeed())
						})
					})
				})
				Context("Concurrent Writes", func() {
					It("Should fail to commit all but one of the writes", func() {
						writerCount := 20
						errors := make([]error, writerCount)
						writers := make([]*domain.Writer, writerCount)
						var wg sync.WaitGroup
						wg.Add(writerCount)
						for i := 0; i < writerCount; i++ {
							writers[i] = MustSucceed(db.NewWriter(ctx, domain.WriterConfig{
								Start: 10 * telem.SecondTS,
							}))
						}
						for i, w := range writers {
							i, w := i, w
							go func(i int, w *domain.Writer) {
								defer wg.Done()
								MustSucceed(w.Write([]byte{1, 2, 3, 4, 5, 6}))
								errors[i] = w.Commit(ctx, 15*telem.SecondTS)
							}(i, w)
						}
						wg.Wait()

						occurred := lo.Filter(errors, func(err error, i int) bool {
							return err != nil
						})
						Expect(occurred).To(HaveLen(writerCount - 1))
						for _, err := range occurred {
							Expect(err).To(HaveOccurredAs(domain.ErrDomainOverlap))
						}
					})
				})
			})
			Describe("AutoPersist", func() {
				It("Should persist to disk every subsequent call after the set time interval", func() {
					By("Opening a writer")
					w := MustSucceed(db.NewWriter(ctx, domain.WriterConfig{Start: 10 * telem.SecondTS, EnableAutoCommit: config.True(), AutoIndexPersistInterval: 50 * telem.Millisecond}))

					modTime := MustSucceed(db.FS.Stat("index.domain")).ModTime()

					By("Writing some data and committing it right after")
					_, err := w.Write([]byte{6, 7, 8, 9, 10})
					Expect(err).ToNot(HaveOccurred())
					Expect(w.Commit(ctx, 20*telem.SecondTS+1)).To(Succeed())

					_, err = w.Write([]byte{11, 12, 13, 14, 15})
					Expect(err).ToNot(HaveOccurred())
					Expect(w.Commit(ctx, 25*telem.SecondTS+1)).To(Succeed())

					By("Asserting that the previous commits have not been persisted")
					s := MustSucceed(db.FS.Stat("index.domain"))
					Expect(s.Size()).To(Equal(int64(0)))

					By("Sleeping for some time")
					time.Sleep(time.Duration(50 * telem.Millisecond))
					_, err = w.Write([]byte{16, 17, 18, 19, 20})
					Expect(err).ToNot(HaveOccurred())
					Expect(w.Commit(ctx, 30*telem.SecondTS+1)).To(Succeed())

					By("Asserting that the commits will be persisted the next time we use the method after the set time interval")
					Eventually(func() time.Time {
						return MustSucceed(db.FS.Stat("index.domain")).ModTime()
					}).ShouldNot(Equal(modTime))

					f := MustSucceed(db.FS.Open("index.domain", os.O_RDONLY))
					p := extractPointer(f)
					Expect(p.End).To(Equal(30*telem.SecondTS + 1))
					Expect(p.length).To(Equal(uint32(15)))
				})

				It("Should persist to disk every time when the interval is set to always persist", func() {
					By("Opening a writer")
					w := MustSucceed(db.NewWriter(ctx, domain.WriterConfig{Start: 10 * telem.SecondTS, EnableAutoCommit: config.True(), AutoIndexPersistInterval: domain.AlwaysIndexPersistOnAutoCommit}))

					By("Writing some data and committing it")
					_, err := w.Write([]byte{1, 2, 3, 4, 5})
					Expect(err).ToNot(HaveOccurred())
					Expect(w.Commit(ctx, 15*telem.SecondTS+1)).To(Succeed())

					By("Asserting that the previous commit has been persisted")
					f := MustSucceed(db.FS.Open("index.domain", os.O_RDONLY))
					p := extractPointer(f)
					Expect(f.Close()).To(Succeed())
					Expect(p.End).To(Equal(15*telem.SecondTS + 1))
					Expect(p.length).To(Equal(uint32(5)))

					By("Writing some data and committing it with auto persist right after")
					_, err = w.Write([]byte{6, 7, 8, 9, 10})
					Expect(w.Commit(ctx, 20*telem.SecondTS+1)).To(Succeed())

					By("Asserting that the previous commit has been persisted")
					f = MustSucceed(db.FS.Open("index.domain", os.O_RDONLY))
					p = extractPointer(f)
					Expect(f.Close()).To(Succeed())
					Expect(p.End).To(Equal(20*telem.SecondTS + 1))
					Expect(p.length).To(Equal(uint32(10)))

					By("Writing some data and committing it with auto persist right after")
					_, err = w.Write([]byte{11, 12, 13, 14, 15})
					Expect(w.Commit(ctx, 25*telem.SecondTS+1)).To(Succeed())

					By("Asserting that the previous commits have not been persisted")
					f = MustSucceed(db.FS.Open("index.domain", os.O_RDONLY))
					p = extractPointer(f)
					Expect(f.Close()).To(Succeed())
					Expect(p.End).To(Equal(25*telem.SecondTS + 1))
					Expect(p.length).To(Equal(uint32(15)))
				})

				It("Should persist any unpersisted, but committed (stranded) data on close", func() {
					By("Opening a writer")
					w := MustSucceed(db.NewWriter(ctx, domain.WriterConfig{Start: 10 * telem.SecondTS, EnableAutoCommit: config.True(), AutoIndexPersistInterval: 10 * telem.Second}))

					By("Writing some data and committing it")
					_, err := w.Write([]byte{1, 2, 3, 4, 5})
					Expect(err).ToNot(HaveOccurred())
					Expect(w.Commit(ctx, 15*telem.SecondTS+1)).To(Succeed())

					By("Writing some data and committing it")
					_, err = w.Write([]byte{6, 7, 8, 9, 10})
					Expect(err).ToNot(HaveOccurred())
					Expect(w.Commit(ctx, 20*telem.SecondTS+1)).To(Succeed())

					By("Closing the writer")
					Expect(w.Close(ctx)).To(Succeed())

					By("Asserting that the commit has been persisted")
					f := MustSucceed(db.FS.Open("index.domain", os.O_RDONLY))
					p := extractPointer(f)
					Expect(f.Close()).To(Succeed())
					Expect(p.End).To(Equal(20*telem.SecondTS + 1))
					Expect(p.length).To(Equal(uint32(10)))
				})

				It("Should always persist if auto commit is not enabled, no matter the interval", func() {
					By("Opening a writer")
					w := MustSucceed(db.NewWriter(ctx, domain.WriterConfig{Start: 10 * telem.SecondTS, AutoIndexPersistInterval: 1 * telem.Hour}))

					By("Writing some data and committing it")
					_, err := w.Write([]byte{1, 2, 3, 4, 5})
					Expect(err).ToNot(HaveOccurred())
					Expect(w.Commit(ctx, 15*telem.SecondTS+1)).To(Succeed())

					By("Writing some data and committing it")
					_, err = w.Write([]byte{6, 7, 8, 9, 10})
					Expect(err).ToNot(HaveOccurred())
					Expect(w.Commit(ctx, 20*telem.SecondTS+1)).To(Succeed())

					By("Asserting that the commit has been persisted")
					f := MustSucceed(db.FS.Open("index.domain", os.O_RDONLY))
					p := extractPointer(f)
					Expect(f.Close()).To(Succeed())
					Expect(p.End).To(Equal(20*telem.SecondTS + 1))
					Expect(p.length).To(Equal(uint32(10)))

					By("Closing the writer")
					Expect(w.Close(ctx)).To(Succeed())
				})
			})
			Describe("Close", func() {
				It("Should not allow operations on a closed writer", func() {
					var (
						w = MustSucceed(db.NewWriter(ctx, domain.WriterConfig{Start: 10 * telem.SecondTS}))
						e = core.EntityClosed("domain.writer")
					)
					Expect(w.Close(ctx)).To(Succeed())
					Expect(w.Commit(ctx, telem.TimeStampMax)).To(MatchError(e))
					_, err := w.Write([]byte{1, 2, 3})
					Expect(err).To(MatchError(e))
					Expect(w.Close(ctx)).To(Succeed())
				})
			})
		})
	}
})
