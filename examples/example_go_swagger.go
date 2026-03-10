package main

import "net/http"

// @Summary Lista todos os usuários
// @Description Retorna uma lista paginada de usuários do sistema
// @Success 200 {array} User Lista de usuários
// @Failure 500 {object} Error Erro interno do servidor
// @Router /api/users [get]
func GetUsersSwagger(w http.ResponseWriter, r *http.Request) {
	// implementação
}

// @Summary Cria um novo usuário
// @Description Adiciona um novo usuário ao sistema
// @Success 201 {object} User Usuário criado com sucesso
// @Failure 400 {object} Error Dados inválidos
// @Failure 409 {object} Error Usuário já existe
// @Router /api/users [post]
func CreateUserSwagger(w http.ResponseWriter, r *http.Request) {
	// implementação
}
