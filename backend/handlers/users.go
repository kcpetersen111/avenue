package handlers

import "github.com/gin-gonic/gin"

func (s *Server) Login(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}

func (s *Server) Logout(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}

func (s *Server) Register(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}

func (s *Server) GetProfile(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}

func (s *Server) UpdateProfile(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}

func (s *Server) UpdatePassword(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}
