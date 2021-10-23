package vendors

type Manager struct {
	dbConnector DBConnector
}

func NewManager(dbConnector DBConnector) *Manager {
	return &Manager{
		dbConnector: dbConnector,
	}
}

type DBConnector interface {
	List(collectionName string) (map[string]interface{}, error)
}

func (m *Manager) List() (map[string]interface{}, error) {
	return m.dbConnector.List("vendors")
}
