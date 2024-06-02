docker run -it --name some-postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=postgres \
  -p 5432:5432 \
  --sig-proxy=true \
  -v postgres_data:/var/lib/postgresql/data postgres:16