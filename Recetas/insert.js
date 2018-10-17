use REST_API;
//var bulk = db.recetas.initializeUnorderedBulkOp();
//bulk.insert({pos: "1", nombre: "Hamburguesa", ingredientes:["pan", "tomate", "carne"], elaboracion: "hacer rica hamburguesa"});
//bulk.execute();
db.recetas.drop();
db.recetas.find().count();