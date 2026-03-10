/**
 * @api {get} /api/orders Lista pedidos
 * @apiSuccess (200) Lista de pedidos retornada com sucesso
 * @apiError (404) Nenhum pedido encontrado
 * @apiError (500) Erro interno do servidor
 */
function getOrders(req, res) {
  // implementação
}

/**
 * @api {post} /api/orders Cria novo pedido
 * @apiSuccess (201) Pedido criado com sucesso
 * @apiError (400) Dados do pedido inválidos
 * @apiError (402) Pagamento necessário
 */
function createOrder(req, res) {
  // implementação
}

/**
 * @api {get} /api/orders/{id} Busca pedido por ID
 * @apiSuccess (200) Pedido encontrado
 * @apiError (404) Pedido não encontrado
 */
function getOrderById(req, res) {
  // implementação
}
