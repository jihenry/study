var cr1ypto = require('crypto');

module.exports = {
    md5: function (data:string) {
        return cr1ypto.createHash('md5').update(data).digest("hex");
    }
};