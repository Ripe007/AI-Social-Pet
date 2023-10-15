docker stop social-pet
docker rm social-pet
docker rmi social-pet-api
./release.bat
docker build -t social-pet-api .
docker run --name social-pet -v /home/app/static:/app/static --restart always -p 8080:8080 -d social-pet-api
docker exec -it social-pet /bin/bash