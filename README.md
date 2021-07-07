## Portto API Service

### Usage

Start service
```
docker-compose up -d
curl http://localhost:3000/blocks\?limit=5
curl http://localhost:3000/blocks/5
curl http://localhost:3000/transaction/:txHash
```

Stop service
```
docker-compose down
docker rmi portto_api_api-service portto_api_indexer-service
```