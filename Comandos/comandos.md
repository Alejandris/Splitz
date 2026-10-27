 Comandos Docker para reiniciar staging
Desde la raíz del proyecto:
docker compose --env-file .env.staging -f docker-compose.staging.yml down

Levantar nuevamente reconstruyendo las imágenes:
docker compose --env-file .env.staging -f docker-compose.staging.yml up -d --build

Verificar el estado:
docker compose --env-file .env.staging -f docker-compose.staging.yml ps

Ver logs del backend:
docker compose --env-file .env.staging -f docker-compose.staging.yml logs -f backend

Ver logs del frontend:
docker compose --env-file .env.staging -f docker-compose.staging.yml logs -f frontend

Validar salud de la API:
curl http://localhost:8080/health

Validar frontend:
curl http://localhost:5174

Para un reinicio rápido sin reconstruir imágenes:
docker compose --env-file .env.staging -f docker-compose.staging.yml restart