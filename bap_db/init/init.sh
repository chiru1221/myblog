mongoimport -u root -p mongo --authenticationDatabase admin --db database --collection profile --file ./docker-entrypoint-initdb.d/profile.json
mongoimport -u root -p mongo --authenticationDatabase admin --db database --collection blog --file ./docker-entrypoint-initdb.d/blog.json
