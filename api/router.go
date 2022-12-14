package api

func (s *Server) registerRoutes() {
	v1 := s.router.Group("/v1")

	{
		v1.POST("/initiateMultipartUpload", s.upload.InitiateMultipartUpload)
		v1.POST("/uploadPart", s.upload.UploadPart)
		v1.POST("/uploadComplete", s.upload.UploadComplete)
		v1.POST("/uploadList", s.upload.UploadList)
	}

}
