const express = require('express');
const app = express();
const router = express.Router();

// Estes endpoints serão detectados automaticamente pelo JSExpressExtractor

app.get('/api/products', (req, res) => {
  // Lista todos os produtos
  res.json({ products: [] });
});

app.post('/api/products', (req, res) => {
  // Cria um novo produto
  res.status(201).json({ message: 'Produto criado' });
});

router.put('/api/products/:id', (req, res) => {
  // Atualiza um produto
  res.json({ message: 'Produto atualizado' });
});

router.delete('/api/products/:id', (req, res) => {
  // Remove um produto
  res.status(204).send();
});

app.use(router);
