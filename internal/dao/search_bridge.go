package dao

import (
	"github.com/Masterminds/semver/v3"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/sirupsen/logrus"
)

func (s *bridgeTweetSearchServant) Name() string {
	return "BridgeTweetSearch"
}

func (s *bridgeTweetSearchServant) Version() *semver.Version {
	return semver.MustParse("v0.1.0")
}

func (s *bridgeTweetSearchServant) IndexName() string {
	return s.ts.IndexName()
}

func (s *bridgeTweetSearchServant) AddDocuments(data core.DocItems, primaryKey ...string) (bool, error) {
	s.updateDocs(&documents{
		primaryKey: primaryKey,
		docItems:   data,
	})
	return true, nil
}

func (s *bridgeTweetSearchServant) DeleteDocuments(identifiers []string) error {
	s.updateDocs(&documents{
		identifiers: identifiers,
	})
	return nil
}

func (s *bridgeTweetSearchServant) Search(q *core.QueryReq, offset, limit int) (*core.QueryResp, error) {
	return s.ts.Search(q, offset, limit)
}

func (s *bridgeTweetSearchServant) updateDocs(doc *documents) {
	select {
	case s.updateDocsCh <- doc:
		logrus.Debugln("addDocuments send documents by chan")
	default:
		go func(ch chan<- *documents, item *documents) {
			s.updateDocsCh <- item
			logrus.Debugln("addDocuments send documents by goroutine")
		}(s.updateDocsCh, doc)
	}
}

func (s *bridgeTweetSearchServant) startUpdateDocs() {
	for doc := range s.updateDocsCh {
		if len(doc.docItems) > 0 {
			if _, err := s.ts.AddDocuments(doc.docItems, doc.primaryKey...); err != nil {
				logrus.Errorf("addDocuments occurs error: %v", err)
			}
		}
		if len(doc.identifiers) > 0 {
			if err := s.ts.DeleteDocuments(doc.identifiers); err != nil {
				logrus.Errorf("deleteDocuments occurs error: %s", err)
			}
		}
	}
}
