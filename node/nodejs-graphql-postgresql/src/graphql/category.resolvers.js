const { checkRolesGql, validateJwtGql } = require('../utils/authGql');
const CategoryService = require('./../services/category.service');
const service = new CategoryService();

const addCategory = async (_, { dto }, context) => {
  const user = await validateJwtGql(context);
  checkRolesGql(user, 'admin');
  return service.create({ ...dto, image: dto.image.href });
};
const getCategory = (_, { id }) => {
    return service.findOne(id);
}

module.exports = { addCategory, getCategory };
