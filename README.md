# Spring-2023-Group 8

CSCI 6234 Object Oriented Design Final Project

# TOJ

An online practice system based on Gin, Gorm, and Vue

Background language: Golang

Framework: Gin, GORM

Front-end framework: Vue, ElementUI

# Integrating Swagger

Reference document: https://github.com/swaggo/gin-swagger interface access address: http://localhost:8080/swagger/index.html

// GetProblemList
// @Tags public method
// @Summary list of questions
// @Param page query int false "page"
// @Param size query int false "size"
// @Success 200 {string} json "{"code":"200","msg","","data":""}"
// @Router /problem-list [get]

# Install jwt

go get github.com/dgrijalva/jwt-go

