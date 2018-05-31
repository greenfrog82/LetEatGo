* postgres
    docker pull postgres
    docker run -d -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password123 --name postgres-db -p 5432:5432 postgres
    
* pgadmin4
    docker pull dpage/pgadmin4
    docker run -it --rm --link postgres-db:postgres-db -e PGADMIN_DEFAULT_EMAIL=user@foo.com -e PGADMIN_DEFAULT_PASSWORD=password123 -p 80:80 dpage/pgadmin4:latest    
