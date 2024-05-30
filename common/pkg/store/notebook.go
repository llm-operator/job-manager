package store

import (
	"fmt"

	v1 "github.com/llm-operator/job-manager/api/v1"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"
)

// NotebookState is the state of a notebook.
type NotebookState string

const (
	// NotebookStateQueued is the state of a notebook that is waiting to be scheduled.
	NotebookStateQueued NotebookState = "queued"
	// NotebookStateRunning is the state of a notebook that is currently running.
	NotebookStateRunning NotebookState = "running"
	// NotebookStateStopping is the state of a notebook that is stopping.
	NotebookStateStopping NotebookState = "stopping"
	// NotebookStateStopped is the state of a notebook that has been stopped.
	NotebookStateStopped NotebookState = "stopped"
	// NotebookStateFailed is the state of a notebook that has failed.
	NotebookStateFailed NotebookState = "failed"
)

// Notebook is a model of notebook.
type Notebook struct {
	gorm.Model

	NotebookID string `gorm:"uniqueIndex"`

	Image string

	// Message is the marshalled JSON of the v1.Notebook.
	Message []byte

	State NotebookState

	TenantID            string
	OrganizationID      string
	ProjectID           string `gorm:"index"`
	KubernetesNamespace string

	Version int
}

// V1Notebook converts a notebook to a v1.Notebook.
func (n *Notebook) V1Notebook() (*v1.Notebook, error) {
	var nbProto v1.Notebook
	if err := proto.Unmarshal(n.Message, &nbProto); err != nil {
		return nil, err
	}
	nbProto.Status = string(n.State)
	return &nbProto, nil
}

// MutateMessage mutates the message field of a notebook.
func (n *Notebook) MutateMessage(mutateFn func(nb *v1.Notebook)) error {
	nbProto, err := n.V1Notebook()
	if err != nil {
		return err
	}
	mutateFn(nbProto)
	msg, err := proto.Marshal(nbProto)
	if err != nil {
		return err
	}
	n.Message = msg
	return nil
}

// CreateNotebook creates a new notebook.
func (s *S) CreateNotebook(nb *Notebook) error {
	return s.db.Create(nb).Error
}

// GetNotebookByIDAndProjectID gets a notebook by its notebook ID and tenant ID.
func (s *S) GetNotebookByIDAndProjectID(id, projectID string) (*Notebook, error) {
	var nb Notebook
	if err := s.db.Where("notebook_id = ? AND project_id = ?", id, projectID).Take(&nb).Error; err != nil {
		return nil, err
	}
	return &nb, nil
}

// ListNotebooksByProjectIDWithPagination finds notebooks with pagination. Notebooks are returned with a descending order of ID.
func (s *S) ListNotebooksByProjectIDWithPagination(projectID string, afterID uint, limit int) ([]*Notebook, bool, error) {
	var nbs []*Notebook
	q := s.db.Where("project_id = ?", projectID)
	if afterID > 0 {
		q = q.Where("id < ?", afterID)
	}
	if err := q.Order("id DESC").Limit(limit + 1).Find(&nbs).Error; err != nil {
		return nil, false, err
	}

	var hasMore bool
	if len(nbs) > limit {
		nbs = nbs[:limit]
		hasMore = true
	}
	return nbs, hasMore, nil
}

// ListQueuedOrStoppingNotebooks finds queued notebooks.
func (s *S) ListQueuedOrStoppingNotebooks() ([]*Notebook, error) {
	var nbs []*Notebook
	if err := s.db.Where("state = ? OR state = ?", NotebookStateQueued, NotebookStateStopping).Find(&nbs).Error; err != nil {
		return nil, err
	}
	return nbs, nil
}

// UpdateNotebookState updates a notebook state.
func (s *S) UpdateNotebookState(id string, currentVersion int, newState NotebookState) error {
	result := s.db.Model(&Notebook{}).
		Where("notebook_id = ?", id).
		Where("version = ?", currentVersion).
		Updates(map[string]interface{}{
			"state":   newState,
			"version": currentVersion + 1,
		})
	if err := result.Error; err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("update notebook: %w", ErrConcurrentUpdate)
	}
	return nil
}

// UpdateNotebookStateAndMessage updates a notebook state and message.
func (s *S) UpdateNotebookStateAndMessage(id string, currentVersion int, newState NotebookState, message []byte) error {
	result := s.db.Model(&Notebook{}).
		Where("notebook_id = ?", id).
		Where("version = ?", currentVersion).
		Updates(map[string]interface{}{
			"state":   newState,
			"message": message,
			"version": currentVersion + 1,
		})
	if err := result.Error; err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("update notebook: %w", ErrConcurrentUpdate)
	}
	return nil
}
