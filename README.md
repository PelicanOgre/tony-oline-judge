# Spring-2023-Group 8

CSCI 6234 Object Oriented Design Final Project

# Deom

[TOJ](http://124.221.199.57/dist/#/questionList)

# TOJ

An online practice system based on Gin, Gorm, and Vue

Background language: Golang
Useful Link: https://go.dev/doc/
Framework: Gin, GORM
Useful Link: https://gin-gonic.com/docs/
Useful Link: https://gorm.io/docs/index.html

Front-end framework: Vue, ElementUI
Useful Link: https://vuejs.org/guide/introduction.html
Useful Link: https://github.com/ElemeFE/element

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

# Configuration

Configuring MailPassword to an environment variable

# System Modules

1. User Module
   Password Login
   Email Registration
   User details
2. Title Management Module
   Question list, question details
   Question generation, question modification
3. Category Management Module
   Category list
   Category creation, category modification, category deletion
4. Judgment Module
   Submission Record List
   Code submission and judgment
5. Ranking Module
   Ranking of the list of cases
