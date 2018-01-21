'use strict';
module.exports = (sequelize, DataTypes) => {
    var Dataalat = sequelize.define('Dataalat', {
        TeganganMikro:DataTypes.STRING,
        ArusMikro:DataTypes.STRING,
        Waktu:DataTypes.STRING,
        PMFrekuensi:DataTypes.STRING,
        PMCospi:DataTypes.STRING,
        PMTegangan:DataTypes.STRING,
        PMArus:DataTypes.STRING,
    });
    return Dataalat;
};