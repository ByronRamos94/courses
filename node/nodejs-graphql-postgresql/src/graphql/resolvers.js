const {
  getProduct,
  getProducts,
  addProduct,
  deleteProduct,
  updateProduct,
getProductsByCategory,
} = require('./product.resolvers');
const { addCategory, getCategory } = require('./category.resolvers');
const { login } = require('./auth.resolvers');
const { RegularExpression } = require('graphql-scalars')

const CategoryNameType = new RegularExpression('CategoryNameType', /^[a-zA-Z0-9]{3,8}$/)

const resolvers = {
  Query: {
    hello: () => 'graphql happy coding!',
    product: getProduct,
    products: getProducts,
    category: getCategory
  },
  Mutation: {
    login,
    addProduct,
    updateProduct,
    deleteProduct,
    addCategory,
  },
  CategoryNameType,
  Category: {
    products: getProductsByCategory
  }
};

module.exports = resolvers;
