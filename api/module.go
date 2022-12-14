package api

import "go-interface/client"

type modules struct {
	fileSystem client.FileSystem
	upload     *UploadMultipart
}

func (m *modules) initModules(s *Server) error {
	var err error
	m.fileSystem, err = client.NewFileSystem(s.config)
	if err != nil {
		return err
	}

	m.upload, err = NewUploadMultipart(m.fileSystem)
	if err != nil {
		return err
	}

	return nil
}
