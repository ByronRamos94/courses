const boom = require('@hapi/boom')

function checkRolesGql(user, ...roles) {
    if (!roles.includes(user.role)){
        throw boom.unauthorized('this role cannot perform this operation')
    }
}
async function validateJwtGql(context) {
    const { user } = await context.authenticate('jwt', { session: false });
    if (!user) {
      throw boom.unauthorized('invalid token');
    }
    return user
}

module.exports = { checkRolesGql, validateJwtGql }